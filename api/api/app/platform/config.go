package api

import (
	"github.com/gogf/gf/v2/frame/g"
)

/*--------获取 开始--------*/
type ConfigGetReq struct {
	g.Meta       `path:"/config/get" method:"post" tags:"APP/配置" sm:"获取"`
	ConfigKeyArr *[]string `c:"config_key_arr,omitempty" json:"config_key_arr" v:"required|distinct|foreach|in:hotSearch,userAgreement,privacyAgreement" dc:"配置项Key列表。传值参考默认返回的字段"`
}

type ConfigGetRes struct {
	Config Config `json:"config" dc:"配置列表"`
}

type Config struct {
	HotSearch        *[]string `json:"hotSearch,omitempty" dc:"热门搜索"`
	UserAgreement    *string   `json:"userAgreement,omitempty" dc:"用户协议"`
	PrivacyAgreement *string   `json:"privacyAgreement,omitempty" dc:"隐私协议"`
}

/*--------获取 结束--------*/
