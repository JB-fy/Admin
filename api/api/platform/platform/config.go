package api

import (
	"github.com/gogf/gf/v2/frame/g"
)

/*--------获取 开始--------*/
type ConfigGetReq struct {
	g.Meta       `path:"/config/get" method:"post" tags:"平台后台/配置中心/平台配置" sm:"获取"`
	ConfigKeyArr *[]string `json:"config_key_arr,omitempty" v:"required|distinct|foreach|min-length:1" dc:"配置Key列表。传值参考默认返回的字段"`
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

	UploadType                       *string `json:"uploadType,omitempty" dc:"上传方式"`
	UploadOfLocalUrl                 *string `json:"uploadOfLocalUrl,omitempty" dc:"本地-上传地址"`
	UploadOfLocalSignKey             *string `json:"uploadOfLocalSignKey,omitempty" dc:"本地-密钥"`
	UploadOfLocalFileSaveDir         *string `json:"uploadOfLocalFileSaveDir,omitempty" dc:"本地-文件保存目录"`
	UploadOfLocalFileUrlPrefix       *string `json:"uploadOfLocalFileUrlPrefix,omitempty" dc:"本地-文件地址前缀"`
	UploadOfAliyunOssHost            *string `json:"uploadOfAliyunOssHost,omitempty" dc:"阿里云OSS-域名"`
	UploadOfAliyunOssBucket          *string `json:"uploadOfAliyunOssBucket,omitempty" dc:"阿里云OSS-Bucket"`
	UploadOfAliyunOssAccessKeyId     *string `json:"uploadOfAliyunOssAccessKeyId,omitempty" dc:"阿里云OSS-AccessKeyId"`
	UploadOfAliyunOssAccessKeySecret *string `json:"uploadOfAliyunOssAccessKeySecret,omitempty" dc:"阿里云OSS-AccessKeySecret"`
	UploadOfAliyunOssCallbackUrl     *string `json:"uploadOfAliyunOssCallbackUrl,omitempty" dc:"阿里云OSS-回调地址"`
	UploadOfAliyunOssEndpoint        *string `json:"uploadOfAliyunOssEndpoint,omitempty" dc:"阿里云OSS-Endpoint"`
	UploadOfAliyunOssRoleArn         *string `json:"uploadOfAliyunOssRoleArn,omitempty" dc:"阿里云OSS-RoleArn"`

	PayOfAliAppId      *string `json:"payOfAliAppId,omitempty" dc:"AppId"`
	PayOfAliPrivateKey *string `json:"payOfAliPrivateKey,omitempty" dc:"私钥"`
	PayOfAliPublicKey  *string `json:"payOfAliPublicKey,omitempty" dc:"公钥"`
	PayOfAliNotifyUrl  *string `json:"payOfAliNotifyUrl,omitempty" dc:"异步回调地址"`
	PayOfAliOpAppId    *string `json:"payOfAliOpAppId,omitempty" dc:"小程序AppId"`

	PayOfWxAppId      *string `json:"payOfWxAppId,omitempty" dc:"AppId"`
	PayOfWxMchid      *string `json:"payOfWxMchid,omitempty" dc:"商户ID"`
	PayOfWxSerialNo   *string `json:"payOfWxSerialNo,omitempty" dc:"证书序列号"`
	PayOfWxApiV3Key   *string `json:"payOfWxApiV3Key,omitempty" dc:"APIV3密钥"`
	PayOfWxPrivateKey *string `json:"payOfWxPrivateKey,omitempty" dc:"私钥"`
	PayOfWxNotifyUrl  *string `json:"payOfWxNotifyUrl,omitempty" dc:"异步回调地址"`

	SmsType                    *string `json:"smsType,omitempty" dc:"短信方式"`
	SmsOfAliyunAccessKeyId     *string `json:"smsOfAliyunAccessKeyId,omitempty" dc:"阿里云SMS-AccessKeyId"`
	SmsOfAliyunAccessKeySecret *string `json:"smsOfAliyunAccessKeySecret,omitempty" dc:"阿里云SMS-AccessKeySecret"`
	SmsOfAliyunEndpoint        *string `json:"smsOfAliyunEndpoint,omitempty" dc:"阿里云SMS-Endpoint"`
	SmsOfAliyunSignName        *string `json:"smsOfAliyunSignName,omitempty" dc:"阿里云SMS-签名"`
	SmsOfAliyunTemplateCode    *string `json:"smsOfAliyunTemplateCode,omitempty" dc:"阿里云SMS-模板标识"`

	EmailType              *string `json:"emailType,omitempty" dc:"邮箱方式"`
	EmailOfCommonSmtpHost  *string `json:"emailOfCommonSmtpHost,omitempty" dc:"通用-SmtpHost"`
	EmailOfCommonSmtpPort  *string `json:"emailOfCommonSmtpPort,omitempty" dc:"通用-SmtpPort"`
	EmailOfCommonFromEmail *string `json:"emailOfCommonFromEmail,omitempty" dc:"通用-邮箱"`
	EmailOfCommonPassword  *string `json:"emailOfCommonPassword,omitempty" dc:"通用-密码"`
	EmailCodeSubject       *string `json:"emailCodeSubject,omitempty" dc:"验证码邮件标题"`
	EmailCodeTemplate      *string `json:"emailCodeTemplate,omitempty" dc:"验证码邮件内容"`

	IdCardType            *string `json:"idCardType,omitempty" dc:"实名认证方式"`
	IdCardOfAliyunHost    *string `json:"idCardOfAliyunHost,omitempty" dc:"阿里云IdCard-域名"`
	IdCardOfAliyunPath    *string `json:"idCardOfAliyunPath,omitempty" dc:"阿里云IdCard-请求路径"`
	IdCardOfAliyunAppcode *string `json:"idCardOfAliyunAppcode,omitempty" dc:"阿里云IdCard-Appcode"`

	OneClickOfWxHost          *string `json:"oneClickOfWxHost,omitempty" dc:"微信-域名"`
	OneClickOfWxAppId         *string `json:"oneClickOfWxAppId,omitempty" dc:"微信-AppId"`
	OneClickOfWxSecret        *string `json:"oneClickOfWxSecret,omitempty" dc:"微信-密钥"`
	OneClickOfYidunSecretId   *string `json:"oneClickOfYidunSecretId,omitempty" dc:"易盾-SecretId"`
	OneClickOfYidunSecretKey  *string `json:"oneClickOfYidunSecretKey,omitempty" dc:"易盾-SecretKey"`
	OneClickOfYidunBusinessId *string `json:"oneClickOfYidunBusinessId,omitempty" dc:"易盾-BusinessId"`

	PushType                 *string `json:"pushType,omitempty" dc:"推送方式"`
	PushOfTxHost             *string `json:"pushOfTxHost,omitempty" dc:"腾讯移动推送-域名"`
	PushOfTxAndroidAccessID  *string `json:"pushOfTxAndroidAccessID,omitempty" dc:"腾讯移动推送-AccessID(安卓)"`
	PushOfTxAndroidSecretKey *string `json:"pushOfTxAndroidSecretKey,omitempty" dc:"腾讯移动推送-SecretKey(安卓)"`
	PushOfTxIosAccessID      *string `json:"pushOfTxIosAccessID,omitempty" dc:"腾讯移动推送-AccessID(苹果)"`
	PushOfTxIosSecretKey     *string `json:"pushOfTxIosSecretKey,omitempty" dc:"腾讯移动推送-SecretKey(苹果)"`
	PushOfTxMacOSAccessID    *string `json:"pushOfTxMacOSAccessID,omitempty" dc:"腾讯移动推送-AccessID(苹果电脑)"`
	PushOfTxMacOSSecretKey   *string `json:"pushOfTxMacOSSecretKey,omitempty" dc:"腾讯移动推送-SecretKey(苹果电脑)"`

	VodType                    *string `json:"vodType,omitempty" dc:"视频点播方式"`
	VodOfAliyunAccessKeyId     *string `json:"vodOfAliyunAccessKeyId,omitempty" dc:"阿里云VOD-AccessKeyId"`
	VodOfAliyunAccessKeySecret *string `json:"vodOfAliyunAccessKeySecret,omitempty" dc:"阿里云VOD-AccessKeySecret"`
	VodOfAliyunEndpoint        *string `json:"vodOfAliyunEndpoint,omitempty" dc:"阿里云VOD-Endpoint"`
	VodOfAliyunRoleArn         *string `json:"vodOfAliyunRoleArn,omitempty" dc:"阿里云VOD-RoleArn"`

	WxGzhHost           *string `json:"wxGzhHost,omitempty" dc:"域名"`
	WxGzhAppId          *string `json:"wxGzhAppId,omitempty" dc:"AppId"`
	WxGzhSecret         *string `json:"wxGzhSecret,omitempty" dc:"密钥"`
	WxGzhToken          *string `json:"wxGzhToken,omitempty" dc:"Token"`
	WxGzhEncodingAESKey *string `json:"wxGzhEncodingAESKey,omitempty" dc:"EncodingAESKey"`
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

	UploadType                       *string `json:"uploadType,omitempty" v:"in:uploadOfLocal,uploadOfAliyunOss" dc:"上传方式"`
	UploadOfLocalUrl                 *string `json:"uploadOfLocalUrl,omitempty" v:"url" dc:"本地-上传地址"`
	UploadOfLocalSignKey             *string `json:"uploadOfLocalSignKey,omitempty" v:"" dc:"本地-密钥"`
	UploadOfLocalFileSaveDir         *string `json:"uploadOfLocalFileSaveDir,omitempty" v:"" dc:"本地-文件保存目录"`
	UploadOfLocalFileUrlPrefix       *string `json:"uploadOfLocalFileUrlPrefix,omitempty" v:"url" dc:"本地-文件地址前缀"`
	UploadOfAliyunOssHost            *string `json:"uploadOfAliyunOssHost,omitempty" v:"url" dc:"阿里云OSS-域名"`
	UploadOfAliyunOssBucket          *string `json:"uploadOfAliyunOssBucket,omitempty" v:"" dc:"阿里云OSS-Bucket"`
	UploadOfAliyunOssAccessKeyId     *string `json:"uploadOfAliyunOssAccessKeyId,omitempty" v:"regex:^[\\p{L}\\p{N}_-]+$" dc:"阿里云OSS-AccessKeyId"`
	UploadOfAliyunOssAccessKeySecret *string `json:"uploadOfAliyunOssAccessKeySecret,omitempty" v:"regex:^[\\p{L}\\p{N}_-]+$" dc:"阿里云OSS-AccessKeySecret"`
	UploadOfAliyunOssCallbackUrl     *string `json:"uploadOfAliyunOssCallbackUrl,omitempty" v:"url" dc:"阿里云OSS-回调地址"`
	UploadOfAliyunOssEndpoint        *string `json:"uploadOfAliyunOssEndpoint,omitempty" v:"" dc:"阿里云OSS-Endpoint"`
	UploadOfAliyunOssRoleArn         *string `json:"uploadOfAliyunOssRoleArn,omitempty" v:"" dc:"阿里云OSS-RoleArn"`

	PayOfAliAppId      *string `json:"payOfAliAppId,omitempty" v:"" dc:"AppId"`
	PayOfAliPrivateKey *string `json:"payOfAliPrivateKey,omitempty" v:"" dc:"私钥"`
	PayOfAliPublicKey  *string `json:"payOfAliPublicKey,omitempty" v:"" dc:"公钥"`
	PayOfAliNotifyUrl  *string `json:"payOfAliNotifyUrl,omitempty" v:"url" dc:"异步回调地址"`
	PayOfAliOpAppId    *string `json:"payOfAliOpAppId,omitempty" v:"" dc:"小程序AppId"`

	PayOfWxAppId      *string `json:"payOfWxAppId,omitempty" v:"" dc:"AppId"`
	PayOfWxMchid      *string `json:"payOfWxMchid,omitempty" v:"" dc:"商户ID"`
	PayOfWxSerialNo   *string `json:"payOfWxSerialNo,omitempty" v:"" dc:"证书序列号"`
	PayOfWxApiV3Key   *string `json:"payOfWxApiV3Key,omitempty" v:"" dc:"APIV3密钥"`
	PayOfWxPrivateKey *string `json:"payOfWxPrivateKey,omitempty" v:"" dc:"私钥"`
	PayOfWxNotifyUrl  *string `json:"payOfWxNotifyUrl,omitempty" v:"url" dc:"异步回调地址"`

	SmsType                    *string `json:"smsType,omitempty" v:"in:smsOfAliyun" dc:"短信方式"`
	SmsOfAliyunAccessKeyId     *string `json:"smsOfAliyunAccessKeyId,omitempty" v:"regex:^[\\p{L}\\p{N}_-]+$" dc:"阿里云SMS-AccessKeyId"`
	SmsOfAliyunAccessKeySecret *string `json:"smsOfAliyunAccessKeySecret,omitempty" v:"regex:^[\\p{L}\\p{N}_-]+$" dc:"阿里云SMS-AccessKeySecret"`
	SmsOfAliyunEndpoint        *string `json:"smsOfAliyunEndpoint,omitempty" v:"" dc:"阿里云SMS-Endpoint"`
	SmsOfAliyunSignName        *string `json:"smsOfAliyunSignName,omitempty" v:"" dc:"阿里云SMS-签名"`
	SmsOfAliyunTemplateCode    *string `json:"smsOfAliyunTemplateCode,omitempty" v:"" dc:"阿里云SMS-模板标识"`

	EmailType              *string `json:"emailType,omitempty" v:"in:emailOfCommon" dc:"邮箱方式"`
	EmailOfCommonSmtpHost  *string `json:"emailOfCommonSmtpHost,omitempty" v:"" dc:"通用-SmtpHost"`
	EmailOfCommonSmtpPort  *string `json:"emailOfCommonSmtpPort,omitempty" v:"" dc:"通用-SmtpPort"`
	EmailOfCommonFromEmail *string `json:"emailOfCommonFromEmail,omitempty" v:"email" dc:"通用-邮箱"`
	EmailOfCommonPassword  *string `json:"emailOfCommonPassword,omitempty" v:"" dc:"通用-密码"`
	EmailCodeSubject       *string `json:"emailCodeSubject,omitempty" v:"" dc:"验证码标题"`
	EmailCodeTemplate      *string `json:"emailCodeTemplate,omitempty" v:"" dc:"验证码模板"`

	IdCardType            *string `json:"idCardType,omitempty" v:"in:idCardOfAliyun" dc:"实名认证方式"`
	IdCardOfAliyunHost    *string `json:"idCardOfAliyunHost,omitempty" v:"url" dc:"阿里云IdCard-域名"`
	IdCardOfAliyunPath    *string `json:"idCardOfAliyunPath,omitempty" v:"" dc:"阿里云IdCard-请求路径"`
	IdCardOfAliyunAppcode *string `json:"idCardOfAliyunAppcode,omitempty" v:"" dc:"阿里云IdCard-Appcode"`

	OneClickOfWxHost          *string `json:"oneClickOfWxHost,omitempty" v:"url" dc:"微信-域名"`
	OneClickOfWxAppId         *string `json:"oneClickOfWxAppId,omitempty" v:"" dc:"微信-AppId"`
	OneClickOfWxSecret        *string `json:"oneClickOfWxSecret,omitempty" v:"" dc:"微信-密钥"`
	OneClickOfYidunSecretId   *string `json:"oneClickOfYidunSecretId,omitempty" v:"" dc:"易盾-SecretId"`
	OneClickOfYidunSecretKey  *string `json:"oneClickOfYidunSecretKey,omitempty" v:"" dc:"易盾-SecretKey"`
	OneClickOfYidunBusinessId *string `json:"oneClickOfYidunBusinessId,omitempty" v:"" dc:"易盾-BusinessId"`

	PushType                 *string `json:"pushType,omitempty" v:"in:pushOfTx" dc:"推送方式"`
	PushOfTxHost             *string `json:"pushOfTxHost,omitempty" v:"url" dc:"腾讯移动推送-域名"`
	PushOfTxAndroidAccessID  *string `json:"pushOfTxAndroidAccessID,omitempty" v:"" dc:"腾讯移动推送-AccessID(安卓)"`
	PushOfTxAndroidSecretKey *string `json:"pushOfTxAndroidSecretKey,omitempty" v:"" dc:"腾讯移动推送-SecretKey(安卓)"`
	PushOfTxIosAccessID      *string `json:"pushOfTxIosAccessID,omitempty" v:"" dc:"腾讯移动推送-AccessID(苹果)"`
	PushOfTxIosSecretKey     *string `json:"pushOfTxIosSecretKey,omitempty" v:"" dc:"腾讯移动推送-SecretKey(苹果)"`
	PushOfTxMacOSAccessID    *string `json:"pushOfTxMacOSAccessID,omitempty" v:"" dc:"腾讯移动推送-AccessID(苹果电脑)"`
	PushOfTxMacOSSecretKey   *string `json:"pushOfTxMacOSSecretKey,omitempty" v:"" dc:"腾讯移动推送-SecretKey(苹果电脑)"`

	VodType                    *string `json:"vodType,omitempty" v:"in:vodOfAliyun" dc:"视频点播方式"`
	VodOfAliyunAccessKeyId     *string `json:"vodOfAliyunAccessKeyId,omitempty" v:"regex:^[\\p{L}\\p{N}_-]+$" dc:"阿里云VOD-AccessKeyId"`
	VodOfAliyunAccessKeySecret *string `json:"vodOfAliyunAccessKeySecret,omitempty" v:"regex:^[\\p{L}\\p{N}_-]+$" dc:"阿里云VOD-AccessKeySecret"`
	VodOfAliyunEndpoint        *string `json:"vodOfAliyunEndpoint,omitempty" v:"" dc:"阿里云VOD-Endpoint"`
	VodOfAliyunRoleArn         *string `json:"vodOfAliyunRoleArn,omitempty" v:"" dc:"阿里云VOD-RoleArn"`

	WxGzhHost           *string `json:"wxGzhHost,omitempty" v:"url" dc:"微信公众号-域名"`
	WxGzhAppId          *string `json:"wxGzhAppId,omitempty" v:"" dc:"微信公众号-AppId"`
	WxGzhSecret         *string `json:"wxGzhSecret,omitempty" v:"" dc:"微信公众号-密钥"`
	WxGzhToken          *string `json:"wxGzhToken,omitempty" v:"" dc:"微信公众号-Token"`
	WxGzhEncodingAESKey *string `json:"wxGzhEncodingAESKey,omitempty" v:"size:43" dc:"微信公众号-EncodingAESKey"`
}

/*--------保存 结束--------*/
