// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package app

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// App is the golang structure of table app for DAO operations like Where/Data.
type App struct {
	g.Meta      `orm:"table:app, do:true"`
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 更新时间
	IsStop      interface{} // 停用：0否 1是
	AppId       interface{} // APPID
	NameType    interface{} // 名称：0APP。有两种以上APP时自行扩展
	AppType     interface{} // 类型：0安卓 1苹果 2PC
	PackageName interface{} // 包名
	PackageFile interface{} // 安装包
	VerNo       interface{} // 版本号
	VerName     interface{} // 版本名称
	VerIntro    interface{} // 版本介绍
	ExtraConfig interface{} // 额外配置
	Remark      interface{} // 备注
	IsForcePrev interface{} // 强制更新：0否 1是。注意：只根据前一个版本来设置，与更早之前的版本无关
}
