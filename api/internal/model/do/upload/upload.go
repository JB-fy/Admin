// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package upload

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Upload is the golang structure of table upload for DAO operations like Where/Data.
type Upload struct {
	g.Meta       `orm:"table:upload, do:true"`
	CreatedAt    *gtime.Time // 创建时间
	UpdatedAt    *gtime.Time // 更新时间
	UploadId     any         // 上传ID
	UploadType   any         // 类型：0本地 1阿里云OSS
	UploadConfig any         // 配置。JSON格式，根据类型设置
	IsDefault    any         // 默认：0否 1是
	Remark       any         // 备注
}
