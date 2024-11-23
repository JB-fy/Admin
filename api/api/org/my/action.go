package my

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
	Id    *string `json:"id,omitempty" dc:"ID"`
	Label *string `json:"label,omitempty" dc:"标签。常用于前端组件"`
}

/*--------列表 结束--------*/
