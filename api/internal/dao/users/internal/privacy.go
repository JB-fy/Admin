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

// PrivacyDao is the data access object for table users_privacy.
type PrivacyDao struct {
	table     string           // table is the underlying table name of the DAO.
	group     string           // group is the database configuration group name of current DAO.
	columns   PrivacyColumns   // columns contains all the column names of Table for convenient usage.
	columnArr *garray.StrArray // 所有字段的数组
}

// PrivacyColumns defines and stores column names for table users_privacy.
type PrivacyColumns struct {
	CreatedAt      string // 创建时间
	UpdatedAt      string // 更新时间
	UserId         string // 用户ID
	Password       string // 密码。md5保存
	Salt           string // 密码盐
	IdCardNo       string // 身份证号码
	IdCardName     string // 身份证姓名
	IdCardGender   string // 身份证性别：0未设置 1男 2女
	IdCardBirthday string // 身份证生日
	IdCardAddress  string // 身份证地址
}

// privacyColumns holds the columns for table users_privacy.
var privacyColumns = PrivacyColumns{
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
	UserId:         "user_id",
	Password:       "password",
	Salt:           "salt",
	IdCardNo:       "id_card_no",
	IdCardName:     "id_card_name",
	IdCardGender:   "id_card_gender",
	IdCardBirthday: "id_card_birthday",
	IdCardAddress:  "id_card_address",
}

// NewPrivacyDao creates and returns a new DAO object for table data access.
func NewPrivacyDao() *PrivacyDao {
	return &PrivacyDao{
		group:   `default`,
		table:   `users_privacy`,
		columns: privacyColumns,
		columnArr: func() *garray.StrArray {
			v := reflect.ValueOf(privacyColumns)
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
func (dao *PrivacyDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *PrivacyDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
// 使用较为频繁。为优化内存考虑，改成返回指针更为合适，但切忌使用过程中不可修改，否则会污染全局
func (dao *PrivacyDao) Columns() *PrivacyColumns {
	return &dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *PrivacyDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *PrivacyDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *PrivacyDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

// 所有字段的数组
func (dao *PrivacyDao) ColumnArr() *garray.StrArray {
	return dao.columnArr
}
