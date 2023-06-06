// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"api/internal/model/dao/auth/internal"
	"context"
	"encoding/json"
	"strconv"
	"strings"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
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

// 解析insert
func (daoMenu *menuDao) ParseInsert(insert []map[string]interface{}, fill ...bool) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		insertData := make([]map[string]interface{}, len(insert))
		for index, item := range insert {
			insertData[index] = map[string]interface{}{}
			for k, v := range item {
				switch k {
				case "id":
					insertData[index][daoMenu.PrimaryKey()] = v
				default:
					//数据库不存在的字段过滤掉，未传值默认true
					if (len(fill) == 0 || fill[0]) && !daoMenu.ColumnArrG().Contains(k) {
						continue
					}
					insertData[index][k] = v
				}
			}
		}
		if len(insertData) == 1 {
			m = m.Data(insertData[0])
		} else {
			m = m.Data(insertData)
		}
		return m
	}
}

// 解析update
func (daoMenu *menuDao) ParseUpdate(update map[string]interface{}, fill ...bool) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		updateData := map[string]interface{}{}
		for k, v := range update {
			switch k {
			case "id":
				updateData[daoMenu.Table()+"."+daoMenu.PrimaryKey()] = v
			case "pidPathOfChild": //更新所有子孙级的pidPath。参数：map[string]string{"newVal": "父级新pidPath", "oldVal":"父级旧pidPath"}
				updateData[daoMenu.Table()+".pidPath"] = gdb.Raw("REPLACE(" + daoMenu.Table() + ".pidPath, '" + v.(map[string]string)["oldVal"] + "', '" + v.(map[string]string)["newVal"] + "')")
			case "levelOfChild": //更新所有子孙级的level。参数：map[string]int{"newVal": 父级新level, "oldVal":父级旧level}
				updateData[daoMenu.Table()+".level"] = gdb.Raw(daoMenu.Table() + ".level + " + strconv.Itoa(v.(map[string]int)["newVal"]-v.(map[string]int)["oldVal"]))
			default:
				//数据库不存在的字段过滤掉，未传值默认true
				if (len(fill) == 0 || fill[0]) && !daoMenu.ColumnArrG().Contains(k) {
					continue
				}
				updateData[daoMenu.Table()+"."+k] = v
			}
		}
		//m = m.Data(updateData) //字段被解析成`table.xxxx`，正确的应该是`table`.`xxxx`
		//解决字段被解析成`table.xxxx`的BUG
		fieldArr := []string{}
		valueArr := []interface{}{}
		for k, v := range updateData {
			fieldArr = append(fieldArr, k+" = ?")
			valueArr = append(valueArr, v)
		}
		data := []interface{}{strings.Join(fieldArr, ",")}
		data = append(data, valueArr...)
		m = m.Data(data...)
		return m
	}
}

// 解析field
func (daoMenu *menuDao) ParseField(field []string, joinTableArr *[]string) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		afterField := []string{}
		for _, v := range field {
			switch v {
			/* case "xxxx":
			m = daoMenu.ParseJoin("xxxx", joinTableArr)(m)
			afterField = append(afterField, v) */
			case "id":
				m = m.Fields(daoMenu.Table() + "." + daoMenu.PrimaryKey() + " AS " + v)
			case "menuTree": //树状需要以下字段和排序方式
				m = m.Fields(daoMenu.Table() + "." + daoMenu.PrimaryKey())
				m = m.Fields(daoMenu.Table() + ".pid")

				m = daoMenu.ParseOrder([][2]string{{"menuTree", ""}}, joinTableArr)(m) //排序方式
			case "showMenu": //前端显示菜单需要以下字段，且title需要转换
				m = m.Fields(daoMenu.Table() + ".menuName")
				m = m.Fields(daoMenu.Table() + ".menuIcon")
				m = m.Fields(daoMenu.Table() + ".menuUrl")
				m = m.Fields(daoMenu.Table() + ".extraData->'$.i18n' AS i18n")
				//m = m.Fields(gdb.Raw("JSON_UNQUOTE(JSON_EXTRACT(extraData, \"$.i18n\")) AS i18n"))//mysql不能直接转成对象返回
				afterField = append(afterField, v)
			case "sceneName":
				m = m.Fields(Scene.Table() + "." + v)
				m = daoMenu.ParseJoin("scene", joinTableArr)(m)
			case "pMenuName":
				m = m.Fields("p_" + daoMenu.Table() + ".menuName AS " + v)
				m = daoMenu.ParseJoin("pMenu", joinTableArr)(m)
			default:
				if daoMenu.ColumnArrG().Contains(v) {
					m = m.Fields(daoMenu.Table() + "." + v)
				} else {
					m = m.Fields(v)
				}
			}
		}
		if len(afterField) > 0 {
			m = m.Hook(daoMenu.AfterField(afterField))
		}
		return m
	}
}

