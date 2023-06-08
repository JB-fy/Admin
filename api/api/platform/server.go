package api

import (
	apiCommon "api/api"
)

type ServerListReq struct {
	apiCommon.CommonListReq
	Filter ServerListFilterReq `p:"filter"`
}

type ServerListFilterReq struct {
	apiCommon.CommonListFilterReq `c:",omitempty"`
	ServerId                      *uint  `c:"serverId,omitempty" p:"serverId" v:"integer|min:1"`
	NetworkIp                     string `c:"networkIp,omitempty" p:"networkIp" v:"ip"`
	LocalIp                       string `c:"localIp,omitempty" p:"localIp" v:"ip"`
}
