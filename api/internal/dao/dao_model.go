package dao

import (
	"context"
	"sync"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

// 定义接口
type DaoInterface interface {
	CtxDaoModel(ctx context.Context, dbOpt ...any) *DaoModel
	ParseDbGroup(ctx context.Context, dbGroupOpt ...any) string
	ParseDbTable(ctx context.Context, dbTableOpt ...any) string
	ParseId(daoModel *DaoModel) string
	ParseLabel(daoModel *DaoModel) string
	ParseFilter(filter map[string]any, daoModel *DaoModel) gdb.ModelHandler
	ParseField(field []string, fieldWithParam map[string]any, daoModel *DaoModel) gdb.ModelHandler
	HandleAfterField(ctx context.Context, record gdb.Record, daoModel *DaoModel)
	HookSelect(daoModel *DaoModel) gdb.HookHandler
	ParseInsert(insert map[string]any, daoModel *DaoModel) gdb.ModelHandler
	HookInsert(daoModel *DaoModel) gdb.HookHandler
	ParseUpdate(update map[string]any, daoModel *DaoModel) gdb.ModelHandler
	HookUpdate(daoModel *DaoModel) gdb.HookHandler
	HookDelete(daoModel *DaoModel) gdb.HookHandler
	ParseGroup(group []string, daoModel *DaoModel) gdb.ModelHandler
	ParseOrder(order []string, daoModel *DaoModel) gdb.ModelHandler
	ParseJoin(joinTable string, daoModel *DaoModel) gdb.ModelHandler
}

type DaoModel struct {
	Ctx context.Context
	dao DaoInterface
	db  gdb.DB
	*gdb.Model
	DbGroup      string // 分库情况下，解析后所确定的库
	DbTable      string // 分表情况下，解析后所确定的表
	JoinTableMap map[string]struct{}
	AfterField   map[string]any
	AfterInsert  map[string]any
	AfterUpdate  map[string]any
	SaveData     map[string]any
	IdArr        []gdb.Value // 新增需要后置处理且主键非自增时 或 更新|删除需要后置处理时 使用。注意：一般在新增|更新|删除方法执行前调用（即在各种sql条件设置完后）
}

// 对象池。性能提醒不明显，暂时不用。确实大幅减少了对象创建和销毁（内存压力减少），但却需要手动增加放入对象池的代码：defer daoModel.PutPool()
var poolDaoModel = sync.Pool{
	New: func() any { return &DaoModel{} },
}

func (daoModelThis *DaoModel) PutPool() {
	daoModelThis.Ctx = nil
	daoModelThis.dao = nil
	daoModelThis.db = nil
	daoModelThis.Model = nil
	daoModelThis.DbGroup = ``
	daoModelThis.DbTable = ``
	daoModelThis.JoinTableMap = nil
	daoModelThis.AfterField = nil
	daoModelThis.AfterInsert = nil
	daoModelThis.AfterUpdate = nil
	daoModelThis.SaveData = nil
	daoModelThis.IdArr = nil
	poolDaoModel.Put(daoModelThis)
}

// 注意：dbOpt存在时，dbOpt[0]解析DbTable，dbOpt[1]索引参数解析DbGroup
func NewDaoModel(ctx context.Context, dao DaoInterface, dbOpt ...any) *DaoModel {
	daoModelObj := &DaoModel{} // poolDaoModel.Get().(*DaoModel)
	daoModelObj.Ctx = ctx
	daoModelObj.dao = dao
	daoModelObj.JoinTableMap = map[string]struct{}{}
	daoModelObj.AfterField = map[string]any{}
	daoModelObj.AfterInsert = map[string]any{}
	daoModelObj.AfterUpdate = map[string]any{}
	daoModelObj.SaveData = map[string]any{}
	switch len(dbOpt) {
	case 1:
		daoModelObj.DbTable = daoModelObj.dao.ParseDbTable(ctx, dbOpt[0])
		daoModelObj.DbGroup = daoModelObj.dao.ParseDbGroup(ctx)
	case 2:
		daoModelObj.DbTable = daoModelObj.dao.ParseDbTable(ctx, dbOpt[0])
		daoModelObj.DbGroup = daoModelObj.dao.ParseDbGroup(ctx, dbOpt[1])
	default:
		daoModelObj.DbTable = daoModelObj.dao.ParseDbTable(ctx)
		daoModelObj.DbGroup = daoModelObj.dao.ParseDbGroup(ctx)
	}
	daoModelObj.db = g.DB(daoModelObj.DbGroup)
	daoModelObj.Model = daoModelObj.newModel()
	return daoModelObj
}

// 缓存暂不考虑。原因1：当key很多时，删除操作会有效率问题。原因2：直接清空全部缓存也很不友好
//	func (daoModelThis *DaoModel) SetCache() *DaoModel {
//		daoModelThis.Cache(gdb.CacheOption{
//			Duration: 0, // 5 * time.Minute
//			Name:     "#`" + daoModelThis.DbTable + "`#All",
//			Force:    true,
//		})
//		/* // daoModelThis.GetCore().ClearCache(daoModelThis.Ctx, daoModelThis.DbTable)
//		cache := daoModelThis.GetDB().GetCache()
//		keyArr := cache.MustKeyStrings(daoModelThis.Ctx)
//		keyArrOfRemove := []any{}
//		for _, key := range keyArr {
//			if gstr.Pos(key, "SelectCache:#`"+daoModelThis.DbTable+"`#") == 0 {
//				keyArrOfRemove = append(keyArrOfRemove, key)
//			}
//		}
//		cache.Removes(daoModelThis.Ctx, keyArrOfRemove) */
//		return daoModelThis
//	}

// 生成模型
func (daoModelThis *DaoModel) newModel() *gdb.Model {
	return daoModelThis.db.Model(daoModelThis.DbTable). /* Safe(). */ Ctx(daoModelThis.Ctx)
}

// 返回当前模型的副本（当外部还需要做特殊处理时使用）
func (daoModelThis *DaoModel) cloneModel() *gdb.Model {
	return daoModelThis.Model.Clone()
}

/*--------业务可能用到的方法 开始--------*/
// 复制新的daoModel（所有属性重置）。作用：对同一个表做多次操作时，不用再解析分库分表
func (daoModelThis *DaoModel) CloneNew() *DaoModel {
	daoModelObj := DaoModel{
		Ctx:          daoModelThis.Ctx,
		dao:          daoModelThis.dao,
		db:           daoModelThis.db,
		DbGroup:      daoModelThis.DbGroup,
		DbTable:      daoModelThis.DbTable,
		JoinTableMap: map[string]struct{}{},
		AfterField:   map[string]any{},
		AfterInsert:  map[string]any{},
		AfterUpdate:  map[string]any{},
		SaveData:     map[string]any{},
	}
	daoModelObj.Model = daoModelObj.newModel()
	return &daoModelObj
}

// 重置daoModel（所有属性重置）。作用：对同一个表做多次操作时，不用再解析分库分表。注意：要在原daoModel已经不用的情况下使用
func (daoModelThis *DaoModel) ResetNew() *DaoModel {
	daoModelThis.JoinTableMap = map[string]struct{}{}
	daoModelThis.AfterField = map[string]any{}
	daoModelThis.AfterInsert = map[string]any{}
	daoModelThis.AfterUpdate = map[string]any{}
	daoModelThis.SaveData = map[string]any{}
	daoModelThis.IdArr = nil
	daoModelThis.Model = daoModelThis.newModel()
	return daoModelThis
}

// 新增需要后置处理且主键非自增时 或 更新|删除需要后置处理时 使用。注意：一般在新增|更新|删除方法执行前调用（即在各种sql条件设置完后）
// 新增需要后置处理且主键非自增时，不使用此方法，直接在dao做了处理
// 该方法只在 更新|删除需要后置处理 时，才使用
func (daoModelThis *DaoModel) SetIdArr(idOrFilterOpt ...any) *DaoModel {
	if len(idOrFilterOpt) == 0 {
		daoModelThis.IdArr, _ = daoModelThis.cloneModel().Distinct().Array(daoModelThis.dao.ParseId(daoModelThis))
		return daoModelThis
	}
	daoModelThis.IdArr = nil
	if filter, ok := idOrFilterOpt[0].(g.Map); ok {
		daoModelThis.Filters(filter)
		if len(filter) != 1 {
			daoModelThis.IdArr, _ = daoModelThis.cloneModel().Distinct().Array(daoModelThis.dao.ParseId(daoModelThis))
			return daoModelThis
		}
		if id, ok := filter[`id`]; ok {
			daoModelThis.IdArr = append(daoModelThis.IdArr, gvar.New(id))
		} else if idArr, ok := filter[`id_arr`]; ok {
			for _, id := range gconv.SliceAny(idArr) {
				daoModelThis.IdArr = append(daoModelThis.IdArr, gvar.New(id))
			}
		} else if idArr, ok := filter[`idArr`]; ok {
			for _, id := range gconv.SliceAny(idArr) {
				daoModelThis.IdArr = append(daoModelThis.IdArr, gvar.New(id))
			}
		} else {
			idField := daoModelThis.dao.ParseId(daoModelThis)
			if id, ok := filter[idField]; ok {
				daoModelThis.IdArr = append(daoModelThis.IdArr, gvar.New(id))
			} else if gstr.Pos(idField, daoModelThis.DbTable+`.`) == 0 {
				idField = gstr.Replace(idField, daoModelThis.DbTable+`.`, ``, 1)
				if id, ok := filter[idField]; ok {
					daoModelThis.IdArr = append(daoModelThis.IdArr, gvar.New(id))
				}
			}
		}
	} else {
		daoModelThis.FilterPri(idOrFilterOpt[0])
		if idVar := gvar.New(idOrFilterOpt[0]); idVar.IsSlice() {
			for _, id := range idVar.Slice() {
				daoModelThis.IdArr = append(daoModelThis.IdArr, gvar.New(id))
			}
		} else {
			daoModelThis.IdArr = append(daoModelThis.IdArr, idVar)
		}
	}
	return daoModelThis
}

// 判断是否联表
func (daoModelThis *DaoModel) IsJoin() bool {
	return len(daoModelThis.JoinTableMap) > 0
}

// 联表时，GroupBy主键
func (daoModelThis *DaoModel) GroupPriOnJoin() *DaoModel {
	if daoModelThis.IsJoin() {
		daoModelThis.Group(`id`)
	}
	return daoModelThis
}

// 列表（联表时，GroupBy主键）
// 为兼容Postgresql，不再默认GroupBy主键（Postgresql联表查询其它表字段时，如果存在GroupBy，其它表字段必须放在GroupBy后面）
func (daoModelThis *DaoModel) ListPri() (gdb.Result, error) {
	return daoModelThis. /* GroupPriOnJoin(). */ All()
}

// 总数（联表时，主键去重）
func (daoModelThis *DaoModel) CountPri() (int, error) {
	if daoModelThis.IsJoin() {
		return daoModelThis.cloneModel().Distinct().CountColumn(daoModelThis.dao.ParseId(daoModelThis))
	}
	return daoModelThis.Count()
}

// 详情（联表时，GroupBy主键）
// 为兼容Postgresql，不再默认GroupBy主键（Postgresql联表查询其它表字段时，如果存在GroupBy，其它表字段必须放在GroupBy后面）
func (daoModelThis *DaoModel) InfoPri() (gdb.Record, error) {
	return daoModelThis. /* GroupPriOnJoin(). */ One()
}

/*--------业务可能用到的方法 结束--------*/

/*--------简化对dao方法的调用 开始--------*/
func (daoModelThis *DaoModel) Filter(key string, val any) *DaoModel {
	return daoModelThis.Filters(map[string]any{key: val})
}

func (daoModelThis *DaoModel) FilterPri(id any) *DaoModel {
	return daoModelThis.Filters(map[string]any{`id`: id})
}

func (daoModelThis *DaoModel) Filters(filter map[string]any) *DaoModel {
	daoModelThis.Handler(daoModelThis.dao.ParseFilter(filter, daoModelThis))
	return daoModelThis
}

func (daoModelThis *DaoModel) Fields(field ...string) *DaoModel {
	daoModelThis.Handler(daoModelThis.dao.ParseField(field, nil, daoModelThis))
	return daoModelThis
}

func (daoModelThis *DaoModel) FieldWithParam(key string, val any) *DaoModel {
	return daoModelThis.FieldsWithParam(map[string]any{key: val})
}

func (daoModelThis *DaoModel) FieldsWithParam(fieldWithParam map[string]any) *DaoModel {
	daoModelThis.Handler(daoModelThis.dao.ParseField(nil, fieldWithParam, daoModelThis))
	return daoModelThis
}

func (daoModelThis *DaoModel) HookInsertOne(key string, val any) *DaoModel {
	return daoModelThis.HookInsert(map[string]any{key: val})
}

func (daoModelThis *DaoModel) HookInsert(data map[string]any) *DaoModel {
	daoModelThis.Handler(daoModelThis.dao.ParseInsert(data, daoModelThis))
	return daoModelThis
}

func (daoModelThis *DaoModel) HookUpdateOne(key string, val any) *DaoModel {
	return daoModelThis.HookUpdate(map[string]any{key: val})
}

func (daoModelThis *DaoModel) HookUpdate(data map[string]any) *DaoModel {
	daoModelThis.Handler(daoModelThis.dao.ParseUpdate(data, daoModelThis))
	return daoModelThis
}

func (daoModelThis *DaoModel) HookDelete(filterOpt ...map[string]any) *DaoModel {
	if len(filterOpt) > 0 {
		daoModelThis.Filters(filterOpt[0])
	}
	daoModelThis.Hook(daoModelThis.dao.HookDelete(daoModelThis))
	return daoModelThis
}

func (daoModelThis *DaoModel) Group(group ...string) *DaoModel {
	daoModelThis.Handler(daoModelThis.dao.ParseGroup(group, daoModelThis))
	return daoModelThis
}

func (daoModelThis *DaoModel) Order(order ...string) *DaoModel {
	daoModelThis.Handler(daoModelThis.dao.ParseOrder(order, daoModelThis))
	return daoModelThis
}

func (daoModelThis *DaoModel) Join(joinTable string) *DaoModel {
	daoModelThis.Handler(daoModelThis.dao.ParseJoin(joinTable, daoModelThis))
	return daoModelThis
}

/*--------简化对dao方法的调用 结束--------*/

/*--------简化对db部分常用方法的调用 开始--------*/
func (daoModelThis *DaoModel) Begin(ctx context.Context) (gdb.TX, error) {
	return daoModelThis.db.Begin(ctx)
}

func (daoModelThis *DaoModel) GetCore() *gdb.Core {
	return daoModelThis.db.GetCore()
}

func (daoModelThis *DaoModel) GetCache() *gcache.Cache {
	return daoModelThis.db.GetCache()
}

/*--------简化对db部分常用方法的调用 结束--------*/

/*--------简化对model方法的调用，并封装部分常用方法 开始--------*/
func (daoModelThis *DaoModel) Transaction(f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return daoModelThis.Model.Transaction(daoModelThis.Ctx, f)
}

func (daoModelThis *DaoModel) TX(tx gdb.TX) *DaoModel {
	daoModelThis.Model = daoModelThis.Model.TX(tx)
	return daoModelThis
}

func (daoModelThis *DaoModel) LockShared() *DaoModel {
	daoModelThis.Model = daoModelThis.Model.LockShared()
	return daoModelThis
}

func (daoModelThis *DaoModel) LockUpdate() *DaoModel {
	daoModelThis.Model = daoModelThis.Model.LockUpdate()
	return daoModelThis
}

func (daoModelThis *DaoModel) Master() *DaoModel {
	daoModelThis.Model = daoModelThis.Model.Master()
	return daoModelThis
}

func (daoModelThis *DaoModel) Slave() *DaoModel {
	daoModelThis.Model = daoModelThis.Model.Slave()
	return daoModelThis
}

func (daoModelThis *DaoModel) Schema(schema string) *DaoModel {
	daoModelThis.Model = daoModelThis.Model.Schema(schema)
	return daoModelThis
}

func (daoModelThis *DaoModel) Cache(option gdb.CacheOption) *DaoModel {
	daoModelThis.Model = daoModelThis.Model.Cache(option)
	return daoModelThis
}

func (daoModelThis *DaoModel) Unscoped() *DaoModel {
	daoModelThis.Model = daoModelThis.Model.Unscoped()
	return daoModelThis
}

func (daoModelThis *DaoModel) Handler(handlers ...gdb.ModelHandler) *DaoModel {
	daoModelThis.Model = daoModelThis.Model.Handler(handlers...)
	return daoModelThis
}

func (daoModelThis *DaoModel) Hook(hook gdb.HookHandler) *DaoModel {
	daoModelThis.Model = daoModelThis.Model.Hook(hook)
	return daoModelThis
}

func (daoModelThis *DaoModel) Data(data ...any) *DaoModel {
	daoModelThis.Model = daoModelThis.Model.Data(data...)
	return daoModelThis
}

func (daoModelThis *DaoModel) OnConflict(onConflict ...any) *DaoModel {
	daoModelThis.Model = daoModelThis.Model.OnConflict(onConflict...)
	return daoModelThis
}

func (daoModelThis *DaoModel) Batch(batch int) *DaoModel {
	daoModelThis.Model = daoModelThis.Model.Batch(batch)
	return daoModelThis
}

func (daoModelThis *DaoModel) Distinct() *DaoModel {
	daoModelThis.Model = daoModelThis.Model.Distinct()
	return daoModelThis
}

func (daoModelThis *DaoModel) Partition(partitions ...string) *DaoModel {
	daoModelThis.Model = daoModelThis.Model.Partition(partitions...)
	return daoModelThis
}

func (daoModelThis *DaoModel) Union(unions ...*gdb.Model) *DaoModel {
	daoModelThis.Model = daoModelThis.Model.Union(unions...)
	return daoModelThis
}

func (daoModelThis *DaoModel) UnionAll(unions ...*gdb.Model) *DaoModel {
	daoModelThis.Model = daoModelThis.Model.UnionAll(unions...)
	return daoModelThis
}

// 以下Where开头的方法通常情况下不建议使用，更建议使用filter方法代替。只在极个别情况下可临时使用
func (daoModelThis *DaoModel) Where(where any, args ...any) *DaoModel {
	daoModelThis.Model = daoModelThis.Model.Where(where, args...)
	return daoModelThis
}

func (daoModelThis *DaoModel) WhereGT(column string, value any) *DaoModel {
	daoModelThis.Model = daoModelThis.Model.WhereGT(column, value)
	return daoModelThis
}

func (daoModelThis *DaoModel) WhereGTE(column string, value any) *DaoModel {
	daoModelThis.Model = daoModelThis.Model.WhereGTE(column, value)
	return daoModelThis
}

func (daoModelThis *DaoModel) WhereLT(column string, value any) *DaoModel {
	daoModelThis.Model = daoModelThis.Model.WhereLT(column, value)
	return daoModelThis
}

func (daoModelThis *DaoModel) WhereLTE(column string, value any) *DaoModel {
	daoModelThis.Model = daoModelThis.Model.WhereLTE(column, value)
	return daoModelThis
}

func (daoModelThis *DaoModel) WhereBetween(column string, min any, max any) *DaoModel {
	daoModelThis.Model = daoModelThis.Model.WhereBetween(column, min, max)
	return daoModelThis
}

func (daoModelThis *DaoModel) WhereIn(column string, in any) *DaoModel {
	daoModelThis.Model = daoModelThis.Model.WhereIn(column, in)
	return daoModelThis
}

func (daoModelThis *DaoModel) WhereLike(column string, like string) *DaoModel {
	daoModelThis.Model = daoModelThis.Model.WhereLike(column, like)
	return daoModelThis
}

func (daoModelThis *DaoModel) WhereNull(columns ...string) *DaoModel {
	daoModelThis.Model = daoModelThis.Model.WhereNull(columns...)
	return daoModelThis
}

func (daoModelThis *DaoModel) WhereNot(column string, value any) *DaoModel {
	daoModelThis.Model = daoModelThis.Model.WhereNot(column, value)
	return daoModelThis
}

func (daoModelThis *DaoModel) WhereNotBetween(column string, min any, max any) *DaoModel {
	daoModelThis.Model = daoModelThis.Model.WhereNotBetween(column, min, max)
	return daoModelThis
}

func (daoModelThis *DaoModel) WhereNotIn(column string, in any) *DaoModel {
	daoModelThis.Model = daoModelThis.Model.WhereNotIn(column, in)
	return daoModelThis
}

func (daoModelThis *DaoModel) WhereNotLike(column string, like string) *DaoModel {
	daoModelThis.Model = daoModelThis.Model.WhereNotLike(column, like)
	return daoModelThis
}

func (daoModelThis *DaoModel) WhereNotNull(columns ...string) *DaoModel {
	daoModelThis.Model = daoModelThis.Model.WhereNotNull(columns...)
	return daoModelThis
}

func (daoModelThis *DaoModel) WhereOr(where any, args ...any) *DaoModel {
	daoModelThis.Model = daoModelThis.Model.WhereOr(where, args...)
	return daoModelThis
}

func (daoModelThis *DaoModel) OrderAsc(column string) *DaoModel {
	daoModelThis.Model = daoModelThis.Model.OrderAsc(column)
	return daoModelThis
}

func (daoModelThis *DaoModel) OrderDesc(column string) *DaoModel {
	daoModelThis.Model = daoModelThis.Model.OrderDesc(column)
	return daoModelThis
}

func (daoModelThis *DaoModel) OrderRandom() *DaoModel {
	daoModelThis.Model = daoModelThis.Model.OrderRandom()
	return daoModelThis
}

func (daoModelThis *DaoModel) Page(page, limit int) *DaoModel {
	if limit == 0 {
		return daoModelThis
	}
	daoModelThis.Model = daoModelThis.Model.Page(page, limit)
	return daoModelThis
}

func (daoModelThis *DaoModel) Offset(offset int) *DaoModel {
	daoModelThis.Model = daoModelThis.Model.Offset(offset)
	return daoModelThis
}

func (daoModelThis *DaoModel) Limit(limit ...int) *DaoModel {
	daoModelThis.Model = daoModelThis.Model.Limit(limit...)
	return daoModelThis
}

func (daoModelThis *DaoModel) InsertAndGetAffected(data ...any) (affected int64, err error) {
	result, err := daoModelThis.Insert(data...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (daoModelThis *DaoModel) IncrementAndGetAffected(column string, amount any) (int64, error) {
	result, err := daoModelThis.Increment(column, amount)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (daoModelThis *DaoModel) DecrementAndGetAffected(column string, amount any) (int64, error) {
	result, err := daoModelThis.Decrement(column, amount)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (daoModelThis *DaoModel) DeleteAndGetAffected(where ...any) (affected int64, err error) {
	result, err := daoModelThis.Delete(where...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (daoModelThis *DaoModel) ArrayStr(fieldsAndWhere ...any) ([]string, error) {
	result, err := daoModelThis.Array(fieldsAndWhere...)
	if err != nil {
		return nil, err
	}
	return gconv.Strings(result), nil
}

func (daoModelThis *DaoModel) ArrayInt(fieldsAndWhere ...any) ([]int, error) {
	result, err := daoModelThis.Array(fieldsAndWhere...)
	if err != nil {
		return nil, err
	}
	return gconv.Ints(result), nil
}

func (daoModelThis *DaoModel) ArrayUint(fieldsAndWhere ...any) ([]uint, error) {
	result, err := daoModelThis.Array(fieldsAndWhere...)
	if err != nil {
		return nil, err
	}
	return gconv.Uints(result), nil
}

func (daoModelThis *DaoModel) Set(fieldsAndWhere ...any) (map[gdb.Value]struct{}, error) {
	arr, err := daoModelThis.Distinct().Array(fieldsAndWhere...)
	if err != nil {
		return nil, err
	}
	arrLen := len(arr)
	if arrLen == 0 {
		return nil, nil
	}
	result := make(map[gdb.Value]struct{}, arrLen)
	for _, v := range arr {
		result[v] = struct{}{}
	}
	return result, nil
}

func (daoModelThis *DaoModel) SetStr(fieldsAndWhere ...any) (map[string]struct{}, error) {
	arr, err := daoModelThis.Distinct().Array(fieldsAndWhere...)
	if err != nil {
		return nil, err
	}
	arrLen := len(arr)
	if arrLen == 0 {
		return nil, nil
	}
	result := make(map[string]struct{}, arrLen)
	for _, v := range arr {
		result[v.String()] = struct{}{}
	}
	return result, nil
}

func (daoModelThis *DaoModel) SetInt(fieldsAndWhere ...any) (map[int]struct{}, error) {
	arr, err := daoModelThis.Distinct().Array(fieldsAndWhere...)
	if err != nil {
		return nil, err
	}
	arrLen := len(arr)
	if arrLen == 0 {
		return nil, nil
	}
	result := make(map[int]struct{}, arrLen)
	for _, v := range arr {
		result[v.Int()] = struct{}{}
	}
	return result, nil
}

func (daoModelThis *DaoModel) SetUint(fieldsAndWhere ...any) (map[uint]struct{}, error) {
	arr, err := daoModelThis.Distinct().Array(fieldsAndWhere...)
	if err != nil {
		return nil, err
	}
	arrLen := len(arr)
	if arrLen == 0 {
		return nil, nil
	}
	result := make(map[uint]struct{}, arrLen)
	for _, v := range arr {
		result[v.Uint()] = struct{}{}
	}
	return result, nil
}

func (daoModelThis *DaoModel) Pluck(key string, field string) (map[gdb.Value]gdb.Value, error) {
	list, err := daoModelThis.Fields(key, field).All()
	if err != nil {
		return nil, err
	}
	if list.IsEmpty() {
		return nil, nil
	}
	result := make(map[gdb.Value]gdb.Value, len(list))
	for _, v := range list {
		result[v[key]] = v[field]
	}
	return result, nil
}

func (daoModelThis *DaoModel) PluckStr(key string, field string) (map[string]gdb.Value, error) {
	list, err := daoModelThis.Fields(key, field).All()
	if err != nil {
		return nil, err
	}
	if list.IsEmpty() {
		return nil, nil
	}
	result := make(gdb.Record, len(list))
	for _, v := range list {
		result[v[key].String()] = v[field]
	}
	return result, nil
}

func (daoModelThis *DaoModel) PluckInt(key string, field string) (map[int]gdb.Value, error) {
	list, err := daoModelThis.Fields(key, field).All()
	if err != nil {
		return nil, err
	}
	if list.IsEmpty() {
		return nil, nil
	}
	result := make(map[int]gdb.Value, len(list))
	for _, v := range list {
		result[v[key].Int()] = v[field]
	}
	return result, nil
}

func (daoModelThis *DaoModel) PluckUint(key string, field string) (map[uint]gdb.Value, error) {
	list, err := daoModelThis.Fields(key, field).All()
	if err != nil {
		return nil, err
	}
	if list.IsEmpty() {
		return nil, nil
	}
	result := make(map[uint]gdb.Value, len(list))
	for _, v := range list {
		result[v[key].Uint()] = v[field]
	}
	return result, nil
}

func (daoModelThis *DaoModel) ValueStr(fieldsAndWhere ...any) (string, error) {
	result, err := daoModelThis.Value(fieldsAndWhere...)
	if err != nil {
		return ``, err
	}
	return result.String(), nil
}

func (daoModelThis *DaoModel) ValueInt(fieldsAndWhere ...any) (int, error) {
	result, err := daoModelThis.Value(fieldsAndWhere...)
	if err != nil {
		return 0, err
	}
	return result.Int(), nil
}

func (daoModelThis *DaoModel) ValueUint(fieldsAndWhere ...any) (uint, error) {
	result, err := daoModelThis.Value(fieldsAndWhere...)
	if err != nil {
		return 0, err
	}
	return result.Uint(), nil
}

func (daoModelThis *DaoModel) ValueInt64(fieldsAndWhere ...any) (int64, error) {
	result, err := daoModelThis.Value(fieldsAndWhere...)
	if err != nil {
		return 0, err
	}
	return result.Int64(), nil
}

func (daoModelThis *DaoModel) ValueMap(fieldsAndWhere ...any) (g.Map, error) {
	result, err := daoModelThis.Value(fieldsAndWhere...)
	if err != nil {
		return nil, err
	}
	return result.Map(), nil
}

/*--------简化对gdb.Model方法的调用，并封装部分常用方法 结束--------*/
