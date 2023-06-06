package api

import (
	apiCommon "api/api"
)

type ServerListReq struct {
	apiCommon.CommonListReq
	Filter ServerListFilterReq `p:"filter"`
}

/* type Server struct {
    ServerId   uint        `json:"serverId"   `// 服务器ID
    NetworkIp  string      `json:"networkIp"  `// 公网IP
    LocalIp    string      `json:"localIp"    `// 内网IP
    UpdateTime *gtime.Time `json:"updateTime" `// 更新时间
    CreateTime *gtime.Time `json:"createTime" `// 创建时间
} */

type ServerListFilterReq struct {
	apiCommon.CommonListFilterReq `c:",omitempty"`
	ServerId                      *uint  `c:"serverId,omitempty" p:"serverId" v:"min:1"`
	ServerName                    string `c:"serverName,omitempty" p:"serverName" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$"`
	IsStop                        *uint  `c:"isStop,omitempty" p:"isStop" v:"in:0,1"`
}

type ServerInfoReq struct {
	apiCommon.CommonInfoReq
}

type ServerCreateReq struct {
	ServerName *string `c:"serverName,omitempty" p:"serverName" v:"required|length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$"`
	IsStop     *uint   `c:"isStop,omitempty" p:"isStop" v:"in:0,1"`
}

type ServerUpdateReq struct {
	apiCommon.CommonUpdateDeleteIdArrReq `c:",omitempty"`
	ServerName                           *string `c:"serverName,omitempty" p:"serverName" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$"`
	IsStop                               *uint   `c:"isStop,omitempty" p:"isStop" v:"in:0,1"`
}

type ServerDeleteReq struct {
	apiCommon.CommonUpdateDeleteIdArrReq
}
