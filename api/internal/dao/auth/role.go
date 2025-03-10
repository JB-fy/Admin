// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package auth

import (
	"api/internal/cache"
	"api/internal/consts"
	daoIndex "api/internal/dao"
	"api/internal/dao/auth/internal"
	daoOrg "api/internal/dao/org/allow"
	"context"
	"database/sql"
	"database/sql/driver"
	"sync"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

// internalRoleDao is internal type for wrapping internal DAO implements.
type internalRoleDao = *internal.RoleDao

// roleDao is the data access object for table auth_role.
// You can define custom methods on it to extend its functionality as you wish.
type roleDao struct {
	internalRoleDao
}

var (
	// Role is globally public accessible object for table auth_role operations.
	Role = roleDao{
		internal.NewRoleDao(),
	}
)

// 获取daoModel
func (daoThis *roleDao) CtxDaoModel(ctx context.Context, dbOpt ...any) *daoIndex.DaoModel {
	return daoIndex.NewDaoModel(ctx, daoThis, dbOpt...)
}

// 解析分库
func (daoThis *roleDao) ParseDbGroup(ctx context.Context, dbGroupOpt ...any) string {
	group := daoThis.Group()
	// 分库逻辑
	/* if len(dbGroupOpt) > 0 {
	} */
	return group
}

// 解析分表
func (daoThis *roleDao) ParseDbTable(ctx context.Context, dbTableOpt ...any) string {
	table := daoThis.Table()
	// 分表逻辑
	/* if len(dbTableOpt) > 0 {
	} */
	return table
}

// 解析Id（未使用代码自动生成，且id字段不在第1个位置时，需手动修改）
func (daoThis *roleDao) ParseId(daoModel *daoIndex.DaoModel) string {
	return daoModel.DbTable + `.` + daoThis.Columns().RoleId
}

// 解析Label（未使用代码自动生成，且id字段不在第2个位置时，需手动修改）
func (daoThis *roleDao) ParseLabel(daoModel *daoIndex.DaoModel) string {
	return daoModel.DbTable + `.` + daoThis.Columns().RoleName
}

// 解析filter
func (daoThis *roleDao) ParseFilter(filter map[string]any, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for k, v := range filter {
			switch k {
			/* case `xxxx`:
			tableXxxx := Xxxx.ParseDbTable(m.GetCtx())
			m = m.Where(tableXxxx+`.`+k, v)
			m = m.Handler(daoThis.ParseJoin(tableXxxx, daoModel)) */
			case `id`, `id_arr`:
				m = m.Where(daoModel.DbTable+`.`+daoThis.Columns().RoleId, v)
			case `exc_id`, `exc_id_arr`:
				if gvar.New(v).IsSlice() {
					m = m.WhereNotIn(daoModel.DbTable+`.`+daoThis.Columns().RoleId, v)
				} else {
					m = m.WhereNot(daoModel.DbTable+`.`+daoThis.Columns().RoleId, v)
				}
			case `label`:
				m = m.WhereLike(daoModel.DbTable+`.`+daoThis.Columns().RoleName, `%`+gconv.String(v)+`%`)
			case daoThis.Columns().RoleName:
				m = m.WhereLike(daoModel.DbTable+`.`+k, `%`+gconv.String(v)+`%`)
			case RoleRelToAction.Columns().ActionId:
				tableRoleRelToAction := RoleRelToAction.ParseDbTable(m.GetCtx())
				m = m.Where(tableRoleRelToAction+`.`+k, v)
				m = m.Handler(daoThis.ParseJoin(tableRoleRelToAction, daoModel))
			case RoleRelToMenu.Columns().MenuId:
				tableRoleRelToMenu := RoleRelToMenu.ParseDbTable(m.GetCtx())
				m = m.Where(tableRoleRelToMenu+`.`+k, v)
				m = m.Handler(daoThis.ParseJoin(tableRoleRelToMenu, daoModel))
			case `time_range_start`:
				m = m.WhereGTE(daoModel.DbTable+`.`+daoThis.Columns().CreatedAt, v)
			case `time_range_end`:
				m = m.WhereLTE(daoModel.DbTable+`.`+daoThis.Columns().CreatedAt, v)
			case `self_role`: //获取当前登录身份可用的角色。参数：map[string]any{`scene_id`: `场景ID`, `login_id`: 登录身份id}
				m = m.Where(daoModel.DbTable+`.`+daoThis.Columns().IsStop, 0)
				val := gconv.Map(v)
				var roleIdArr []*gvar.Var
				switch gconv.String(val[`scene_id`]) {
				case `platform`:
					// 方式1：非联表查询
					roleIdArr, _ = RoleRelOfPlatformAdmin.CtxDaoModel(m.GetCtx()).Filter(RoleRelOfPlatformAdmin.Columns().AdminId, val[`login_id`]).Array(RoleRelOfPlatformAdmin.Columns().RoleId)
					/* // 方式2：联表查询（不推荐。原因：auth_role及其关联表，后期表数据只会越来越大，故不建议联表）
					tableRoleRelOfPlatformAdmin := RoleRelOfPlatformAdmin.ParseDbTable(m.GetCtx())
					m = m.Where(tableRoleRelOfPlatformAdmin+`.`+RoleRelOfPlatformAdmin.Columns().AdminId, val[`login_id`])
					m = m.Handler(daoThis.ParseJoin(tableRoleRelOfPlatformAdmin, daoModel))
					continue */
				case `org`:
					// 方式1：非联表查询
					roleIdArr, _ = RoleRelOfOrgAdmin.CtxDaoModel(m.GetCtx()).Filter(RoleRelOfOrgAdmin.Columns().AdminId, val[`login_id`]).Array(RoleRelOfOrgAdmin.Columns().RoleId)
					/* // 方式2：联表查询（不推荐。原因：auth_role及其关联表，后期表数据只会越来越大，故不建议联表）
					tableRoleRelOfOrgAdmin := RoleRelOfOrgAdmin.ParseDbTable(m.GetCtx())
					m = m.Where(tableRoleRelOfOrgAdmin+`.`+RoleRelOfOrgAdmin.Columns().AdminId, val[`login_id`])
					m = m.Handler(daoThis.ParseJoin(tableRoleRelOfOrgAdmin, daoModel))
					continue */
				}
				if len(roleIdArr) == 0 {
					m = m.Where(`1 = 0`)
					continue
				}
				m = m.Where(daoModel.DbTable+`.`+daoThis.Columns().RoleId, roleIdArr)
			default:
				if daoThis.Contains(k) {
					m = m.Where(daoModel.DbTable+`.`+k, v)
				} else {
					m = m.Where(k, v)
				}
			}
		}
		return m
	}
}

