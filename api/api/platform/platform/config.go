package platform

import (
	"api/api"

	"github.com/gogf/gf/v2/frame/g"
)

/*--------获取 开始--------*/
type ConfigGetReq struct {
	g.Meta `path:"/config/get" method:"post" tags:"平台后台/配置中心/平台配置" sm:"获取"`
	api.CommonPlatformHeaderReq
	ConfigKeyArr *[]string `json:"config_key_arr,omitempty" v:"required|distinct|foreach|length:1,30" dc:"配置键列表。传值参考默认返回的字段"`
}

type ConfigGetRes struct {
	Config Config `json:"config" dc:"配置列表"`
}

type Config struct {
	HotSearch        *[]string `json:"hot_search,omitempty" dc:"热门搜索"`
	UserAgreement    *string   `json:"user_agreement,omitempty" dc:"用户协议"`
	PrivacyAgreement *string   `json:"privacy_agreement,omitempty" dc:"隐私协议"`

	RoleIdArrOfPlatformDef *[]uint `json:"role_id_arr_of_platform_def,omitempty" dc:"默认角色(注册)"`

	RoleIdArrOfOrgDef *[]uint `json:"role_id_arr_of_org_def,omitempty" dc:"默认角色(注册)"`

	SmsType     *string `json:"sms_type,omitempty" dc:"短信方式"`
	SmsOfAliyun *struct {
		AccessKeyId     *string `json:"access_key_id,omitempty" dc:"阿里云-AccessKeyId"`
		AccessKeySecret *string `json:"access_key_secret,omitempty" dc:"阿里云-AccessKeySecret"`
		Endpoint        *string `json:"endpoint,omitempty" dc:"阿里云-Endpoint"`
		SignName        *string `json:"sign_name,omitempty" dc:"阿里云-签名"`
		TemplateCode    *string `json:"template_code,omitempty" dc:"阿里云-模板标识"`
	} `json:"sms_of_aliyun,omitempty" dc:"短信配置-阿里云"`

	EmailCode *struct {
		Subject  *string `json:"subject,omitempty" dc:"标题"`
		Template *string `json:"template,omitempty" dc:"内容"`
	} `json:"email_code,omitempty" dc:"验证码邮件配置"`
	EmailType     *string `json:"email_type,omitempty" dc:"邮箱方式"`
	EmailOfCommon *struct {
		SmtpHost  *string `json:"smtp_host,omitempty" dc:"通用-SmtpHost"`
		SmtpPort  *string `json:"smtp_port,omitempty" dc:"通用-SmtpPort"`
		FromEmail *string `json:"from_email,omitempty" dc:"通用-邮箱"`
		Password  *string `json:"password,omitempty" dc:"通用-密码"`
	} `json:"email_of_common,omitempty" dc:"邮箱配置-通用"`

	IdCardType     *string `json:"id_card_type,omitempty" dc:"实名认证方式"`
	IdCardOfAliyun *struct {
		Url     *string `json:"url,omitempty" dc:"阿里云-请求地址"`
		Appcode *string `json:"appcode,omitempty" dc:"阿里云-Appcode"`
	} `json:"id_card_of_aliyun,omitempty" dc:"实名认证配置-阿里云"`

	OneClickOfWx *struct {
		Host   *string `json:"host,omitempty" dc:"微信-域名"`
		AppId  *string `json:"app_id,omitempty" dc:"微信-AppId"`
		Secret *string `json:"secret,omitempty" dc:"微信-密钥"`
	} `json:"one_click_of_wx,omitempty" dc:"一键登录配置-微信"`
	OneClickOfYidun *struct {
		SecretId   *string `json:"secret_id,omitempty" dc:"易盾-SecretId"`
		SecretKey  *string `json:"secret_key,omitempty" dc:"易盾-SecretKey"`
		BusinessId *string `json:"business_id,omitempty" dc:"易盾-BusinessId"`
	} `json:"one_click_of_yidun,omitempty" dc:"一键登录配置-易盾"`

	PushType *string `json:"push_type,omitempty" dc:"推送方式"`
	PushOfTx *struct {
		Host               *string `json:"host,omitempty" dc:"腾讯移动推送-域名"`
		AccessIDOfAndroid  *string `json:"access_id_of_android,omitempty" dc:"腾讯移动推送-AccessID(安卓)"`
		SecretKeyOfAndroid *string `json:"secret_key_of_android,omitempty" dc:"腾讯移动推送-SecretKey(安卓)"`
		AccessIDOfIos      *string `json:"access_id_of_ios,omitempty" dc:"腾讯移动推送-AccessID(苹果)"`
		SecretKeyOfIos     *string `json:"secret_key_of_ios,omitempty" dc:"腾讯移动推送-SecretKey(苹果)"`
		AccessIDOfMacOS    *string `json:"access_id_of_mac_os,omitempty" dc:"腾讯移动推送-AccessID(苹果电脑)"`
		SecretKeyOfMacOS   *string `json:"secret_key_of_mac_os,omitempty" dc:"腾讯移动推送-SecretKey(苹果电脑)"`
	} `json:"push_of_tx,omitempty" dc:"推送配置-腾讯移动推送"`

	VodType     *string `json:"vod_type,omitempty" dc:"视频点播方式"`
	VodOfAliyun *struct {
		AccessKeyId     *string `json:"access_key_id,omitempty" dc:"阿里云-AccessKeyId"`
		AccessKeySecret *string `json:"access_key_secret,omitempty" dc:"阿里云-AccessKeySecret"`
		Endpoint        *string `json:"endpoint,omitempty" dc:"阿里云-Endpoint"`
		RoleArn         *string `json:"role_arn,omitempty" dc:"阿里云-RoleArn"`
	} `json:"vod_of_aliyun,omitempty" dc:"视频点播配置-阿里云"`

	WxGzh *struct {
		Host           *string `json:"host,omitempty" dc:"公众号-域名"`
		AppId          *string `json:"app_id,omitempty" dc:"公众号-AppId"`
		Secret         *string `json:"secret,omitempty" dc:"公众号-密钥"`
		Token          *string `json:"token,omitempty" dc:"公众号-Token"`
		EncodingAESKey *string `json:"encoding_aes_key,omitempty" dc:"公众号-EncodingAESKey"`
	} `json:"wx_gzh,omitempty" dc:"微信配置-公众号"`
}