// 解析filter
func (daoMenu *menuDao) ParseFilter(filter map[string]interface{}, joinTableArr *[]string) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for k, v := range filter {
			switch k {
			case "id", "idArr":
				m = m.Where(daoMenu.Table()+"."+daoMenu.PrimaryKey(), v)
			case "excId":
				m = m.WhereNot(daoMenu.Table()+"."+daoMenu.PrimaryKey(), v)
			case "excIdArr":
				m = m.WhereNotIn(daoMenu.Table()+"."+daoMenu.PrimaryKey(), v)
			case "startTime":
				m = m.WhereGTE(daoMenu.Table()+".createTime", v)
			case "endTime":
				m = m.WhereLTE(daoMenu.Table()+".createTime", v)
			case "keyword":
				keywordField := strings.ReplaceAll(daoMenu.PrimaryKey(), "Id", "Name")
				switch v := v.(type) {
				case *string:
					m = m.WhereLike(daoMenu.Table()+"."+keywordField, *v)
				case string:
					m = m.WhereLike(daoMenu.Table()+"."+keywordField, v)
				default:
					m = m.Where(daoMenu.Table()+"."+keywordField, v)
				}
			case "selfMenu": //获取当前登录身份可用的菜单。参数：map[string]interface{}{"sceneCode": "场景标识", "loginId": 登录身份id}
				val := v.(map[string]interface{})
				ctx := m.GetCtx()
				sceneInfo := m.GetCtx().Value("sceneInfo").(gdb.Record)
				sceneId := 0
				if len(sceneInfo) == 0 {
					sceneIdTmp, _ := Scene.Ctx(ctx).Where("sceneCode", val["sceneCode"]).Value("sceneId")
					sceneId = sceneIdTmp.Int()
				} else {
					sceneId = sceneInfo["sceneId"].Int()
				}

				m = m.Where(daoMenu.Table()+".sceneId", sceneId)
				m = m.Where(daoMenu.Table()+".isStop", 0)
				switch val["sceneCode"].(string) {
				case "platformAdmin":
					if gconv.Int(val["loginId"]) == g.Cfg().MustGet(ctx, "superPlatformAdminId").Int() { //平台超级管理员，不再需要其他条件
						return m
					}
					m = m.Where(Role.Table()+".isStop", 0)
					m = m.Where(RoleRelOfPlatformAdmin.Table()+".adminId", val["loginId"])

					m = daoMenu.ParseJoin("roleRelToMenu", joinTableArr)(m)
					m = daoMenu.ParseJoin("role", joinTableArr)(m)
					m = daoMenu.ParseJoin("roleRelOfPlatformAdmin", joinTableArr)(m)
				}
				m = daoMenu.ParseGroup([]string{"id"}, joinTableArr)(m)
			default:
				kArr := strings.Split(k, " ")
				if daoMenu.ColumnArrG().Contains(kArr[0]) {
					m = m.Where(daoMenu.Table()+"."+k, v)
				} else {
					m = m.Where(k, v)
				}
			}
		}
		return m
	}
}

// 解析group
func (daoMenu *menuDao) ParseGroup(group []string, joinTableArr *[]string) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for _, v := range group {
			switch v {
			case "id":
				m = m.Group(daoMenu.Table() + "." + daoMenu.PrimaryKey())
			default:
				if daoMenu.ColumnArrG().Contains(v) {
					m = m.Group(daoMenu.Table() + "." + v)
				} else {
					m = m.Group(v)
				}
			}
		}
		return m
	}
}

// 解析order
func (daoMenu *menuDao) ParseOrder(order [][2]string, joinTableArr *[]string) func(m *gdb.Model) *gdb.Model {
	return func(m *gdb.Model) *gdb.Model {
		for _, v := range order {
			switch v[0] {
			case "id":
				m = m.Order(daoMenu.Table()+"."+daoMenu.PrimaryKey(), v[1])
			case "menuTree":
				m = m.Order(daoMenu.Table()+".pid", "ASC")
				m = m.Order(daoMenu.Table()+".sort", "ASC")
				m = m.Order(daoMenu.Table()+".menuId", "ASC")
			default:
				if daoMenu.ColumnArrG().Contains(v[0]) {
					m = m.Order(daoMenu.Table()+"."+v[0], v[1])
				} else {
					m = m.Order(v[0], v[1])
				}
			}
		}
		return m
	}
}

