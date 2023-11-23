package api

import (
	"github.com/gogf/gf/v2/frame/g"
)

/*--------获取 开始--------*/
type ConfigGetReq struct {
	g.Meta       `path:"/config/get" method:"post" tags:"平台后台/配置中心/平台配置" sm:"获取"`
	ConfigKeyArr *[]string `json:"configKeyArr,omitempty" v:"required|distinct|foreach|min-length:1" dc:"配置Key列表。传值参考默认返回的字段"`
}

type ConfigGetRes struct {
	Config Config `json:"config" dc:"配置列表"`
}

type Config struct {
	HotSearch        *[]string `json:"hotSearch,omitempty" dc:"热门搜索"`
	UserAgreement    *string   `json:"userAgreement,omitempty" dc:"用户协议"`
	PrivacyAgreement *string   `json:"privacyAgreement,omitempty" dc:"隐私协议"`

	PackageUrlOfAndroid    *string `json:"packageUrlOfAndroid,omitempty" dc:"安装包(安卓)"`
	PackageSizeOfAndroid   *uint   `json:"packageSizeOfAndroid,omitempty" dc:"包大小(安卓)"`
	PackageNameOfAndroid   *string `json:"packageNameOfAndroid,omitempty" dc:"包名(安卓)"`
	IsForceUpdateOfAndroid *uint   `json:"isForceUpdateOfAndroid,omitempty" dc:"强制更新(安卓)"`
	VersionNumberOfAndroid *uint   `json:"versionNumberOfAndroid,omitempty" dc:"版本号(安卓)"`
	VersionNameOfAndroid   *string `json:"versionNameOfAndroid,omitempty" dc:"版本名称(安卓)"`
	VersionIntroOfAndroid  *string `json:"versionIntroOfAndroid,omitempty" dc:"版本介绍(安卓)"`

	PackageUrlOfIos    *string `json:"packageUrlOfIos,omitempty" dc:"安装包(苹果)"`
	PackageSizeOfIos   *uint   `json:"packageSizeOfIos,omitempty" dc:"包大小(苹果)"`
	PackageNameOfIos   *string `json:"packageNameOfIos,omitempty" dc:"包名(苹果)"`
	IsForceUpdateOfIos *uint   `json:"isForceUpdateOfIos,omitempty" dc:"强制更新(苹果)"`
	VersionNumberOfIos *uint   `json:"versionNumberOfIos,omitempty" dc:"版本号(苹果)"`
	VersionNameOfIos   *string `json:"versionNameOfIos,omitempty" dc:"版本名称(苹果)"`
	VersionIntroOfIos  *string `json:"versionIntroOfIos,omitempty" dc:"版本介绍(苹果)"`
	PlistUrlOfIos      *string `json:"plistUrlOfIos,omitempty" dc:"plist文件(苹果)"`

	UploadType               *string `json:"uploadType,omitempty" dc:"上传方式"`
	LocalUploadUrl           *string `json:"localUploadUrl,omitempty" dc:"本地-上传地址"`
	LocalUploadSignKey       *string `json:"localUploadSignKey,omitempty" dc:"本地-密钥"`
	LocalUploadFileSaveDir   *string `json:"localUploadFileSaveDir,omitempty" dc:"本地-文件保存目录"`
	LocalUploadFileUrlPrefix *string `json:"localUploadFileUrlPrefix,omitempty" dc:"本地-文件地址前缀"`
	AliyunOssHost            *string `json:"aliyunOssHost,omitempty" dc:"阿里云OSS-域名"`
	AliyunOssBucket          *string `json:"aliyunOssBucket,omitempty" dc:"阿里云OSS-Bucket"`
	AliyunOssAccessKeyId     *string `json:"aliyunOssAccessKeyId,omitempty" dc:"阿里云OSS-AccessKeyId"`
	AliyunOssAccessKeySecret *string `json:"aliyunOssAccessKeySecret,omitempty" dc:"阿里云OSS-AccessKeySecret"`
	AliyunOssCallbackUrl     *string `json:"aliyunOssCallbackUrl,omitempty" dc:"阿里云OSS-回调地址"`
	AliyunOssEndpoint        *string `json:"aliyunOssEndpoint,omitempty" dc:"阿里云OSS-Endpoint"`
	AliyunOssRoleArn         *string `json:"aliyunOssRoleArn,omitempty" dc:"阿里云OSS-RoleArn"`

	SmsType                  *string `json:"smsType,omitempty" dc:"短信方式"`
	AliyunSmsAccessKeyId     *string `json:"aliyunSmsAccessKeyId,omitempty" dc:"阿里云SMS-AccessKeyId"`
	AliyunSmsAccessKeySecret *string `json:"aliyunSmsAccessKeySecret,omitempty" dc:"阿里云SMS-AccessKeySecret"`
	AliyunSmsEndpoint        *string `json:"aliyunSmsEndpoint,omitempty" dc:"阿里云SMS-Endpoint"`
	AliyunSmsSignName        *string `json:"aliyunSmsSignName,omitempty" dc:"阿里云SMS-签名"`
	AliyunSmsTemplateCode    *string `json:"aliyunSmsTemplateCode,omitempty" dc:"阿里云SMS-模板标识"`

	IdCardType          *string `json:"idCardType,omitempty" dc:"实名认证方式"`
	AliyunIdCardHost    *string `json:"aliyunIdCardHost,omitempty" dc:"阿里云IdCard-域名"`
	AliyunIdCardPath    *string `json:"aliyunIdCardPath,omitempty" dc:"阿里云IdCard-请求路径"`
	AliyunIdCardAppcode *string `json:"aliyunIdCardAppcode,omitempty" dc:"阿里云IdCard-Appcode"`

	PushType                 *string `json:"pushType,omitempty" dc:"推送方式"`
	TxTpnsHost               *string `json:"txTpnsHost,omitempty" dc:"腾讯移动推送-域名"`
	TxTpnsAccessIDOfAndroid  *string `json:"txTpnsAccessIDOfAndroid,omitempty" dc:"腾讯移动推送-AccessID(安卓)"`
	TxTpnsSecretKeyOfAndroid *string `json:"txTpnsSecretKeyOfAndroid,omitempty" dc:"腾讯移动推送-SecretKey(安卓)"`
	TxTpnsAccessIDOfIos      *string `json:"txTpnsAccessIDOfIos,omitempty" dc:"腾讯移动推送-AccessID(苹果)"`
	TxTpnsSecretKeyOfIos     *string `json:"txTpnsSecretKeyOfIos,omitempty" dc:"腾讯移动推送-SecretKey(苹果)"`
	TxTpnsAccessIDOfMacOS    *string `json:"txTpnsAccessIDOfMacOS,omitempty" dc:"腾讯移动推送-AccessID(苹果电脑)"`
	TxTpnsSecretKeyOfMacOS   *string `json:"txTpnsSecretKeyOfMacOS,omitempty" dc:"腾讯移动推送-SecretKey(苹果电脑)"`

	VodType                  *string `json:"vodType,omitempty" dc:"视频点播方式"`
	AliyunVodAccessKeyId     *string `json:"aliyunVodAccessKeyId,omitempty" dc:"阿里云VOD-AccessKeyId"`
	AliyunVodAccessKeySecret *string `json:"aliyunVodAccessKeySecret,omitempty" dc:"阿里云VOD-AccessKeySecret"`
	AliyunVodEndpoint        *string `json:"aliyunVodEndpoint,omitempty" dc:"阿里云VOD-Endpoint"`
	AliyunVodRoleArn         *string `json:"aliyunVodRoleArn,omitempty" dc:"阿里云VOD-RoleArn"`
}

