// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package auth

import (
	daoIndex "api/internal/dao"
	"api/internal/dao/auth/internal"
	"context"
	"database/sql"
	"database/sql/driver"
	"reflect"
	"sync"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/text/gstr"
)

// internalRoleRelToActionDao is internal type for wrapping internal DAO implements.
type internalRoleRelToActionDao = *internal.RoleRelToActionDao

// roleRelToActionDao is the data access object for table auth_role_rel_to_action.
// You can define custom methods on it to extend its functionality as you wish.
type roleRelToActionDao struct {
	internalRoleRelToActionDao
}

var (
	// RoleRelToAction is globally public accessible object for table auth_role_rel_to_action operations.
	RoleRelToAction = roleRelToActionDao{
		internal.NewRoleRelToActionDao(),
	}
)

// 获取daoModel
func (daoThis *roleRelToActionDao) CtxDaoModel(ctx context.Context, dbOpt ...any) *daoIndex.DaoModel {
	return daoIndex.NewDaoModel(ctx, daoThis, dbOpt...)
}

// 解析分库
func (daoThis *roleRelToActionDao) ParseDbGroup(ctx context.Context, dbGroupOpt ...any) string {
	group := daoThis.Group()
	// 分库逻辑
	/* if len(dbGroupOpt) > 0 {
	} */
	return group
}

// 解析分表
func (daoThis *roleRelToActionDao) ParseDbTable(ctx context.Context, dbTableOpt ...any) string {
	table := daoThis.Table()
	// 分表逻辑
	/* if len(dbTableOpt) > 0 {
	} */
	return table
}

// 解析Id（未使用代码自动生成，且id字段不在第1个位置时，需手动修改）
func (daoThis *roleRelToActionDao) ParseId(daoModel *daoIndex.DaoModel) string {
	return daoModel.DbTable + `.` + reflect.ValueOf(daoThis.Columns()).Field(0).String()
}

// 解析Label（未使用代码自动生成，且id字段不在第2个位置时，需手动修改）
func (daoThis *roleRelToActionDao) ParseLabel(daoModel *daoIndex.DaoModel) string {
	return daoModel.DbTable + `.` + reflect.ValueOf(daoThis.Columns()).Field(1).String()
}

// 解析filter
func (daoThis *roleRelToActionDao) ParseFilter(filter map[string]any, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for k, v := range filter {
			switch k {
			/* case `xxxx`:
			tableXxxx := Xxxx.ParseDbTable(m.GetCtx())
			m = m.Where(tableXxxx+`.`+k, v)
			m = m.Handler(daoThis.ParseJoin(tableXxxx, daoModel)) */
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
func (daoThis *roleRelToActionDao) ParseField(field []string, fieldWithParam map[string]any, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for _, v := range field {
			switch v {
			/* case `xxxx`:
			tableXxxx := Xxxx.ParseDbTable(m.GetCtx())
			m = m.Fields(tableXxxx + `.` + v)
			m = m.Handler(daoThis.ParseJoin(tableXxxx, daoModel))
			daoModel.AfterField[v] = struct{}{} */
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
				daoModel.AfterField[k] = v
			}
		}
		if len(daoModel.AfterField) > 0 {
			m = m.Hook(daoThis.HookSelect(daoModel))
		}
		return m
	}
}

// 处理afterField
func (daoThis *roleRelToActionDao) HandleAfterField(ctx context.Context, record gdb.Record, daoModel *daoIndex.DaoModel) {
	for k, v := range daoModel.AfterField {
		switch k {
		default:
			if v == struct{}{} {
				record[k] = gvar.New(nil)
			} else {
				record[k] = gvar.New(v)
			}
		}
	}
}

// hook select
func (daoThis *roleRelToActionDao) HookSelect(daoModel *daoIndex.DaoModel) gdb.HookHandler {
	return gdb.HookHandler{
		Select: func(ctx context.Context, in *gdb.HookSelectInput) (result gdb.Result, err error) {
			result, err = in.Next(ctx)
			if err != nil || len(result) == 0 {
				return
			}

			var wg sync.WaitGroup
			wg.Add(len(result))
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
func (daoThis *roleRelToActionDao) ParseInsert(insert map[string]any, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		insertData := map[string]any{}
		for k, v := range insert {
			switch k {
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
func (daoThis *roleRelToActionDao) HookInsert(daoModel *daoIndex.DaoModel) gdb.HookHandler {
	return gdb.HookHandler{
		Insert: func(ctx context.Context, in *gdb.HookInsertInput) (result sql.Result, err error) {
			result, err = in.Next(ctx)
			if err != nil {
				return
			}
			// id, _ := result.LastInsertId()

			/* for k, v := range daoModel.AfterInsert {
				switch k {
				case `xxxx`:
					daoModel.CloneNew().FilterPri(id).HookUpdate(g.Map{k: v}).Update()
				}
			} */
			return
		},
	}
}

// 解析update
func (daoThis *roleRelToActionDao) ParseUpdate(update map[string]any, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		updateData := map[string]any{}
		for k, v := range update {
			switch k {
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
func (daoThis *roleRelToActionDao) HookUpdate(daoModel *daoIndex.DaoModel) gdb.HookHandler {
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
func (daoThis *roleRelToActionDao) HookDelete(daoModel *daoIndex.DaoModel) gdb.HookHandler {
	return gdb.HookHandler{
		Delete: func(ctx context.Context, in *gdb.HookDeleteInput) (result sql.Result, err error) { //有软删除字段时需改成Update事件
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
func (daoThis *roleRelToActionDao) ParseGroup(group []string, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for _, v := range group {
			switch v {
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
func (daoThis *roleRelToActionDao) ParseOrder(order []string, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for _, v := range order {
			v = gstr.Trim(v)
			kArr := gstr.Split(v, `,`)
			k := gstr.Split(kArr[0], ` `)[0]
			switch k {
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
func (daoThis *roleRelToActionDao) ParseJoin(joinTable string, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		if _, ok := daoModel.JoinTableMap[joinTable]; ok {
			return m
		}
		daoModel.JoinTableMap[joinTable] = struct{}{}
		switch joinTable {
		/* case Xxxx.ParseDbTable(m.GetCtx()):
		m = m.LeftJoin(joinTable, joinTable+`.`+Xxxx.Columns().XxxxId+` = `+daoModel.DbTable+`.`+daoThis.Columns().XxxxId)
		// m = m.LeftJoin(Xxxx.ParseDbTable(m.GetCtx())+` AS `+joinTable, joinTable+`.`+Xxxx.Columns().XxxxId+` = `+daoModel.DbTable+`.`+daoThis.Columns().XxxxId) */
		}
		return m
	}
}

// Fill with you ideas below.
