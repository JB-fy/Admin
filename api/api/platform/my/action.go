package api

import (
	"github.com/gogf/gf/v2/frame/g"
)

/*--------列表 开始--------*/
type ActionListReq struct {
	g.Meta `path:"/action/list" method:"post" tags:"平台后台/我的" sm:"操作列表"`
}

type ActionListRes struct {
	List []ActionListItem `json:"list" dc:"列表"`
}

type ActionListItem struct {
	Id         *uint   `json:"id,omitempty" dc:"ID"`
	Label      *string `json:"label,omitempty" dc:"标签。常用于前端组件"`
	ActionId   *uint   `json:"action_id,omitempty" dc:"操作ID"`
	ActionName *string `json:"action_name,omitempty" dc:"操作名称"`
	ActionCode *string `json:"action_code,omitempty" dc:"操作标识"`
}

/*--------列表 结束--------*/
