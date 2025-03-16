// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package platform

import (
	daoIndex "api/internal/dao"
	daoAuth "api/internal/dao/auth"
	"api/internal/dao/platform/internal"
	"context"
	"database/sql"
	"database/sql/driver"
	"sync"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
)

// internalAdminDao is internal type for wrapping internal DAO implements.
type internalAdminDao = *internal.AdminDao

// adminDao is the data access object for table platform_admin.
// You can define custom methods on it to extend its functionality as you wish.
type adminDao struct {
	internalAdminDao
}

var (
	// Admin is globally public accessible object for table platform_admin operations.
	Admin = adminDao{
		internal.NewAdminDao(),
	}
)

// 获取daoModel
func (daoThis *adminDao) CtxDaoModel(ctx context.Context, dbOpt ...any) *daoIndex.DaoModel {
	return daoIndex.NewDaoModel(ctx, daoThis, dbOpt...)
}

// 解析分库
func (daoThis *adminDao) ParseDbGroup(ctx context.Context, dbGroupOpt ...any) string {
	group := daoThis.Group()
	// 分库逻辑
	/* if len(dbGroupOpt) > 0 {
	} */
	return group
}

// 解析分表
func (daoThis *adminDao) ParseDbTable(ctx context.Context, dbTableOpt ...any) string {
	table := daoThis.Table()
	// 分表逻辑
	/* if len(dbTableOpt) > 0 {
	} */
	return table
}

// 解析Id（未使用代码自动生成，且id字段不在第1个位置时，需手动修改）
func (daoThis *adminDao) ParseId(daoModel *daoIndex.DaoModel) string {
	return daoModel.DbTable + `.` + daoThis.Columns().AdminId
}

// 解析Label（未使用代码自动生成，且id字段不在第2个位置时，需手动修改）
func (daoThis *adminDao) ParseLabel(daoModel *daoIndex.DaoModel) string {
	return `COALESCE(NULLIF(` + daoModel.DbTable + `.` + daoThis.Columns().Phone + `, ''), NULLIF(` + daoModel.DbTable + `.` + daoThis.Columns().Email + `, ''), NULLIF(` + daoModel.DbTable + `.` + daoThis.Columns().Account + `, ''), NULLIF(` + daoModel.DbTable + `.` + daoThis.Columns().Nickname + `, ''))`
}

