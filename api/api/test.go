package api

import "github.com/gogf/gf/v2/frame/g"

type TestMetaReq struct {
	g.Meta `path:"/testMeta" method:"get,post" mime:"application/json" deprecated:"废弃标记" tags:"标签，用于分类" sm:"概要描述" dc:"详细描述"`
	Test   string `json:"test" v:"required|length:4,30" in:"query" default:"默认值（嵌套结构体二级不起作用）" dc:"详细描述"`
}

type TestList struct {
	Test string `json:"test" dc:"测试"`
}

type TestMetaRes struct {
	g.Meta `mime:"text/html" example:"string"`
	Test   string   `json:"test" dc:"测试"`
	List   TestList `json:"list" dc:"列表"`
}

type TestReq struct {
	Test string `json:"test" v:"required|length:4,30#请输入账号|账号长度为:{min}到:{max}位"`
}
