package dao

import (
	"context"
	"database/sql"
	"sync"

	"github.com/gogf/gf/v2/container/gset"
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
	Ctx                 context.Context
	dao                 DaoInterface
	db                  gdb.DB
	model               *gdb.Model
	DbGroup             string // 分库情况下，解析后所确定的库
	DbTable             string // 分表情况下，解析后所确定的表
	AfterField          *gset.StrSet
	AfterFieldSlice     []string // 后置处理前，将AfterField转换成AfterFieldSlice，减少列表后置处理多次调用AfterField.Slice()转换
	AfterFieldWithParam map[string]any
	IdArr               []*gvar.Var // 新增需要后置处理且主键非自增时 或 更新|删除需要后置处理时 使用。注意：一般在新增|更新|删除方法执行前调用（即在各种sql条件设置完后）
	AfterInsert         map[string]any
	IsOnlyAfterUpdate   bool // 更新时，用于判断是否只做后置更新
	AfterUpdate         map[string]any
	JoinTableSet        *gset.StrSet
}

// 对象池。性能提醒不明显，暂时不用。确实大幅减少了对象创建和销毁（内存压力减少），但却需要手动增加放入对象池的代码：defer daoModel.PutPool()
var poolDaoModel = sync.Pool{
	New: func() any {
		return &DaoModel{}
	},
}

func (daoModelThis *DaoModel) PutPool() {
	//以下属性从池中取出对象时，会重新赋值
	daoModelThis.Ctx = nil
	daoModelThis.dao = nil
	daoModelThis.AfterField = nil
	daoModelThis.AfterFieldWithParam = nil
	daoModelThis.AfterInsert = nil
	daoModelThis.AfterUpdate = nil
	daoModelThis.JoinTableSet = nil
	daoModelThis.DbTable = ``
	daoModelThis.DbGroup = ``
	daoModelThis.db = nil
	daoModelThis.model = nil

	daoModelThis.AfterFieldSlice = nil
	daoModelThis.IdArr = nil
	daoModelThis.IsOnlyAfterUpdate = false
	daoModelThis.IsOnlyAfterUpdate = false

	poolDaoModel.Put(daoModelThis)
}

