package api

import (
	"github.com/gogf/gf/v2/frame/g"
)

type TestReq struct {
	// g.Meta `path:"/test/*path" method:"get,post" mime:"application/json" deprecated:"废弃标记" tags:"测试。标签，用于分类" sm:"接口名称" dc:"详细描述"`
	g.Meta `path:"/test" method:"get" mime:"application/json" deprecated:"true" tags:"测试。标签，用于分类" sm:"接口名称" dc:"详细描述"`
	//校验规则v在转换规则json之前执行，所以当结构体字段名与转换规则json中字段名不一致时，校验规则v中必须写别名，即此示例：test_test@，否则校验会报错
	Test string `json:"test_test" v:"test_test@required|length:4,30#请输入测试字段|长度在4~30个字符之间" in:"query" default:"默认值（嵌套结构体二级不起作用）" dc:"测试字段"`
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
