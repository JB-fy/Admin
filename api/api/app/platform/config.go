package platform

import (
	"api/api"

	"github.com/gogf/gf/v2/frame/g"
)

/*--------获取 开始--------*/
type ConfigGetReq struct {
	g.Meta `path:"/config/get" method:"post" tags:"APP/配置" sm:"获取"`
	api.CommonAppHeaderReq
	ConfigKeyArr *[]string `c:"config_key_arr,omitempty" json:"config_key_arr" v:"required|distinct|foreach|in:hot_search,user_agreement,privacy_agreement" dc:"配置项Key列表。传值参考默认返回的字段"`
}

type ConfigGetRes struct {
	Config Config `json:"config" dc:"配置列表"`
}

type Config struct {
	HotSearch        *[]string `json:"hot_search,omitempty" dc:"热门搜索"`
	UserAgreement    *string   `json:"user_agreement,omitempty" dc:"用户协议"`
	PrivacyAgreement *string   `json:"privacy_agreement,omitempty" dc:"隐私协议"`
}

/*--------获取 结束--------*/
