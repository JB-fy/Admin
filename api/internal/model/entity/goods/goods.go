// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Goods is the golang structure for table goods.
type Goods struct {
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" ` // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" ` // 更新时间
	IsStop    uint        `json:"isStop"    orm:"is_stop"    ` // 停用：0否 1是
	GoodsId   uint        `json:"goodsId"   orm:"goods_id"   ` // 商品ID
	GoodsName string      `json:"goodsName" orm:"goods_name" ` // 名称
	OrgId     uint        `json:"orgId"     orm:"org_id"     ` // 机构ID
	GoodsNo   string      `json:"goodsNo"   orm:"goods_no"   ` // 编号
	Image     string      `json:"image"     orm:"image"      ` // 图片
	AttrShow  string      `json:"attrShow"  orm:"attr_show"  ` // 展示属性。JSON格式：[{"name":"属性名","val":"属性值"},...]
	AttrOpt   string      `json:"attrOpt"   orm:"attr_opt"   ` // 可选属性。通常由不会影响价格和库存的属性组成。JSON格式：[{"name":"属性名","val_arr":["属性值1","属性值2",...]},...]
	Status    uint        `json:"status"    orm:"status"     ` // 状态：0上架 1下架
	Sort      uint        `json:"sort"      orm:"sort"       ` // 排序值。从大到小排序
}
