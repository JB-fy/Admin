package app

import (
	"github.com/gogf/gf/v2/frame/g"
)

// 共用详情。list,info,tree等接口返回时用，但返回默认字段有差异。可根据需要在controller对应的defaultField中补充所需字段
type AppInfo struct {
	PackageName      *string `json:"package_name,omitempty" dc:"包名"`
	PackageFile      *string `json:"package_file,omitempty" dc:"安装包"`
	VerNo            *uint   `json:"ver_no,omitempty" dc:"版本号"`
	VerName          *string `json:"ver_name,omitempty" dc:"版本名称"`
	VerIntro         *string `json:"ver_intro,omitempty" dc:"版本介绍"`
	IsForce          *uint   `json:"is_force,omitempty" dc:"强制更新：0否 1是"`
	DownloadUrlToApp *string `json:"download_url_to_app,omitempty" dc:"下载地址。APP内部用"`
	DownloadUrlToH5  *string `json:"download_url_to_h5,omitempty" dc:"下载地址。H5用"`
}

/*--------详情 开始--------*/
type AppInfoReq struct {
	g.Meta       `path:"/app/info" method:"post" tags:"APP/APP" sm:"更新和下载"`
	AppType      *uint `json:"app_type,omitempty" v:"required|in:0,1,2" dc:"类型：0安卓 1苹果 2PC"`
	CurrentVerNo *uint `json:"current_ver_no,omitempty" v:"between:0,4294967295" dc:"当前版本号。作用：判断APP是否需要更新，不需要返回的info为空"`
}

type AppInfoRes struct {
	Info AppInfo `json:"info" dc:"详情。当不需要更新时，返回空对象"`
}

/*--------详情 结束--------*/