/*--------获取 结束--------*/

/*--------保存 开始--------*/
type ConfigSaveReq struct {
	g.Meta `path:"/config/save" method:"post" tags:"平台后台/配置中心/平台配置" sm:"保存"`
	api.CommonPlatformHeaderReq
	HotSearch        *[]string `json:"hot_search,omitempty" v:"distinct|foreach|min-length:1" dc:"热门搜索"`
	UserAgreement    *string   `json:"user_agreement,omitempty" v:"" dc:"用户协议"`
	PrivacyAgreement *string   `json:"privacy_agreement,omitempty" v:"" dc:"隐私协议"`

	RoleIdArrOfPlatformDef *[]uint `json:"role_id_arr_of_platform_def,omitempty" v:"distinct|foreach|between:1,4294967295" dc:"默认角色(注册)"`

	RoleIdArrOfOrgDef *[]uint `json:"role_id_arr_of_org_def,omitempty" v:"distinct|foreach|between:1,4294967295" dc:"默认角色(注册)"`

	SmsType     *string `json:"sms_type,omitempty" v:"in:sms_of_aliyun" dc:"短信方式"`
	SmsOfAliyun *struct {
		AccessKeyId     *string `json:"access_key_id,omitempty" v:"" dc:"阿里云-AccessKeyId"`
		AccessKeySecret *string `json:"access_key_secret,omitempty" v:"" dc:"阿里云-AccessKeySecret"`
		Endpoint        *string `json:"endpoint,omitempty" v:"" dc:"阿里云-Endpoint"`
		SignName        *string `json:"sign_name,omitempty" v:"" dc:"阿里云-签名"`
		TemplateCode    *string `json:"template_code,omitempty" v:"" dc:"阿里云-模板标识"`
	} `json:"sms_of_aliyun,omitempty" v:"required-if:SmsType,sms_of_aliyun" dc:"短信配置-阿里云"`

	EmailCode *struct {
		Subject  *string `json:"subject,omitempty" v:"" dc:"标题"`
		Template *string `json:"template,omitempty" v:"" dc:"内容"`
	} `json:"email_code,omitempty" v:"" dc:"验证码邮件配置"`
	EmailType     *string `json:"email_type,omitempty" v:"in:email_of_common" dc:"邮箱方式"`
	EmailOfCommon *struct {
		SmtpHost  *string `json:"smtp_host,omitempty" v:"" dc:"通用-SmtpHost"`
		SmtpPort  *string `json:"smtp_port,omitempty" v:"" dc:"通用-SmtpPort"`
		FromEmail *string `json:"from_email,omitempty" v:"email" dc:"通用-邮箱"`
		Password  *string `json:"password,omitempty" v:"" dc:"通用-密码"`
	} `json:"email_of_common,omitempty" v:"required-if:EmailType,email_of_common" dc:"邮箱配置-通用"`

	IdCardType     *string `json:"id_card_type,omitempty" v:"in:id_card_of_aliyun" dc:"实名认证方式"`
	IdCardOfAliyun *struct {
		Url     *string `json:"url,omitempty" v:"url" dc:"阿里云-请求地址"`
		Appcode *string `json:"appcode,omitempty" v:"" dc:"阿里云-Appcode"`
	} `json:"id_card_of_aliyun,omitempty" v:"required-if:IdCardType,id_card_of_aliyun" dc:"实名认证配置-阿里云"`

	OneClickOfWx *struct {
		Host   *string `json:"host,omitempty" v:"url" dc:"微信-域名"`
		AppId  *string `json:"app_id,omitempty" v:"" dc:"微信-AppId"`
		Secret *string `json:"secret,omitempty" v:"" dc:"微信-密钥"`
	} `json:"one_click_of_wx,omitempty" v:"" dc:"一键登录配置-微信"`
	OneClickOfYidun *struct {
		SecretId   *string `json:"secret_id,omitempty" v:"" dc:"易盾-SecretId"`
		SecretKey  *string `json:"secret_key,omitempty" v:"" dc:"易盾-SecretKey"`
		BusinessId *string `json:"business_id,omitempty" v:"" dc:"易盾-BusinessId"`
	} `json:"one_click_of_yidun,omitempty" v:"" dc:"一键登录配置-易盾"`

	PushType *string `json:"push_type,omitempty" v:"in:push_of_tx" dc:"推送方式"`
	PushOfTx *struct {
		Host               *string `json:"host,omitempty" v:"url" dc:"腾讯移动推送-域名"`
		AccessIDOfAndroid  *string `json:"access_id_of_android,omitempty" v:"" dc:"腾讯移动推送-AccessID(安卓)"`
		SecretKeyOfAndroid *string `json:"secret_key_of_android,omitempty" v:"" dc:"腾讯移动推送-SecretKey(安卓)"`
		AccessIDOfIos      *string `json:"access_id_of_ios,omitempty" v:"" dc:"腾讯移动推送-AccessID(苹果)"`
		SecretKeyOfIos     *string `json:"secret_key_of_ios,omitempty" v:"" dc:"腾讯移动推送-SecretKey(苹果)"`
		AccessIDOfMacOS    *string `json:"access_id_of_mac_os,omitempty" v:"" dc:"腾讯移动推送-AccessID(苹果电脑)"`
		SecretKeyOfMacOS   *string `json:"secret_key_of_mac_os,omitempty" v:"" dc:"腾讯移动推送-SecretKey(苹果电脑)"`
	} `json:"push_of_tx,omitempty" v:"required-if:PushType,push_of_tx" dc:"推送配置-腾讯移动推送"`

	VodType     *string `json:"vod_type,omitempty" v:"in:vod_of_aliyun" dc:"视频点播方式"`
	VodOfAliyun *struct {
		AccessKeyId     *string `json:"access_key_id,omitempty" v:"" dc:"阿里云-AccessKeyId"`
		AccessKeySecret *string `json:"access_key_secret,omitempty" v:"" dc:"阿里云-AccessKeySecret"`
		Endpoint        *string `json:"endpoint,omitempty" v:"" dc:"阿里云-Endpoint"`
		RoleArn         *string `json:"role_arn,omitempty" v:"" dc:"阿里云-RoleArn"`
	} `json:"vod_of_aliyun,omitempty" v:"required-if:VodType,vod_of_aliyun" dc:"视频点播配置-阿里云"`

	WxGzh *struct {
		Host           *string `json:"host,omitempty" v:"url" dc:"公众号-域名"`
		AppId          *string `json:"app_id,omitempty" v:"" dc:"公众号-AppId"`
		Secret         *string `json:"secret,omitempty" v:"" dc:"公众号-密钥"`
		Token          *string `json:"token,omitempty" v:"" dc:"公众号-Token"`
		EncodingAESKey *string `json:"encoding_aes_key,omitempty" v:"size:43" dc:"公众号-EncodingAESKey"`
	} `json:"wx_gzh,omitempty" v:"" dc:"微信配置-公众号"`
}

/*--------保存 结束--------*/
