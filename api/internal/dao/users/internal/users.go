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

// UsersDao is the data access object for table users.
type UsersDao struct {
	table     string           // table is the underlying table name of the DAO.
	group     string           // group is the database configuration group name of current DAO.
	columns   UsersColumns     // columns contains all the column names of Table for convenient usage.
	columnArr *garray.StrArray // 所有字段的数组
}

// UsersColumns defines and stores column names for table users.
type UsersColumns struct {
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
	IsStop    string // 停用：0否 1是
	UserId    string // 用户ID
	Nickname  string // 昵称
	Avatar    string // 头像
	Gender    string // 性别：0未设置 1男 2女
	Birthday  string // 生日
	Address   string // 地址
	Phone     string // 手机
	Email     string // 邮箱
	Account   string // 账号
	WxOpenId  string // 微信openId
	WxUnionId string // 微信unionId
}

// usersColumns holds the columns for table users.
var usersColumns = UsersColumns{
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	IsStop:    "is_stop",
	UserId:    "user_id",
	Nickname:  "nickname",
	Avatar:    "avatar",
	Gender:    "gender",
	Birthday:  "birthday",
	Address:   "address",
	Phone:     "phone",
	Email:     "email",
	Account:   "account",
	WxOpenId:  "wx_open_id",
	WxUnionId: "wx_union_id",
}

// NewUsersDao creates and returns a new DAO object for table data access.
func NewUsersDao() *UsersDao {
	return &UsersDao{
		group:   `default`,
		table:   `users`,
		columns: usersColumns,
		columnArr: func() *garray.StrArray {
			v := reflect.ValueOf(usersColumns)
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
func (dao *UsersDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UsersDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
// 使用较为频繁。为优化内存考虑，改成返回指针更为合适，但切忌使用过程中不可修改，否则会污染全局
func (dao *UsersDao) Columns() *UsersColumns {
	return &dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UsersDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UsersDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UsersDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

// 所有字段的数组
func (dao *UsersDao) ColumnArr() *garray.StrArray {
	return dao.columnArr
}
