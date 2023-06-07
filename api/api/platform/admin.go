package api

import (
	apiCommon "api/api"
)

type AdminListReq struct {
	apiCommon.CommonListReq
	Filter AdminListFilterReq `p:"filter"`
}

type AdminListFilterReq struct {
	apiCommon.CommonListFilterReq `c:",omitempty"`
	AdminId                       *uint  `c:"adminId,omitempty" p:"adminId" v:"min:1"`
	Account                       string `c:"account,omitempty" p:"account" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$"`
	Phone                         string `c:"phone,omitempty" p:"phone" v:"phone"`
	RoleId                        *uint  `c:"roleId,omitempty" p:"roleId" v:"min:1"`
	IsStop                        *uint  `c:"isStop,omitempty" p:"isStop" v:"in:0,1"`
}

type AdminInfoReq struct {
	apiCommon.CommonInfoReq
}

type AdminCreateReq struct {
	Account   *string `c:"account,omitempty" p:"account" v:"required-without:Phone|length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$"`
	Phone     *string `c:"phone,omitempty" p:"phone" v:"required-without:Account|phone"`
	Password  *string `c:"password,omitempty" p:"password" v:"required|size:32|regex:^[\\p{L}\\p{N}_-]+$"`
	RoleIdArr *[]uint `c:"roleIdArr,omitempty" p:"roleIdArr" v:"required|foreach|min:1"`
	Nickname  *string `c:"nickname,omitempty" p:"nickname" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$"`
	Avatar    *string `c:"avatar,omitempty" p:"avatar" v:"url|length:1,120"`
	IsStop    *uint   `c:"isStop,omitempty" p:"isStop" v:"in:0,1"`
}

type AdminUpdateReq struct {
	apiCommon.CommonUpdateDeleteIdArrReq `c:",omitempty"`
	Account                              *string `c:"account,omitempty" p:"account" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$"`
	Phone                                *string `c:"phone,omitempty" p:"phone" v:"phone"`
	Password                             *string `c:"password,omitempty" p:"password" v:"size:32|regex:^[\\p{L}\\p{N}_-]+$"`
	RoleIdArr                            *[]uint `c:"roleIdArr,omitempty" p:"roleIdArr" v:"foreach|min:1"`
	Nickname                             *string `c:"nickname,omitempty" p:"nickname" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$"`
	Avatar                               *string `c:"avatar,omitempty" p:"avatar" v:"url|length:1,120"`
	IsStop                               *uint   `c:"isStop,omitempty" p:"isStop" v:"in:0,1"`
}

type AdminDeleteReq struct {
	apiCommon.CommonUpdateDeleteIdArrReq
}

type AdminUpdateSelfReq struct {
	Account       *string `c:"account,omitempty" p:"account" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$"`
	Phone         *string `c:"phone,omitempty" p:"phone" v:"phone"`
	Nickname      *string `c:"nickname,omitempty" p:"nickname" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$"`
	Avatar        *string `c:"avatar,omitempty" p:"avatar" v:"url|length:1,120"`
	Password      *string `c:"password,omitempty" p:"password" v:"size:32|regex:^[\\p{L}\\p{N}_-]+$|different:CheckPassword"`
	CheckPassword *string `c:"checkPassword,omitempty" p:"checkPassword" v:"required-with:account,phone,password|size:32|regex:^[\\p{L}\\p{N}_-]+$"`
}
