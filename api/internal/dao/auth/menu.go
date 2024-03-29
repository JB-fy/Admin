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

// 获取daoModel
func (daoThis *menuDao) CtxDaoModel(ctx context.Context, dbOpt ...map[string]interface{}) *daoIndex.DaoModel {
	return daoIndex.NewDaoModel(ctx, daoThis, dbOpt...)
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

// 解析filter
func (daoThis *menuDao) ParseFilter(filter map[string]interface{}, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
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
				m = m.WhereLike(daoModel.DbTable+`.`+daoThis.Columns().MenuName, `%`+gconv.String(v)+`%`)
			case daoThis.Columns().MenuName:
				m = m.WhereLike(daoModel.DbTable+`.`+k, `%`+gconv.String(v)+`%`)
			case `pIdPathOfOld`: //父级IdPath（旧）
				m = m.WhereLike(daoModel.DbTable+`.`+daoThis.Columns().IdPath, gconv.String(v)+`-%`)
			case `timeRangeStart`:
				m = m.WhereGTE(daoModel.DbTable+`.`+daoThis.Columns().CreatedAt, v)
			case `timeRangeEnd`:
				m = m.WhereLTE(daoModel.DbTable+`.`+daoThis.Columns().CreatedAt, v)
			case `selfMenu`: //获取当前登录身份可用的菜单。参数：map[string]interface{}{`sceneCode`: `场景标识`, `sceneId`: 场景id, `loginId`: 登录身份id}
				val := gconv.Map(v)
				m = m.Where(daoModel.DbTable+`.`+daoThis.Columns().SceneId, val[`sceneId`])
				m = m.Where(daoModel.DbTable+`.`+daoThis.Columns().IsStop, 0)
				switch gconv.String(val[`sceneCode`]) {
				case `platform`:
					if gconv.Uint(val[`loginId`]) == g.Cfg().MustGet(m.GetCtx(), `superPlatformAdminId`).Uint() { //平台超级管理员，不再需要其它条件
						continue
					}
					tableRole := Role.ParseDbTable(m.GetCtx())
					tableRoleRelToMenu := RoleRelToMenu.ParseDbTable(m.GetCtx())
					m = m.Where(tableRole+`.`+Role.Columns().IsStop, 0)
					m = m.Handler(daoThis.ParseJoin(tableRoleRelToMenu, daoModel))
					m = m.Handler(daoThis.ParseJoin(tableRole, daoModel))

					tableRoleRelOfPlatformAdmin := RoleRelOfPlatformAdmin.ParseDbTable(m.GetCtx())
					m = m.Where(tableRoleRelOfPlatformAdmin+`.`+RoleRelOfPlatformAdmin.Columns().AdminId, val[`loginId`])
					m = m.Handler(daoThis.ParseJoin(tableRoleRelOfPlatformAdmin, daoModel))
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
func (daoThis *menuDao) ParseField(field []string, fieldWithParam map[string]interface{}, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
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
				m = m.Fields(daoModel.DbTable + `.` + daoThis.Columns().MenuName + ` AS ` + v)
			case Scene.Columns().SceneName:
				tableScene := Scene.ParseDbTable(m.GetCtx())
				m = m.Fields(tableScene + `.` + v)
				m = m.Handler(daoThis.ParseJoin(tableScene, daoModel))
			case `pMenuName`:
				tableP := `p_` + daoModel.DbTable
				m = m.Fields(tableP + `.` + daoThis.Columns().MenuName + ` AS ` + v)
				m = m.Handler(daoThis.ParseJoin(tableP, daoModel))
			case `tree`:
				m = m.Fields(daoModel.DbTable + `.` + daoThis.PrimaryKey())
				m = m.Fields(daoModel.DbTable + `.` + daoThis.Columns().Pid)
				m = m.Handler(daoThis.ParseOrder([]string{`tree`}, daoModel))
			case `showMenu`: //前端显示菜单需要以下字段，且title需要转换
				m = m.Fields(daoModel.DbTable + `.` + daoThis.Columns().MenuName)
				m = m.Fields(daoModel.DbTable + `.` + daoThis.Columns().MenuIcon)
				m = m.Fields(daoModel.DbTable + `.` + daoThis.Columns().MenuUrl)
				m = m.Fields(daoModel.DbTable + `.` + daoThis.Columns().ExtraData)
				// m = m.Fields(daoModel.DbTable + `.` + daoThis.Columns().ExtraData + `->'$.i18n' AS i18n`)	//mysql5.6版本不支持
				// m = m.Fields(gdb.Raw(`JSON_UNQUOTE(JSON_EXTRACT(` + daoThis.Columns().ExtraData + `, \`$.i18n\`)) AS i18n`))	//mysql不能直接转成对象返回
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
func (daoThis *menuDao) HookSelect(daoModel *daoIndex.DaoModel) gdb.HookHandler {
	return gdb.HookHandler{
		Select: func(ctx context.Context, in *gdb.HookSelectInput) (result gdb.Result, err error) {
			result, err = in.Next(ctx)
			if err != nil {
				return
			}
			for _, record := range result {
				for _, v := range daoModel.AfterField.Slice() {
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
func (daoThis *menuDao) ParseInsert(insert map[string]interface{}, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		if _, ok := insert[daoThis.Columns().Pid]; !ok {
			insert[daoThis.Columns().Pid] = 0
		}
		insertData := map[string]interface{}{}
		for k, v := range insert {
			switch k {
			case daoThis.Columns().Pid:
				insertData[k] = v
				if gconv.Uint(v) > 0 {
					pInfo, _ := daoThis.CtxDaoModel(m.GetCtx()).Filter(daoThis.PrimaryKey(), v).One()
					daoModel.AfterInsert[`pIdPath`] = pInfo[daoThis.Columns().IdPath].String()
					daoModel.AfterInsert[`pLevel`] = pInfo[daoThis.Columns().Level].Uint()
				} else {
					daoModel.AfterInsert[`pIdPath`] = `0`
					daoModel.AfterInsert[`pLevel`] = 0
				}
			case daoThis.Columns().ExtraData:
				insertData[k] = v
				if gconv.String(v) == `` {
					insertData[k] = nil
				}
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
func (daoThis *menuDao) HookInsert(daoModel *daoIndex.DaoModel) gdb.HookHandler {
	return gdb.HookHandler{
		Insert: func(ctx context.Context, in *gdb.HookInsertInput) (result sql.Result, err error) {
			result, err = in.Next(ctx)
			if err != nil {
				return
			}
			id, _ := result.LastInsertId()

			updateSelfData := map[string]interface{}{}
			for k, v := range daoModel.AfterInsert {
				switch k {
				case `pIdPath`:
					updateSelfData[daoThis.Columns().IdPath] = gconv.String(v) + `-` + gconv.String(id)
				case `pLevel`:
					updateSelfData[daoThis.Columns().Level] = gconv.Uint(v) + 1
				}
			}
			if len(updateSelfData) > 0 {
				daoModel.CloneNew().Filter(daoThis.PrimaryKey(), id).HookUpdate(updateSelfData).Update()
			}
			return
		},
	}
}

// 解析update
func (daoThis *menuDao) ParseUpdate(update map[string]interface{}, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		updateData := map[string]interface{}{}
		for k, v := range update {
			switch k {
			case daoThis.Columns().Pid:
				updateData[daoModel.DbTable+`.`+k] = v
				pIdPath := `0`
				var pLevel uint = 0
				if gconv.Uint(v) > 0 {
					pInfo, _ := daoThis.CtxDaoModel(m.GetCtx()).Filter(daoThis.PrimaryKey(), v).One()
					pIdPath = pInfo[daoThis.Columns().IdPath].String()
					pLevel = pInfo[daoThis.Columns().Level].Uint()
				}
				updateData[daoModel.DbTable+`.`+daoThis.Columns().IdPath] = gdb.Raw(`CONCAT('` + pIdPath + `-', ` + daoThis.PrimaryKey() + `)`)
				updateData[daoModel.DbTable+`.`+daoThis.Columns().Level] = pLevel + 1
				//更新所有子孙级的idPath和level
				updateChildIdPathAndLevelList := []map[string]interface{}{}
				oldList, _ := daoThis.CtxDaoModel(m.GetCtx()).Filter(daoThis.PrimaryKey(), daoModel.IdArr).All()
				for _, oldInfo := range oldList {
					if gconv.Uint(v) != oldInfo[daoThis.Columns().Pid].Uint() {
						updateChildIdPathAndLevelList = append(updateChildIdPathAndLevelList, map[string]interface{}{
							`pIdPathOfOld`: oldInfo[daoThis.Columns().IdPath],
							`pIdPathOfNew`: pIdPath + `-` + oldInfo[daoThis.PrimaryKey()].String(),
							`pLevelOfOld`:  oldInfo[daoThis.Columns().Level],
							`pLevelOfNew`:  pLevel + 1,
						})
					}
				}
				if len(updateChildIdPathAndLevelList) > 0 {
					daoModel.AfterUpdate[`updateChildIdPathAndLevelList`] = updateChildIdPathAndLevelList
				}
			case `childIdPath`: //更新所有子孙级的idPath。参数：map[string]interface{}{`pIdPathOfOld`: `父级IdPath（旧）`, `pIdPathOfNew`: `父级IdPath（新）`}
				val := gconv.Map(v)
				pIdPathOfOld := gconv.String(val[`pIdPathOfOld`])
				pIdPathOfNew := gconv.String(val[`pIdPathOfNew`])
				updateData[daoModel.DbTable+`.`+daoThis.Columns().IdPath] = gdb.Raw(`REPLACE(` + daoThis.Columns().IdPath + `, '` + pIdPathOfOld + `', '` + pIdPathOfNew + `')`)
			case `childLevel`: //更新所有子孙级的level。参数：map[string]interface{}{`pLevelOfOld`: `父级Level（旧）`, `pLevelOfNew`: `父级Level（新）`}
				val := gconv.Map(v)
				pLevelOfOld := gconv.Uint(val[`pLevelOfOld`])
				pLevelOfNew := gconv.Uint(val[`pLevelOfNew`])
				updateData[daoModel.DbTable+`.`+daoThis.Columns().Level] = gdb.Raw(daoModel.DbTable + `.` + daoThis.Columns().Level + ` + ` + gconv.String(pLevelOfNew-pLevelOfOld))
				if pLevelOfNew < pLevelOfOld {
					updateData[daoModel.DbTable+`.`+daoThis.Columns().Level] = gdb.Raw(daoModel.DbTable + `.` + daoThis.Columns().Level + ` - ` + gconv.String(pLevelOfOld-pLevelOfNew))
				}
			case daoThis.Columns().ExtraData:
				updateData[daoModel.DbTable+`.`+k] = gvar.New(v)
				if gconv.String(v) == `` {
					updateData[daoModel.DbTable+`.`+k] = nil
				}
			default:
				if daoThis.ColumnArr().Contains(k) {
					updateData[daoModel.DbTable+`.`+k] = gvar.New(v) //因下面bug处理方式，json类型字段传参必须是gvar变量，否则不会自动生成json格式
				}
			}
		}
		//m = m.Data(updateData) //字段被解析成`table.xxxx`，正确的应该是`table`.`xxxx`
		//解决字段被解析成`table.xxxx`的BUG
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
func (daoThis *menuDao) HookUpdate(daoModel *daoIndex.DaoModel) gdb.HookHandler {
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

			for k, v := range daoModel.AfterUpdate {
				switch k {
				case `updateChildIdPathAndLevelList`: //修改pid时，更新所有子孙级的idPath和level。参数：[]map[string]interface{}{`pIdPathOfOld`: `父级IdPath（旧）`, `pIdPathOfNew`: `父级IdPath（新）`, `pLevelOfOld`: `父级Level（旧）`, `pLevelOfNew`: `父级Level（新）`}
					val := v.([]map[string]interface{})
					for _, v1 := range val {
						daoModel.CloneNew().Filter(`pIdPathOfOld`, v1[`pIdPathOfOld`]).HookUpdate(g.Map{
							`childIdPath`: g.Map{
								`pIdPathOfOld`: v1[`pIdPathOfOld`],
								`pIdPathOfNew`: v1[`pIdPathOfNew`],
							},
							`childLevel`: g.Map{
								`pLevelOfOld`: v1[`pLevelOfOld`],
								`pLevelOfNew`: v1[`pLevelOfNew`],
							},
						}).Update()
					}
				}
			}
			return
		},
	}
}

// hook delete
func (daoThis *menuDao) HookDelete(daoModel *daoIndex.DaoModel) gdb.HookHandler {
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

// 解析group
func (daoThis *menuDao) ParseGroup(group []string, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
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
func (daoThis *menuDao) ParseOrder(order []string, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for _, v := range order {
			v = gstr.Trim(v)
			kArr := gstr.Split(v, `,`)
			k := gstr.Split(kArr[0], ` `)[0]
			switch k {
			case `id`:
				m = m.Order(daoModel.DbTable + `.` + gstr.Replace(v, k, daoThis.PrimaryKey(), 1))
			case `tree`:
				m = m.OrderAsc(daoModel.DbTable + `.` + daoThis.Columns().Pid)
				m = m.OrderAsc(daoModel.DbTable + `.` + daoThis.Columns().Sort)
				m = m.OrderAsc(daoModel.DbTable + `.` + daoThis.PrimaryKey())
			case daoThis.Columns().Level:
				m = m.Order(daoModel.DbTable + `.` + v)
				m = m.OrderDesc(daoModel.DbTable + `.` + daoThis.PrimaryKey())
			case daoThis.Columns().Sort:
				m = m.Order(daoModel.DbTable + `.` + v)
				m = m.OrderDesc(daoModel.DbTable + `.` + daoThis.PrimaryKey())
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
func (daoThis *menuDao) ParseJoin(joinTable string, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
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
		case `p_` + daoModel.DbTable:
			m = m.LeftJoin(daoModel.DbTable+` AS `+joinTable, joinTable+`.`+daoThis.PrimaryKey()+` = `+daoModel.DbTable+`.`+daoThis.Columns().Pid)
		case Role.ParseDbTable(m.GetCtx()):
			m = m.LeftJoin(joinTable, joinTable+`.`+Role.PrimaryKey()+` = `+RoleRelToMenu.ParseDbTable(m.GetCtx())+`.`+RoleRelToMenu.Columns().RoleId)
		case RoleRelOfPlatformAdmin.ParseDbTable(m.GetCtx()):
			m = m.LeftJoin(joinTable, joinTable+`.`+RoleRelOfPlatformAdmin.Columns().RoleId+` = `+RoleRelToMenu.ParseDbTable(m.GetCtx())+`.`+RoleRelToMenu.Columns().RoleId)
		/* case RoleRelToMenu.ParseDbTable(m.GetCtx()):
		m = m.LeftJoin(joinCode+` AS `+joinCode, joinCode+`.`+RoleRelToMenu.Columns().MenuId+` = `+daoModel.DbTable+`.`+daoThis.PrimaryKey()) */
		default:
			m = m.LeftJoin(joinTable, joinTable+`.`+daoThis.PrimaryKey()+` = `+daoModel.DbTable+`.`+daoThis.PrimaryKey())
		}
		return m
	}
}

// Fill with you ideas below.