// 解析field
func (daoThis *roleDao) ParseField(field []string, fieldWithParam map[string]any, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for _, v := range field {
			switch v {
			/* case `xxxx`:
			tableXxxx := Xxxx.ParseDbTable(m.GetCtx())
			m = m.Fields(tableXxxx + `.` + v)
			m = m.Handler(daoThis.ParseJoin(tableXxxx, daoModel))
			daoModel.AfterField.Add(v) */
			case `id`:
				m = m.Fields(daoThis.ParseId(daoModel) + ` AS ` + v)
			case `label`:
				m = m.Fields(daoThis.ParseLabel(daoModel) + ` AS ` + v)
			case Scene.Columns().SceneName:
				tableScene := Scene.ParseDbTable(m.GetCtx())
				m = m.Fields(tableScene + `.` + v)
				m = m.Handler(daoThis.ParseJoin(tableScene, daoModel))
			case `action_id_arr`:
				m = m.Fields(daoModel.DbTable + `.` + daoThis.Columns().RoleId)
				daoModel.AfterField.Add(v)
			case `menu_id_arr`:
				m = m.Fields(daoModel.DbTable + `.` + daoThis.Columns().RoleId)
				daoModel.AfterField.Add(v)
			case `rel_name`:
				m = m.Fields(daoModel.DbTable + `.` + daoThis.Columns().SceneId)
				m = m.Fields(daoModel.DbTable + `.` + daoThis.Columns().RelId)
				daoModel.AfterField.Add(v)
			default:
				if daoThis.Contains(v) {
					m = m.Fields(daoModel.DbTable + `.` + v)
				} else {
					m = m.Fields(v)
				}
			}
		}
		for k, v := range fieldWithParam {
			switch k {
			default:
				daoModel.AfterFieldWithParam[k] = v
			}
		}
		if daoModel.AfterField.Size() > 0 || len(daoModel.AfterFieldWithParam) > 0 {
			m = m.Hook(daoThis.HookSelect(daoModel))
		}
		return m
	}
}

