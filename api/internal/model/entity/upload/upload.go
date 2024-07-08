// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Upload is the golang structure for table upload.
type Upload struct {
	CreatedAt    *gtime.Time `json:"createdAt"    orm:"created_at"    ` // 创建时间
	UpdatedAt    *gtime.Time `json:"updatedAt"    orm:"updated_at"    ` // 更新时间
	IsStop       uint        `json:"isStop"       orm:"is_stop"       ` // 停用：0否 1是
	UploadId     uint        `json:"uploadId"     orm:"upload_id"     ` // 上传ID
	UploadType   uint        `json:"uploadType"   orm:"upload_type"   ` // 类型：0本地 1阿里云OSS
	UploadConfig string      `json:"uploadConfig" orm:"upload_config" ` // 配置。根据upload_type类型设置
	IsDefault    uint        `json:"isDefault"    orm:"is_default"    ` // 默认：0否 1是
	Remark       string      `json:"remark"       orm:"remark"        ` // 备注
}
