package api

import (
	"github.com/gogf/gf/v2/frame/g"
)

/*--------获取 开始--------*/
type ConfigGetReq struct {
	g.Meta       `path:"/config/get" method:"post" tags:"平台后台/配置中心/平台配置" sm:"获取"`
	ConfigKeyArr *[]string `json:"config_key_arr,omitempty" v:"required|distinct|foreach|min-length:1" dc:"配置键列表。传值参考默认返回的字段"`
}

type ConfigGetRes struct {
	Config Config `json:"config" dc:"配置列表"`
}

type Config struct {
	HotSearch        *[]string `json:"hotSearch,omitempty" dc:"热门搜索"`
	UserAgreement    *string   `json:"userAgreement,omitempty" dc:"用户协议"`
	PrivacyAgreement *string   `json:"privacyAgreement,omitempty" dc:"隐私协议"`

	SmsType     *string `json:"smsType,omitempty" dc:"短信方式"`
	SmsOfAliyun *struct {
		AccessKeyId     *string `json:"accessKeyId,omitempty" dc:"阿里云SMS-AccessKeyId"`
		AccessKeySecret *string `json:"accessKeySecret,omitempty" dc:"阿里云SMS-AccessKeySecret"`
		Endpoint        *string `json:"endpoint,omitempty" dc:"阿里云SMS-Endpoint"`
		SignName        *string `json:"signName,omitempty" dc:"阿里云SMS-签名"`
		TemplateCode    *string `json:"templateCode,omitempty" dc:"阿里云SMS-模板标识"`
	} `json:"smsOfAliyun,omitempty" dc:"阿里云SMS配置"`

	EmailCodeSubject  *string `json:"emailCodeSubject,omitempty" dc:"验证码邮件标题"`
	EmailCodeTemplate *string `json:"emailCodeTemplate,omitempty" dc:"验证码邮件内容"`
	EmailType         *string `json:"emailType,omitempty" dc:"邮箱方式"`
	EmailOfCommon     *struct {
		SmtpHost  *string `json:"smtpHost,omitempty" dc:"通用-SmtpHost"`
		SmtpPort  *string `json:"smtpPort,omitempty" dc:"通用-SmtpPort"`
		FromEmail *string `json:"fromEmail,omitempty" dc:"通用-邮箱"`
		Password  *string `json:"password,omitempty" dc:"通用-密码"`
	} `json:"emailOfCommon,omitempty" dc:"通用配置"`

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

	SmsType     *string `json:"smsType,omitempty" v:"in:smsOfAliyun" dc:"短信方式"`
	SmsOfAliyun *struct {
		AccessKeyId     *string `json:"accessKeyId,omitempty" v:"" dc:"阿里云SMS-AccessKeyId"`
		AccessKeySecret *string `json:"accessKeySecret,omitempty" v:"" dc:"阿里云SMS-AccessKeySecret"`
		Endpoint        *string `json:"endpoint,omitempty" v:"" dc:"阿里云SMS-Endpoint"`
		SignName        *string `json:"signName,omitempty" v:"" dc:"阿里云SMS-签名"`
		TemplateCode    *string `json:"templateCode,omitempty" v:"" dc:"阿里云SMS-模板标识"`
	} `json:"smsOfAliyun,omitempty" v:"required-if:SmsType,smsOfAliyun|json" dc:"阿里云SMS配置"`

	EmailCodeSubject  *string `json:"emailCodeSubject,omitempty" v:"" dc:"验证码标题"`
	EmailCodeTemplate *string `json:"emailCodeTemplate,omitempty" v:"" dc:"验证码模板"`
	EmailType         *string `json:"emailType,omitempty" v:"in:emailOfCommon" dc:"邮箱方式"`
	EmailOfCommon     *struct {
		SmtpHost  *string `json:"smtpHost,omitempty" v:"required" dc:"通用-SmtpHost"`
		SmtpPort  *string `json:"smtpPort,omitempty" v:"required" dc:"通用-SmtpPort"`
		FromEmail *string `json:"fromEmail,omitempty" v:"required|email" dc:"通用-邮箱"`
		Password  *string `json:"password,omitempty" v:"required" dc:"通用-密码"`
	} `json:"emailOfCommon,omitempty" v:"required-if:EmailType,emailOfCommon|json" dc:"通用配置"`

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
