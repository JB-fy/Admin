// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	daoIndex "api/internal/dao"
	"api/internal/dao/platform/internal"
	"context"
	"database/sql"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

// internalConfigDao is internal type for wrapping internal DAO implements.
type internalConfigDao = *internal.ConfigDao

// configDao is the data access object for table platform_config.
// You can define custom methods on it to extend its functionality as you wish.
type configDao struct {
	internalConfigDao
}

var (
	// Config is globally public accessible object for table platform_config operations.
	Config = configDao{
		internal.NewConfigDao(),
	}
)

// 获取daoModel
func (daoThis *configDao) CtxDaoModel(ctx context.Context, dbOpt ...map[string]interface{}) *daoIndex.DaoModel {
	return daoIndex.NewDaoModel(ctx, daoThis, dbOpt...)
}

// 解析分库
func (daoThis *configDao) ParseDbGroup(ctx context.Context, dbGroupOpt ...map[string]interface{}) string {
	group := daoThis.Group()
	// 分库逻辑
	/* if len(dbGroupOpt) > 0 {
	} */
	return group
}

// 解析分表
func (daoThis *configDao) ParseDbTable(ctx context.Context, dbTableOpt ...map[string]interface{}) string {
	table := daoThis.Table()
	// 分表逻辑
	/* if len(dbTableOpt) > 0 {
	} */
	return table
}

// 解析filter
func (daoThis *configDao) ParseFilter(filter map[string]interface{}, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for k, v := range filter {
			switch k {
			/* case `xxxx`:
			tableXxxx := Xxxx.ParseDbTable(m.GetCtx())
			m = m.Where(tableXxxx+`.`+k, v)
			m = m.Handler(daoThis.ParseJoin(tableXxxx, daoModel)) */
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
func (daoThis *configDao) ParseField(field []string, fieldWithParam map[string]interface{}, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for _, v := range field {
			switch v {
			/* case `xxxx`:
			tableXxxx := Xxxx.ParseDbTable(m.GetCtx())
			m = m.Fields(tableXxxx + `.` + v)
			m = m.Handler(daoThis.ParseJoin(tableXxxx, daoModel))
			daoModel.AfterField.Add(v) */
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
func (daoThis *configDao) HookSelect(daoModel *daoIndex.DaoModel) gdb.HookHandler {
	return gdb.HookHandler{
		Select: func(ctx context.Context, in *gdb.HookSelectInput) (result gdb.Result, err error) {
			result, err = in.Next(ctx)
			if err != nil {
				return
			}
			for _, record := range result {
				for _, v := range daoModel.AfterField.Slice() {
					switch v {
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
func (daoThis *configDao) ParseInsert(insert map[string]interface{}, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		insertData := map[string]interface{}{}
		for k, v := range insert {
			switch k {
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
func (daoThis *configDao) HookInsert(daoModel *daoIndex.DaoModel) gdb.HookHandler {
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
					daoModel.CloneNew().Filter(daoThis.PrimaryKey(), id).HookUpdate(g.Map{k: v}).Update()
				}
			} */
			return
		},
	}
}

// 解析update
func (daoThis *configDao) ParseUpdate(update map[string]interface{}, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		updateData := map[string]interface{}{}
		for k, v := range update {
			switch k {
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
func (daoThis *configDao) HookUpdate(daoModel *daoIndex.DaoModel) gdb.HookHandler {
	return gdb.HookHandler{
		Update: func(ctx context.Context, in *gdb.HookUpdateInput) (result sql.Result, err error) {
			result, err = in.Next(ctx)
			if err != nil {
				return
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
func (daoThis *configDao) HookDelete(daoModel *daoIndex.DaoModel) gdb.HookHandler {
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
func (daoThis *configDao) ParseGroup(group []string, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for _, v := range group {
			switch v {
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
func (daoThis *configDao) ParseOrder(order []string, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for _, v := range order {
			v = gstr.Trim(v)
			kArr := gstr.Split(v, `,`)
			k := gstr.Split(kArr[0], ` `)[0]
			switch k {
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
func (daoThis *configDao) ParseJoin(joinTable string, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		if daoModel.JoinTableSet.Contains(joinTable) {
			return m
		}
		daoModel.JoinTableSet.Add(joinTable)
		switch joinTable {
		/* case Xxxx.ParseDbTable(m.GetCtx()):
		m = m.LeftJoin(joinTable, joinTable+`.`+Xxxx.Columns().XxxxId+` = `+daoModel.DbTable+`.`+daoThis.PrimaryKey())
		// m = m.LeftJoin(Xxxx.ParseDbTable(m.GetCtx())+` AS `+joinTable, joinTable+`.`+Xxxx.Columns().XxxxId+` = `+daoModel.DbTable+`.`+daoThis.PrimaryKey()) */
		default:
			m = m.LeftJoin(joinTable, joinTable+`.`+daoThis.PrimaryKey()+` = `+daoModel.DbTable+`.`+daoThis.PrimaryKey())
		}
		return m
	}
}

// Fill with you ideas below.

// 获取配置
func (daoThis *configDao) Get(ctx context.Context, configKeyArr []string) (config gdb.Record, err error) {
	return daoThis.CtxDaoModel(ctx).Filter(daoThis.Columns().ConfigKey, configKeyArr).PluckStr(daoThis.Columns().ConfigValue, daoThis.Columns().ConfigKey)
}

// 保存配置
func (daoThis *configDao) Save(ctx context.Context, config map[string]interface{}) (err error) {
	daoModelThis := daoThis.CtxDaoModel(ctx)
	err = daoModelThis.Transaction(func(ctx context.Context, tx gdb.TX) (err error) {
		for k, v := range config {
			_, err = tx.Model(daoModelThis.DbTable).Data(g.Map{
				daoThis.Columns().ConfigKey:   k,
				daoThis.Columns().ConfigValue: v,
			}).Save()
			if err != nil {
				return
			}
		}
		return
	})
	return
}
