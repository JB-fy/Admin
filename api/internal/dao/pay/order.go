// =================================================================================
// This file is auto-generated by the GoFrame CLI tool. You may modify it as needed.
// =================================================================================

package pay

import (
	daoIndex "api/internal/dao"
	"api/internal/dao/pay/internal"
	"context"
	"database/sql"
	"database/sql/driver"
	"sync"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

// orderDao is the data access object for the table pay_order.
// You can define custom methods on it to extend its functionality as needed.
type orderDao struct {
	*internal.OrderDao
}

var (
	// Order is a globally accessible object for table pay_order operations.
	Order = orderDao{internal.NewOrderDao()}
)

// 获取daoModel
func (daoThis *orderDao) CtxDaoModel(ctx context.Context, dbOpt ...any) *daoIndex.DaoModel {
	return daoIndex.NewDaoModel(ctx, daoThis, dbOpt...)
}

// 解析分库
func (daoThis *orderDao) ParseDbGroup(ctx context.Context, dbGroupOpt ...any) string {
	group := daoThis.Group()
	// 分库逻辑
	/* if len(dbGroupOpt) > 0 {
	} */
	return group
}

// 解析分表
func (daoThis *orderDao) ParseDbTable(ctx context.Context, dbTableOpt ...any) string {
	table := daoThis.Table()
	// 分表逻辑
	/* if len(dbTableOpt) > 0 {
	} */
	return table
}

// 解析Id（未使用代码自动生成，且id字段不在第1个位置时，需手动修改）
func (daoThis *orderDao) ParseId(daoModel *daoIndex.DaoModel) string {
	return daoModel.DbTable + `.` + daoThis.Columns().OrderId
}

// 解析Label（未使用代码自动生成，且id字段不在第2个位置时，需手动修改）
func (daoThis *orderDao) ParseLabel(daoModel *daoIndex.DaoModel) string {
	return daoModel.DbTable + `.` + daoThis.Columns().OrderNo
}

// 解析filter
func (daoThis *orderDao) ParseFilter(filter map[string]any, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for k, v := range filter {
			switch k {
			/* case `xxxx`:
			tableXxxx := Xxxx.ParseDbTable(m.GetCtx())
			m = m.Where(tableXxxx+`.`+k, v)
			m = m.Handler(daoThis.ParseJoin(tableXxxx, daoModel)) */
			case `id`, `id_arr`:
				m = m.Where(daoModel.DbTable+`.`+daoThis.Columns().OrderId, v)
			case `exc_id`, `exc_id_arr`:
				if gvar.New(v).IsSlice() {
					m = m.WhereNotIn(daoModel.DbTable+`.`+daoThis.Columns().OrderId, v)
				} else {
					m = m.WhereNot(daoModel.DbTable+`.`+daoThis.Columns().OrderId, v)
				}
			case `label`:
				m = m.WhereLike(daoModel.DbTable+`.`+daoThis.Columns().OrderNo, `%`+gconv.String(v)+`%`)
			case OrderRel.Columns().RelOrderType, OrderRel.Columns().RelOrderId, OrderRel.Columns().RelOrderUserId:
				tableOrderRel := OrderRel.ParseDbTable(m.GetCtx())
				m = m.Where(tableOrderRel+`.`+k, v)
				m = m.Handler(daoThis.ParseJoin(tableOrderRel, daoModel))
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
func (daoThis *orderDao) ParseField(field []string, fieldWithParam map[string]any, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
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
			case Pay.Columns().PayName:
				tablePay := Pay.ParseDbTable(m.GetCtx())
				m = m.Fields(tablePay + `.` + v)
				m = m.Handler(daoThis.ParseJoin(tablePay, daoModel))
			case Channel.Columns().ChannelName:
				tableChannel := Channel.ParseDbTable(m.GetCtx())
				m = m.Fields(tableChannel + `.` + v)
				m = m.Handler(daoThis.ParseJoin(tableChannel, daoModel))
			case `order_rel_list`:
				m = m.Fields(daoModel.DbTable + `.` + daoThis.Columns().OrderId)
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
func (daoThis *orderDao) HandleAfterField(ctx context.Context, record gdb.Record, daoModel *daoIndex.DaoModel) {
	for k, v := range daoModel.AfterField {
		switch k {
		case `order_rel_list`:
			orderRelList, _ := OrderRel.CtxDaoModel(ctx).Filter(OrderRel.Columns().OrderId, record[daoThis.Columns().OrderId]). /* OrderAsc(OrderRel.Columns().CreatedAt). */ All() // 有顺序要求时使用，自定义OrderAsc
			record[k] = gvar.New(gjson.MustEncodeString(orderRelList))                                                                                                              //转成json字符串，控制器中list.Structs(&res.List)和info.Struct(&res.Info)才有效
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
func (daoThis *orderDao) HookSelect(daoModel *daoIndex.DaoModel) gdb.HookHandler {
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
func (daoThis *orderDao) ParseInsert(insert map[string]any, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for k, v := range insert {
			switch k {
			case `order_rel_list`:
				daoModel.AfterInsert[k] = v
			default:
				if daoThis.Contains(k) {
					daoModel.SaveData[k] = v
				}
			}
		}
		m = m.Data(daoModel.SaveData)
		if len(daoModel.AfterInsert) > 0 {
			m = m.Hook(daoThis.HookInsert(daoModel))
		}
		return m
	}
}

// hook insert
func (daoThis *orderDao) HookInsert(daoModel *daoIndex.DaoModel) gdb.HookHandler {
	return gdb.HookHandler{
		Insert: func(ctx context.Context, in *gdb.HookInsertInput) (result sql.Result, err error) {
			result, err = in.Next(ctx)
			if err != nil {
				return
			}
			id, _ := result.LastInsertId()

			for k, v := range daoModel.AfterInsert {
				switch k {
				case `order_rel_list`:
					vList := gconv.Maps(v)
					insertList := make([]map[string]any, len(vList))
					for index, item := range vList {
						insertItem := gjson.New(gjson.MustEncodeString(item)).Map()
						insertItem[OrderRel.Columns().OrderId] = id
						insertList[index] = insertItem
					}
					OrderRel.CtxDaoModel(ctx).Data(insertList).Insert()
				}
			}
			return
		},
	}
}

// 解析update
func (daoThis *orderDao) ParseUpdate(update map[string]any, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for k, v := range update {
			switch k {
			case `order_rel_list`:
				daoModel.AfterUpdate[k] = v
			default:
				if daoThis.Contains(k) {
					daoModel.SaveData[k] = v
				}
			}
		}
		m = m.Data(daoModel.SaveData)
		if len(daoModel.AfterUpdate) > 0 {
			m = m.Hook(daoThis.HookUpdate(daoModel))
		}
		return m
	}
}

// hook update
func (daoThis *orderDao) HookUpdate(daoModel *daoIndex.DaoModel) gdb.HookHandler {
	return gdb.HookHandler{
		Update: func(ctx context.Context, in *gdb.HookUpdateInput) (result sql.Result, err error) {
			if len(daoModel.SaveData) == 0 {
				result = driver.RowsAffected(0)
			} else {
				result, err = in.Next(ctx)
				if err != nil {
					return
				}
			}

			for k, v := range daoModel.AfterUpdate {
				switch k {
				case `order_rel_list`:
					valList := gconv.Maps(v)
					daoIndex.SaveListRelManyWithSort(ctx, &OrderRel, OrderRel.Columns().OrderId, gconv.SliceAny(daoModel.IdArr), valList)
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
func (daoThis *orderDao) HookDelete(daoModel *daoIndex.DaoModel) gdb.HookHandler {
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

			// 对并发有要求时，可使用以下代码解决情形1。并发说明请参考：api/internal/dao/auth/scene.go中HookDelete方法内的注释
			OrderRel.CtxDaoModel(ctx).Filter(OrderRel.Columns().OrderId, daoModel.IdArr).Delete()
			return
		},
	}
}

// 解析group
func (daoThis *orderDao) ParseGroup(group []string, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for _, v := range group {
			switch v {
			case `id`:
				m = m.Group(daoModel.DbTable + `.` + daoThis.Columns().OrderId)
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
func (daoThis *orderDao) ParseOrder(order []string, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for _, v := range order {
			v = gstr.Trim(v)
			kArr := gstr.Split(v, `,`)
			k := gstr.Split(kArr[0], ` `)[0]
			switch k {
			case `id`:
				m = m.Order(daoModel.DbTable + `.` + gstr.Replace(v, k, daoThis.Columns().OrderId, 1))
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
func (daoThis *orderDao) ParseJoin(joinTable string, daoModel *daoIndex.DaoModel) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		if _, ok := daoModel.JoinTableMap[joinTable]; ok {
			return m
		}
		daoModel.JoinTableMap[joinTable] = struct{}{}
		switch joinTable {
		/* case Xxxx.ParseDbTable(m.GetCtx()):
		m = m.LeftJoin(joinTable, joinTable+`.`+Xxxx.Columns().XxxxId+` = `+daoModel.DbTable+`.`+daoThis.Columns().XxxxId)
		// m = m.LeftJoin(Xxxx.ParseDbTable(m.GetCtx())+` AS `+joinTable, joinTable+`.`+Xxxx.Columns().XxxxId+` = `+daoModel.DbTable+`.`+daoThis.Columns().XxxxId) */
		case Pay.ParseDbTable(m.GetCtx()):
			m = m.LeftJoin(joinTable, joinTable+`.`+Pay.Columns().PayId+` = `+daoModel.DbTable+`.`+daoThis.Columns().PayId)
		case Channel.ParseDbTable(m.GetCtx()):
			m = m.LeftJoin(joinTable, joinTable+`.`+Channel.Columns().ChannelId+` = `+daoModel.DbTable+`.`+daoThis.Columns().ChannelId)
		case OrderRel.ParseDbTable(m.GetCtx()):
			m = m.LeftJoin(joinTable, joinTable+`.`+OrderRel.Columns().OrderId+` = `+daoModel.DbTable+`.`+daoThis.Columns().OrderId)
		default:
			m = m.LeftJoin(joinTable, joinTable+`.`+daoThis.Columns().OrderId+` = `+daoModel.DbTable+`.`+daoThis.Columns().OrderId)
		}
		return m
	}
}

// Add your custom methods and functionality below.
