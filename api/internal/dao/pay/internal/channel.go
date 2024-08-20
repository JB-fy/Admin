// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"
	"reflect"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ChannelDao is the data access object for table pay_channel.
type ChannelDao struct {
	table     string           // table is the underlying table name of the DAO.
	group     string           // group is the database configuration group name of current DAO.
	columns   ChannelColumns   // columns contains all the column names of Table for convenient usage.
	columnArr *garray.StrArray // 所有字段的数组
}

// ChannelColumns defines and stores column names for table pay_channel.
type ChannelColumns struct {
	CreatedAt   string // 创建时间
	UpdatedAt   string // 更新时间
	IsStop      string // 停用：0否 1是
	ChannelId   string // 通道ID
	ChannelName string // 名称
	ChannelIcon string // 图标
	SceneId     string // 场景ID
	PayId       string // 支付ID
	Method      string // 支付方式：1APP支付 2H5支付 3扫码支付 4小程序支付
	Sort        string // 排序值。从大到小排序
	TotalAmount string // 总额
}

// channelColumns holds the columns for table pay_channel.
var channelColumns = ChannelColumns{
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	IsStop:      "is_stop",
	ChannelId:   "channel_id",
	ChannelName: "channel_name",
	ChannelIcon: "channel_icon",
	SceneId:     "scene_id",
	PayId:       "pay_id",
	Method:      "method",
	Sort:        "sort",
	TotalAmount: "total_amount",
}

// NewChannelDao creates and returns a new DAO object for table data access.
func NewChannelDao() *ChannelDao {
	return &ChannelDao{
		group:   `default`,
		table:   `pay_channel`,
		columns: channelColumns,
		columnArr: func() *garray.StrArray {
			v := reflect.ValueOf(channelColumns)
			count := v.NumField()
			column := make([]string, count)
			for i := 0; i < count; i++ {
				column[i] = v.Field(i).String()
			}
			return garray.NewStrArrayFrom(column)
		}(),
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ChannelDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ChannelDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
// 使用较为频繁。为优化内存考虑，改成返回指针更为合适，但切忌使用过程中不可修改，否则会污染全局
func (dao *ChannelDao) Columns() *ChannelColumns {
	return &dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ChannelDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ChannelDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ChannelDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

// 所有字段的数组
func (dao *ChannelDao) ColumnArr() *garray.StrArray {
	return dao.columnArr
}
