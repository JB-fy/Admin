package api

import (
	"github.com/gogf/gf/v2/frame/g"
)

type TestReq struct {
	// g.Meta `path:"/test/*path" method:"get,post" mime:"application/json" deprecated:"废弃标记" tags:"标签。用于分类" sm:"接口名称" dc:"详细描述"`
	g.Meta `path:"/test" method:"get" deprecated:"true" tags:"测试" sm:"接口名称" dc:"详细描述"`
	//当结构体字段名与参数传入字段名不一致时，验证规则必须设置属性别名，否则会默认以结构体字段名去寻找字段值做验证，形成隐患BUG
	// Xxxx string `json:"xxxx,omitempty" v:"[属性别名@]校验规则1|校验规则2[#校验规则1错误提示|校验规则2错误提示]" in:"header/path/query/cookie" d:"默认值（嵌套结构体二级不起作用）" dc:"字段说明"`
	Test string `json:"test_test,omitempty" v:"test_test@length:4,30#长度在4~30个字符之间" dc:"测试"`
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
