package api

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type TestReq struct {
	g.Meta `path:"/test" method:"get,post" mime:"application/json" deprecated:"废弃标记" tags:"测试" sm:"接口名称" dc:"详细描述"`
	// 当结构体字段名与参数传入字段名不一致时，验证规则必须设置属性别名，否则会默认以结构体字段名去寻找字段值做验证，形成隐患BUG
	// Xxxx string `json:"xxxx,omitempty" v:"[属性别名@]校验规则1|校验规则2[#校验规则1错误提示|校验规则2错误提示]" in:"header/path/query/cookie" d:"默认值（嵌套结构体二级不起作用）" dc:"字段说明"`
	Test   string     `json:"test,omitempty" v:"test@required|length:4,30#请输入测试字段|长度在4~30个字符之间" d:"test" dc:"测试"`
	Filter TestFilter `json:"filter" d:"{\"timeRangeStart\": \"2006-01-02 15:04:05\"}" dc:"过滤条件"`
	/* Filter struct {
		FilterTimeRangeStart *gtime.Time `json:"timeRangeStart,omitempty" v:"date-format:Y-m-d H:i:s" dc:"开始时间：YYYY-mm-dd HH:ii:ss"`
	} `json:"where" dc:"查询条件"` */
}

type TestFilter struct {
	Id             *uint       `json:"id,omitempty" v:"min:1" dc:"ID"`
	IdArr          []uint      `json:"idArr,omitempty" v:"distinct|foreach|min:1" dc:"ID数组"`
	ExcId          *uint       `json:"excId,omitempty" v:"min:1" dc:"排除ID"`
	ExcIdArr       []uint      `json:"excIdArr,omitempty" v:"distinct|foreach|min:1" dc:"排除ID数组"`
	Label          string      `json:"label,omitempty" v:"max-length:30|regex:^[\\p{L}\\p{N}_-]+$" dc:"标签。常用于前端组件"`
	TimeRangeStart *gtime.Time `json:"timeRangeStart,omitempty" v:"date-format:Y-m-d H:i:s" dc:"开始时间：YYYY-mm-dd HH:ii:ss"`
	TimeRangeEnd   *gtime.Time `json:"timeRangeEnd,omitempty" v:"date-format:Y-m-d H:i:s|after-equal:TimeRangeStart" dc:"结束时间：YYYY-mm-dd HH:ii:ss"`
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
