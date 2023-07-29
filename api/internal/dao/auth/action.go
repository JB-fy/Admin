// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"api/internal/dao/auth/internal"
	"api/internal/utils"
	"context"
	"database/sql"
	"strings"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gregex"
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

// 解析分库
func (daoThis *actionDao) ParseDbGroup(dbGroupSeldata map[string]interface{}) string {
	group := daoThis.Group()
	/* if len(dbGroupSeldata) > 0 { //分库逻辑
	} */
	return group
}

// 解析分表
func (daoThis *actionDao) ParseDbTable(dbTableSelData map[string]interface{}) string {
	table := daoThis.Table()
	/* if len(dbTableSelData) > 0 { //分表逻辑
	} */
	return table
}

// 解析分库分表（对外暴露使用）
func (daoThis *actionDao) ParseDbCtx(ctx context.Context, dbSelDataList ...map[string]interface{}) *gdb.Model {
	switch len(dbSelDataList) {
	case 1:
		return g.DB(daoThis.ParseDbGroup(dbSelDataList[0])).Model(daoThis.Table()).Safe().Ctx(ctx)
	case 2:
		return g.DB(daoThis.ParseDbGroup(dbSelDataList[0])).Model(daoThis.ParseDbTable(dbSelDataList[1])).Safe().Ctx(ctx)
	default:
		return daoThis.Ctx(ctx)
	}
}

// 解析insert
func (daoThis *actionDao) ParseInsert(insert map[string]interface{}, fill ...bool) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		insertData := map[string]interface{}{}
		hookData := map[string]interface{}{}
		for k, v := range insert {
			switch k {
			case `id`:
				insertData[daoThis.PrimaryKey()] = v
			case `sceneIdArr`:
				hookData[k] = v
			/*--------ParseInsert自动代码生成锚点（不允许修改和删除，否则将不能自动生成代码）--------*/
			default:
				//数据库不存在的字段过滤掉，未传值默认true
				if (len(fill) == 0 || fill[0]) && !daoThis.ColumnArrG().Contains(k) {
					continue
				}
				insertData[k] = v
			}
		}
		m = m.Data(insertData).Hook(daoThis.HookInsert(hookData))
		return m
	}
}

// hook insert
func (daoThis *actionDao) HookInsert(data map[string]interface{}) gdb.HookHandler {
	return gdb.HookHandler{
		Insert: func(ctx context.Context, in *gdb.HookInsertInput) (result sql.Result, err error) {
			result, err = in.Next(ctx)
			if err != nil {
				match, _ := gregex.MatchString(`1062.*Duplicate.*\.([^']*)'`, err.Error())
				if len(match) > 0 {
					err = utils.NewErrorCode(ctx, 29991062, ``, map[string]interface{}{`errField`: match[1]})
				}
				return
			}
			id, _ := result.LastInsertId()

			for k, v := range data {
				switch k {
				case `sceneIdArr`:
					daoThis.SaveRelScene(ctx, gconv.SliceInt(v), int(id))
				}
			}
			return
		},
	}
}

// 解析update
func (daoThis *actionDao) ParseUpdate(update map[string]interface{}, fill ...bool) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		updateData := map[string]interface{}{}
		for k, v := range update {
			switch k {
			case `id`:
				updateData[daoThis.Table()+`.`+daoThis.PrimaryKey()] = v
			/*--------ParseUpdate自动代码生成锚点（不允许修改和删除，否则将不能自动生成代码）--------*/
			default:
				//数据库不存在的字段过滤掉，未传值默认true
				if (len(fill) == 0 || fill[0]) && !daoThis.ColumnArrG().Contains(k) {
					continue
				}
				updateData[daoThis.Table()+`.`+k] = gvar.New(v) //因下面bug处理方式，json类型字段传参必须是gvar变量，否则不会自动生成json格式
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
		data := []interface{}{strings.Join(fieldArr, `,`)}
		data = append(data, valueArr...)
		m = m.Data(data...)
		return m
	}
}

