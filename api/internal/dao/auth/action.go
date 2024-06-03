// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	daoIndex "api/internal/dao"
	"api/internal/dao/auth/internal"
	"context"
	"database/sql"
	"database/sql/driver"
	"sync"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

// internalActionDao is internal type for wrapping internal DAO implements.
type internalActionDao = *internal.ActionDao

// actionDao is the data access object for table auth_action.
// You can define custom methods on it to extend its functionality as you wish.
type actionDao struct {
	internalActionDao
}

var (
	// Action is globally public accessible object for table auth_action operations.
	Action = actionDao{
		internal.NewActionDao(),
	}
)

// 获取daoModel
func (daoThis *actionDao) CtxDaoModel(ctx context.Context, dbOpt ...map[string]any) *daoIndex.DaoModel {
	return daoIndex.NewDaoModel(ctx, daoThis, dbOpt...)
}

// 解析分库
func (daoThis *actionDao) ParseDbGroup(ctx context.Context, dbGroupOpt ...map[string]any) string {
	group := daoThis.Group()
	// 分库逻辑
	/* if len(dbGroupOpt) > 0 {
	} */
	return group
}

// 解析分表
func (daoThis *actionDao) ParseDbTable(ctx context.Context, dbTableOpt ...map[string]any) string {
	table := daoThis.Table()
	// 分表逻辑
	/* if len(dbTableOpt) > 0 {
	} */
	return table
}

// 解析Id（未使用代码自动生成，且id字段不在第1个位置时，需手动修改）
func (daoThis *actionDao) ParseId(daoModel *daoIndex.DaoModel) string {
	return daoModel.DbTable + `.` + daoThis.Columns().ActionId
}

// 解析Label（未使用代码自动生成，且id字段不在第2个位置时，需手动修改）
func (daoThis *actionDao) ParseLabel(daoModel *daoIndex.DaoModel) string {
	return daoModel.DbTable + `.` + daoThis.Columns().ActionName
}