// 注意：dbOpt存在时，dbOpt[0]解析DbTable，dbOpt[1]索引参数解析DbGroup
func NewDaoModel(ctx context.Context, dao DaoInterface, dbOpt ...any) *DaoModel {
	daoModelObj := &DaoModel{} // poolDaoModel.Get().(*DaoModel)
	daoModelObj.Ctx = ctx
	daoModelObj.dao = dao
	daoModelObj.AfterField = gset.NewStrSet()
	daoModelObj.AfterFieldWithParam = map[string]any{}
	daoModelObj.AfterInsert = map[string]any{}
	daoModelObj.AfterUpdate = map[string]any{}
	daoModelObj.JoinTableSet = gset.NewStrSet()
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
	daoModelObj.db = daoModelObj.NewDB()
	daoModelObj.model = daoModelObj.NewModel()
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

/*--------业务可能用到的方法 开始--------*/
// 复制新的daoModel（所有属性重置）。作用：对同一个表做多次操作时，不用再解析分库分表
func (daoModelThis *DaoModel) CloneNew() *DaoModel {
	daoModelObj := DaoModel{
		Ctx:                 daoModelThis.Ctx,
		dao:                 daoModelThis.dao,
		db:                  daoModelThis.db,
		DbGroup:             daoModelThis.DbGroup,
		DbTable:             daoModelThis.DbTable,
		AfterField:          gset.NewStrSet(),
		AfterFieldWithParam: map[string]any{},
		AfterInsert:         map[string]any{},
		AfterUpdate:         map[string]any{},
		JoinTableSet:        gset.NewStrSet(),
	}
	daoModelObj.model = daoModelObj.NewModel()
	return &daoModelObj
}

// 生成数据库
func (daoModelThis *DaoModel) NewDB() gdb.DB {
	return g.DB(daoModelThis.DbGroup)
}

// 返回当前数据库
func (daoModelThis *DaoModel) GetDB() gdb.DB {
	return daoModelThis.db
}

// 生成模型
func (daoModelThis *DaoModel) NewModel() *gdb.Model {
	return daoModelThis.GetDB().Model(daoModelThis.DbTable). /* Safe(). */ Ctx(daoModelThis.Ctx)
}

// 返回当前模型的副本（当外部还需要做特殊处理时使用）
func (daoModelThis *DaoModel) CloneModel() *gdb.Model {
	return daoModelThis.GetModel().Clone()
}

// 返回当前模型（当外部还需要做特殊处理时使用）
func (daoModelThis *DaoModel) GetModel() *gdb.Model {
	return daoModelThis.model
}

// 新增需要后置处理且主键非自增时 或 更新|删除需要后置处理时 使用。注意：一般在新增|更新|删除方法执行前调用（即在各种sql条件设置完后）
// 新增需要后置处理且主键非自增时，不使用此方法，直接在dao做了处理
// 该方法只在 更新|删除需要后置处理 时，才使用
func (daoModelThis *DaoModel) SetIdArr(idOrFilterOpt ...any) *DaoModel {
	if len(idOrFilterOpt) == 0 {
		daoModelThis.IdArr, _ = daoModelThis.CloneModel().Distinct().Array(daoModelThis.dao.ParseId(daoModelThis))
		return daoModelThis
	}
	daoModelThis.IdArr = nil
	if filter, ok := idOrFilterOpt[0].(g.Map); ok {
		daoModelThis.Filters(filter)
		if len(filter) != 1 {
			daoModelThis.IdArr, _ = daoModelThis.CloneModel().Distinct().Array(daoModelThis.dao.ParseId(daoModelThis))
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
	return daoModelThis.JoinTableSet.Size() > 0
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
		return daoModelThis.CloneModel().Distinct().CountColumn(daoModelThis.dao.ParseId(daoModelThis))
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
	daoModelThis.Handler(daoModelThis.dao.ParseField(field, map[string]any{}, daoModelThis))
	return daoModelThis
}

func (daoModelThis *DaoModel) FieldWithParam(key string, val any) *DaoModel {
	return daoModelThis.FieldsWithParam(map[string]any{key: val})
}

func (daoModelThis *DaoModel) FieldsWithParam(fieldWithParam map[string]any) *DaoModel {
	daoModelThis.Handler(daoModelThis.dao.ParseField([]string{}, fieldWithParam, daoModelThis))
	return daoModelThis
}

func (daoModelThis *DaoModel) HandleAfterField(result ...gdb.Record) {
	daoModelThis.AfterFieldSlice = daoModelThis.AfterField.Slice()
	for _, record := range result {
		daoModelThis.dao.HandleAfterField(daoModelThis.Ctx, record, daoModelThis)
	}
}

/* func (daoModelThis *DaoModel) HookSelect() *DaoModel {
	if daoModelThis.AfterField.Size() > 0 || len(daoModelThis.AfterFieldWithParam) > 0 {
		daoModelThis.Hook(daoModelThis.dao.HookSelect(daoModelThis))
	}
	return daoModelThis
} */

func (daoModelThis *DaoModel) HookInsert(data map[string]any) *DaoModel {
	daoModelThis.Handler(daoModelThis.dao.ParseInsert(data, daoModelThis))
	return daoModelThis
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
	return daoModelThis.model.Transaction(daoModelThis.Ctx, f)
}

func (daoModelThis *DaoModel) TX(tx gdb.TX) *DaoModel {
	daoModelThis.model = daoModelThis.model.TX(tx)
	return daoModelThis
}

func (daoModelThis *DaoModel) LockShared() *DaoModel {
	daoModelThis.model = daoModelThis.model.LockShared()
	return daoModelThis
}

func (daoModelThis *DaoModel) LockUpdate() *DaoModel {
	daoModelThis.model = daoModelThis.model.LockUpdate()
	return daoModelThis
}

func (daoModelThis *DaoModel) Master() *DaoModel {
	daoModelThis.model = daoModelThis.model.Master()
	return daoModelThis
}

func (daoModelThis *DaoModel) Slave() *DaoModel {
	daoModelThis.model = daoModelThis.model.Slave()
	return daoModelThis
}

func (daoModelThis *DaoModel) Schema(schema string) *DaoModel {
	daoModelThis.model = daoModelThis.model.Schema(schema)
	return daoModelThis
}

func (daoModelThis *DaoModel) Cache(option gdb.CacheOption) *DaoModel {
	daoModelThis.model = daoModelThis.model.Cache(option)
	return daoModelThis
}

func (daoModelThis *DaoModel) Unscoped() *DaoModel {
	daoModelThis.model = daoModelThis.model.Unscoped()
	return daoModelThis
}

func (daoModelThis *DaoModel) Handler(handlers ...gdb.ModelHandler) *DaoModel {
	daoModelThis.model = daoModelThis.model.Handler(handlers...)
	return daoModelThis
}

func (daoModelThis *DaoModel) Hook(hook gdb.HookHandler) *DaoModel {
	daoModelThis.model = daoModelThis.model.Hook(hook)
	return daoModelThis
}

func (daoModelThis *DaoModel) Data(data ...any) *DaoModel {
	daoModelThis.model = daoModelThis.model.Data(data...)
	return daoModelThis
}

func (daoModelThis *DaoModel) OnConflict(onConflict ...any) *DaoModel {
	daoModelThis.model = daoModelThis.model.OnConflict(onConflict...)
	return daoModelThis
}

func (daoModelThis *DaoModel) Batch(batch int) *DaoModel {
	daoModelThis.model = daoModelThis.model.Batch(batch)
	return daoModelThis
}

func (daoModelThis *DaoModel) Distinct() *DaoModel {
	daoModelThis.model = daoModelThis.model.Distinct()
	return daoModelThis
}

func (daoModelThis *DaoModel) Partition(partitions ...string) *DaoModel {
	daoModelThis.model = daoModelThis.model.Partition(partitions...)
	return daoModelThis
}

func (daoModelThis *DaoModel) Union(unions ...*gdb.Model) *DaoModel {
	daoModelThis.model = daoModelThis.model.Union(unions...)
	return daoModelThis
}

func (daoModelThis *DaoModel) UnionAll(unions ...*gdb.Model) *DaoModel {
	daoModelThis.model = daoModelThis.model.UnionAll(unions...)
	return daoModelThis
}

// 以下Where开头的方法通常情况下不建议使用，更建议使用filter方法代替。只在极个别情况下可临时使用
func (daoModelThis *DaoModel) Where(where interface{}, args ...interface{}) *DaoModel {
	daoModelThis.model = daoModelThis.model.Where(where, args...)
	return daoModelThis
}

func (daoModelThis *DaoModel) WhereGT(column string, value interface{}) *DaoModel {
	daoModelThis.model = daoModelThis.model.WhereGT(column, value)
	return daoModelThis
}

func (daoModelThis *DaoModel) WhereGTE(column string, value interface{}) *DaoModel {
	daoModelThis.model = daoModelThis.model.WhereGTE(column, value)
	return daoModelThis
}

func (daoModelThis *DaoModel) WhereLT(column string, value interface{}) *DaoModel {
	daoModelThis.model = daoModelThis.model.WhereLT(column, value)
	return daoModelThis
}

func (daoModelThis *DaoModel) WhereLTE(column string, value interface{}) *DaoModel {
	daoModelThis.model = daoModelThis.model.WhereLTE(column, value)
	return daoModelThis
}

/* func (daoModelThis *DaoModel) WhereBetween(column string, min interface{}, max interface{}) *DaoModel {
	daoModelThis.model = daoModelThis.model.WhereBetween(column, min, max)
	return daoModelThis
}

func (daoModelThis *DaoModel) WhereIn(column string, in interface{}) *DaoModel {
	daoModelThis.model = daoModelThis.model.WhereIn(column, in)
	return daoModelThis
}

func (daoModelThis *DaoModel) WhereLike(column string, like string) *DaoModel {
	daoModelThis.model = daoModelThis.model.WhereLike(column, like)
	return daoModelThis
}

func (daoModelThis *DaoModel) WhereNull(columns ...string) *DaoModel {
	daoModelThis.model = daoModelThis.model.WhereNull(columns...)
	return daoModelThis
}

func (daoModelThis *DaoModel) WhereNot(column string, value interface{}) *DaoModel {
	daoModelThis.model = daoModelThis.model.WhereNot(column, value)
	return daoModelThis
}

func (daoModelThis *DaoModel) WhereNotBetween(column string, min interface{}, max interface{}) *DaoModel {
	daoModelThis.model = daoModelThis.model.WhereNotBetween(column, min, max)
	return daoModelThis
}

func (daoModelThis *DaoModel) WhereNotIn(column string, in interface{}) *DaoModel {
	daoModelThis.model = daoModelThis.model.WhereNotIn(column, in)
	return daoModelThis
}

func (daoModelThis *DaoModel) WhereNotLike(column string, like string) *DaoModel {
	daoModelThis.model = daoModelThis.model.WhereNotLike(column, like)
	return daoModelThis
}

func (daoModelThis *DaoModel) WhereNotNull(columns ...string) *DaoModel {
	daoModelThis.model = daoModelThis.model.WhereNotNull(columns...)
	return daoModelThis
}

func (daoModelThis *DaoModel) WhereOr(where interface{}, args ...interface{}) *DaoModel {
	daoModelThis.model = daoModelThis.model.WhereOr(where, args...)
	return daoModelThis
} */

func (daoModelThis *DaoModel) OrderAsc(column string) *DaoModel {
	daoModelThis.model = daoModelThis.model.OrderAsc(column)
	return daoModelThis
}

func (daoModelThis *DaoModel) OrderDesc(column string) *DaoModel {
	daoModelThis.model = daoModelThis.model.OrderDesc(column)
	return daoModelThis
}

func (daoModelThis *DaoModel) OrderRandom() *DaoModel {
	daoModelThis.model = daoModelThis.model.OrderRandom()
	return daoModelThis
}

func (daoModelThis *DaoModel) Page(page, limit int) *DaoModel {
	if limit == 0 {
		return daoModelThis
	}
	daoModelThis.model = daoModelThis.model.Page(page, limit)
	return daoModelThis
}

func (daoModelThis *DaoModel) Offset(offset int) *DaoModel {
	daoModelThis.model = daoModelThis.model.Offset(offset)
	return daoModelThis
}

func (daoModelThis *DaoModel) Limit(limit ...int) *DaoModel {
	daoModelThis.model = daoModelThis.model.Limit(limit...)
	return daoModelThis
}

func (daoModelThis *DaoModel) Save(data ...any) (result sql.Result, err error) {
	return daoModelThis.model.Save(data...)
}

func (daoModelThis *DaoModel) Replace(data ...any) (result sql.Result, err error) {
	return daoModelThis.model.Replace(data...)
}

func (daoModelThis *DaoModel) Insert(data ...any) (result sql.Result, err error) {
	return daoModelThis.model.Insert(data...)
}

func (daoModelThis *DaoModel) InsertAndGetId(data ...any) (lastInsertId int64, err error) {
	return daoModelThis.model.InsertAndGetId(data...)
}

func (daoModelThis *DaoModel) InsertIgnore(data ...any) (result sql.Result, err error) {
	return daoModelThis.model.InsertIgnore(data...)
}

// 封装常用方法
func (daoModelThis *DaoModel) InsertAndGetAffected(data ...any) (affected int64, err error) {
	result, err := daoModelThis.model.Insert(data...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (daoModelThis *DaoModel) Update(dataAndWhere ...any) (result sql.Result, err error) {
	return daoModelThis.model.Update(dataAndWhere...)
}

func (daoModelThis *DaoModel) UpdateAndGetAffected(dataAndWhere ...any) (affected int64, err error) {
	return daoModelThis.model.UpdateAndGetAffected(dataAndWhere...)
}

func (daoModelThis *DaoModel) Increment(column string, amount any) (sql.Result, error) {
	return daoModelThis.model.Increment(column, amount)
}

// 封装常用方法
func (daoModelThis *DaoModel) IncrementAndGetAffected(column string, amount any) (int64, error) {
	result, err := daoModelThis.model.Increment(column, amount)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (daoModelThis *DaoModel) Decrement(column string, amount any) (sql.Result, error) {
	return daoModelThis.model.Decrement(column, amount)
}

// 封装常用方法
func (daoModelThis *DaoModel) DecrementAndGetAffected(column string, amount any) (int64, error) {
	result, err := daoModelThis.model.Decrement(column, amount)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (daoModelThis *DaoModel) Delete(where ...any) (result sql.Result, err error) {
	return daoModelThis.model.Delete(where...)
}

// 封装常用方法
func (daoModelThis *DaoModel) DeleteAndGetAffected(where ...any) (affected int64, err error) {
	result, err := daoModelThis.Delete(where...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (daoModelThis *DaoModel) Chunk(size int, handler gdb.ChunkHandler) {
	daoModelThis.model.Chunk(size, handler)
}

func (daoModelThis *DaoModel) Scan(pointer any, where ...any) error {
	return daoModelThis.model.Scan(pointer, where...)
}

func (daoModelThis *DaoModel) ScanAndCount(pointer any, totalCount *int, useFieldForCount bool) (err error) {
	return daoModelThis.model.ScanAndCount(pointer, totalCount, useFieldForCount)
}

func (daoModelThis *DaoModel) ScanList(structSlicePointer any, bindToAttrName string, relationAttrNameAndFields ...string) (err error) {
	return daoModelThis.model.ScanList(structSlicePointer, bindToAttrName, relationAttrNameAndFields...)
}

func (daoModelThis *DaoModel) All() (gdb.Result, error) {
	return daoModelThis.model.All()
}

func (daoModelThis *DaoModel) AllAndCount(useFieldForCount bool) (result gdb.Result, totalCount int, err error) {
	return daoModelThis.model.AllAndCount(useFieldForCount)
}

func (daoModelThis *DaoModel) One(where ...any) (gdb.Record, error) {
	return daoModelThis.model.One(where...)
}

func (daoModelThis *DaoModel) Array(fieldsAndWhere ...any) ([]gdb.Value, error) {
	return daoModelThis.model.Array(fieldsAndWhere...)
}

// 封装常用方法
func (daoModelThis *DaoModel) ArrayStr(fieldsAndWhere ...any) ([]string, error) {
	result, err := daoModelThis.Array(fieldsAndWhere...)
	if err != nil {
		return nil, err
	}
	return gconv.Strings(result), nil
}

// 封装常用方法
func (daoModelThis *DaoModel) ArrayUint(fieldsAndWhere ...any) ([]uint, error) {
	result, err := daoModelThis.Array(fieldsAndWhere...)
	if err != nil {
		return nil, err
	}
	return gconv.Uints(result), nil
}

// 封装常用方法
func (daoModelThis *DaoModel) ArrayInt(fieldsAndWhere ...any) ([]int, error) {
	result, err := daoModelThis.Array(fieldsAndWhere...)
	if err != nil {
		return nil, err
	}
	return gconv.Ints(result), nil
}

// 封装常用方法
func (daoModelThis *DaoModel) Pluck(field string, key string) (map[gdb.Value]gdb.Value, error) {
	list, err := daoModelThis.Fields(field, key).All()
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

// 封装常用方法
func (daoModelThis *DaoModel) PluckStr(field string, key string) (gdb.Record, error) {
	list, err := daoModelThis.Fields(field, key).All()
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

// 封装常用方法
func (daoModelThis *DaoModel) PluckStrStr(field string, key string) (g.MapStrStr, error) {
	list, err := daoModelThis.Fields(field, key).All()
	if err != nil {
		return nil, err
	}
	if list.IsEmpty() {
		return nil, nil
	}
	result := make(g.MapStrStr, len(list))
	for _, v := range list {
		result[v[key].String()] = v[field].String()
	}
	return result, nil
}

// 封装常用方法
func (daoModelThis *DaoModel) PluckUint(field string, key string) (map[uint]gdb.Value, error) {
	list, err := daoModelThis.Fields(field, key).All()
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

// 封装常用方法
func (daoModelThis *DaoModel) PluckInt(field string, key string) (map[int]gdb.Value, error) {
	list, err := daoModelThis.Fields(field, key).All()
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

func (daoModelThis *DaoModel) Value(fieldsAndWhere ...any) (gdb.Value, error) {
	return daoModelThis.model.Value(fieldsAndWhere...)
}

// 封装常用方法
func (daoModelThis *DaoModel) ValueStr(fieldsAndWhere ...any) (string, error) {
	result, err := daoModelThis.Value(fieldsAndWhere...)
	if err != nil {
		return ``, err
	}
	return result.String(), nil
}

// 封装常用方法
func (daoModelThis *DaoModel) ValueUint(fieldsAndWhere ...any) (uint, error) {
	result, err := daoModelThis.Value(fieldsAndWhere...)
	if err != nil {
		return 0, err
	}
	return result.Uint(), nil
}

// 封装常用方法
func (daoModelThis *DaoModel) ValueInt(fieldsAndWhere ...any) (int, error) {
	result, err := daoModelThis.Value(fieldsAndWhere...)
	if err != nil {
		return 0, err
	}
	return result.Int(), nil
}

// 封装常用方法
func (daoModelThis *DaoModel) ValueInt64(fieldsAndWhere ...any) (int64, error) {
	result, err := daoModelThis.Value(fieldsAndWhere...)
	if err != nil {
		return 0, err
	}
	return result.Int64(), nil
}

// 封装常用方法
func (daoModelThis *DaoModel) ValueMap(fieldsAndWhere ...any) (g.Map, error) {
	result, err := daoModelThis.Value(fieldsAndWhere...)
	if err != nil {
		return nil, err
	}
	return result.Map(), nil
}

func (daoModelThis *DaoModel) HasField(field string) (bool, error) {
	return daoModelThis.model.HasField(field)
}

func (daoModelThis *DaoModel) Count(where ...any) (int, error) {
	return daoModelThis.model.Count(where...)
}

func (daoModelThis *DaoModel) CountColumn(column string) (int, error) {
	return daoModelThis.model.CountColumn(column)
}

func (daoModelThis *DaoModel) Sum(column string) (float64, error) {
	return daoModelThis.model.Sum(column)
}

func (daoModelThis *DaoModel) Avg(column string) (float64, error) {
	return daoModelThis.model.Avg(column)
}

func (daoModelThis *DaoModel) Min(column string) (float64, error) {
	return daoModelThis.model.Min(column)
}

func (daoModelThis *DaoModel) Max(column string) (float64, error) {
	return daoModelThis.model.Max(column)
}

/*--------简化对gdb.Model方法的调用，并封装部分常用方法 结束--------*/