// hook update
func (daoThis *actionDao) HookUpdate(data map[string]interface{}, idArr ...int) gdb.HookHandler {
	return gdb.HookHandler{
		Update: func(ctx context.Context, in *gdb.HookUpdateInput) (result sql.Result, err error) {
			/* //不能这样拿idArr，联表时会有bug
			var idArr []*gvar.Var
			if len(data) > 0 {
				idArr, _ = daoThis.ParseDbCtx(ctx).Where(in.Condition, in.Args[len(in.Args)-gstr.Count(in.Condition, `?`):]...).Array(daoThis.PrimaryKey())
			} */
			result, err = in.Next(ctx)
			if err != nil {
				match, _ := gregex.MatchString(`1062.*Duplicate.*\.([^']*)'`, err.Error())
				if len(match) > 0 {
					err = utils.NewErrorCode(ctx, 29991062, ``, map[string]interface{}{`errField`: match[1]})
				}
				return
			}
			// row, _ := result.RowsAffected()

			for k, v := range data {
				switch k {
				case `sceneIdArr`:
					relIdArr := gconv.SliceInt(v)
					for _, id := range idArr {
						daoThis.SaveRelScene(ctx, relIdArr, id)
					}
				}
			}

			/* if row == 0 {
				// err = utils.NewErrorCode(ctx, 99999999, ``)
				return
			} */
			return
		},
	}
}

// hook delete
func (daoThis *actionDao) HookDelete(idArr ...int) gdb.HookHandler {
	return gdb.HookHandler{
		Delete: func(ctx context.Context, in *gdb.HookDeleteInput) (result sql.Result, err error) {
			result, err = in.Next(ctx)
			if err != nil {
				return
			}
			row, _ := result.RowsAffected()
			if row == 0 {
				// err = utils.NewErrorCode(ctx, 99999999, ``)
				return
			}

			ActionRelToScene.ParseDbCtx(ctx).Where(ActionRelToScene.Columns().ActionId, idArr).Delete()
			RoleRelToAction.ParseDbCtx(ctx).Where(RoleRelToAction.Columns().ActionId, idArr).Delete()
			return
		},
	}
}

// 解析field
func (daoThis *actionDao) ParseField(field []string, joinTableArr *[]string) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		afterField := []string{}
		for _, v := range field {
			switch v {
			/* case `xxxx`:
			m = daoThis.ParseJoin(Xxxx.Table(), joinTableArr)(m)
			afterField = append(afterField, v) */
			case `id`:
				m = m.Fields(daoThis.Table() + `.` + daoThis.PrimaryKey() + ` AS ` + v)
			case `label`:
				m = m.Fields(daoThis.Table() + `.` + daoThis.Columns().ActionName + ` AS ` + v)
			case `sceneIdArr`:
				//需要id字段
				m = m.Fields(daoThis.Table() + `.` + daoThis.PrimaryKey())
				afterField = append(afterField, v)
			/*--------ParseField自动代码生成锚点（不允许修改和删除，否则将不能自动生成代码）--------*/
			default:
				if daoThis.ColumnArrG().Contains(v) {
					m = m.Fields(daoThis.Table() + `.` + v)
				} else {
					m = m.Fields(v)
				}
			}
		}
		if len(afterField) > 0 {
			m = m.Hook(daoThis.HookSelect(afterField))
		}
		return m
	}
}

// hook select
func (daoThis *actionDao) HookSelect(afterField []string) gdb.HookHandler {
	return gdb.HookHandler{
		Select: func(ctx context.Context, in *gdb.HookSelectInput) (result gdb.Result, err error) {
			result, err = in.Next(ctx)
			if err != nil {
				return
			}
			for index, record := range result {
				for _, v := range afterField {
					switch v {
					/* case `xxxx`:
					record[v] = gvar.New(``) */
					case `sceneIdArr`:
						idArr, _ := ActionRelToScene.ParseDbCtx(ctx).Where(daoThis.PrimaryKey(), record[daoThis.PrimaryKey()]).Array(ActionRelToScene.Columns().SceneId)
						record[v] = gvar.New(idArr)
					/*--------HookSelect自动代码生成锚点（不允许修改和删除，否则将不能自动生成代码）--------*/
					default:
					}
				}
				result[index] = record
			}
			return
		},
	}
}

