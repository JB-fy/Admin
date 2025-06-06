// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package app

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Pkg is the golang structure for table pkg.
type Pkg struct {
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"    ` // 创建时间
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"    ` // 更新时间
	IsStop      uint        `json:"isStop"      orm:"is_stop"       ` // 停用：0否 1是
	PkgId       uint        `json:"pkgId"       orm:"pkg_id"        ` // 安装包ID
	AppId       string      `json:"appId"       orm:"app_id"        ` // APPID
	PkgType     uint        `json:"pkgType"     orm:"pkg_type"      ` // 类型：0安卓 1苹果 2PC
	PkgName     string      `json:"pkgName"     orm:"pkg_name"      ` // 包名
	PkgFile     string      `json:"pkgFile"     orm:"pkg_file"      ` // 安装包
	VerNo       uint        `json:"verNo"       orm:"ver_no"        ` // 版本号
	VerName     string      `json:"verName"     orm:"ver_name"      ` // 版本名称
	VerIntro    string      `json:"verIntro"    orm:"ver_intro"     ` // 版本介绍
	ExtraConfig string      `json:"extraConfig" orm:"extra_config"  ` // 额外配置。JSON格式，需要时设置
	Remark      string      `json:"remark"      orm:"remark"        ` // 备注
	IsForcePrev uint        `json:"isForcePrev" orm:"is_force_prev" ` // 强制更新：0否 1是。注意：只根据前一个版本来设置，与更早之前的版本无关
}