// 处理afterField
func (daoThis *roleDao) HandleAfterField(ctx context.Context, record gdb.Record, daoModel *daoIndex.DaoModel) {
	for _, v := range daoModel.AfterFieldSlice {
		switch v {
		case `action_id_arr`:
			actionIdArr, _ := RoleRelToAction.CtxDaoModel(ctx).Filter(RoleRelToAction.Columns().RoleId, record[daoThis.Columns().RoleId]).Array(RoleRelToAction.Columns().ActionId)
			record[v] = gvar.New(actionIdArr)
		case `menu_id_arr`:
			menuIdArr, _ := RoleRelToMenu.CtxDaoModel(ctx).Filter(RoleRelToMenu.Columns().RoleId, record[daoThis.Columns().RoleId]).Array(RoleRelToMenu.Columns().MenuId)
			record[v] = gvar.New(menuIdArr)
		case `rel_name`:
			relName := ``
			if record[daoThis.Columns().RelId].Uint() == 0 {
				relName = `平台`
			} else {
				switch record[Scene.Columns().SceneId].String() {
				// case `platform`:	// 平台都是0
				case `org`:
					relName, _ = daoOrg.Org.CtxDaoModel(ctx).FilterPri(record[daoThis.Columns().RelId]).ValueStr(daoOrg.Org.Columns().OrgName)
				}
			}
			record[v] = gvar.New(relName)
		default:
			record[v] = gvar.New(nil)
		}
	}
	/* for k, v := range daoModel.AfterFieldWithParam {
		switch k {
		case `xxxx`:
			record[k] = gvar.New(v)
		}
	} */
}

// hook select
func (daoThis *roleDao) HookSelect(daoModel *daoIndex.DaoModel) gdb.HookHandler {
	return gdb.HookHandler{
		Select: func(ctx context.Context, in *gdb.HookSelectInput) (result gdb.Result, err error) {
			result, err = in.Next(ctx)
			if err != nil || len(result) == 0 {
				return
			}

			var wg sync.WaitGroup
			wg.Add(len(result))
			daoModel.AfterFieldSlice = daoModel.AfterField.Slice()
			for _, record := range result {
				go func(record gdb.Record) {
					defer wg.Done()
					daoThis.HandleAfterField(ctx, record, daoModel)
				}(record)
			}
			wg.Wait()
			return
		},
	}
}

// 解析insert
func (daoThis *roleDao) ParseInsert(insert map[string]any, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		insertData := map[string]any{}
		for k, v := range insert {
			switch k {
			case `action_id_arr`:
				daoModel.AfterInsert[k] = v
			case `menu_id_arr`:
				daoModel.AfterInsert[k] = v
			default:
				if daoThis.Contains(k) {
					insertData[k] = v
				}
			}
		}
		m = m.Data(insertData)
		if len(daoModel.AfterInsert) > 0 {
			m = m.Hook(daoThis.HookInsert(daoModel))
		}
		return m
	}
}

// hook insert
func (daoThis *roleDao) HookInsert(daoModel *daoIndex.DaoModel) gdb.HookHandler {
	return gdb.HookHandler{
		Insert: func(ctx context.Context, in *gdb.HookInsertInput) (result sql.Result, err error) {
			result, err = in.Next(ctx)
			if err != nil {
				return
			}
			id, _ := result.LastInsertId()

			for k, v := range daoModel.AfterInsert {
				switch k {
				case `action_id_arr`:
					vArr := gconv.SliceAny(v)
					insertList := make([]map[string]any, len(vArr))
					for index, item := range vArr {
						insertList[index] = map[string]any{RoleRelToAction.Columns().RoleId: id, RoleRelToAction.Columns().ActionId: item}
					}
					RoleRelToAction.CtxDaoModel(ctx).Data(insertList).Insert()
				case `menu_id_arr`:
					vArr := gconv.SliceAny(v)
					insertList := make([]map[string]any, len(vArr))
					for index, item := range vArr {
						insertList[index] = map[string]any{RoleRelToMenu.Columns().RoleId: id, RoleRelToMenu.Columns().MenuId: item}
					}
					RoleRelToMenu.CtxDaoModel(ctx).Data(insertList).Insert()
				}
			}
			return
		},
	}
}

// 解析update
func (daoThis *roleDao) ParseUpdate(update map[string]any, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		updateData := map[string]any{}
		for k, v := range update {
			switch k {
			case `action_id_arr`:
				daoModel.AfterUpdate[k] = v
			case `menu_id_arr`:
				daoModel.AfterUpdate[k] = v
			default:
				if daoThis.Contains(k) {
					updateData[k] = v
				}
			}
		}
		m = m.Data(updateData)
		if len(daoModel.AfterUpdate) == 0 {
			return m
		}
		m = m.Hook(daoThis.HookUpdate(daoModel))
		if len(updateData) == 0 {
			daoModel.IsOnlyAfterUpdate = true
		}
		return m
	}
}

