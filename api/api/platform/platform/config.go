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
		AccessKeyId     *string `json:"accessKeyId,omitempty" dc:"阿里云-AccessKeyId"`
		AccessKeySecret *string `json:"accessKeySecret,omitempty" dc:"阿里云-AccessKeySecret"`
		Endpoint        *string `json:"endpoint,omitempty" dc:"阿里云-Endpoint"`
		SignName        *string `json:"signName,omitempty" dc:"阿里云-签名"`
		TemplateCode    *string `json:"templateCode,omitempty" dc:"阿里云-模板标识"`
	} `json:"smsOfAliyun,omitempty" dc:"短信配置-阿里云"`

	EmailCodeSubject  *string `json:"emailCodeSubject,omitempty" dc:"验证码邮件标题"`
	EmailCodeTemplate *string `json:"emailCodeTemplate,omitempty" dc:"验证码邮件内容"`
	EmailType         *string `json:"emailType,omitempty" dc:"邮箱方式"`
	EmailOfCommon     *struct {
		SmtpHost  *string `json:"smtpHost,omitempty" dc:"通用-SmtpHost"`
		SmtpPort  *string `json:"smtpPort,omitempty" dc:"通用-SmtpPort"`
		FromEmail *string `json:"fromEmail,omitempty" dc:"通用-邮箱"`
		Password  *string `json:"password,omitempty" dc:"通用-密码"`
	} `json:"emailOfCommon,omitempty" dc:"邮箱配置-通用"`

	IdCardType     *string `json:"idCardType,omitempty" dc:"实名认证方式"`
	IdCardOfAliyun *struct {
		Host    *string `json:"host,omitempty" dc:"阿里云-域名"`
		Path    *string `json:"path,omitempty" dc:"阿里云-请求路径"`
		Appcode *string `json:"appcode,omitempty" dc:"阿里云-Appcode"`
	} `json:"idCardOfAliyun,omitempty" dc:"实名认证配置-阿里云"`

	OneClickOfWx *struct {
		Host   *string `json:"host,omitempty" dc:"微信-域名"`
		AppId  *string `json:"appId,omitempty" dc:"微信-AppId"`
		Secret *string `json:"secret,omitempty" dc:"微信-密钥"`
	} `json:"oneClickOfWx,omitempty" dc:"一键登录配置-微信"`
	OneClickOfYidun *struct {
		SecretId   *string `json:"secretId,omitempty" dc:"易盾-SecretId"`
		SecretKey  *string `json:"secretKey,omitempty" dc:"易盾-SecretKey"`
		BusinessId *string `json:"businessId,omitempty" dc:"易盾-BusinessId"`
	} `json:"oneClickOfYidun,omitempty" dc:"一键登录配置-易盾"`

	PushType *string `json:"pushType,omitempty" dc:"推送方式"`
	PushOfTx *struct {
		Host               *string `json:"host,omitempty" dc:"腾讯移动推送-域名"`
		AccessIDOfAndroid  *string `json:"accessIDOfAndroid,omitempty" dc:"腾讯移动推送-AccessID(安卓)"`
		SecretKeyOfAndroid *string `json:"secretKeyOfAndroid,omitempty" dc:"腾讯移动推送-SecretKey(安卓)"`
		AccessIDOfIos      *string `json:"accessIDOfIos,omitempty" dc:"腾讯移动推送-AccessID(苹果)"`
		SecretKeyOfIos     *string `json:"secretKeyOfIos,omitempty" dc:"腾讯移动推送-SecretKey(苹果)"`
		AccessIDOfMacOS    *string `json:"accessIDOfMacOS,omitempty" dc:"腾讯移动推送-AccessID(苹果电脑)"`
		SecretKeyOfMacOS   *string `json:"secretKeyOfMacOS,omitempty" dc:"腾讯移动推送-SecretKey(苹果电脑)"`
	} `json:"pushOfTx,omitempty" dc:"推送配置-腾讯移动推送"`

	VodType     *string `json:"vodType,omitempty" dc:"视频点播方式"`
	VodOfAliyun *struct {
		AccessKeyId     *string `json:"accessKeyId,omitempty" dc:"阿里云-AccessKeyId"`
		AccessKeySecret *string `json:"accessKeySecret,omitempty" dc:"阿里云-AccessKeySecret"`
		Endpoint        *string `json:"endpoint,omitempty" dc:"阿里云-Endpoint"`
		RoleArn         *string `json:"roleArn,omitempty" dc:"阿里云-RoleArn"`
	} `json:"vodOfAliyun,omitempty" dc:"视频点播配置-阿里云"`

	WxGzh *struct {
		Host           *string `json:"host,omitempty" dc:"公众号-域名"`
		AppId          *string `json:"appId,omitempty" dc:"公众号-AppId"`
		Secret         *string `json:"secret,omitempty" dc:"公众号-密钥"`
		Token          *string `json:"token,omitempty" dc:"公众号-Token"`
		EncodingAESKey *string `json:"encodingAESKey,omitempty" dc:"公众号-EncodingAESKey"`
	} `json:"wxGzh,omitempty" dc:"微信配置-公众号"`
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
		AccessKeyId     *string `json:"accessKeyId,omitempty" v:"" dc:"阿里云-AccessKeyId"`
		AccessKeySecret *string `json:"accessKeySecret,omitempty" v:"" dc:"阿里云-AccessKeySecret"`
		Endpoint        *string `json:"endpoint,omitempty" v:"" dc:"阿里云-Endpoint"`
		SignName        *string `json:"signName,omitempty" v:"" dc:"阿里云-签名"`
		TemplateCode    *string `json:"templateCode,omitempty" v:"" dc:"阿里云-模板标识"`
	} `json:"smsOfAliyun,omitempty" v:"required-if:SmsType,smsOfAliyun" dc:"短信配置-阿里云"`

	EmailCodeSubject  *string `json:"emailCodeSubject,omitempty" v:"" dc:"验证码标题"`
	EmailCodeTemplate *string `json:"emailCodeTemplate,omitempty" v:"" dc:"验证码模板"`
	EmailType         *string `json:"emailType,omitempty" v:"in:emailOfCommon" dc:"邮箱方式"`
	EmailOfCommon     *struct {
		SmtpHost  *string `json:"smtpHost,omitempty" v:"" dc:"通用-SmtpHost"`
		SmtpPort  *string `json:"smtpPort,omitempty" v:"" dc:"通用-SmtpPort"`
		FromEmail *string `json:"fromEmail,omitempty" v:"email" dc:"通用-邮箱"`
		Password  *string `json:"password,omitempty" v:"" dc:"通用-密码"`
	} `json:"emailOfCommon,omitempty" v:"required-if:EmailType,emailOfCommon" dc:"邮箱配置-通用"`

	IdCardType     *string `json:"idCardType,omitempty" v:"in:idCardOfAliyun" dc:"实名认证方式"`
	IdCardOfAliyun *struct {
		Host    *string `json:"host,omitempty" v:"url" dc:"阿里云-域名"`
		Path    *string `json:"path,omitempty" v:"" dc:"阿里云-请求路径"`
		Appcode *string `json:"appcode,omitempty" v:"" dc:"阿里云-Appcode"`
	} `json:"idCardOfAliyun,omitempty" v:"required-if:IdCardType,idCardOfAliyun" dc:"实名认证配置-阿里云"`

	OneClickOfWx *struct {
		Host   *string `json:"host,omitempty" v:"url" dc:"微信-域名"`
		AppId  *string `json:"appId,omitempty" v:"" dc:"微信-AppId"`
		Secret *string `json:"secret,omitempty" v:"" dc:"微信-密钥"`
	} `json:"oneClickOfWx,omitempty" v:"" dc:"一键登录配置-微信"`
	OneClickOfYidun *struct {
		SecretId   *string `json:"secretId,omitempty" v:"" dc:"易盾-SecretId"`
		SecretKey  *string `json:"secretKey,omitempty" v:"" dc:"易盾-SecretKey"`
		BusinessId *string `json:"businessId,omitempty" v:"" dc:"易盾-BusinessId"`
	} `json:"oneClickOfYidun,omitempty" v:"" dc:"一键登录配置-易盾"`

	PushType *string `json:"pushType,omitempty" v:"in:pushOfTx" dc:"推送方式"`
	PushOfTx *struct {
		Host               *string `json:"host,omitempty" v:"url" dc:"腾讯移动推送-域名"`
		AccessIDOfAndroid  *string `json:"accessIDOfAndroid,omitempty" v:"" dc:"腾讯移动推送-AccessID(安卓)"`
		SecretKeyOfAndroid *string `json:"secretKeyOfAndroid,omitempty" v:"" dc:"腾讯移动推送-SecretKey(安卓)"`
		AccessIDOfIos      *string `json:"accessIDOfIos,omitempty" v:"" dc:"腾讯移动推送-AccessID(苹果)"`
		SecretKeyOfIos     *string `json:"secretKeyOfIos,omitempty" v:"" dc:"腾讯移动推送-SecretKey(苹果)"`
		AccessIDOfMacOS    *string `json:"accessIDOfMacOS,omitempty" v:"" dc:"腾讯移动推送-AccessID(苹果电脑)"`
		SecretKeyOfMacOS   *string `json:"secretKeyOfMacOS,omitempty" v:"" dc:"腾讯移动推送-SecretKey(苹果电脑)"`
	} `json:"pushOfTx,omitempty" v:"required-if:PushType,pushOfTx" dc:"推送配置-腾讯移动推送"`

	VodType     *string `json:"vodType,omitempty" v:"in:vodOfAliyun" dc:"视频点播方式"`
	VodOfAliyun *struct {
		AccessKeyId     *string `json:"accessKeyId,omitempty" v:"" dc:"阿里云-AccessKeyId"`
		AccessKeySecret *string `json:"accessKeySecret,omitempty" v:"" dc:"阿里云-AccessKeySecret"`
		Endpoint        *string `json:"endpoint,omitempty" v:"" dc:"阿里云-Endpoint"`
		RoleArn         *string `json:"roleArn,omitempty" v:"" dc:"阿里云-RoleArn"`
	} `json:"vodOfAliyun,omitempty" v:"required-if:VodType,vodOfAliyun" dc:"视频点播配置-阿里云"`

	WxGzh *struct {
		Host           *string `json:"host,omitempty" v:"url" dc:"公众号-域名"`
		AppId          *string `json:"appId,omitempty" v:"" dc:"公众号-AppId"`
		Secret         *string `json:"secret,omitempty" v:"" dc:"公众号-密钥"`
		Token          *string `json:"token,omitempty" v:"" dc:"公众号-Token"`
		EncodingAESKey *string `json:"encodingAESKey,omitempty" v:"size:43" dc:"公众号-EncodingAESKey"`
	} `json:"wxGzh,omitempty" v:"" dc:"微信配置-公众号"`
}

/*--------保存 结束--------*/
