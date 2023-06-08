package api

import (
	apiCommon "api/api"
)

type RequestListReq struct {
	apiCommon.CommonListReq
	Filter RequestListFilterReq `p:"filter"`
}

type RequestListFilterReq struct {
	apiCommon.CommonListFilterReq `c:",omitempty"`
	LogId                         *uint    `c:"logId,omitempty" p:"logId" v:"integer|min:1"`
	RequestUrl                    string   `c:"requestUrl,omitempty" p:"requestUrl" v:"url"`
	MinRunTime                    *float64 `c:"minRunTime,omitempty" p:"minRunTime" v:"float|min:0"`
	MaxRunTime                    *float64 `c:"maxRunTime,omitempty" p:"maxRunTime" v:"float|min:0|gte:MinRunTime"`
}
