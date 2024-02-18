// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	daoIndex "api/internal/dao"
	"api/internal/dao/auth/internal"
	"context"
	"database/sql"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

// internalMenuDao is internal type for wrapping internal DAO implements.
type internalMenuDao = *internal.MenuDao

// menuDao is the data access object for table auth_menu.
// You can define custom methods on it to extend its functionality as you wish.
type menuDao struct {
	internalMenuDao
}

var (
	// Menu is globally public accessible object for table auth_menu operations.
	Menu = menuDao{
		internal.NewMenuDao(),
	}
)

// 获取daoHandler
func (daoThis *menuDao) HandlerCtx(ctx context.Context, dbOpt ...map[string]interface{}) *daoIndex.DaoHandler {
	return daoIndex.NewDaoHandler(ctx, daoThis, dbOpt...)
}

// 解析分库
func (daoThis *menuDao) ParseDbGroup(ctx context.Context, dbGroupOpt ...map[string]interface{}) string {
	group := daoThis.Group()
	// 分库逻辑
	/* if len(dbGroupOpt) > 0 {
	} */
	return group
}

// 解析分表
func (daoThis *menuDao) ParseDbTable(ctx context.Context, dbTableOpt ...map[string]interface{}) string {
	table := daoThis.Table()
	// 分表逻辑
	/* if len(dbTableOpt) > 0 {
	} */
	return table
}

// 解析分库分表（对外暴露使用）
func (daoThis *menuDao) ParseDbCtx(ctx context.Context, dbOpt ...map[string]interface{}) *gdb.Model {
	switch len(dbOpt) {
	case 1:
		return g.DB(daoThis.ParseDbGroup(ctx, dbOpt[0])).Model(daoThis.ParseDbTable(ctx)). /* Safe(). */ Ctx(ctx)
	case 2:
		return g.DB(daoThis.ParseDbGroup(ctx, dbOpt[0])).Model(daoThis.ParseDbTable(ctx, dbOpt[1])). /* Safe(). */ Ctx(ctx)
	default:
		return g.DB(daoThis.ParseDbGroup(ctx)).Model(daoThis.ParseDbTable(ctx)). /* Safe(). */ Ctx(ctx)
	}
}

// 解析insert
func (daoThis *menuDao) ParseInsert(insert map[string]interface{}, daoHandler *daoIndex.DaoHandler) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		insertData := map[string]interface{}{}
		for k, v := range insert {
			switch k {
			case `id`:
				insertData[daoThis.PrimaryKey()] = v
			case daoThis.Columns().Pid:
				insertData[k] = v
				if gconv.Uint(v) > 0 {
					pInfo, _ := daoThis.ParseDbCtx(m.GetCtx()).Where(daoThis.PrimaryKey(), v).One()
					daoHandler.AfterInsert[`pIdPath`] = pInfo[daoThis.Columns().IdPath].String()
					daoHandler.AfterInsert[`pLevel`] = pInfo[daoThis.Columns().Level].Uint()
				} else {
					daoHandler.AfterInsert[`pIdPath`] = `0`
					daoHandler.AfterInsert[`pLevel`] = 0
				}
			case daoThis.Columns().ExtraData:
				insertData[k] = v
				if gconv.String(v) == `` {
					insertData[k] = nil
				}
			default:
				if daoThis.ColumnArrG().Contains(k) {
					insertData[k] = v
				}
			}
		}
		m = m.Data(insertData)
		if len(daoHandler.AfterInsert) > 0 {
			m = m.Hook(daoThis.HookInsert(daoHandler))
		}
		return m
	}
}

// hook insert
func (daoThis *menuDao) HookInsert(daoHandler *daoIndex.DaoHandler) gdb.HookHandler {
	return gdb.HookHandler{
		Insert: func(ctx context.Context, in *gdb.HookInsertInput) (result sql.Result, err error) {
			result, err = in.Next(ctx)
			if err != nil {
				return
			}
			id, _ := result.LastInsertId()

			updateSelfData := map[string]interface{}{}
			for k, v := range daoHandler.AfterInsert {
				switch k {
				case `pIdPath`:
					updateSelfData[daoThis.Columns().IdPath] = gconv.String(v) + `-` + gconv.String(id)
				case `pLevel`:
					updateSelfData[daoThis.Columns().Level] = gconv.Uint(v) + 1
				}
			}
			if len(updateSelfData) > 0 {
				daoThis.ParseDbCtx(ctx).Where(daoThis.PrimaryKey(), id).Data(updateSelfData).Update()
			}
			return
		},
	}
}