// 解析filter
func (daoThis *adminDao) ParseFilter(filter map[string]any, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for k, v := range filter {
			switch k {
			/* case `xxxx`:
			tableXxxx := Xxxx.ParseDbTable(m.GetCtx())
			m = m.Where(tableXxxx+`.`+k, v)
			m = m.Handler(daoThis.ParseJoin(tableXxxx, daoModel)) */
			case `id`, `id_arr`:
				m = m.Where(daoModel.DbTable+`.`+daoThis.Columns().AdminId, v)
			case `exc_id`, `exc_id_arr`:
				if gvar.New(v).IsSlice() {
					m = m.WhereNotIn(daoModel.DbTable+`.`+daoThis.Columns().AdminId, v)
				} else {
					m = m.WhereNot(daoModel.DbTable+`.`+daoThis.Columns().AdminId, v)
				}
			case `label`:
				m = m.Where(m.Builder().WhereLike(daoModel.DbTable+`.`+daoThis.Columns().Phone, `%`+gconv.String(v)+`%`).WhereOrLike(daoModel.DbTable+`.`+daoThis.Columns().Email, `%`+gconv.String(v)+`%`).WhereOrLike(daoModel.DbTable+`.`+daoThis.Columns().Account, `%`+gconv.String(v)+`%`).WhereOrLike(daoModel.DbTable+`.`+daoThis.Columns().Nickname, `%`+gconv.String(v)+`%`))
			case daoAuth.RoleRelOfPlatformAdmin.Columns().RoleId:
				tableAuthRoleRelOfPlatformAdmin := daoAuth.RoleRelOfPlatformAdmin.ParseDbTable(m.GetCtx())
				m = m.Where(tableAuthRoleRelOfPlatformAdmin+`.`+k, v)
				m = m.Handler(daoThis.ParseJoin(tableAuthRoleRelOfPlatformAdmin, daoModel))
			case `time_range_start`:
				m = m.WhereGTE(daoModel.DbTable+`.`+daoThis.Columns().CreatedAt, v)
			case `time_range_end`:
				m = m.WhereLTE(daoModel.DbTable+`.`+daoThis.Columns().CreatedAt, v)
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
func (daoThis *adminDao) ParseField(field []string, fieldWithParam map[string]any, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for _, v := range field {
			switch v {
			/* case `xxxx`:
			tableXxxx := Xxxx.ParseDbTable(m.GetCtx())
			m = m.Fields(tableXxxx + `.` + v)
			m = m.Handler(daoThis.ParseJoin(tableXxxx, daoModel))
			daoModel.AfterField[v] = struct{}{} */
			case `id`:
				m = m.Fields(daoThis.ParseId(daoModel) + ` AS ` + v)
			case `label`:
				m = m.Fields(daoThis.ParseLabel(daoModel) + ` AS ` + v)
			case `role_id_arr`:
				m = m.Fields(daoModel.DbTable + `.` + daoThis.Columns().AdminId)
				daoModel.AfterField[v] = struct{}{}
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
func (daoThis *adminDao) HandleAfterField(ctx context.Context, record gdb.Record, daoModel *daoIndex.DaoModel) {
	for k, v := range daoModel.AfterField {
		switch k {
		case `role_id_arr`:
			roleIdArr, _ := daoAuth.RoleRelOfPlatformAdmin.CtxDaoModel(ctx).Filter(daoAuth.RoleRelOfPlatformAdmin.Columns().AdminId, record[daoThis.Columns().AdminId]).Array(daoAuth.RoleRelOfPlatformAdmin.Columns().RoleId)
			record[k] = gvar.New(roleIdArr)
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
func (daoThis *adminDao) HookSelect(daoModel *daoIndex.DaoModel) gdb.HookHandler {
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
func (daoThis *adminDao) ParseInsert(insert map[string]any, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		insertData := map[string]any{}
		for k, v := range insert {
			switch k {
			case daoThis.Columns().Phone:
				if gconv.String(v) == `` {
					v = nil
				}
				insertData[k] = v
			case daoThis.Columns().Email:
				if gconv.String(v) == `` {
					v = nil
				}
				insertData[k] = v
			case daoThis.Columns().Account:
				if gconv.String(v) == `` {
					v = nil
				}
				insertData[k] = v
			case daoThis.Columns().Password:
				password := gconv.String(v)
				if len(password) != 32 {
					password = gmd5.MustEncrypt(password)
				}
				salt := grand.S(8)
				insertData[daoThis.Columns().Salt] = salt
				password = gmd5.MustEncrypt(password + salt)
				insertData[k] = password
			case `role_id_arr`:
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
func (daoThis *adminDao) HookInsert(daoModel *daoIndex.DaoModel) gdb.HookHandler {
	return gdb.HookHandler{
		Insert: func(ctx context.Context, in *gdb.HookInsertInput) (result sql.Result, err error) {
			result, err = in.Next(ctx)
			if err != nil {
				return
			}
			id, _ := result.LastInsertId()

			for k, v := range daoModel.AfterInsert {
				switch k {
				case `role_id_arr`:
					vArr := gconv.SliceAny(v)
					insertList := make([]map[string]any, len(vArr))
					for index, item := range vArr {
						insertList[index] = map[string]any{daoAuth.RoleRelOfPlatformAdmin.Columns().AdminId: id, daoAuth.RoleRelOfPlatformAdmin.Columns().RoleId: item}
					}
					daoAuth.RoleRelOfPlatformAdmin.CtxDaoModel(ctx).Data(insertList).Insert()
				}
			}
			return
		},
	}
}

// 解析update
func (daoThis *adminDao) ParseUpdate(update map[string]any, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		updateData := map[string]any{}
		for k, v := range update {
			switch k {
			case daoThis.Columns().Phone:
				if gconv.String(v) == `` {
					v = nil
				}
				updateData[k] = v
			case daoThis.Columns().Email:
				if gconv.String(v) == `` {
					v = nil
				}
				updateData[k] = v
			case daoThis.Columns().Account:
				if gconv.String(v) == `` {
					v = nil
				}
				updateData[k] = v
			case daoThis.Columns().Password:
				password := gconv.String(v)
				if len(password) != 32 {
					password = gmd5.MustEncrypt(password)
				}
				salt := grand.S(8)
				updateData[daoThis.Columns().Salt] = salt
				password = gmd5.MustEncrypt(password + salt)
				updateData[k] = password
			case `role_id_arr`:
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
func (daoThis *adminDao) HookUpdate(daoModel *daoIndex.DaoModel) gdb.HookHandler {
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
				case `role_id_arr`:
					// daoIndex.SaveArrRelManyWithSort(ctx, &daoAuth.RoleRelOfPlatformAdmin, daoAuth.RoleRelOfPlatformAdmin.Columns().AdminId, daoAuth.RoleRelOfPlatformAdmin.Columns().RoleId, gconv.SliceAny(daoModel.IdArr), gconv.SliceAny(v)) // 有顺序要求时使用，同时注释下面代码
					valArr := gconv.Strings(v)
					for _, id := range daoModel.IdArr {
						daoIndex.SaveArrRelMany(ctx, &daoAuth.RoleRelOfPlatformAdmin, daoAuth.RoleRelOfPlatformAdmin.Columns().AdminId, daoAuth.RoleRelOfPlatformAdmin.Columns().RoleId, id, valArr)
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
						daoModel.CloneNew().FilterPri(id).HookUpdateOne(k, v).Update()
					}
				}
			} */
			return
		},
	}
}

// hook delete
func (daoThis *adminDao) HookDelete(daoModel *daoIndex.DaoModel) gdb.HookHandler {
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

			daoAuth.RoleRelOfPlatformAdmin.CtxDaoModel(ctx).Filter(daoAuth.RoleRelOfPlatformAdmin.Columns().AdminId, daoModel.IdArr).Delete()
			return
		},
	}
}

// 解析group
func (daoThis *adminDao) ParseGroup(group []string, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for _, v := range group {
			switch v {
			case `id`:
				m = m.Group(daoModel.DbTable + `.` + daoThis.Columns().AdminId)
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
func (daoThis *adminDao) ParseOrder(order []string, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for _, v := range order {
			v = gstr.Trim(v)
			kArr := gstr.Split(v, `,`)
			k := gstr.Split(kArr[0], ` `)[0]
			switch k {
			case `id`:
				m = m.Order(daoModel.DbTable + `.` + gstr.Replace(v, k, daoThis.Columns().AdminId, 1))
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
func (daoThis *adminDao) ParseJoin(joinTable string, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		if _, ok := daoModel.JoinTableMap[joinTable]; ok {
			return m
		}
		daoModel.JoinTableMap[joinTable] = struct{}{}
		switch joinTable {
		/* case Xxxx.ParseDbTable(m.GetCtx()):
		m = m.LeftJoin(joinTable, joinTable+`.`+Xxxx.Columns().XxxxId+` = `+daoModel.DbTable+`.`+daoThis.Columns().XxxxId)
		// m = m.LeftJoin(Xxxx.ParseDbTable(m.GetCtx())+` AS `+joinTable, joinTable+`.`+Xxxx.Columns().XxxxId+` = `+daoModel.DbTable+`.`+daoThis.Columns().XxxxId) */
		case daoAuth.RoleRelOfPlatformAdmin.ParseDbTable(m.GetCtx()):
			m = m.LeftJoin(joinTable, joinTable+`.`+daoAuth.RoleRelOfPlatformAdmin.Columns().AdminId+` = `+daoModel.DbTable+`.`+daoThis.Columns().AdminId)
		default:
			m = m.LeftJoin(joinTable, joinTable+`.`+daoThis.Columns().AdminId+` = `+daoModel.DbTable+`.`+daoThis.Columns().AdminId)
		}
		return m
	}
}

// Fill with you ideas below.