// hook update
func (daoThis *roleDao) HookUpdate(daoModel *daoIndex.DaoModel) gdb.HookHandler {
	return gdb.HookHandler{
		Update: func(ctx context.Context, in *gdb.HookUpdateInput) (result sql.Result, err error) {
			if daoModel.IsOnlyAfterUpdate {
				result = driver.RowsAffected(0)
			} else {
				result, err = in.Next(ctx)
				if err != nil {
					return
				}
			}

			for k, v := range daoModel.AfterUpdate {
				switch k {
				case `action_id_arr`:
					// daoIndex.SaveArrRelManyWithSort(ctx, &RoleRelToAction, RoleRelToAction.Columns().RoleId, RoleRelToAction.Columns().ActionId, gconv.SliceAny(daoModel.IdArr), gconv.SliceAny(v)) // 有顺序要求时使用，同时注释下面代码
					valArr := gconv.Strings(v)
					for _, id := range daoModel.IdArr {
						daoIndex.SaveArrRelMany(ctx, &RoleRelToAction, RoleRelToAction.Columns().RoleId, RoleRelToAction.Columns().ActionId, id, valArr)
					}
				case `menu_id_arr`:
					// daoIndex.SaveArrRelManyWithSort(ctx, &RoleRelToMenu, RoleRelToMenu.Columns().RoleId, RoleRelToMenu.Columns().MenuId, gconv.SliceAny(daoModel.IdArr), gconv.SliceAny(v)) // 有顺序要求时使用，同时注释下面代码
					valArr := gconv.Strings(v)
					for _, id := range daoModel.IdArr {
						daoIndex.SaveArrRelMany(ctx, &RoleRelToMenu, RoleRelToMenu.Columns().RoleId, RoleRelToMenu.Columns().MenuId, id, valArr)
					}
				}
			}

			daoThis.CacheDeleteInfo(ctx, gconv.Uints(daoModel.IdArr)...)

			/* row, _ := result.RowsAffected()
			if row == 0 {
				return
			} */

			/* for k, v := range daoModel.AfterUpdate {
				switch k {
				case `xxxx`:
					for _, id := range daoModel.IdArr {
						daoModel.CloneNew().FilterPri(id).HookUpdate(g.Map{k: v}).Update()
					}
				}
			} */
			return
		},
	}
}

// hook delete
func (daoThis *roleDao) HookDelete(daoModel *daoIndex.DaoModel) gdb.HookHandler {
	return gdb.HookHandler{
		Delete: func(ctx context.Context, in *gdb.HookDeleteInput) (result sql.Result, err error) { //有软删除字段时需改成Update事件
			result, err = in.Next(ctx)
			if err != nil {
				return
			}

			row, _ := result.RowsAffected()
			if row == 0 {
				return
			}

			daoThis.CacheDeleteInfo(ctx, gconv.Uints(daoModel.IdArr)...)

			RoleRelToAction.CtxDaoModel(ctx).Filter(RoleRelToAction.Columns().RoleId, daoModel.IdArr).Delete()
			RoleRelToMenu.CtxDaoModel(ctx).Filter(RoleRelToMenu.Columns().RoleId, daoModel.IdArr).Delete()
			/* // 对并发有要求时，可使用以下代码解决情形1。并发说明请参考：api/internal/dao/auth/scene.go中HookDelete方法内的注释
			RoleRelOfOrgAdmin.CtxDaoModel(ctx).Filter(RoleRelOfOrgAdmin.Columns().RoleId, daoModel.IdArr).Delete()
			RoleRelOfPlatformAdmin.CtxDaoModel(ctx).Filter(RoleRelOfPlatformAdmin.Columns().RoleId, daoModel.IdArr).Delete() */
			return
		},
	}
}

// 解析group
func (daoThis *roleDao) ParseGroup(group []string, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for _, v := range group {
			switch v {
			case `id`:
				m = m.Group(daoModel.DbTable + `.` + daoThis.Columns().RoleId)
			default:
				if daoThis.Contains(v) {
					m = m.Group(daoModel.DbTable + `.` + v)
				} else {
					m = m.Group(v)
				}
			}
		}
		return m
	}
}

// 解析order
func (daoThis *roleDao) ParseOrder(order []string, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for _, v := range order {
			v = gstr.Trim(v)
			kArr := gstr.Split(v, `,`)
			k := gstr.Split(kArr[0], ` `)[0]
			switch k {
			case `id`:
				m = m.Order(daoModel.DbTable + `.` + gstr.Replace(v, k, daoThis.Columns().RoleId, 1))
			default:
				if daoThis.Contains(k) {
					m = m.Order(daoModel.DbTable + `.` + v)
				} else {
					m = m.Order(v)
				}
			}
		}
		return m
	}
}