// 解析update
func (daoThis *menuDao) ParseUpdate(update map[string]interface{}, daoHandler *daoIndex.DaoHandler) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		updateData := map[string]interface{}{}
		for k, v := range update {
			switch k {
			case `id`:
				updateData[daoHandler.DbTable+`.`+daoThis.PrimaryKey()] = v
			case daoThis.Columns().Pid:
				updateData[daoHandler.DbTable+`.`+k] = v
				pIdPath := `0`
				var pLevel uint = 0
				if gconv.Uint(v) > 0 {
					pInfo, _ := daoThis.ParseDbCtx(m.GetCtx()).Where(daoThis.PrimaryKey(), v).One()
					pIdPath = pInfo[daoThis.Columns().IdPath].String()
					pLevel = pInfo[daoThis.Columns().Level].Uint()
				}
				updateData[daoHandler.DbTable+`.`+daoThis.Columns().IdPath] = gdb.Raw(`CONCAT('` + pIdPath + `-', ` + daoThis.PrimaryKey() + `)`)
				updateData[daoHandler.DbTable+`.`+daoThis.Columns().Level] = pLevel + 1
				//更新所有子孙级的idPath和level
				updateChildIdPathAndLevelList := []map[string]interface{}{}
				oldList, _ := daoThis.ParseDbCtx(m.GetCtx()).Where(daoThis.PrimaryKey(), daoHandler.IdArr).All()
				for _, oldInfo := range oldList {
					if gconv.Uint(v) != oldInfo[daoThis.Columns().Pid].Uint() {
						updateChildIdPathAndLevelList = append(updateChildIdPathAndLevelList, map[string]interface{}{
							`newIdPath`: pIdPath + `-` + oldInfo[daoThis.PrimaryKey()].String(),
							`oldIdPath`: oldInfo[daoThis.Columns().IdPath],
							`newLevel`:  pLevel + 1,
							`oldLevel`:  oldInfo[daoThis.Columns().Level],
						})
					}
				}
				if len(updateChildIdPathAndLevelList) > 0 {
					daoHandler.AfterUpdate[`updateChildIdPathAndLevelList`] = updateChildIdPathAndLevelList
				}
			case daoThis.Columns().ExtraData:
				updateData[daoHandler.DbTable+`.`+k] = gvar.New(v)
				if gconv.String(v) == `` {
					updateData[daoHandler.DbTable+`.`+k] = nil
				}
			default:
				if daoThis.ColumnArrG().Contains(k) {
					updateData[daoHandler.DbTable+`.`+k] = gvar.New(v) //因下面bug处理方式，json类型字段传参必须是gvar变量，否则不会自动生成json格式
				}
			}
		}
		//m = m.Data(updateData) //字段被解析成`table.xxxx`，正确的应该是`table`.`xxxx`
		//解决字段被解析成`table.xxxx`的BUG
		fieldArr := []string{}
		valueArr := []interface{}{}
		for k, v := range updateData {
			_, ok := v.(gdb.Raw)
			if ok {
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
func (daoThis *menuDao) HookUpdate(daoHandler *daoIndex.DaoHandler) gdb.HookHandler {
	return gdb.HookHandler{
		Update: func(ctx context.Context, in *gdb.HookUpdateInput) (result sql.Result, err error) {
			result, err = in.Next(ctx)
			if err != nil {
				return
			}

			row, _ := result.RowsAffected()
			if row == 0 {
				return
			}

			for k, v := range daoHandler.AfterUpdate {
				switch k {
				case `updateChildIdPathAndLevelList`: //修改pid时，更新所有子孙级的idPath和level。参数：[]map[string]interface{}{newIdPath: 父级新idPath, oldIdPath: 父级旧idPath, newLevel: 父级新level, oldLevel: 父级旧level}
					val := v.([]map[string]interface{})
					for _, v1 := range val {
						daoThis.UpdateChildIdPathAndLevel(ctx, gconv.String(v1[`newIdPath`]), gconv.String(v1[`oldIdPath`]), gconv.Uint(v1[`newLevel`]), gconv.Uint(v1[`oldLevel`]))
					}
				}
			}
			return
		},
	}
}

// hook delete
func (daoThis *menuDao) HookDelete(daoHandler *daoIndex.DaoHandler) gdb.HookHandler {
	return gdb.HookHandler{
		Delete: func(ctx context.Context, in *gdb.HookDeleteInput) (result sql.Result, err error) {
			result, err = in.Next(ctx)
			if err != nil {
				return
			}
			/* row, _ := result.RowsAffected()
			if row == 0 {
				return
			} */
			return
		},
	}
}

// 解析field
func (daoThis *menuDao) ParseField(field []string, fieldWithParam map[string]interface{}, daoHandler *daoIndex.DaoHandler) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for _, v := range field {
			switch v {
			/* case `xxxx`:
			m = m.Handler(daoThis.ParseJoin(Xxxx.ParseDbTable(m.GetCtx()), daoHandler))
			daoHandler.AfterField = append(daoHandler.AfterField, v) */
			case `id`:
				m = m.Fields(daoHandler.DbTable + `.` + daoThis.PrimaryKey() + ` AS ` + v)
			case `label`:
				m = m.Fields(daoHandler.DbTable + `.` + daoThis.Columns().MenuName + ` AS ` + v)
			case Scene.Columns().SceneName: //因前端页面已用该字段名显示，故不存在时改成`sceneName`（控制器也要改）。同时下面Fields方法改成m = m.Fields(tableScene + `.` + Scene.Columns().Xxxx + ` AS ` + v)
				tableScene := Scene.ParseDbTable(m.GetCtx())
				m = m.Fields(tableScene + `.` + v)
				m = m.Handler(daoThis.ParseJoin(tableScene, daoHandler))
			case `pMenuName`:
				tableP := `p_` + daoHandler.DbTable
				m = m.Fields(tableP + `.` + daoThis.Columns().MenuName + ` AS ` + v)
				m = m.Handler(daoThis.ParseJoin(tableP, daoHandler))
			case `tree`:
				m = m.Fields(daoHandler.DbTable + `.` + daoThis.PrimaryKey())
				m = m.Fields(daoHandler.DbTable + `.` + daoThis.Columns().Pid)
				m = m.Handler(daoThis.ParseOrder([]string{`tree`}, daoHandler))
			case `showMenu`: //前端显示菜单需要以下字段，且title需要转换
				m = m.Fields(daoHandler.DbTable + `.` + daoThis.Columns().MenuName)
				m = m.Fields(daoHandler.DbTable + `.` + daoThis.Columns().MenuIcon)
				m = m.Fields(daoHandler.DbTable + `.` + daoThis.Columns().MenuUrl)
				m = m.Fields(daoHandler.DbTable + `.` + daoThis.Columns().ExtraData)
				// m = m.Fields(daoHandler.DbTable + `.` + daoThis.Columns().ExtraData + `->'$.i18n' AS i18n`)	//mysql5.6版本不支持
				// m = m.Fields(gdb.Raw(`JSON_UNQUOTE(JSON_EXTRACT(` + daoThis.Columns().ExtraData + `, \`$.i18n\`)) AS i18n`))	//mysql不能直接转成对象返回
				daoHandler.AfterField = append(daoHandler.AfterField, v)
			default:
				if daoThis.ColumnArrG().Contains(v) {
					m = m.Fields(daoHandler.DbTable + `.` + v)
				} else {
					m = m.Fields(v)
				}
			}
		}
		for k, v := range fieldWithParam {
			switch k {
			default:
				daoHandler.AfterFieldWithParam[k] = v
			}
		}
		return m
	}
}

