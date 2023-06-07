package api

import (
	apiCommon "api/api"
)

type ActionListReq struct {
	apiCommon.CommonListReq
	Filter ActionListFilterReq `p:"filter"`
}

type ActionListFilterReq struct {
	apiCommon.CommonListFilterReq `c:",omitempty"`
	ActionId                      *uint   `c:"actionId,omitempty" p:"actionId" v:"min:1"`
	ActionName                    string  `c:"actionName,omitempty" p:"actionName" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$"`
	ActionCode                    *string `c:"actionCode,omitempty" p:"actionCode" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$"`
	SceneId                       *uint   `c:"sceneId,omitempty" p:"sceneId" v:"min:1"`
	IsStop                        *uint   `c:"isStop,omitempty" p:"isStop" v:"in:0,1"`
}

type ActionInfoReq struct {
	apiCommon.CommonInfoReq
}

type ActionCreateReq struct {
	ActionName *string `c:"actionName,omitempty" p:"actionName" v:"required|length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$"`
	ActionCode *string `c:"actionCode,omitempty" p:"actionCode" v:"required|length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$"`
	SceneIdArr *[]uint `c:"sceneIdArr,omitempty" p:"sceneIdArr" v:"required|foreach|min:1"`
	Remark     *[]uint `c:"remark,omitempty" p:"remark" v:"length:1,120"`
	IsStop     *uint   `c:"isStop,omitempty" p:"isStop" v:"in:0,1"`
}

type ActionUpdateReq struct {
	apiCommon.CommonUpdateDeleteIdArrReq `c:",omitempty"`
	ActionName                           *string `c:"actionName,omitempty" p:"actionName" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$"`
	ActionCode                           *string `c:"actionCode,omitempty" p:"actionCode" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$"`
	SceneIdArr                           *[]uint `c:"sceneIdArr,omitempty" p:"sceneIdArr" v:"foreach|min:1"`
	Remark                               *[]uint `c:"remark,omitempty" p:"remark" v:"length:1,120"`
	IsStop                               *uint   `c:"isStop,omitempty" p:"isStop" v:"in:0,1"`
}

type ActionDeleteReq struct {
	apiCommon.CommonUpdateDeleteIdArrReq
}