// 解析join
func (daoThis *roleDao) ParseJoin(joinTable string, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		if daoModel.JoinTableSet.Contains(joinTable) {
			return m
		}
		daoModel.JoinTableSet.Add(joinTable)
		switch joinTable {
		/* case Xxxx.ParseDbTable(m.GetCtx()):
		m = m.LeftJoin(joinTable, joinTable+`.`+Xxxx.Columns().XxxxId+` = `+daoModel.DbTable+`.`+daoThis.Columns().XxxxId)
		// m = m.LeftJoin(Xxxx.ParseDbTable(m.GetCtx())+` AS `+joinTable, joinTable+`.`+Xxxx.Columns().XxxxId+` = `+daoModel.DbTable+`.`+daoThis.Columns().XxxxId) */
		case Scene.ParseDbTable(m.GetCtx()):
			m = m.LeftJoin(joinTable, joinTable+`.`+Scene.Columns().SceneId+` = `+daoModel.DbTable+`.`+daoThis.Columns().SceneId)
		case RoleRelToAction.ParseDbTable(m.GetCtx()):
			m = m.LeftJoin(joinTable, joinTable+`.`+RoleRelToAction.Columns().RoleId+` = `+daoModel.DbTable+`.`+daoThis.Columns().RoleId)
		case RoleRelToMenu.ParseDbTable(m.GetCtx()):
			m = m.LeftJoin(joinTable, joinTable+`.`+RoleRelToMenu.Columns().RoleId+` = `+daoModel.DbTable+`.`+daoThis.Columns().RoleId)
		case RoleRelOfPlatformAdmin.ParseDbTable(m.GetCtx()):
			m = m.LeftJoin(joinTable, joinTable+`.`+RoleRelOfPlatformAdmin.Columns().RoleId+` = `+daoModel.DbTable+`.`+daoThis.Columns().RoleId)
		case RoleRelOfOrgAdmin.ParseDbTable(m.GetCtx()):
			m = m.LeftJoin(joinTable, joinTable+`.`+RoleRelOfOrgAdmin.Columns().RoleId+` = `+daoModel.DbTable+`.`+daoThis.Columns().RoleId)
		default:
			m = m.LeftJoin(joinTable, joinTable+`.`+daoThis.Columns().RoleId+` = `+daoModel.DbTable+`.`+daoThis.Columns().RoleId)
		}
		return m
	}
}

// Fill with you ideas below.

func (daoThis *roleDao) CacheGetInfo(ctx context.Context, id uint) (info gdb.Record, err error) {
	value, err := cache.DbData.GetOrSet(ctx, daoThis.CtxDaoModel(ctx), id, consts.CACHE_TIME_DEFAULT, append(daoThis.ColumnArr(), `action_id_arr`, `menu_id_arr`)...)
	if err != nil {
		return
	}
	value.Scan(&info)
	return
}

func (daoThis *roleDao) CacheDeleteInfo(ctx context.Context, idArr ...uint) (row int64, err error) {
	row, err = cache.DbData.Del(ctx, daoThis.CtxDaoModel(ctx), gconv.SliceAny(idArr)...)
	return
}

func (daoThis *roleDao) CacheGetActionIdArr(ctx context.Context, idArr ...uint) (actionIdArr []string, err error) {
	for _, id := range idArr {
		var info gdb.Record
		info, err = daoThis.CacheGetInfo(ctx, id)
		if err != nil {
			return
		}
		if !info.IsEmpty() {
			actionIdArr = append(actionIdArr, info[`action_id_arr`].Strings()...)
		}
	}
	actionIdArr = garray.NewStrArrayFrom(actionIdArr).Unique().Slice()
	return
}

func (daoThis *roleDao) CacheGetMenuIdArr(ctx context.Context, idArr ...uint) (menuIdArr []uint, err error) {
	menuIdArrG := garray.New()
	for _, id := range idArr {
		var info gdb.Record
		info, err = daoThis.CacheGetInfo(ctx, id)
		if err != nil {
			return
		}
		if !info.IsEmpty() {
			menuIdArrG.Append(info[`menu_id_arr`].Interfaces()...)
		}
	}
	menuIdArr = gconv.Uints(menuIdArrG.Unique())
	return
}

func (daoThis *roleDao) GetRoleIdArrOfSelf(ctx context.Context, sceneId string, loginId *gvar.Var) (roleIdArr []uint, err error) {
	roleIdArr, err = daoThis.CtxDaoModel(ctx).Filter(`self_role`, g.Map{`scene_id`: sceneId, `login_id`: loginId}).ArrayUint(daoThis.Columns().RoleId)
	return
}