/*--------获取 结束--------*/

/*--------保存 开始--------*/
type ConfigSaveReq struct {
	g.Meta `path:"/config/save" method:"post" tags:"平台后台/配置中心/平台配置" sm:"保存"`

	HotSearch        *[]string `json:"hotSearch,omitempty" v:"distinct|foreach|min-length:1" dc:"热门搜索"`
	UserAgreement    *string   `json:"userAgreement,omitempty" v:"" dc:"用户协议"`
	PrivacyAgreement *string   `json:"privacyAgreement,omitempty" v:"" dc:"隐私协议"`

	PackageUrlOfAndroid    *string `json:"packageUrlOfAndroid,omitempty" v:"url" dc:"安装包(安卓)"`
	PackageSizeOfAndroid   *uint   `json:"packageSizeOfAndroid,omitempty" v:"" dc:"包大小(安卓)"`
	PackageNameOfAndroid   *string `json:"packageNameOfAndroid,omitempty" v:"" dc:"包名(安卓)"`
	IsForceUpdateOfAndroid *uint   `json:"isForceUpdateOfAndroid,omitempty" v:"integer|in:0,1" dc:"强制更新(安卓)"`
	VersionNumberOfAndroid *uint   `json:"versionNumberOfAndroid,omitempty" v:"" dc:"版本号(安卓)"`
	VersionNameOfAndroid   *string `json:"versionNameOfAndroid,omitempty" v:"" dc:"版本名称(安卓)"`
	VersionIntroOfAndroid  *string `json:"versionIntroOfAndroid,omitempty" v:"" dc:"版本介绍(安卓)"`

	PackageUrlOfIos    *string `json:"packageUrlOfIos,omitempty" v:"url" dc:"安装包(苹果)"`
	PackageSizeOfIos   *uint   `json:"packageSizeOfIos,omitempty" v:"" dc:"包大小(苹果)"`
	PackageNameOfIos   *string `json:"packageNameOfIos,omitempty" v:"" dc:"包名(苹果)"`
	IsForceUpdateOfIos *uint   `json:"isForceUpdateOfIos,omitempty" v:"integer|in:0,1" dc:"强制更新(苹果)"`
	VersionNumberOfIos *uint   `json:"versionNumberOfIos,omitempty" v:"" dc:"版本号(苹果)"`
	VersionNameOfIos   *string `json:"versionNameOfIos,omitempty" v:"" dc:"版本名称(苹果)"`
	VersionIntroOfIos  *string `json:"versionIntroOfIos,omitempty" v:"" dc:"版本介绍(苹果)"`
	PlistUrlOfIos      *string `json:"plistUrlOfIos,omitempty" v:"url" dc:"plist文件(苹果)"`

	UploadType               *string `json:"uploadType,omitempty" v:"in:local,aliyunOss" dc:"上传方式"`
	LocalUploadUrl           *string `json:"localUploadUrl,omitempty" v:"url" dc:"本地-上传地址"`
	LocalUploadSignKey       *string `json:"localUploadSignKey,omitempty" v:"" dc:"本地-密钥"`
	LocalUploadFileSaveDir   *string `json:"localUploadFileSaveDir,omitempty" v:"" dc:"本地-文件保存目录"`
	LocalUploadFileUrlPrefix *string `json:"localUploadFileUrlPrefix,omitempty" v:"url" dc:"本地-文件地址前缀"`
	AliyunOssHost            *string `json:"aliyunOssHost,omitempty" v:"url" dc:"阿里云OSS-域名"`
	AliyunOssBucket          *string `json:"aliyunOssBucket,omitempty" v:"" dc:"阿里云OSS-Bucket"`
	AliyunOssAccessKeyId     *string `json:"aliyunOssAccessKeyId,omitempty" v:"regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"阿里云OSS-AccessKeyId"`
	AliyunOssAccessKeySecret *string `json:"aliyunOssAccessKeySecret,omitempty" v:"regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"阿里云OSS-AccessKeySecret"`
	AliyunOssCallbackUrl     *string `json:"aliyunOssCallbackUrl,omitempty" v:"url" dc:"阿里云OSS-回调地址"`
	AliyunOssEndpoint        *string `json:"aliyunOssEndpoint,omitempty" v:"" dc:"阿里云OSS-Endpoint"`
	AliyunOssRoleArn         *string `json:"aliyunOssRoleArn,omitempty" v:"" dc:"阿里云OSS-RoleArn"`

	SmsType                  *string `json:"smsType,omitempty" v:"in:aliyunSms" dc:"短信方式"`
	AliyunSmsAccessKeyId     *string `json:"aliyunSmsAccessKeyId,omitempty" v:"regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"阿里云SMS-AccessKeyId"`
	AliyunSmsAccessKeySecret *string `json:"aliyunSmsAccessKeySecret,omitempty" v:"regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"阿里云SMS-AccessKeySecret"`
	AliyunSmsEndpoint        *string `json:"aliyunSmsEndpoint,omitempty" v:"" dc:"阿里云SMS-Endpoint"`
	AliyunSmsSignName        *string `json:"aliyunSmsSignName,omitempty" v:"" dc:"阿里云SMS-签名"`
	AliyunSmsTemplateCode    *string `json:"aliyunSmsTemplateCode,omitempty" v:"" dc:"阿里云SMS-模板标识"`

	IdCardType          *string `json:"idCardType,omitempty" v:"in:aliyunIdCard" dc:"实名认证方式"`
	AliyunIdCardHost    *string `json:"aliyunIdCardHost,omitempty" v:"url" dc:"阿里云IdCard-域名"`
	AliyunIdCardPath    *string `json:"aliyunIdCardPath,omitempty" v:"" dc:"阿里云IdCard-请求路径"`
	AliyunIdCardAppcode *string `json:"aliyunIdCardAppcode,omitempty" v:"" dc:"阿里云IdCard-Appcode"`

	PushType                 *string `json:"pushType,omitempty" v:"in:txTpns" dc:"推送方式"`
	TxTpnsHost               *string `json:"txTpnsHost,omitempty" v:"url" dc:"腾讯移动推送-域名"`
	TxTpnsAccessIDOfAndroid  *string `json:"txTpnsAccessIDOfAndroid,omitempty" v:"" dc:"腾讯移动推送-AccessID(安卓)"`
	TxTpnsSecretKeyOfAndroid *string `json:"txTpnsSecretKeyOfAndroid,omitempty" v:"" dc:"腾讯移动推送-SecretKey(安卓)"`
	TxTpnsAccessIDOfIos      *string `json:"txTpnsAccessIDOfIos,omitempty" v:"" dc:"腾讯移动推送-AccessID(苹果)"`
	TxTpnsSecretKeyOfIos     *string `json:"txTpnsSecretKeyOfIos,omitempty" v:"" dc:"腾讯移动推送-SecretKey(苹果)"`
	TxTpnsAccessIDOfMacOS    *string `json:"txTpnsAccessIDOfMacOS,omitempty" v:"" dc:"腾讯移动推送-AccessID(苹果电脑)"`
	TxTpnsSecretKeyOfMacOS   *string `json:"txTpnsSecretKeyOfMacOS,omitempty" v:"" dc:"腾讯移动推送-SecretKey(苹果电脑)"`

	VodType                  *string `json:"vodType,omitempty" v:"in:aliyunVod" dc:"视频点播方式"`
	AliyunVodAccessKeyId     *string `json:"aliyunVodAccessKeyId,omitempty" v:"regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"阿里云VOD-AccessKeyId"`
	AliyunVodAccessKeySecret *string `json:"aliyunVodAccessKeySecret,omitempty" v:"regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"阿里云VOD-AccessKeySecret"`
	AliyunVodEndpoint        *string `json:"aliyunVodEndpoint,omitempty" v:"" dc:"阿里云VOD-Endpoint"`
	AliyunVodRoleArn         *string `json:"aliyunVodRoleArn,omitempty" v:"" dc:"阿里云VOD-RoleArn"`
}

/*--------保存 结束--------*/
