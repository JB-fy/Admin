package org

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// 共用详情。list,info,tree等接口返回时用，但返回默认字段有差异。可根据需要在controller对应的defaultField中补充所需字段
type AdminInfo struct {
	Id      *uint   `json:"id,omitempty" dc:"ID"`
	Label   *string `json:"label,omitempty" dc:"标签。常用于前端组件"`
	AdminId *uint   `json:"admin_id,omitempty" dc:"管理员ID"`
	// OrgId     *uint       `json:"org_id,omitempty" dc:"机构ID"`
	Nickname  *string     `json:"nickname,omitempty" dc:"昵称"`
	Avatar    *string     `json:"avatar,omitempty" dc:"头像"`
	Phone     *string     `json:"phone,omitempty" dc:"手机"`
	Email     *string     `json:"email,omitempty" dc:"邮箱"`
	Account   *string     `json:"account,omitempty" dc:"账号"`
	IsSuper   *uint       `json:"is_super,omitempty" dc:"超管：0否 1是"`
	RoleIdArr []uint      `json:"role_id_arr,omitempty" dc:"角色ID"`
	IsStop    *uint       `json:"is_stop,omitempty" dc:"停用：0否 1是"`
	UpdatedAt *gtime.Time `json:"updated_at,omitempty" dc:"更新时间"`
	CreatedAt *gtime.Time `json:"created_at,omitempty" dc:"创建时间"`
	// OrgName   *string     `json:"org_name,omitempty" dc:"机构"`
}

type AdminFilter struct {
	Id             *uint       `json:"id,omitempty" v:"between:1,4294967295" dc:"ID"`
	IdArr          []uint      `json:"id_arr,omitempty" v:"distinct|foreach|between:1,4294967295" dc:"ID数组"`
	ExcId          *uint       `json:"exc_id,omitempty" v:"between:1,4294967295" dc:"排除ID"`
	ExcIdArr       []uint      `json:"exc_id_arr,omitempty" v:"distinct|foreach|between:1,4294967295" dc:"排除ID数组"`
	Label          string      `json:"label,omitempty" v:"max-length:30|regex:^[\\p{L}\\p{N}_-]+$" dc:"标签。常用于前端组件"`
	TimeRangeStart *gtime.Time `json:"time_range_start,omitempty" v:"date-format:Y-m-d H:i:s" dc:"开始时间：YYYY-mm-dd HH:ii:ss"`
	TimeRangeEnd   *gtime.Time `json:"time_range_end,omitempty" v:"date-format:Y-m-d H:i:s|after-equal:TimeRangeStart" dc:"结束时间：YYYY-mm-dd HH:ii:ss"`
	AdminId        *uint       `json:"admin_id,omitempty" v:"between:1,4294967295" dc:"管理员ID"`
	// OrgId          *uint       `json:"org_id,omitempty" v:"between:1,4294967295" dc:"机构ID"`
	Nickname string `json:"nickname,omitempty" v:"max-length:30" dc:"昵称"`
	Phone    string `json:"phone,omitempty" v:"max-length:20|phone" dc:"手机"`
	Account  string `json:"account,omitempty" v:"max-length:20|regex:^[\\p{L}][\\p{L}\\p{N}_]{3,}$" dc:"账号"`
	IsSuper  *uint  `json:"is_super,omitempty" v:"in:0,1" dc:"超管：0否 1是"`
	RoleId   *uint  `json:"role_id,omitempty" v:"between:1,4294967295" dc:"角色ID"`
	IsStop   *uint  `json:"is_stop,omitempty" v:"in:0,1" dc:"停用：0否 1是"`
}

/*--------列表 开始--------*/
type AdminListReq struct {
	g.Meta `path:"/admin/list" method:"post" tags:"机构后台/权限管理/管理员" sm:"列表"`
	Filter AdminFilter `json:"filter" dc:"过滤条件"`
	Field  []string    `json:"field" v:"distinct|foreach|min-length:1" dc:"查询字段，传值参考返回的字段名，默认返回常用字段，如果所需字段较少或需特别字段时，可使用。特别注意：所需字段较少时使用，可大幅减轻数据库压力"`
	Sort   string      `json:"sort" default:"id DESC" dc:"排序"`
	Page   int         `json:"page" v:"min:1" default:"1" dc:"页码"`
	Limit  int         `json:"limit" v:"min:0" default:"10" dc:"每页数量。可传0取全部"`
}

type AdminListRes struct {
	Count int         `json:"count" dc:"总数"`
	List  []AdminInfo `json:"list" dc:"列表"`
}

/*--------列表 结束--------*/

