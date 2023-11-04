package api

import (
	"github.com/gogf/gf/v2/frame/g"
)

type TestReq struct {
	// g.Meta `path:"/test/*path" method:"get,post" mime:"application/json" deprecated:"废弃标记" tags:"测试。标签，用于分类" sm:"接口名称" dc:"详细描述"`
	g.Meta `path:"/test" method:"get" mime:"application/json" deprecated:"true" tags:"测试。标签，用于分类" sm:"接口名称" dc:"详细描述"`
	Test   string `json:"test" v:"required|length:4,30" in:"query" default:"默认值（嵌套结构体二级不起作用）" dc:"详细描述"`
}

type TestRes struct {
	g.Meta `mime:"text/html" example:"string"`
	Info   *TestInfo `json:"info,omitempty" dc:"详情"`
}

type TestInfo struct {
	Id     *uint   `json:"id,omitempty" dc:"ID"`
	Label  *string `json:"label,omitempty" dc:"标签。常用于前端组件"`
	TestId *uint   `json:"testId,omitempty" dc:"测试ID"`
}
