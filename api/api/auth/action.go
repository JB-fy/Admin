package api

import (
	apiCommon "api/api"
)

type ActionListReq struct {
	apiCommon.CommonListReq `c:",omitempty"`
	Filter                  ActionListFilterReq `p:"filter"`
}

/* type Action struct {
    ActionId   uint        `json:"actionId"   `// 权限操作ID
    ActionName string      `json:"actionName" `// 名称
    ActionCode string      `json:"actionCode" `// 标识（代码中用于判断权限）
    Remark     string      `json:"remark"     `// 备注
    IsStop     uint        `json:"isStop"     `// 是否停用：0否 1是
    UpdateTime *gtime.Time `json:"updateTime" `// 更新时间
    CreateTime *gtime.Time `json:"createTime" `// 创建时间
} */

type ActionListFilterReq struct {
	apiCommon.CommonListFilterReq `c:",omitempty"`
	ActionId                      *uint  `c:"actionId,omitempty" p:"actionId" v:"min:1"`
	ActionName                    string `c:"actionName,omitempty" p:"actionName" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$"`
	IsStop                        *uint  `c:"isStop,omitempty" p:"isStop" v:"in:0,1"`
}

type ActionInfoReq struct {
	apiCommon.CommonInfoReq `c:",omitempty"`
}

type ActionCreateReq struct {
	ActionName *string `c:"actionName,omitempty" p:"actionName" v:"required|length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$"`
	IsStop     *uint   `c:"isStop,omitempty" p:"isStop" v:"in:0,1"`
}

type ActionUpdateReq struct {
	apiCommon.CommonUpdateDeleteIdArrReq `c:",omitempty"`
	ActionName                           *string `c:"actionName,omitempty" p:"actionName" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$"`
	IsStop                               *uint   `c:"isStop,omitempty" p:"isStop" v:"in:0,1"`
}

type ActionDeleteReq struct {
	apiCommon.CommonUpdateDeleteIdArrReq `c:",omitempty"`
}
