package api

import "github.com/gogf/gf/v2/frame/g"

type TestMetaReq struct {
	g.Meta `path:"/testMeta" method:"get" tags:"测试" summary:"测试"`
	Test   string `json:"test" v:"required|length:4,30" dc:"测试"`
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
