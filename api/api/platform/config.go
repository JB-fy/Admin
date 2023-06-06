package api

import (
	apiCommon "api/api"
)

type ConfigListReq struct {
	apiCommon.CommonListReq
	Filter ConfigListFilterReq `p:"filter"`
}

/* type Config struct {
    ConfigId    uint        `json:"configId"    `// 配置ID
    ConfigKey   string      `json:"configKey"   `// 配置项Key
    ConfigValue string      `json:"configValue" `// 配置项值（设置大点。以后可能需要保存富文本内容，如公司简介或协议等等）
    UpdateTime  *gtime.Time `json:"updateTime"  `// 更新时间
    CreateTime  *gtime.Time `json:"createTime"  `// 创建时间
} */

type ConfigListFilterReq struct {
	apiCommon.CommonListFilterReq `c:",omitempty"`
	ConfigId                      *uint  `c:"configId,omitempty" p:"configId" v:"min:1"`
	ConfigName                    string `c:"configName,omitempty" p:"configName" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$"`
	IsStop                        *uint  `c:"isStop,omitempty" p:"isStop" v:"in:0,1"`
}

type ConfigInfoReq struct {
	apiCommon.CommonInfoReq
}

type ConfigCreateReq struct {
	ConfigName *string `c:"configName,omitempty" p:"configName" v:"required|length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$"`
	IsStop     *uint   `c:"isStop,omitempty" p:"isStop" v:"in:0,1"`
}

type ConfigUpdateReq struct {
	apiCommon.CommonUpdateDeleteIdArrReq `c:",omitempty"`
	ConfigName                           *string `c:"configName,omitempty" p:"configName" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$"`
	IsStop                               *uint   `c:"isStop,omitempty" p:"isStop" v:"in:0,1"`
}

type ConfigDeleteReq struct {
	apiCommon.CommonUpdateDeleteIdArrReq
}
