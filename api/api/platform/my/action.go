package api

import (
	"github.com/gogf/gf/v2/frame/g"
)

/*--------列表 开始--------*/
type ActionListReq struct {
	g.Meta `path:"/action/list" method:"post" tags:"平台后台/我的" sm:"操作列表"`
}

type ActionListRes struct {
	List []ActionItem `json:"list" dc:"列表"`
}

type ActionItem struct {
	Id         uint   `json:"id" dc:"ID"`
	Label      string `json:"label" dc:"标签。常用于前端组件"`
	ActionId   uint   `json:"actionId" dc:"操作ID"`
	ActionName string `json:"actionName" dc:"操作名称"`
}

/*--------列表 结束--------*/
