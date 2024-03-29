// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	daoIndex "api/internal/dao"
	daoAuth "api/internal/dao/auth"
	"api/internal/dao/platform/internal"
	"context"
	"database/sql"

	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
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
func (daoThis *adminDao) CtxDaoModel(ctx context.Context, dbOpt ...map[string]interface{}) *daoIndex.DaoModel {
	return daoIndex.NewDaoModel(ctx, daoThis, dbOpt...)
}

// 解析分库
func (daoThis *adminDao) ParseDbGroup(ctx context.Context, dbGroupOpt ...map[string]interface{}) string {
	group := daoThis.Group()
	// 分库逻辑
	/* if len(dbGroupOpt) > 0 {
	} */
	return group
}

// 解析分表
func (daoThis *adminDao) ParseDbTable(ctx context.Context, dbTableOpt ...map[string]interface{}) string {
	table := daoThis.Table()
	// 分表逻辑
	/* if len(dbTableOpt) > 0 {
	} */
	return table
}

// 解析filter
func (daoThis *adminDao) ParseFilter(filter map[string]interface{}, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
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
				m = m.Where(m.Builder().WhereLike(daoModel.DbTable+`.`+daoThis.Columns().Phone, `%`+gconv.String(v)+`%`).WhereOrLike(daoModel.DbTable+`.`+daoThis.Columns().Account, `%`+gconv.String(v)+`%`).WhereOrLike(daoModel.DbTable+`.`+daoThis.Columns().Nickname, `%`+gconv.String(v)+`%`))
			case `timeRangeStart`:
				m = m.WhereGTE(daoModel.DbTable+`.`+daoThis.Columns().CreatedAt, v)
			case `timeRangeEnd`:
				m = m.WhereLTE(daoModel.DbTable+`.`+daoThis.Columns().CreatedAt, v)
			case daoAuth.RoleRelOfPlatformAdmin.Columns().RoleId:
				tableRoleRelOfPlatformAdmin := daoAuth.RoleRelOfPlatformAdmin.ParseDbTable(m.GetCtx())
				m = m.Where(tableRoleRelOfPlatformAdmin+`.`+k, v)
				m = m.Handler(daoThis.ParseJoin(tableRoleRelOfPlatformAdmin, daoModel))
			case `loginName`:
				if g.Validator().Rules(`required|phone`).Data(v).Run(m.GetCtx()) == nil {
					m = m.Where(daoModel.DbTable+`.`+daoThis.Columns().Phone, v)
				} else {
					m = m.Where(daoModel.DbTable+`.`+daoThis.Columns().Account, v)
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
func (daoThis *adminDao) ParseField(field []string, fieldWithParam map[string]interface{}, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
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
				m = m.Fields(`IF(IFNULL(` + daoModel.DbTable + `.` + daoThis.Columns().Phone + `, '') != '', ` + daoModel.DbTable + `.` + daoThis.Columns().Phone + `, IF(IFNULL(` + daoModel.DbTable + `.` + daoThis.Columns().Account + `, '') != '', ` + daoModel.DbTable + `.` + daoThis.Columns().Account + `, ` + daoModel.DbTable + `.` + daoThis.Columns().Nickname + `)) AS ` + v)
			case `roleIdArr`:
				m = m.Fields(daoModel.DbTable + `.` + daoThis.PrimaryKey())
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
func (daoThis *adminDao) HookSelect(daoModel *daoIndex.DaoModel) gdb.HookHandler {
	return gdb.HookHandler{
		Select: func(ctx context.Context, in *gdb.HookSelectInput) (result gdb.Result, err error) {
			result, err = in.Next(ctx)
			if err != nil {
				return
			}
			for _, record := range result {
				for _, v := range daoModel.AfterField.Slice() {
					switch v {
					case `roleIdArr`:
						idArr, _ := daoAuth.RoleRelOfPlatformAdmin.CtxDaoModel(ctx).Filter(daoThis.PrimaryKey(), record[daoThis.PrimaryKey()]).Array(daoAuth.RoleRelOfPlatformAdmin.Columns().RoleId)
						record[v] = gvar.New(idArr)
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
func (daoThis *adminDao) ParseInsert(insert map[string]interface{}, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		insertData := map[string]interface{}{}
		for k, v := range insert {
			switch k {
			case daoThis.Columns().Phone:
				insertData[k] = v
				if gconv.String(v) == `` {
					insertData[k] = nil
				}
			case daoThis.Columns().Account:
				insertData[k] = v
				if gconv.String(v) == `` {
					insertData[k] = nil
				}
			case daoThis.Columns().Password:
				password := gconv.String(v)
				if len(password) != 32 {
					password = gmd5.MustEncrypt(password)
				}
				salt := grand.S(8)
				insertData[daoThis.Columns().Salt] = salt
				password = gmd5.MustEncrypt(password + salt)
				insertData[k] = password
			case `roleIdArr`:
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
				case `roleIdArr`:
					daoThis.SaveRelRole(ctx, gconv.SliceUint(v), uint(id))
				}
			}
			return
		},
	}
}

// 解析update
func (daoThis *adminDao) ParseUpdate(update map[string]interface{}, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		updateData := map[string]interface{}{}
		for k, v := range update {
			switch k {
			case daoThis.Columns().Phone:
				updateData[daoModel.DbTable+`.`+k] = v
				if gconv.String(v) == `` {
					updateData[daoModel.DbTable+`.`+k] = nil
				}
			case daoThis.Columns().Account:
				updateData[daoModel.DbTable+`.`+k] = v
				if gconv.String(v) == `` {
					updateData[daoModel.DbTable+`.`+k] = nil
				}
			case daoThis.Columns().Password:
				password := gconv.String(v)
				if len(password) != 32 {
					password = gmd5.MustEncrypt(password)
				}
				salt := grand.S(8)
				updateData[daoModel.DbTable+`.`+daoThis.Columns().Salt] = salt
				password = gmd5.MustEncrypt(password + salt)
				updateData[daoModel.DbTable+`.`+k] = password
			case `roleIdArr`:
				daoModel.AfterUpdate[k] = v
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
func (daoThis *adminDao) HookUpdate(daoModel *daoIndex.DaoModel) gdb.HookHandler {
	return gdb.HookHandler{
		Update: func(ctx context.Context, in *gdb.HookUpdateInput) (result sql.Result, err error) {
			result, err = in.Next(ctx)
			if err != nil {
				return
			}

			for k, v := range daoModel.AfterUpdate {
				switch k {
				case `roleIdArr`:
					relIdArr := gconv.SliceUint(v)
					for _, id := range daoModel.IdArr {
						daoThis.SaveRelRole(ctx, relIdArr, id)
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
func (daoThis *adminDao) HookDelete(daoModel *daoIndex.DaoModel) gdb.HookHandler {
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
func (daoThis *adminDao) ParseOrder(order []string, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
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
func (daoThis *adminDao) ParseJoin(joinTable string, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		if daoModel.JoinTableSet.Contains(joinTable) {
			return m
		}
		daoModel.JoinTableSet.Add(joinTable)
		switch joinTable {
		/* case Xxxx.ParseDbTable(m.GetCtx()):
		m = m.LeftJoin(joinTable, joinTable+`.`+Xxxx.Columns().XxxxId+` = `+daoModel.DbTable+`.`+daoThis.PrimaryKey())
		// m = m.LeftJoin(Xxxx.ParseDbTable(m.GetCtx())+` AS `+joinTable, joinTable+`.`+Xxxx.Columns().XxxxId+` = `+daoModel.DbTable+`.`+daoThis.PrimaryKey()) */
		/* case daoAuth.RoleRelOfPlatformAdmin.ParseDbTable(m.GetCtx()):
		m = m.LeftJoin(joinTable, joinTable+`.`+daoAuth.RoleRelOfPlatformAdmin.Columns().AdminId+` = `+daoModel.DbTable+`.`+daoThis.PrimaryKey()) */
		default:
			m = m.LeftJoin(joinTable, joinTable+`.`+daoThis.PrimaryKey()+` = `+daoModel.DbTable+`.`+daoThis.PrimaryKey())
		}
		return m
	}
}

// Fill with you ideas below.

// 保存关联角色
func (daoThis *adminDao) SaveRelRole(ctx context.Context, relIdArr []uint, id uint) {
	relDao := daoAuth.RoleRelOfPlatformAdmin
	priKey := relDao.Columns().AdminId
	relKey := relDao.Columns().RoleId
	relIdArrOfOld, _ := relDao.CtxDaoModel(ctx).Filter(priKey, id).ArrayUint(relKey)

	/**----新增关联 开始----**/
	insertRelIdArr := gset.NewFrom(relIdArr).Diff(gset.NewFrom(relIdArrOfOld)).Slice()
	if len(insertRelIdArr) > 0 {
		insertList := []map[string]interface{}{}
		for _, v := range insertRelIdArr {
			insertList = append(insertList, map[string]interface{}{
				priKey: id,
				relKey: v,
			})
		}
		relDao.CtxDaoModel(ctx).Data(insertList).Insert()
	}
	/**----新增关联 结束----**/

	/**----删除关联 开始----**/
	deleteRelIdArr := gset.NewFrom(relIdArrOfOld).Diff(gset.NewFrom(relIdArr)).Slice()
	if len(deleteRelIdArr) > 0 {
		relDao.CtxDaoModel(ctx).Filters(g.Map{
			priKey: id,
			relKey: deleteRelIdArr,
		}).Delete()
	}
	/**----删除关联 结束----**/
}