// 解析join
func (daoMenu *menuDao) ParseJoin(joinCode string, joinTableArr *[]string) func(m *gdb.Model) *gdb.Model {
	return func(m *gdb.Model) *gdb.Model {
		switch joinCode {
		/* case "xxxx":
		xxxxTable := xxxx.Table()
		if !garray.NewStrArrayFrom(*joinTableArr).Contains(xxxxTable) {
			*joinTableArr = append(*joinTableArr, xxxxTable)
			m = m.LeftJoin(xxxxTable, xxxxTable+"."+daoMenu.PrimaryKey()+" = "+daoMenu.Table()+"."+daoMenu.PrimaryKey())
		} */
		case "scene":
			sceneTable := Scene.Table()
			if !garray.NewStrArrayFrom(*joinTableArr).Contains(sceneTable) {
				*joinTableArr = append(*joinTableArr, sceneTable)
				m = m.LeftJoin(sceneTable, sceneTable+"."+Scene.PrimaryKey()+" = "+daoMenu.Table()+"."+Scene.PrimaryKey())
			}
		case "pMenu":
			pMenuTable := "p_" + daoMenu.Table()
			if !garray.NewStrArrayFrom(*joinTableArr).Contains(pMenuTable) {
				*joinTableArr = append(*joinTableArr, pMenuTable)
				m = m.LeftJoin(daoMenu.Table()+" AS "+pMenuTable, pMenuTable+"."+daoMenu.PrimaryKey()+" = "+daoMenu.Table()+".pid")
			}
		case "roleRelToMenu":
			roleRelToMenuTable := RoleRelToMenu.Table()
			if !garray.NewStrArrayFrom(*joinTableArr).Contains(roleRelToMenuTable) {
				*joinTableArr = append(*joinTableArr, roleRelToMenuTable)
				m = m.LeftJoin(roleRelToMenuTable, roleRelToMenuTable+"."+daoMenu.PrimaryKey()+" = "+daoMenu.Table()+"."+daoMenu.PrimaryKey())
			}
		case "role":
			roleTable := Role.Table()
			if !garray.NewStrArrayFrom(*joinTableArr).Contains(roleTable) {
				*joinTableArr = append(*joinTableArr, roleTable)
				roleRelToMenuTable := RoleRelToMenu.Table()
				m = m.LeftJoin(roleTable, roleTable+"."+Role.PrimaryKey()+" = "+roleRelToMenuTable+"."+Role.PrimaryKey())
			}
		case "roleRelOfPlatformAdmin":
			roleRelOfPlatformAdminTable := RoleRelOfPlatformAdmin.Table()
			if !garray.NewStrArrayFrom(*joinTableArr).Contains(roleRelOfPlatformAdminTable) {
				*joinTableArr = append(*joinTableArr, roleRelOfPlatformAdminTable)
				roleRelToMenuTable := RoleRelToMenu.Table()
				m = m.LeftJoin(roleRelOfPlatformAdminTable, roleRelOfPlatformAdminTable+"."+Role.PrimaryKey()+" = "+roleRelToMenuTable+"."+Role.PrimaryKey())
			}
		}
		return m
	}
}

// 获取数据后，再处理的字段
func (daoMenu *menuDao) AfterField(afterField []string) gdb.HookHandler {
	return gdb.HookHandler{
		Select: func(ctx context.Context, in *gdb.HookSelectInput) (result gdb.Result, err error) {
			result, err = in.Next(ctx)
			if err != nil {
				return
			}
			for i, record := range result {
				for _, v := range afterField {
					switch v {
					/* case "xxxx":
					record[v] = gvar.New("") */
					case "showMenu":
						if record["i18n"] == nil {
							record["i18n"] = gvar.New(map[string]interface{}{"title": map[string]interface{}{"zh-cn": record["menuName"]}})
						} else {
							i18n := map[string]interface{}{}
							json.Unmarshal([]byte(record["i18n"].String()), &i18n)
							record["i18n"] = gvar.New(i18n)
						}
					}
				}
				result[i] = record
			}
			return
		},
	}
}

// 详情
func (daoMenu *menuDao) Info(ctx context.Context, filter map[string]interface{}, field []string, order ...[2]string) (info gdb.Record, err error) {
	joinTableArr := []string{}
	model := daoMenu.Ctx(ctx)
	if len(field) > 0 {
		model = model.Handler(daoMenu.ParseField(field, &joinTableArr))
	}
	if len(filter) > 0 {
		model = model.Handler(daoMenu.ParseFilter(filter, &joinTableArr))
	}
	if len(order) > 0 {
		model = model.Handler(daoMenu.ParseOrder(order, &joinTableArr))
	}
	info, err = model.One()
	return
}

// 列表
func (daoMenu *menuDao) List(ctx context.Context, filter map[string]interface{}, field []string, order ...[2]string) (list gdb.Result, err error) {
	joinTableArr := []string{}
	model := daoMenu.Ctx(ctx)
	if len(field) > 0 {
		model = model.Handler(daoMenu.ParseField(field, &joinTableArr))
	}
	if len(filter) > 0 {
		model = model.Handler(daoMenu.ParseFilter(filter, &joinTableArr))
	}
	if len(order) > 0 {
		model = model.Handler(daoMenu.ParseOrder(order, &joinTableArr))
	}
	list, err = model.All()
	return
}

// Fill with you ideas below.
