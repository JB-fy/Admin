// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"
	"reflect"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ChannelDao is the data access object for the table pay_channel.
type ChannelDao struct {
	table     string              // table is the underlying table name of the DAO.
	group     string              // group is the database configuration group name of the current DAO.
	columns   ChannelColumns      // columns contains all the column names of Table for convenient usage.
	handlers  []gdb.ModelHandler  // handlers for customized model modification.
	columnArr []string            // 字段数组
	columnMap map[string]struct{} // 字段map
}

// ChannelColumns defines and stores column names for the table pay_channel.
type ChannelColumns struct {
	CreatedAt   string // 创建时间
	UpdatedAt   string // 更新时间
	IsStop      string // 停用：0否 1是
	ChannelId   string // 通道ID
	ChannelName string // 名称
	ChannelIcon string // 图标
	SceneId     string // 场景ID
	PayId       string // 支付ID
	PayMethod   string // 支付方法：0APP支付 1H5支付 2扫码支付 3小程序支付
	Sort        string // 排序值。从大到小排序
	TotalAmount string // 总额
}

// channelColumns holds the columns for the table pay_channel.
var channelColumns = ChannelColumns{
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	IsStop:      "is_stop",
	ChannelId:   "channel_id",
	ChannelName: "channel_name",
	ChannelIcon: "channel_icon",
	SceneId:     "scene_id",
	PayId:       "pay_id",
	PayMethod:   "pay_method",
	Sort:        "sort",
	TotalAmount: "total_amount",
}

// NewChannelDao creates and returns a new DAO object for table data access.
func NewChannelDao(handlers ...gdb.ModelHandler) *ChannelDao {
	dao := &ChannelDao{
		group:    "default",
		table:    "pay_channel",
		columns:  channelColumns,
		handlers: handlers,
	}
	v := reflect.ValueOf(dao.columns)
	count := v.NumField()
	dao.columnArr = make([]string, count)
	dao.columnMap = make(map[string]struct{}, count)
	for i := range count {
		dao.columnArr[i] = v.Field(i).String()
		dao.columnMap[v.Field(i).String()] = struct{}{}
	}
	return dao
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ChannelDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ChannelDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
// 使用较为频繁。为优化内存考虑，改成返回指针更为合适，但切忌使用过程中不可修改，否则会污染全局
func (dao *ChannelDao) Columns() *ChannelColumns {
	return &dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ChannelDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ChannelDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *ChannelDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

// 字段数组
func (dao *ChannelDao) ColumnArr() []string {
	return dao.columnArr
}

// 字段map
func (dao *ChannelDao) ColumnMap() map[string]struct{} {
	return dao.columnMap
}

// 判断字段是否存在
func (dao *ChannelDao) Contains(column string) (ok bool) {
	_, ok = dao.columnMap[column]
	return
}
