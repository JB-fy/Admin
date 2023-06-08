package api

import (
	apiCommon "api/api"
)

type RoleListReq struct {
	apiCommon.CommonListReq
	Filter RoleListFilterReq `p:"filter"`
}

type RoleListFilterReq struct {
	apiCommon.CommonListFilterReq `c:",omitempty"`
	RoleId                        *uint  `c:"roleId,omitempty" p:"roleId" v:"integer|min:1"`
	RoleName                      string `c:"roleName,omitempty" p:"roleName" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$"`
	SceneId                       *uint  `c:"sceneId,omitempty" p:"sceneId" v:"integer|min:1"`
	IsStop                        *uint  `c:"isStop,omitempty" p:"isStop" v:"integer|in:0,1"`
}

type RoleInfoReq struct {
	apiCommon.CommonInfoReq
}

type RoleCreateReq struct {
	RoleName    *string `c:"roleName,omitempty" p:"roleName" v:"required|length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$"`
	SceneId     *uint   `c:"sceneId,omitempty" p:"required|sceneId" v:"integer|min:1"`
	MenuIdArr   *[]uint `c:"menuIdArr,omitempty" p:"menuIdArr" v:"required|foreach|integer|foreach|min:1"`
	ActionIdArr *[]uint `c:"actionIdArr,omitempty" p:"actionIdArr" v:"required|foreach|integer|foreach|min:1"`
	IsStop      *uint   `c:"isStop,omitempty" p:"isStop" v:"integer|in:0,1"`
}

type RoleUpdateReq struct {
	apiCommon.CommonUpdateDeleteIdArrReq `c:",omitempty"`
	RoleName                             *string `c:"roleName,omitempty" p:"roleName" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$"`
	SceneId                              *uint   `c:"sceneId,omitempty" p:"sceneId" v:"integer|min:1"`
	MenuIdArr                            *[]uint `c:"menuIdArr,omitempty" p:"menuIdArr" v:"foreach|integer|foreach|min:1"`
	ActionIdArr                          *[]uint `c:"actionIdArr,omitempty" p:"actionIdArr" v:"foreach|integer|foreach|min:1"`
	IsStop                               *uint   `c:"isStop,omitempty" p:"isStop" v:"integer|in:0,1"`
}

type RoleDeleteReq struct {
	apiCommon.CommonUpdateDeleteIdArrReq
}
