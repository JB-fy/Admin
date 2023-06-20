package log

import (
	apiCommon "api/api"
)

type HttpListReq struct {
	apiCommon.CommonListReq
	Filter HttpListFilterReq `p:"filter"`
}

type HttpListFilterReq struct {
	apiCommon.CommonListFilterReq `c:",omitempty"`
	HttpId                        *uint    `c:"httpId,omitempty" p:"httpId" v:"integer|min:1"`
	Url                           string   `c:"url,omitempty" p:"url" v:"url"`
	MinRunTime                    *float64 `c:"minRunTime,omitempty" p:"minRunTime" v:"float|min:0"`
	MaxRunTime                    *float64 `c:"maxRunTime,omitempty" p:"maxRunTime" v:"float|min:0|gte:MinRunTime"`
}
