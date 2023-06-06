// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"api/internal/model/dao/platform/internal"
	"context"
	"strings"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

// internalServerDao is internal type for wrapping internal DAO implements.
type internalServerDao = *internal.ServerDao

// serverDao is the data access object for table platform_server.
// You can define custom methods on it to extend its functionality as you wish.
type serverDao struct {
	internalServerDao
}

var (
	// Server is globally public accessible object for table platform_server operations.
	Server = serverDao{
		internal.NewServerDao(),
	}
)

// 解析insert
func (daoThis *serverDao) ParseInsert(insert []map[string]interface{}, fill ...bool) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		insertData := make([]map[string]interface{}, len(insert))
		for index, item := range insert {
			insertData[index] = map[string]interface{}{}
			for k, v := range item {
				switch k {
				case "id":
					insertData[index][daoThis.PrimaryKey()] = v
				default:
					//数据库不存在的字段过滤掉，未传值默认true
					if (len(fill) == 0 || fill[0]) && !daoThis.ColumnArrG().Contains(k) {
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
func (daoThis *serverDao) ParseUpdate(update map[string]interface{}, fill ...bool) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		updateData := map[string]interface{}{}
		for k, v := range update {
			switch k {
			case "id":
				updateData[daoThis.Table()+"."+daoThis.PrimaryKey()] = v
			default:
				//数据库不存在的字段过滤掉，未传值默认true
				if (len(fill) == 0 || fill[0]) && !daoThis.ColumnArrG().Contains(k) {
					continue
				}
				updateData[daoThis.Table()+"."+k] = v
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
func (daoThis *serverDao) ParseField(field []string, joinTableArr *[]string) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		afterField := []string{}
		for _, v := range field {
			switch v {
			/* case "xxxx":
			m = daoThis.ParseJoin("xxxx", joinTableArr)(m)
			afterField = append(afterField, v) */
			case "id":
				m = m.Fields(daoThis.Table() + "." + daoThis.PrimaryKey() + " AS " + v)
			default:
				if daoThis.ColumnArrG().Contains(v) {
					m = m.Fields(daoThis.Table() + "." + v)
				} else {
					m = m.Fields(v)
				}
			}
		}
		if len(afterField) > 0 {
			m = m.Hook(daoThis.AfterField(afterField))
		}
		return m
	}
}

// 解析filter
func (daoThis *serverDao) ParseFilter(filter map[string]interface{}, joinTableArr *[]string) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for k, v := range filter {
			kArr := strings.Split(k, " ") //为支持"id > ?"的key
			switch kArr[0] {
			case "id":
				val := gvar.New(v)
				if val.IsSlice() && len(val.Slice()) == 1 {
					m = m.Where(daoThis.Table()+"."+daoThis.PrimaryKey(), val.Slice()[0])
				} else {
					m = m.Where(daoThis.Table()+"."+daoThis.PrimaryKey(), v)
				}
			case "excId":
				val := gvar.New(v)
				if val.IsSlice() {
					if len(val.Slice()) == 1 {
						m = m.WhereNot(daoThis.Table()+"."+daoThis.PrimaryKey(), val.Slice()[0])
					} else {
						m = m.WhereNotIn(daoThis.Table()+"."+daoThis.PrimaryKey(), v)
					}
				} else {
					m = m.WhereNot(daoThis.Table()+"."+daoThis.PrimaryKey(), v)
				}
			case "startTime":
				m = m.WhereGTE(daoThis.Table()+".createTime", v)
			case "endTime":
				m = m.WhereLTE(daoThis.Table()+".createTime", v)
			case "keyword":
				keywordField := strings.ReplaceAll(daoThis.PrimaryKey(), "Id", "Name")
				m = m.WhereLike(daoThis.Table()+"."+keywordField, gconv.String(v))
			default:
				if daoThis.ColumnArrG().Contains(kArr[0]) {
					if gstr.ToLower(gstr.SubStr(kArr[0], -2)) == "id" {
						val := gvar.New(v)
						if val.IsSlice() && len(val.Slice()) == 1 {
							m = m.Where(daoThis.Table()+"."+k, val.Slice()[0])
						} else {
							m = m.Where(daoThis.Table()+"."+k, v)
						}
					} else {
						m = m.Where(daoThis.Table()+"."+k, v)
					}
				} else {
					m = m.Where(k, v)
				}
			}
		}
		return m
	}
}

// 解析group
func (daoThis *serverDao) ParseGroup(group []string, joinTableArr *[]string) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for _, v := range group {
			switch v {
			case "id":
				m = m.Group(daoThis.Table() + "." + daoThis.PrimaryKey())
			default:
				if daoThis.ColumnArrG().Contains(v) {
					m = m.Group(daoThis.Table() + "." + v)
				} else {
					m = m.Group(v)
				}
			}
		}
		return m
	}
}

// 解析order
func (daoThis *serverDao) ParseOrder(order [][2]string, joinTableArr *[]string) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for _, v := range order {
			switch v[0] {
			case "id":
				m = m.Order(daoThis.Table()+"."+daoThis.PrimaryKey(), v[1])
			default:
				if daoThis.ColumnArrG().Contains(v[0]) {
					m = m.Order(daoThis.Table()+"."+v[0], v[1])
				} else {
					m = m.Order(v[0], v[1])
				}
			}
		}
		return m
	}
}

// 解析join
func (daoThis *serverDao) ParseJoin(joinCode string, joinTableArr *[]string) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		switch joinCode {
		/* case "xxxx":
		xxxxTable := xxxx.Table()
		if !garray.NewStrArrayFrom(*joinTableArr).Contains(xxxxTable) {
			*joinTableArr = append(*joinTableArr, xxxxTable)
			m = m.LeftJoin(xxxxTable, xxxxTable+"."+daoThis.PrimaryKey()+" = "+daoThis.Table()+"."+daoThis.PrimaryKey())
		} */
		}
		return m
	}
}

// 获取数据后，再处理的字段
func (daoThis *serverDao) AfterField(afterField []string) gdb.HookHandler {
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
					}
				}
				result[i] = record
			}
			return
		},
	}
}

// 常用方法（用filter和field查询）
func (daoThis *serverDao) CommonModel(ctx context.Context, filter map[string]interface{}, field []string, joinTableArr ...*[]string) *gdb.Model {
	if len(joinTableArr) == 0 {
		joinTableArr = []*[]string{{}}
	}
	model := daoThis.Ctx(ctx)
	model = model.Handler(daoThis.ParseFilter(filter, joinTableArr[0]))
	if len(field) > 0 {
		model = model.Handler(daoThis.ParseField(field, joinTableArr[0]))
	}
	return model
}

// Fill with you ideas below.
