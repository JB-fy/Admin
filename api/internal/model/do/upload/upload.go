// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Upload is the golang structure of table upload for DAO operations like Where/Data.
type Upload struct {
	g.Meta       `orm:"table:upload, do:true"`
	CreatedAt    *gtime.Time // 创建时间
	UpdatedAt    *gtime.Time // 更新时间
	UploadId     interface{} // 上传ID
	UploadType   interface{} // 类型：0本地 1阿里云OSS
	UploadConfig interface{} // 配置。根据upload_type类型设置
	IsDefault    interface{} // 默认：0否 1是
	Remark       interface{} // 备注
}
