// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Goods is the golang structure of table goods for DAO operations like Where/Data.
type Goods struct {
	g.Meta    `orm:"table:goods, do:true"`
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	IsStop    interface{} // 停用：0否 1是
	GoodsId   interface{} // 商品ID
	GoodsName interface{} // 名称
	OrgId     interface{} // 机构ID
	GoodsNo   interface{} // 编号
	Image     interface{} // 图片
	AttrShow  interface{} // 展示属性。JSON格式：[{"name":"属性名","val":"属性值"},...]
	AttrOpt   interface{} // 可选属性。通常由不会影响价格和库存的属性组成。JSON格式：[{"name":"属性名","val_arr":["属性值1","属性值2",...]},...]
	Status    interface{} // 状态：0上架 1下架
	Sort      interface{} // 排序值。从大到小排序
}
