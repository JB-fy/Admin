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
	UserAgreement    *string `json:"userAgreement,omitempty" dc:"用户协议"`
	PrivacyAgreement *string `json:"privacyAgreement,omitempty" dc:"隐私协议"`

	UploadType               *string `json:"uploadType,omitempty" dc:"上传方式"`
	LocalUploadUrl           *string `json:"localUploadUrl,omitempty" dc:"本地-上传地址"`
	LocalUploadSignKey       *string `json:"localUploadSignKey,omitempty" dc:"本地-密钥"`
	LocalUploadFileSaveDir   *string `json:"localUploadFileSaveDir,omitempty" dc:"本地-文件保存目录"`
	LocalUploadFileUrlPrefix *string `json:"localUploadFileUrlPrefix,omitempty" dc:"本地-文件地址前缀"`
	AliyunOssHost            *string `json:"aliyunOssHost,omitempty" dc:"阿里云OSS-域名"`
	AliyunOssBucket          *string `json:"aliyunOssBucket,omitempty" dc:"阿里云OSS-Bucket"`
	AliyunOssAccessKeyId     *string `json:"aliyunOssAccessKeyId,omitempty" dc:"阿里云OSS-AccessKeyId"`
	AliyunOssAccessKeySecret *string `json:"aliyunOssAccessKeySecret,omitempty" dc:"阿里云OSS-AccessKeySecret"`
	AliyunOssRoleArn         *string `json:"aliyunOssRoleArn,omitempty" dc:"阿里云OSS-RoleArn"`
	AliyunOssCallbackUrl     *string `json:"aliyunOssCallbackUrl,omitempty" dc:"阿里云OSS-回调地址"`

	SmsType                  *string `json:"smsType,omitempty" dc:"短信方式"`
	AliyunSmsAccessKeyId     *string `json:"aliyunSmsAccessKeyId,omitempty" dc:"阿里云短信-AccessKeyId"`
	AliyunSmsAccessKeySecret *string `json:"aliyunSmsAccessKeySecret,omitempty" dc:"阿里云短信-AccessKeySecret"`
	AliyunSmsSignName        *string `json:"aliyunSmsSignName,omitempty" dc:"阿里云短信-签名"`
	AliyunSmsTemplateCode    *string `json:"aliyunSmsTemplateCode,omitempty" dc:"阿里云短信-模板标识"`
}

/*--------获取 结束--------*/

/*--------保存 开始--------*/
type ConfigSaveReq struct {
	g.Meta `path:"/config/save" method:"post" tags:"平台后台/配置" sm:"保存"`

	UserAgreement    *string `json:"userAgreement,omitempty" v:"" dc:"用户协议"`
	PrivacyAgreement *string `json:"privacyAgreement,omitempty" v:"" dc:"隐私协议"`

	UploadType               *string `json:"uploadType,omitempty" v:"in:local,aliyunOss" dc:"上传方式"`
	LocalUploadUrl           *string `json:"localUploadUrl,omitempty" v:"url" dc:"本地-上传地址"`
	LocalUploadSignKey       *string `json:"localUploadSignKey,omitempty" v:"" dc:"本地-密钥"`
	LocalUploadFileSaveDir   *string `json:"localUploadFileSaveDir,omitempty" v:"" dc:"本地-文件保存目录"`
	LocalUploadFileUrlPrefix *string `json:"localUploadFileUrlPrefix,omitempty" v:"url" dc:"本地-文件地址前缀"`
	AliyunOssHost            *string `json:"aliyunOssHost,omitempty" v:"url" dc:"阿里云OSS-域名"`
	AliyunOssBucket          *string `json:"aliyunOssBucket,omitempty" v:"" dc:"阿里云OSS-Bucket"`
	AliyunOssAccessKeyId     *string `json:"aliyunOssAccessKeyId,omitempty" v:"regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"阿里云OSS-AccessKeyId"`
	AliyunOssAccessKeySecret *string `json:"aliyunOssAccessKeySecret,omitempty" v:"regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"阿里云OSS-AccessKeySecret"`
	AliyunOssRoleArn         *string `json:"aliyunOssRoleArn,omitempty" v:"" dc:"阿里云OSS-RoleArn"`
	AliyunOssCallbackUrl     *string `json:"aliyunOssCallbackUrl,omitempty" v:"url" dc:"阿里云OSS-回调地址"`

	SmsType                  *string `json:"smsType,omitempty" v:"in:aliyunSms" dc:"短信方式"`
	AliyunSmsAccessKeyId     *string `json:"aliyunSmsAccessKeyId,omitempty" v:"regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"阿里云短信-AccessKeyId"`
	AliyunSmsAccessKeySecret *string `json:"aliyunSmsAccessKeySecret,omitempty" v:"regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"阿里云短信-AccessKeySecret"`
	AliyunSmsSignName        *string `json:"aliyunSmsSignName,omitempty" v:"" dc:"阿里云短信-签名"`
	AliyunSmsTemplateCode    *string `json:"aliyunSmsTemplateCode,omitempty" v:"" dc:"阿里云短信-模板标识"`
}

/*--------保存 结束--------*/
