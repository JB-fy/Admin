package org

import (
	"github.com/gogf/gf/v2/frame/g"
)

/*--------获取 开始--------*/
type ConfigGetReq struct {
	g.Meta       `path:"/config/get" method:"post" tags:"机构后台/配置中心/机构配置" sm:"获取"`
	ConfigKeyArr *[]string `json:"config_key_arr,omitempty" v:"required|distinct|foreach|length:1,30" dc:"配置键列表。传值参考默认返回的字段"`
}

type ConfigGetRes struct {
	Config Config `json:"config" dc:"配置列表"`
}

type Config struct {
	HotSearch *[]string `json:"hotSearch,omitempty" dc:"热门搜索"`
}

/*--------获取 结束--------*/

/*--------保存 开始--------*/
type ConfigSaveReq struct {
	g.Meta `path:"/config/save" method:"post" tags:"机构后台/配置中心/机构配置" sm:"保存"`

	HotSearch *[]string `json:"hotSearch,omitempty" v:"distinct|foreach|min-length:1" dc:"热门搜索"`
}

/*--------保存 结束--------*/