// hook select
func (daoThis *menuDao) HookSelect(daoHandler *daoIndex.DaoHandler) gdb.HookHandler {
	return gdb.HookHandler{
		Select: func(ctx context.Context, in *gdb.HookSelectInput) (result gdb.Result, err error) {
			result, err = in.Next(ctx)
			if err != nil {
				return
			}
			for _, record := range result {
				for _, v := range daoHandler.AfterField {
					switch v {
					case `showMenu`:
						extraDataJson := gjson.New(record[daoThis.Columns().ExtraData])
						record[`i18n`] = extraDataJson.Get(`i18n`)
						if record[`i18n`] == nil {
							record[`i18n`] = gvar.New(map[string]interface{}{`title`: map[string]interface{}{`zh-cn`: record[`menuName`]}})
						}
					default:
						record[v] = gvar.New(nil)
					}
				}
				/* for k, v := range daoHandler.AfterFieldWithParam {
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

// 解析filter
func (daoThis *menuDao) ParseFilter(filter map[string]interface{}, daoHandler *daoIndex.DaoHandler) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for k, v := range filter {
			switch k {
			case `excId`, `excIdArr`:
				if gvar.New(v).IsSlice() {
					m = m.WhereNotIn(daoHandler.DbTable+`.`+daoThis.PrimaryKey(), v)
				} else {
					m = m.WhereNot(daoHandler.DbTable+`.`+daoThis.PrimaryKey(), v)
				}
			case `id`, `idArr`:
				m = m.Where(daoHandler.DbTable+`.`+daoThis.PrimaryKey(), v)
			case `label`:
				m = m.WhereLike(daoHandler.DbTable+`.`+daoThis.Columns().MenuName, `%`+gconv.String(v)+`%`)
			case daoThis.Columns().MenuName:
				m = m.WhereLike(daoHandler.DbTable+`.`+k, `%`+gconv.String(v)+`%`)
			case `timeRangeStart`:
				m = m.WhereGTE(daoHandler.DbTable+`.`+daoThis.Columns().CreatedAt, v)
			case `timeRangeEnd`:
				m = m.WhereLTE(daoHandler.DbTable+`.`+daoThis.Columns().CreatedAt, v)
			case `selfMenu`: //获取当前登录身份可用的菜单。参数：map[string]interface{}{`sceneCode`: `场景标识`, `sceneId`: 场景id, `loginId`: 登录身份id}
				val := gconv.Map(v)
				m = m.Where(daoHandler.DbTable+`.`+daoThis.Columns().SceneId, val[`sceneId`])
				m = m.Where(daoHandler.DbTable+`.`+daoThis.Columns().IsStop, 0)
				switch gconv.String(val[`sceneCode`]) {
				case `platform`:
					if gconv.Uint(val[`loginId`]) == g.Cfg().MustGet(m.GetCtx(), `superPlatformAdminId`).Uint() { //平台超级管理员，不再需要其它条件
						continue
					}
					tableRole := Role.ParseDbTable(m.GetCtx())
					tableRoleRelToMenu := RoleRelToMenu.ParseDbTable(m.GetCtx())
					m = m.Where(tableRole+`.`+Role.Columns().IsStop, 0)
					m = m.Handler(daoThis.ParseJoin(tableRoleRelToMenu, daoHandler))
					m = m.Handler(daoThis.ParseJoin(tableRole, daoHandler))

					tableRoleRelOfPlatformAdmin := RoleRelOfPlatformAdmin.ParseDbTable(m.GetCtx())
					m = m.Where(tableRoleRelOfPlatformAdmin+`.`+RoleRelOfPlatformAdmin.Columns().AdminId, val[`loginId`])
					m = m.Handler(daoThis.ParseJoin(tableRoleRelOfPlatformAdmin, daoHandler))
				default:
					m = m.Where(`1 = 0`)
				}
			default:
				if daoThis.ColumnArrG().Contains(k) {
					m = m.Where(daoHandler.DbTable+`.`+k, v)
				} else {
					m = m.Where(k, v)
				}
			}
		}
		return m
	}
}

// 解析group
func (daoThis *menuDao) ParseGroup(group []string, daoHandler *daoIndex.DaoHandler) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for _, v := range group {
			switch v {
			case `id`:
				m = m.Group(daoHandler.DbTable + `.` + daoThis.PrimaryKey())
			default:
				if daoThis.ColumnArrG().Contains(v) {
					m = m.Group(daoHandler.DbTable + `.` + v)
				} else {
					m = m.Group(v)
				}
			}
		}
		return m
	}
}

// 解析order
func (daoThis *menuDao) ParseOrder(order []string, daoHandler *daoIndex.DaoHandler) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for _, v := range order {
			v = gstr.Trim(v)
			k := gstr.Split(v, ` `)[0]
			switch k {
			case `id`:
				m = m.Order(daoHandler.DbTable + `.` + gstr.Replace(v, k, daoThis.PrimaryKey(), 1))
			case `tree`:
				m = m.OrderAsc(daoHandler.DbTable + `.` + daoThis.Columns().Pid)
				m = m.OrderAsc(daoHandler.DbTable + `.` + daoThis.Columns().Sort)
				m = m.OrderAsc(daoHandler.DbTable + `.` + daoThis.PrimaryKey())
			case daoThis.Columns().Level:
				m = m.Order(daoHandler.DbTable + `.` + v)
				m = m.OrderDesc(daoHandler.DbTable + `.` + daoThis.PrimaryKey())
			case daoThis.Columns().Sort:
				m = m.Order(daoHandler.DbTable + `.` + v)
				m = m.OrderDesc(daoHandler.DbTable + `.` + daoThis.PrimaryKey())
			default:
				if daoThis.ColumnArrG().Contains(k) {
					m = m.Order(daoHandler.DbTable + `.` + v)
				} else {
					m = m.Order(v)
				}
			}
		}
		return m
	}
}

// 解析join
func (daoThis *menuDao) ParseJoin(joinTable string, daoHandler *daoIndex.DaoHandler) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		if garray.NewStrArrayFrom(daoHandler.JoinTableArr).Contains(joinTable) {
			return m
		}
		daoHandler.JoinTableArr = append(daoHandler.JoinTableArr, joinTable)
		switch joinTable {
		/* case Xxxx.ParseDbTable(m.GetCtx()):
		m = m.LeftJoin(joinTable, joinTable+`.`+Xxxx.Columns().XxxxId+` = `+daoHandler.DbTable+`.`+daoThis.PrimaryKey())
		// m = m.LeftJoin(Xxxx.ParseDbTable(m.GetCtx())+` AS `+joinTable, joinTable+`.`+Xxxx.Columns().XxxxId+` = `+daoHandler.DbTable+`.`+daoThis.PrimaryKey()) */
		case Scene.ParseDbTable(m.GetCtx()):
			m = m.LeftJoin(joinTable, joinTable+`.`+Scene.PrimaryKey()+` = `+daoHandler.DbTable+`.`+daoThis.Columns().SceneId)
		case `p_` + daoHandler.DbTable:
			m = m.LeftJoin(daoHandler.DbTable+` AS `+joinTable, joinTable+`.`+daoThis.PrimaryKey()+` = `+daoHandler.DbTable+`.`+daoThis.Columns().Pid)
		case Role.ParseDbTable(m.GetCtx()):
			m = m.LeftJoin(joinTable, joinTable+`.`+Role.PrimaryKey()+` = `+RoleRelToMenu.ParseDbTable(m.GetCtx())+`.`+RoleRelToMenu.Columns().RoleId)
		case RoleRelOfPlatformAdmin.ParseDbTable(m.GetCtx()):
			m = m.LeftJoin(joinTable, joinTable+`.`+RoleRelOfPlatformAdmin.Columns().RoleId+` = `+RoleRelToMenu.ParseDbTable(m.GetCtx())+`.`+RoleRelToMenu.Columns().RoleId)
		/* case RoleRelToMenu.ParseDbTable(m.GetCtx()):
		m = m.LeftJoin(joinCode+` AS `+joinCode, joinCode+`.`+RoleRelToMenu.Columns().MenuId+` = `+daoHandler.DbTable+`.`+daoThis.PrimaryKey()) */
		default:
			m = m.LeftJoin(joinTable, joinTable+`.`+daoThis.PrimaryKey()+` = `+daoHandler.DbTable+`.`+daoThis.PrimaryKey())
		}
		return m
	}
}

// Fill with you ideas below.

// 修改pid时，更新所有子孙级的idPath和level
func (daoThis *menuDao) UpdateChildIdPathAndLevel(ctx context.Context, newIdPath string, oldIdPath string, newLevel uint, oldLevel uint) {
	data := g.Map{
		daoThis.Columns().IdPath: gdb.Raw(`REPLACE(` + daoThis.Columns().IdPath + `, '` + oldIdPath + `', '` + newIdPath + `')`),
		daoThis.Columns().Level:  gdb.Raw(daoThis.Columns().Level + ` + ` + gconv.String(newLevel-oldLevel)),
	}
	if newLevel < oldLevel {
		data[daoThis.Columns().Level] = gdb.Raw(daoThis.Columns().Level + ` - ` + gconv.String(oldLevel-newLevel))
	}
	daoThis.ParseDbCtx(ctx).WhereLike(daoThis.Columns().IdPath, oldIdPath+`-%`).Data(data).Update()
}