/*--------详情 开始--------*/
type AdminInfoReq struct {
	g.Meta `path:"/admin/info" method:"post" tags:"机构后台/权限管理/管理员" sm:"详情"`
	Field  []string `json:"field" v:"distinct|foreach|min-length:1" dc:"查询字段，传值参考返回的字段名，默认返回常用字段，如果所需字段较少或需特别字段时，可使用。特别注意：所需字段较少时使用，可大幅减轻数据库压力"`
	Id     uint     `json:"id" v:"required|between:1,4294967295" dc:"ID"`
}

type AdminInfoRes struct {
	Info AdminInfo `json:"info" dc:"详情"`
}

/*--------详情 结束--------*/

/*--------新增 开始--------*/
type AdminCreateReq struct {
	g.Meta `path:"/admin/create" method:"post" tags:"机构后台/权限管理/管理员" sm:"新增"`
	// OrgId     *uint   `json:"org_id,omitempty" v:"between:0,4294967295" dc:"机构ID"`
	Nickname *string `json:"nickname,omitempty" v:"max-length:30" dc:"昵称"`
	Avatar   *string `json:"avatar,omitempty" v:"max-length:200|url" dc:"头像"`
	Phone    *string `json:"phone,omitempty" v:"max-length:20|phone" dc:"手机"`
	Email    *string `json:"email,omitempty" v:"max-length:60|email" dc:"邮箱"`
	Account  *string `json:"account,omitempty" v:"max-length:20|regex:^[\\p{L}][\\p{L}\\p{N}_]{3,}$" dc:"账号"`
	Password *string `json:"password,omitempty" v:"required|size:32" dc:"密码。md5保存"`
	// IsSuper   *uint   `json:"is_super,omitempty" v:"in:0,1" dc:"超管：0否 1是"`
	RoleIdArr *[]uint `json:"role_id_arr,omitempty" v:"required|distinct|foreach|between:1,4294967295" dc:"角色ID"`
	IsStop    *uint   `json:"is_stop,omitempty" v:"in:0,1" dc:"停用：0否 1是"`
}

/*--------新增 结束--------*/

/*--------修改 开始--------*/
type AdminUpdateReq struct {
	g.Meta `path:"/admin/update" method:"post" tags:"机构后台/权限管理/管理员" sm:"修改"`
	Id     uint   `json:"id,omitempty" filter:"id,omitempty" data:"-" v:"required-without:IdArr|between:1,4294967295" dc:"ID"`
	IdArr  []uint `json:"id_arr,omitempty" filter:"id_arr,omitempty" data:"-" v:"required-without:Id|distinct|foreach|between:1,4294967295" dc:"ID数组"`
	// OrgId     *uint   `json:"org_id,omitempty" filter:"-" data:"org_id,omitempty" v:"between:0,4294967295" dc:"机构ID"`
	Nickname *string `json:"nickname,omitempty" filter:"-" data:"nickname,omitempty" v:"max-length:30" dc:"昵称"`
	Avatar   *string `json:"avatar,omitempty" filter:"-" data:"avatar,omitempty" v:"max-length:200|url" dc:"头像"`
	Phone    *string `json:"phone,omitempty" filter:"-" data:"phone,omitempty" v:"max-length:20|phone" dc:"手机"`
	Email    *string `json:"email,omitempty" filter:"-" data:"email,omitempty" v:"max-length:60|email" dc:"邮箱"`
	Account  *string `json:"account,omitempty" filter:"-" data:"account,omitempty" v:"max-length:20|regex:^[\\p{L}][\\p{L}\\p{N}_]{3,}$" dc:"账号"`
	Password *string `json:"password,omitempty" filter:"-" data:"password,omitempty" v:"size:32" dc:"密码。md5保存"`
	// IsSuper   *uint   `json:"is_super,omitempty" filter:"-" data:"is_super,omitempty" v:"in:0,1" dc:"超管：0否 1是"`
	RoleIdArr *[]uint `json:"role_id_arr,omitempty" v:"distinct|foreach|between:1,4294967295" dc:"角色ID"`
	IsStop    *uint   `json:"is_stop,omitempty" filter:"-" data:"is_stop,omitempty" v:"in:0,1" dc:"停用：0否 1是"`
}

/*--------修改 结束--------*/

/*--------删除 开始--------*/
type AdminDeleteReq struct {
	g.Meta `path:"/admin/del" method:"post" tags:"机构后台/权限管理/管理员" sm:"删除"`
	Id     uint   `json:"id,omitempty" v:"required-without:IdArr|between:1,4294967295" dc:"ID"`
	IdArr  []uint `json:"id_arr,omitempty" v:"required-without:Id|distinct|foreach|between:1,4294967295" dc:"ID数组"`
}

/*--------删除 结束--------*/
