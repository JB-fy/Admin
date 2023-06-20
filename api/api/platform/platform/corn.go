package platform

import (
	apiCommon "api/api"
)

type CornListReq struct {
	apiCommon.CommonListReq
	Filter CornListFilterReq `p:"filter"`
}

type CornListFilterReq struct {
	apiCommon.CommonListFilterReq `c:",omitempty"`
	CornId                        *uint  `c:"cornId,omitempty" p:"cornId" v:"integer|min:1"`                                   // 定时器ID
	CornCode                      string `c:"cornCode,omitempty" p:"cornCode" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$"` // 标识
	CornName                      string `c:"cornName,omitempty" p:"cornName" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$"` // 名称
	IsStop                        *uint  `c:"isStop,omitempty" p:"isStop" v:"integer|in:0,1"`                                  // 是否停用：0否 1是
}

type CornInfoReq struct {
	apiCommon.CommonInfoReq
}

type CornCreateReq struct {
	CornCode    *string `c:"cornCode,omitempty" p:"cornCode" v:"required|length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$"` // 标识
	CornName    *string `c:"cornName,omitempty" p:"cornName" v:"required|length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$"` // 名称
	CornPattern *string `c:"cornPattern,omitempty" p:"cornPattern" v:"required|length:1,30"`                           // 表达式
	Remark      *string `c:"remark,omitempty" p:"remark" v:"length:1,120"`                                             // 备注
	IsStop      *uint   `c:"isStop,omitempty" p:"isStop" v:"integer|in:0,1"`                                           // 是否停用：0否 1是
}

type CornUpdateReq struct {
	apiCommon.CommonUpdateDeleteIdArrReq `c:",omitempty"`
	CornCode                             *string `c:"cornCode,omitempty" p:"cornCode" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$"` // 标识
	CornName                             *string `c:"cornName,omitempty" p:"cornName" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$"` // 名称
	CornPattern                          *string `c:"cornPattern,omitempty" p:"cornPattern" v:"length:1,30"`                           // 表达式
	Remark                               *string `c:"remark,omitempty" p:"remark" v:"length:1,120"`                                    // 备注
	IsStop                               *uint   `c:"isStop,omitempty" p:"isStop" v:"integer|in:0,1"`                                  // 是否停用：0否 1是
}

type CornDeleteReq struct {
	apiCommon.CommonUpdateDeleteIdArrReq
}