// 解析filter
func (daoThis *actionDao) ParseFilter(filter map[string]interface{}, joinTableArr *[]string) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for k, v := range filter {
			switch k {
			case `id`, `idArr`:
				m = m.Where(daoThis.Table()+`.`+daoThis.PrimaryKey(), v)
			case `excId`, `excIdArr`:
				val := gconv.SliceInt(v)
				switch len(val) {
				case 0: //gconv.SliceInt会把0转换成[]int{}，故不能用转换后的val。必须用原始数据v
					m = m.WhereNot(daoThis.Table()+`.`+daoThis.PrimaryKey(), v)
				case 1:
					m = m.WhereNot(daoThis.Table()+`.`+daoThis.PrimaryKey(), val[0])
				default:
					m = m.WhereNotIn(daoThis.Table()+`.`+daoThis.PrimaryKey(), v)
				}
			case `startTime`:
				m = m.WhereGTE(daoThis.Table()+`.`+daoThis.Columns().CreatedAt, v)
			case `endTime`:
				m = m.WhereLTE(daoThis.Table()+`.`+daoThis.Columns().CreatedAt, v)
			case `label`:
				m = m.WhereLike(daoThis.Table()+`.`+daoThis.Columns().ActionName, `%`+gconv.String(v)+`%`)
			case `sceneId`:
				m = m.Where(ActionRelToScene.Table()+`.`+k, v)

				m = daoThis.ParseJoin(ActionRelToScene.Table(), joinTableArr)(m)
			case `selfAction`: //获取当前登录身份可用的操作。参数：map[string]interface{}{`sceneCode`: `场景标识`, `sceneId`=>场景id, `loginId`: 登录身份id}
				val := v.(map[string]interface{})

				m = m.Where(daoThis.Table()+`.`+daoThis.Columns().IsStop, 0)
				m = m.Where(ActionRelToScene.Table()+`.`+ActionRelToScene.Columns().SceneId, val[`sceneId`])
				m = daoThis.ParseJoin(ActionRelToScene.Table(), joinTableArr)(m)
				switch val[`sceneCode`].(string) {
				case `platform`:
					if gconv.Int(val[`loginId`]) == g.Cfg().MustGet(m.GetCtx(), `superPlatformAdminId`).Int() { //平台超级管理员，不再需要其他条件
						continue
					}
					m = m.Where(Role.Table()+`.`+Role.Columns().IsStop, 0)
					m = m.Where(RoleRelOfPlatformAdmin.Table()+`.`+RoleRelOfPlatformAdmin.Columns().AdminId, val[`loginId`])

					m = daoThis.ParseJoin(RoleRelToAction.Table(), joinTableArr)(m)
					m = daoThis.ParseJoin(Role.Table(), joinTableArr)(m)
					m = daoThis.ParseJoin(RoleRelOfPlatformAdmin.Table(), joinTableArr)(m)
				}
				m = m.Group(daoThis.Table() + `.` + daoThis.PrimaryKey())
			/*--------ParseFilter自动代码生成锚点（不允许修改和删除，否则将不能自动生成代码）--------*/
			default:
				kArr := strings.Split(k, ` `) //支持`id > ?`等k
				if !daoThis.ColumnArrG().Contains(kArr[0]) {
					m = m.Where(k, v)
					continue
				}
				if len(kArr) == 1 {
					if gstr.SubStr(gstr.CaseCamel(kArr[0]), -4) == `Name` {
						m = m.WhereLike(daoThis.Table()+`.`+k, `%`+gconv.String(v)+`%`)
						continue
					}
				}
				m = m.Where(daoThis.Table()+`.`+k, v)
			}
		}
		return m
	}
}

// 解析group
func (daoThis *actionDao) ParseGroup(group []string, joinTableArr *[]string) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for _, v := range group {
			switch v {
			case `id`:
				m = m.Group(daoThis.Table() + `.` + daoThis.PrimaryKey())
			default:
				if daoThis.ColumnArrG().Contains(v) {
					m = m.Group(daoThis.Table() + `.` + v)
				} else {
					m = m.Group(v)
				}
			}
		}
		return m
	}
}

// 解析order
func (daoThis *actionDao) ParseOrder(order []string, joinTableArr *[]string) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for _, v := range order {
			kArr := strings.Split(v, ` `)
			if len(kArr) == 1 {
				kArr = append(kArr, `ASC`)
			}
			switch kArr[0] {
			case `id`:
				m = m.Order(daoThis.Table()+`.`+daoThis.PrimaryKey(), kArr[1])
			default:
				if daoThis.ColumnArrG().Contains(kArr[0]) {
					m = m.Order(daoThis.Table()+`.`+kArr[0], kArr[1])
				} else {
					m = m.Order(kArr[0], kArr[1])
				}
			}
		}
		return m
	}
}

