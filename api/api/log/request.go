package api

import (
	apiCommon "api/api"
)

type RequestListReq struct {
	apiCommon.CommonListReq
	Filter RequestListFilterReq `p:"filter"`
}

/* type Request struct {
    LogId         uint        `json:"logId"         `// 请求日志ID
    RequestUrl    string      `json:"requestUrl"    `// 请求地址
    RequestHeader string      `json:"requestHeader" `// 请求头
    RequestData   string      `json:"requestData"   `// 请求数据
    ResponseBody  string      `json:"responseBody"  `// 响应体
    RunTime       float64     `json:"runTime"       `// 运行时间（单位：毫秒）
    UpdateTime    *gtime.Time `json:"updateTime"    `// 更新时间
    CreateTime    *gtime.Time `json:"createTime"    `// 创建时间
} */

type RequestListFilterReq struct {
	apiCommon.CommonListFilterReq `c:",omitempty"`
	RequestId                     *uint  `c:"requestId,omitempty" p:"requestId" v:"min:1"`
	RequestName                   string `c:"requestName,omitempty" p:"requestName" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$"`
	IsStop                        *uint  `c:"isStop,omitempty" p:"isStop" v:"in:0,1"`
}

type RequestInfoReq struct {
	apiCommon.CommonInfoReq
}

type RequestCreateReq struct {
	RequestName *string `c:"requestName,omitempty" p:"requestName" v:"required|length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$"`
	IsStop      *uint   `c:"isStop,omitempty" p:"isStop" v:"in:0,1"`
}

type RequestUpdateReq struct {
	apiCommon.CommonUpdateDeleteIdArrReq `c:",omitempty"`
	RequestName                          *string `c:"requestName,omitempty" p:"requestName" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$"`
	IsStop                               *uint   `c:"isStop,omitempty" p:"isStop" v:"in:0,1"`
}

type RequestDeleteReq struct {
	apiCommon.CommonUpdateDeleteIdArrReq
}
