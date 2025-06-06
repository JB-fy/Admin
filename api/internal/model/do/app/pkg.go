// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package app

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Pkg is the golang structure of table app_pkg for DAO operations like Where/Data.
type Pkg struct {
	g.Meta      `orm:"table:app_pkg, do:true"`
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 更新时间
	IsStop      interface{} // 停用：0否 1是
	PkgId       interface{} // 安装包ID
	AppId       interface{} // APPID
	PkgType     interface{} // 类型：0安卓 1苹果 2PC
	PkgName     interface{} // 包名
	PkgFile     interface{} // 安装包
	VerNo       interface{} // 版本号
	VerName     interface{} // 版本名称
	VerIntro    interface{} // 版本介绍
	ExtraConfig interface{} // 额外配置。JSON格式，需要时设置
	Remark      interface{} // 备注
	IsForcePrev interface{} // 强制更新：0否 1是。注意：只根据前一个版本来设置，与更早之前的版本无关
}
