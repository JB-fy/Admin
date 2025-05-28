package api

import (
	"github.com/gogf/gf/v2/frame/g"
)

// 共用详情。list,info,tree等接口返回时用，但返回默认字段有差异。可根据需要在controller对应的defaultField中补充所需字段
type AppPkgInfo struct {
	PkgName          *string `json:"pkg_name,omitempty" dc:"包名"`
	PkgFile          *string `json:"pkg_file,omitempty" dc:"安装包"`
	VerNo            *uint   `json:"ver_no,omitempty" dc:"版本号"`
	VerName          *string `json:"ver_name,omitempty" dc:"版本名称"`
	VerIntro         *string `json:"ver_intro,omitempty" dc:"版本介绍"`
	IsForce          *uint   `json:"is_force,omitempty" dc:"强制更新：0否 1是"`
	DownloadUrlToApp *string `json:"download_url_to_app,omitempty" dc:"下载地址。APP内部用"`
	DownloadUrlToH5  *string `json:"download_url_to_h5,omitempty" dc:"下载地址。H5用"`
}

/*--------获取下载地址 开始--------*/
type AppPkgInfoReq struct {
	g.Meta         `path:"/pkg/info" method:"get,post" tags:"APP安装包" sm:"获取下载地址"`
	AppId          *string `json:"app_id,omitempty" v:"required|max-length:15" dc:"APPID"`
	PkgType        *uint   `json:"pkg_type,omitempty" v:"required|in:0,1,2" dc:"类型：0安卓 1苹果 2PC"`
	VerNoOfCurrent *uint   `json:"ver_no_of_current,omitempty" v:"between:0,4294967295" dc:"客户端APP当前版本号。客户端APP判断是否需要更新时使用，不需要更新时返回数据info为空"`
}

type AppPkgInfoRes struct {
	Info AppPkgInfo `json:"info" dc:"详情。当不需要更新时，返回空对象"`
}

/*--------获取下载地址 结束--------*/