// 解析join
func (daoThis *actionDao) ParseJoin(joinCode string, joinTableArr *[]string) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		switch joinCode {
		/* case Xxxx.Table():
		relTable := Xxxx.Table()
		if !garray.NewStrArrayFrom(*joinTableArr).Contains(relTable) {
			*joinTableArr = append(*joinTableArr, relTable)
			m = m.LeftJoin(relTable, relTable+`.`+daoThis.PrimaryKey()+` = `+daoThis.Table()+`.`+daoThis.PrimaryKey())
		} */
		case ActionRelToScene.Table():
			relTable := ActionRelToScene.Table()
			if !garray.NewStrArrayFrom(*joinTableArr).Contains(relTable) {
				*joinTableArr = append(*joinTableArr, relTable)
				m = m.LeftJoin(relTable, relTable+`.`+daoThis.PrimaryKey()+` = `+daoThis.Table()+`.`+daoThis.PrimaryKey())
			}
		case RoleRelToAction.Table():
			relTable := RoleRelToAction.Table()
			if !garray.NewStrArrayFrom(*joinTableArr).Contains(relTable) {
				*joinTableArr = append(*joinTableArr, relTable)
				m = m.LeftJoin(relTable, relTable+`.`+daoThis.PrimaryKey()+` = `+daoThis.Table()+`.`+daoThis.PrimaryKey())
			}
		case Role.Table():
			relTable := Role.Table()
			if !garray.NewStrArrayFrom(*joinTableArr).Contains(relTable) {
				*joinTableArr = append(*joinTableArr, relTable)
				m = m.LeftJoin(relTable, relTable+`.`+Role.PrimaryKey()+` = `+RoleRelToAction.Table()+`.`+Role.PrimaryKey())
			}
		case RoleRelOfPlatformAdmin.Table():
			relTable := RoleRelOfPlatformAdmin.Table()
			if !garray.NewStrArrayFrom(*joinTableArr).Contains(relTable) {
				*joinTableArr = append(*joinTableArr, relTable)
				m = m.LeftJoin(relTable, relTable+`.`+Role.PrimaryKey()+` = `+RoleRelToAction.Table()+`.`+Role.PrimaryKey())
			}
		}
		return m
	}
}

// Fill with you ideas below.

// 保存关联场景
func (daoThis *actionDao) SaveRelScene(ctx context.Context, relIdArr []int, id int) {
	relKey := ActionRelToScene.Columns().SceneId
	priKey := daoThis.PrimaryKey()
	relIdArrOfOldTmp, _ := ActionRelToScene.ParseDbCtx(ctx).Where(priKey, id).Array(relKey)
	relIdArrOfOld := gconv.SliceInt(relIdArrOfOldTmp)

	/**----新增关联场景 开始----**/
	insertRelIdArr := gset.NewIntSetFrom(relIdArr).Diff(gset.NewIntSetFrom(relIdArrOfOld)).Slice()
	if len(insertRelIdArr) > 0 {
		insertList := []map[string]interface{}{}
		for _, v := range insertRelIdArr {
			insertList = append(insertList, map[string]interface{}{
				priKey: id,
				relKey: v,
			})
		}
		ActionRelToScene.ParseDbCtx(ctx).Data(insertList).Insert()
	}
	/**----新增关联场景 结束----**/

	/**----删除关联场景 开始----**/
	deleteRelIdArr := gset.NewIntSetFrom(relIdArrOfOld).Diff(gset.NewIntSetFrom(relIdArr)).Slice()
	if len(deleteRelIdArr) > 0 {
		ActionRelToScene.ParseDbCtx(ctx).Where(priKey, id).Where(relKey, deleteRelIdArr).Delete()
	}
	/**----删除关联场景 结束----**/
}

// 自动生成操作
func (daoThis *actionDao) MyGenAction(ctx context.Context, sceneId int, actionCode string, actionName string) {
	idVar, _ := daoThis.ParseDbCtx(ctx).Where(daoThis.Columns().ActionCode, actionCode).Value(daoThis.PrimaryKey())
	id := idVar.Int64()
	if id == 0 {
		id, _ = daoThis.ParseDbCtx(ctx).Data(map[string]interface{}{
			daoThis.Columns().ActionCode: actionCode,
			daoThis.Columns().ActionName: actionName,
		}).InsertAndGetId()
	} else {
		daoThis.ParseDbCtx(ctx).Where(daoThis.PrimaryKey(), id).Data(daoThis.Columns().ActionName, actionName).Update()
	}
	ActionRelToScene.ParseDbCtx(ctx).Data(map[string]interface{}{
		ActionRelToScene.Columns().ActionId: id,
		ActionRelToScene.Columns().SceneId:  sceneId,
	}).Save()
}