// 解析filter
func (daoThis *actionDao) ParseFilter(filter map[string]any, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for k, v := range filter {
			switch k {
			/* case `xxxx`:
			tableXxxx := Xxxx.ParseDbTable(m.GetCtx())
			m = m.Where(tableXxxx+`.`+k, v)
			m = m.Handler(daoThis.ParseJoin(tableXxxx, daoModel)) */
			case `id`, `id_arr`:
				m = m.Where(daoModel.DbTable+`.`+daoThis.Columns().ActionId, v)
			case `exc_id`, `exc_id_arr`:
				if gvar.New(v).IsSlice() {
					m = m.WhereNotIn(daoModel.DbTable+`.`+daoThis.Columns().ActionId, v)
				} else {
					m = m.WhereNot(daoModel.DbTable+`.`+daoThis.Columns().ActionId, v)
				}
			case `label`:
				m = m.WhereLike(daoModel.DbTable+`.`+daoThis.Columns().ActionName, `%`+gconv.String(v)+`%`)
			case daoThis.Columns().ActionName:
				m = m.WhereLike(daoModel.DbTable+`.`+k, `%`+gconv.String(v)+`%`)
			case `time_range_start`:
				m = m.WhereGTE(daoModel.DbTable+`.`+daoThis.Columns().CreatedAt, v)
			case `time_range_end`:
				m = m.WhereLTE(daoModel.DbTable+`.`+daoThis.Columns().CreatedAt, v)
			case ActionRelToScene.Columns().SceneId:
				tableActionRelToScene := ActionRelToScene.ParseDbTable(m.GetCtx())
				m = m.Where(tableActionRelToScene+`.`+k, v)
				m = m.Handler(daoThis.ParseJoin(tableActionRelToScene, daoModel))
			case `self_action`: //获取当前登录身份可用的操作。参数：map[string]any{`scene_code`: `场景标识`, `login_id`: 登录身份id, `scene_id`: 场景id（平台超级管理员用，不用再做一次查询）}
				m = m.Where(daoModel.DbTable+`.`+daoThis.Columns().IsStop, 0)
				val := gconv.Map(v)
				switch gconv.String(val[`scene_code`]) {
				case `platform`:
					if gconv.Uint(val[`login_id`]) == g.Cfg().MustGet(m.GetCtx(), `superPlatformAdminId`).Uint() { //平台超级管理员
						tableActionRelToScene := ActionRelToScene.ParseDbTable(m.GetCtx())
						m = m.Where(tableActionRelToScene+`.`+ActionRelToScene.Columns().SceneId, val[`scene_id`])
						m = m.Handler(daoThis.ParseJoin(tableActionRelToScene, daoModel))
						continue
					}
					roleIdArr, _ := Role.CtxDaoModel(m.GetCtx()).Fields(Role.Columns().RoleId).Filter(`self_role`, val).Array()
					if len(roleIdArr) == 0 {
						m = m.Where(`1 = 0`)
						continue
					}
					/* // 方式一：联表查询（不推荐。原因：auth_role及其关联表，后期表数据只会越来越大，故不建议联表）
					tableRoleRelToAction := RoleRelToAction.ParseDbTable(m.GetCtx())
					m = m.Where(tableRoleRelToAction+`.`+RoleRelToAction.Columns().RoleId, roleIdArr)
					m = m.Handler(daoThis.ParseJoin(tableRoleRelToAction, daoModel))
					m = m.Group(daoModel.DbTable + `.` + daoThis.Columns().ActionId) */
					// 方式二：非联表查询
					actionIdArr, _ := RoleRelToAction.CtxDaoModel(m.GetCtx()).Filter(RoleRelToAction.Columns().RoleId, roleIdArr).Distinct().Array(RoleRelToAction.Columns().ActionId)
					if len(actionIdArr) == 0 {
						m = m.Where(`1 = 0`)
						continue
					}
					m = m.Where(daoModel.DbTable+`.`+daoThis.Columns().ActionId, actionIdArr)
				default:
					m = m.Where(`1 = 0`)
				}
			default:
				if daoThis.ColumnArr().Contains(k) {
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
func (daoThis *actionDao) ParseField(field []string, fieldWithParam map[string]any, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
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
			case `scene_id_arr`:
				m = m.Fields(daoModel.DbTable + `.` + daoThis.Columns().ActionId)
				daoModel.AfterField.Add(v)
			default:
				if daoThis.ColumnArr().Contains(v) {
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
		return m
	}
}

// hook select
func (daoThis *actionDao) HookSelect(daoModel *daoIndex.DaoModel) gdb.HookHandler {
	return gdb.HookHandler{
		Select: func(ctx context.Context, in *gdb.HookSelectInput) (result gdb.Result, err error) {
			result, err = in.Next(ctx)
			if err != nil || len(result) == 0 {
				return
			}

			var wg sync.WaitGroup
			wg.Add(len(result))
			afterFieldHandleFunc := func(record gdb.Record) {
				defer wg.Done()
				for _, v := range daoModel.AfterField.Slice() {
					switch v {
					case `scene_id_arr`:
						sceneIdArr, _ := ActionRelToScene.CtxDaoModel(ctx).Filter(ActionRelToScene.Columns().ActionId, record[daoThis.Columns().ActionId]).Array(ActionRelToScene.Columns().SceneId)
						record[v] = gvar.New(sceneIdArr)
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
			for _, record := range result {
				go afterFieldHandleFunc(record)
			}
			wg.Wait()
			return
		},
	}
}

// 解析insert
func (daoThis *actionDao) ParseInsert(insert map[string]any, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		insertData := map[string]any{}
		for k, v := range insert {
			switch k {
			case `scene_id_arr`:
				daoModel.AfterInsert[k] = v
			default:
				if daoThis.ColumnArr().Contains(k) {
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
func (daoThis *actionDao) HookInsert(daoModel *daoIndex.DaoModel) gdb.HookHandler {
	return gdb.HookHandler{
		Insert: func(ctx context.Context, in *gdb.HookInsertInput) (result sql.Result, err error) {
			result, err = in.Next(ctx)
			if err != nil {
				return
			}
			id, _ := result.LastInsertId()

			for k, v := range daoModel.AfterInsert {
				switch k {
				case `scene_id_arr`:
					insertList := []map[string]any{}
					for _, item := range gconv.SliceAny(v) {
						insertList = append(insertList, map[string]any{
							ActionRelToScene.Columns().ActionId: id,
							ActionRelToScene.Columns().SceneId:  item,
						})
					}
					ActionRelToScene.CtxDaoModel(ctx).Data(insertList).Insert()
				}
			}
			return
		},
	}
}

// 解析update
func (daoThis *actionDao) ParseUpdate(update map[string]any, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		updateData := map[string]any{}
		for k, v := range update {
			switch k {
			case `scene_id_arr`:
				daoModel.AfterUpdate[k] = v
			default:
				if daoThis.ColumnArr().Contains(k) {
					updateData[k] = v
				}
			}
		}
		if len(updateData) > 0 {
			m = m.Data(updateData)
		} else if len(daoModel.AfterUpdate) > 0 {
			daoModel.IsOnlyAfterUpdate = true
		}
		return m
	}
}

// hook update
func (daoThis *actionDao) HookUpdate(daoModel *daoIndex.DaoModel) gdb.HookHandler {
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
				case `scene_id_arr`:
					// daoIndex.SaveArrRelManyWithSort(ctx, &ActionRelToScene, ActionRelToScene.Columns().ActionId, ActionRelToScene.Columns().SceneId, gconv.SliceAny(daoModel.IdArr), gconv.SliceAny(v)) // 有顺序要求时使用，同时注释下面代码
					valArr := gconv.SliceStr(v)
					for _, id := range daoModel.IdArr {
						daoIndex.SaveArrRelMany(ctx, &ActionRelToScene, ActionRelToScene.Columns().ActionId, ActionRelToScene.Columns().SceneId, id, valArr)
					}
				}
			}

			/* row, _ := result.RowsAffected()
			if row == 0 {
				return
			} */

			/* for k, v := range daoModel.AfterUpdate {
				switch k {
				case `xxxx`:
					for _, id := range daoModel.IdArr {
						daoModel.CloneNew().Filter(`id`, id).HookUpdate(g.Map{k: v}).Update()
					}
				}
			} */
			return
		},
	}
}

// hook delete
func (daoThis *actionDao) HookDelete(daoModel *daoIndex.DaoModel) gdb.HookHandler {
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

			ActionRelToScene.CtxDaoModel(ctx).Filter(ActionRelToScene.Columns().ActionId, daoModel.IdArr).Delete()
			/* // 对并发有要求时，可使用以下代码解决情形1。并发说明请参考：api/internal/dao/auth/scene.go中HookDelete方法内的注释
			RoleRelToAction.CtxDaoModel(ctx).Filter(RoleRelToAction.Columns().ActionId, daoModel.IdArr).Delete() */
			return
		},
	}
}

// 解析group
func (daoThis *actionDao) ParseGroup(group []string, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for _, v := range group {
			switch v {
			case `id`:
				m = m.Group(daoModel.DbTable + `.` + daoThis.Columns().ActionId)
			default:
				if daoThis.ColumnArr().Contains(v) {
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
func (daoThis *actionDao) ParseOrder(order []string, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for _, v := range order {
			v = gstr.Trim(v)
			kArr := gstr.Split(v, `,`)
			k := gstr.Split(kArr[0], ` `)[0]
			switch k {
			case `id`:
				m = m.Order(daoModel.DbTable + `.` + gstr.Replace(v, k, daoThis.Columns().ActionId, 1))
			default:
				if daoThis.ColumnArr().Contains(k) {
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
func (daoThis *actionDao) ParseJoin(joinTable string, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		if daoModel.JoinTableSet.Contains(joinTable) {
			return m
		}
		daoModel.JoinTableSet.Add(joinTable)
		switch joinTable {
		/* case Xxxx.ParseDbTable(m.GetCtx()):
		m = m.LeftJoin(joinTable, joinTable+`.`+Xxxx.Columns().XxxxId+` = `+daoModel.DbTable+`.`+daoThis.Columns().XxxxId)
		// m = m.LeftJoin(Xxxx.ParseDbTable(m.GetCtx())+` AS `+joinTable, joinTable+`.`+Xxxx.Columns().XxxxId+` = `+daoModel.DbTable+`.`+daoThis.Columns().XxxxId) */
		case ActionRelToScene.ParseDbTable(m.GetCtx()):
			m = m.LeftJoin(joinTable, joinTable+`.`+ActionRelToScene.Columns().ActionId+` = `+daoModel.DbTable+`.`+daoThis.Columns().ActionId)
		case RoleRelToAction.ParseDbTable(m.GetCtx()):
			m = m.LeftJoin(joinTable, joinTable+`.`+RoleRelToAction.Columns().ActionId+` = `+daoModel.DbTable+`.`+daoThis.Columns().ActionId)
		default:
			m = m.LeftJoin(joinTable, joinTable+`.`+daoThis.Columns().ActionId+` = `+daoModel.DbTable+`.`+daoThis.Columns().ActionId)
		}
		return m
	}
}

// Fill with you ideas below.
