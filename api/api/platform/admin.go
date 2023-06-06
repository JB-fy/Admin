package api

import (
	apiCommon "api/api"
)

type AdminListReq struct {
	apiCommon.CommonListReq
	Filter AdminListFilterReq `p:"filter"`
}

/* type Admin struct {
    AdminId    uint        `json:"adminId"    `// 管理员ID
    Account    string      `json:"account"    `// 账号
    Phone      string      `json:"phone"      `// 电话号码
    Password   string      `json:"password"   `// 密码（md5保存）
    Nickname   string      `json:"nickname"   `// 昵称
    Avatar     string      `json:"avatar"     `// 头像
    IsStop     uint        `json:"isStop"     `// 是否停用：0否 1是
    UpdateTime *gtime.Time `json:"updateTime" `// 更新时间
    CreateTime *gtime.Time `json:"createTime" `// 创建时间
} */

type AdminListFilterReq struct {
	apiCommon.CommonListFilterReq `c:",omitempty"`
	AdminId                       *uint  `c:"adminId,omitempty" p:"adminId" v:"min:1"`
	AdminName                     string `c:"adminName,omitempty" p:"adminName" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$"`
	IsStop                        *uint  `c:"isStop,omitempty" p:"isStop" v:"in:0,1"`
}

type AdminInfoReq struct {
	apiCommon.CommonInfoReq
}

type AdminCreateReq struct {
	AdminName *string `c:"adminName,omitempty" p:"adminName" v:"required|length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$"`
	IsStop    *uint   `c:"isStop,omitempty" p:"isStop" v:"in:0,1"`
}

type AdminUpdateReq struct {
	apiCommon.CommonUpdateDeleteIdArrReq `c:",omitempty"`
	AdminName                            *string `c:"adminName,omitempty" p:"adminName" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$"`
	IsStop                               *uint   `c:"isStop,omitempty" p:"isStop" v:"in:0,1"`
}

type AdminDeleteReq struct {
	apiCommon.CommonUpdateDeleteIdArrReq
}
