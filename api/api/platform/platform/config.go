package api

import (
	"github.com/gogf/gf/v2/frame/g"
)

/*--------获取 开始--------*/
type ConfigGetReq struct {
	g.Meta       `path:"/config/get" method:"post" tags:"平台后台/配置" sm:"获取"`
	ConfigKeyArr *[]string `json:"configKeyArr,omitempty" v:"required|distinct|foreach|min-length:1" dc:"配置项Key列表。传值参考默认返回的字段"`
}

type ConfigGetRes struct {
	Config Config `json:"config" dc:"配置列表"`
}

type Config struct {
	AliyunOssHost            *string `json:"aliyunOssHost,omitempty" dc:"阿里云存储-域名"`
	AliyunOssBucket          *string `json:"aliyunOssBucket,omitempty" dc:"阿里云存储-Bucket"`
	AliyunOssAccessKeyId     *string `json:"aliyunOssAccessKeyId,omitempty" dc:"阿里云存储-AccessKeyId"`
	AliyunOssAccessKeySecret *string `json:"aliyunOssAccessKeySecret,omitempty" dc:"阿里云存储-AccessKeySecret"`
	AliyunOssRoleArn         *string `json:"aliyunOssRoleArn,omitempty" dc:"阿里云存储-RoleArn"`
	AliyunOssCallbackUrl     *string `json:"aliyunOssCallbackUrl,omitempty" v:"" dc:"阿里云存储-回调地址"`
}

/*--------获取 结束--------*/

/*--------保存 开始--------*/
type ConfigSaveReq struct {
	g.Meta                   `path:"/config/save" method:"post" tags:"平台后台/配置" sm:"保存"`
	AliyunOssHost            *string `json:"aliyunOssHost,omitempty" v:"url" dc:"阿里云存储-域名"`
	AliyunOssBucket          *string `json:"aliyunOssBucket,omitempty" v:"regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"阿里云存储-Bucket"`
	AliyunOssAccessKeyId     *string `json:"aliyunOssAccessKeyId,omitempty" v:"regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"阿里云存储-AccessKeyId"`
	AliyunOssAccessKeySecret *string `json:"aliyunOssAccessKeySecret,omitempty" v:"regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"阿里云存储-AccessKeySecret"`
	AliyunOssRoleArn         *string `json:"aliyunOssRoleArn,omitempty" v:"" dc:"阿里云存储-RoleArn"`
	AliyunOssCallbackUrl     *string `json:"aliyunOssCallbackUrl,omitempty" v:"" dc:"阿里云存储-回调地址"`
}

/*--------保存 结束--------*/
