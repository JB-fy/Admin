// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	daoIndex "api/internal/dao"
	"api/internal/dao/auth/internal"
	"context"
	"database/sql"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
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
func (daoThis *roleDao) CtxDaoModel(ctx context.Context, dbOpt ...map[string]interface{}) *daoIndex.DaoModel {
	return daoIndex.NewDaoModel(ctx, daoThis, dbOpt...)
}

// 解析分库
func (daoThis *roleDao) ParseDbGroup(ctx context.Context, dbGroupOpt ...map[string]interface{}) string {
	group := daoThis.Group()
	// 分库逻辑
	/* if len(dbGroupOpt) > 0 {
	} */
	return group
}

// 解析分表
func (daoThis *roleDao) ParseDbTable(ctx context.Context, dbTableOpt ...map[string]interface{}) string {
	table := daoThis.Table()
	// 分表逻辑
	/* if len(dbTableOpt) > 0 {
	} */
	return table
}

// 解析filter
func (daoThis *roleDao) ParseFilter(filter map[string]interface{}, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for k, v := range filter {
			switch k {
			/* case `xxxx`:
			tableXxxx := Xxxx.ParseDbTable(m.GetCtx())
			m = m.Where(tableXxxx+`.`+k, v)
			m = m.Handler(daoThis.ParseJoin(tableXxxx, daoModel)) */
			case `id`, `idArr`:
				m = m.Where(daoModel.DbTable+`.`+daoThis.PrimaryKey(), v)
			case `excId`, `excIdArr`:
				if gvar.New(v).IsSlice() {
					m = m.WhereNotIn(daoModel.DbTable+`.`+daoThis.PrimaryKey(), v)
				} else {
					m = m.WhereNot(daoModel.DbTable+`.`+daoThis.PrimaryKey(), v)
				}
			case `label`:
				m = m.WhereLike(daoModel.DbTable+`.`+daoThis.Columns().RoleName, `%`+gconv.String(v)+`%`)
			case daoThis.Columns().RoleName:
				m = m.WhereLike(daoModel.DbTable+`.`+k, `%`+gconv.String(v)+`%`)
			case `timeRangeStart`:
				m = m.WhereGTE(daoModel.DbTable+`.`+daoThis.Columns().CreatedAt, v)
			case `timeRangeEnd`:
				m = m.WhereLTE(daoModel.DbTable+`.`+daoThis.Columns().CreatedAt, v)
			case RoleRelToAction.Columns().ActionId:
				tableRoleRelToAction := RoleRelToAction.ParseDbTable(m.GetCtx())
				m = m.Where(tableRoleRelToAction+`.`+k, v)
				m = m.Handler(daoThis.ParseJoin(tableRoleRelToAction, daoModel))
			case RoleRelToMenu.Columns().MenuId:
				tableRoleRelToMenu := RoleRelToMenu.ParseDbTable(m.GetCtx())
				m = m.Where(tableRoleRelToMenu+`.`+k, v)
				m = m.Handler(daoThis.ParseJoin(tableRoleRelToMenu, daoModel))
			case Scene.Columns().SceneCode:
				tableScene := Scene.ParseDbTable(m.GetCtx())
				m = m.Where(tableScene+`.`+k, v)
				m = m.Handler(daoThis.ParseJoin(tableScene, daoModel))
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
func (daoThis *roleDao) ParseField(field []string, fieldWithParam map[string]interface{}, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for _, v := range field {
			switch v {
			/* case `xxxx`:
			tableXxxx := Xxxx.ParseDbTable(m.GetCtx())
			m = m.Fields(tableXxxx + `.` + v)
			m = m.Handler(daoThis.ParseJoin(tableXxxx, daoModel))
			daoModel.AfterField.Add(v) */
			case `id`:
				m = m.Fields(daoModel.DbTable + `.` + daoThis.PrimaryKey() + ` AS ` + v)
			case `label`:
				m = m.Fields(daoModel.DbTable + `.` + daoThis.Columns().RoleName + ` AS ` + v)
			case Scene.Columns().SceneName:
				tableScene := Scene.ParseDbTable(m.GetCtx())
				m = m.Fields(tableScene + `.` + v)
				m = m.Handler(daoThis.ParseJoin(tableScene, daoModel))
			case `actionIdArr`:
				m = m.Fields(daoModel.DbTable + `.` + daoThis.PrimaryKey())
				daoModel.AfterField.Add(v)
			case `menuIdArr`:
				m = m.Fields(daoModel.DbTable + `.` + daoThis.PrimaryKey())
				daoModel.AfterField.Add(v)
			case `tableName`:
				m = m.Fields(daoModel.DbTable + `.` + daoThis.Columns().TableId)
				tableScene := Scene.ParseDbTable(m.GetCtx())
				m = m.Fields(tableScene + `.` + Scene.Columns().SceneCode)
				m = m.Handler(daoThis.ParseJoin(tableScene, daoModel))
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
func (daoThis *roleDao) HookSelect(daoModel *daoIndex.DaoModel) gdb.HookHandler {
	return gdb.HookHandler{
		Select: func(ctx context.Context, in *gdb.HookSelectInput) (result gdb.Result, err error) {
			result, err = in.Next(ctx)
			if err != nil {
				return
			}
			for _, record := range result {
				for _, v := range daoModel.AfterField.Slice() {
					switch v {
					case `actionIdArr`:
						actionIdArr, _ := RoleRelToAction.CtxDaoModel(ctx).Filter(RoleRelToAction.Columns().RoleId, record[daoThis.PrimaryKey()]).Array(RoleRelToAction.Columns().ActionId)
						record[v] = gvar.New(actionIdArr)
					case `menuIdArr`:
						menuIdArr, _ := RoleRelToMenu.CtxDaoModel(ctx).Filter(RoleRelToMenu.Columns().RoleId, record[daoThis.PrimaryKey()]).Array(RoleRelToMenu.Columns().MenuId)
						record[v] = gvar.New(menuIdArr)
					case `tableName`:
						if record[daoThis.Columns().TableId].Uint() == 0 {
							record[v] = gvar.New(`平台`)
							continue
						}
						switch record[Scene.Columns().SceneCode].String() {
						case `platform`:
						}
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
			return
		},
	}
}

// 解析insert
func (daoThis *roleDao) ParseInsert(insert map[string]interface{}, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		insertData := map[string]interface{}{}
		for k, v := range insert {
			switch k {
			case `actionIdArr`:
				daoModel.AfterInsert[k] = v
			case `menuIdArr`:
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
				case `actionIdArr`:
					insertList := []map[string]interface{}{}
					for _, item := range gconv.SliceAny(v) {
						insertList = append(insertList, map[string]interface{}{
							RoleRelToAction.Columns().RoleId:   id,
							RoleRelToAction.Columns().ActionId: item,
						})
					}
					RoleRelToAction.CtxDaoModel(ctx).Data(insertList).Insert()
				case `menuIdArr`:
					insertList := []map[string]interface{}{}
					for _, item := range gconv.SliceAny(v) {
						insertList = append(insertList, map[string]interface{}{
							RoleRelToMenu.Columns().RoleId: id,
							RoleRelToMenu.Columns().MenuId: item,
						})
					}
					RoleRelToMenu.CtxDaoModel(ctx).Data(insertList).Insert()
				}
			}
			return
		},
	}
}

// 解析update
func (daoThis *roleDao) ParseUpdate(update map[string]interface{}, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		updateData := map[string]interface{}{}
		for k, v := range update {
			switch k {
			case `actionIdArr`:
				daoModel.AfterUpdate[k] = v
			case `menuIdArr`:
				daoModel.AfterUpdate[k] = v
			default:
				if daoThis.ColumnArr().Contains(k) {
					updateData[daoModel.DbTable+`.`+k] = gvar.New(v) //json类型字段传参必须是gvar变量（原因：下面BUG解决方式导致map类型数据更新时，不会自动转换json）
				}
			}
		}
		// m = m.Data(updateData) // 2.5某版本之前，字段被解析成`table.xxxx`，正确的应该是`table`.`xxxx`	// 2.6版本开始更过分，居然直接把字段过滤掉不做更新，报错都没有
		// 上面方法的BUG解决方式
		fieldArr := []string{}
		valueArr := []interface{}{}
		for k, v := range updateData {
			if _, ok := v.(gdb.Raw); ok {
				fieldArr = append(fieldArr, k+` = `+gconv.String(v))
			} else {
				fieldArr = append(fieldArr, k+` = ?`)
				valueArr = append(valueArr, v)
			}
		}
		data := []interface{}{gstr.Join(fieldArr, `,`)}
		data = append(data, valueArr...)
		m = m.Data(data...)
		return m
	}
}

// hook update
func (daoThis *roleDao) HookUpdate(daoModel *daoIndex.DaoModel) gdb.HookHandler {
	return gdb.HookHandler{
		Update: func(ctx context.Context, in *gdb.HookUpdateInput) (result sql.Result, err error) {
			result, err = in.Next(ctx)
			if err != nil {
				return
			}

			for k, v := range daoModel.AfterUpdate {
				switch k {
				case `actionIdArr`:
					// daoIndex.SaveArrRelManyWithSort(ctx, &RoleRelToAction, RoleRelToAction.Columns().RoleId, RoleRelToAction.Columns().ActionId, gconv.SliceAny(daoModel.IdArr), gconv.SliceAny(v)) // 有顺序要求时使用，同时注释下面代码
					valArr := gconv.SliceStr(v)
					for _, id := range daoModel.IdArr {
						daoIndex.SaveArrRelMany(ctx, &RoleRelToAction, RoleRelToAction.Columns().RoleId, RoleRelToAction.Columns().ActionId, id, valArr)
					}
				case `menuIdArr`:
					// daoIndex.SaveArrRelManyWithSort(ctx, &RoleRelToMenu, RoleRelToMenu.Columns().RoleId, RoleRelToMenu.Columns().MenuId, gconv.SliceAny(daoModel.IdArr), gconv.SliceAny(v)) // 有顺序要求时使用，同时注释下面代码
					valArr := gconv.SliceStr(v)
					for _, id := range daoModel.IdArr {
						daoIndex.SaveArrRelMany(ctx, &RoleRelToMenu, RoleRelToMenu.Columns().RoleId, RoleRelToMenu.Columns().MenuId, id, valArr)
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
						daoModel.CloneNew().Filter(daoThis.PrimaryKey(), id).HookUpdate(g.Map{k: v}).Update()
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
		Delete: func(ctx context.Context, in *gdb.HookDeleteInput) (result sql.Result, err error) {
			result, err = in.Next(ctx)
			if err != nil {
				return
			}
			row, _ := result.RowsAffected()
			if row == 0 {
				return
			}

			RoleRelToAction.CtxDaoModel(ctx).Filter(RoleRelToAction.Columns().RoleId, daoModel.IdArr).Delete()
			RoleRelToMenu.CtxDaoModel(ctx).Filter(RoleRelToMenu.Columns().RoleId, daoModel.IdArr).Delete()
			RoleRelOfPlatformAdmin.CtxDaoModel(ctx).Filter(RoleRelOfPlatformAdmin.Columns().RoleId, daoModel.IdArr).Delete()
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
				m = m.Group(daoModel.DbTable + `.` + daoThis.PrimaryKey())
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
func (daoThis *roleDao) ParseOrder(order []string, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for _, v := range order {
			v = gstr.Trim(v)
			kArr := gstr.Split(v, `,`)
			k := gstr.Split(kArr[0], ` `)[0]
			switch k {
			case `id`:
				m = m.Order(daoModel.DbTable + `.` + gstr.Replace(v, k, daoThis.PrimaryKey(), 1))
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
func (daoThis *roleDao) ParseJoin(joinTable string, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		if daoModel.JoinTableSet.Contains(joinTable) {
			return m
		}
		daoModel.JoinTableSet.Add(joinTable)
		switch joinTable {
		/* case Xxxx.ParseDbTable(m.GetCtx()):
		m = m.LeftJoin(joinTable, joinTable+`.`+Xxxx.Columns().XxxxId+` = `+daoModel.DbTable+`.`+daoThis.PrimaryKey())
		// m = m.LeftJoin(Xxxx.ParseDbTable(m.GetCtx())+` AS `+joinTable, joinTable+`.`+Xxxx.Columns().XxxxId+` = `+daoModel.DbTable+`.`+daoThis.PrimaryKey()) */
		case Scene.ParseDbTable(m.GetCtx()):
			m = m.LeftJoin(joinTable, joinTable+`.`+Scene.PrimaryKey()+` = `+daoModel.DbTable+`.`+daoThis.Columns().SceneId)
		case RoleRelToAction.ParseDbTable(m.GetCtx()):
			m = m.LeftJoin(joinTable, joinTable+`.`+RoleRelToAction.Columns().RoleId+` = `+daoModel.DbTable+`.`+daoThis.PrimaryKey())
		case RoleRelToMenu.ParseDbTable(m.GetCtx()):
			m = m.LeftJoin(joinTable, joinTable+`.`+RoleRelToMenu.Columns().RoleId+` = `+daoModel.DbTable+`.`+daoThis.PrimaryKey())
		default:
			m = m.LeftJoin(joinTable, joinTable+`.`+daoThis.PrimaryKey()+` = `+daoModel.DbTable+`.`+daoThis.PrimaryKey())
		}
		return m
	}
}

// Fill with you ideas below.
