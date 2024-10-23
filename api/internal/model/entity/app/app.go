// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package app

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// App is the golang structure for table app.
type App struct {
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"    ` // 创建时间
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"    ` // 更新时间
	IsStop      uint        `json:"isStop"      orm:"is_stop"       ` // 停用：0否 1是
	AppId       uint        `json:"appId"       orm:"app_id"        ` // APPID
	AppType     uint        `json:"appType"     orm:"app_type"      ` // 类型：0安卓 1苹果
	PackageName string      `json:"packageName" orm:"package_name"  ` // 包名
	PackageFile string      `json:"packageFile" orm:"package_file"  ` // 安装包
	VerNo       uint        `json:"verNo"       orm:"ver_no"        ` // 版本号
	VerName     string      `json:"verName"     orm:"ver_name"      ` // 版本名称
	VerIntro    string      `json:"verIntro"    orm:"ver_intro"     ` // 版本介绍
	ExtraConfig string      `json:"extraConfig" orm:"extra_config"  ` // 额外配置
	Remark      string      `json:"remark"      orm:"remark"        ` // 备注
	IsForcePrev uint        `json:"isForcePrev" orm:"is_force_prev" ` // 强制更新：0否 1是。注意：只根据前一个版本来设置，与更早之前的版本无关
}
