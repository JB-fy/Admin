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

// GoodsDao is the data access object for table goods.
type GoodsDao struct {
	table     string           // table is the underlying table name of the DAO.
	group     string           // group is the database configuration group name of current DAO.
	columns   GoodsColumns     // columns contains all the column names of Table for convenient usage.
	columnArr *garray.StrArray // 所有字段的数组
}

// GoodsColumns defines and stores column names for table goods.
type GoodsColumns struct {
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
	IsStop    string // 停用：0否 1是
	GoodsId   string // 商品ID
	GoodsName string // 名称
	OrgId     string // 机构ID
	GoodsNo   string // 编号
	Image     string // 图片
	AttrShow  string // 展示属性。JSON格式：[{"name":"属性名","val":"属性值"},...]
	AttrOpt   string // 可选属性。通常由不会影响价格和库存的属性组成。JSON格式：[{"name":"属性名","val_arr":["属性值1","属性值2",...]},...]
	Status    string // 状态：0上架 1下架
	Sort      string // 排序值。从大到小排序
}

// goodsColumns holds the columns for table goods.
var goodsColumns = GoodsColumns{
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	IsStop:    "is_stop",
	GoodsId:   "goods_id",
	GoodsName: "goods_name",
	OrgId:     "org_id",
	GoodsNo:   "goods_no",
	Image:     "image",
	AttrShow:  "attr_show",
	AttrOpt:   "attr_opt",
	Status:    "status",
	Sort:      "sort",
}

// NewGoodsDao creates and returns a new DAO object for table data access.
func NewGoodsDao() *GoodsDao {
	return &GoodsDao{
		group:   `default`,
		table:   `goods`,
		columns: goodsColumns,
		columnArr: func() *garray.StrArray {
			v := reflect.ValueOf(goodsColumns)
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
func (dao *GoodsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *GoodsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
// 使用较为频繁。为优化内存考虑，改成返回指针更为合适，但切忌使用过程中不可修改，否则会污染全局
func (dao *GoodsDao) Columns() *GoodsColumns {
	return &dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *GoodsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *GoodsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *GoodsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

// 所有字段的数组
func (dao *GoodsDao) ColumnArr() *garray.StrArray {
	return dao.columnArr
}
