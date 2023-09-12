package api

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

/*--------上传本地 开始--------*/
type UploadUploadReq struct {
	g.Meta `path:"/upload" method:"post" tags:"上传" sm:"上传本地"`
	Dir    string            `json:"dir" v:"required" dc:"上传目录"`
	Expire string            `json:"expire" v:"required" dc:"过期时间"`
	Rand   string            `json:"rand" v:"required" dc:"随机字符串"`
	Sign   string            `json:"sign" v:"required" dc:"签名"`
	Key    string            `json:"key" v:"" dc:"文件名称"`
	File   *ghttp.UploadFile `json:"file" v:"required" dc:"上传文件"`
}

type UploadUploadRes struct {
	Url string `json:"url" dc:"地址"`
}

/*--------上传本地 结束--------*/

/*--------获取签名 开始--------*/
type UploadSignReq struct {
	g.Meta `path:"/sign" method:"post" tags:"上传" sm:"获取签名(web端直传用)"`
	Type   string `json:"type" v:"" dc:"类型(暂时没用)"`
}

type UploadSignRes struct {
	UploadUrl  string                 `json:"uploadUrl,omitempty" dc:"上传地址"`
	UploadData map[string]interface{} `json:"uploadData,omitempty" dc:"上传数据"`
	Dir        string                 `json:"dir,omitempty" dc:"上传目录"`
	Expire     uint                   `json:"expire,omitempty" dc:"过期时间"`
	Host       string                 `json:"host,omitempty" dc:"站点域名（当上传无响应数据，前端组件用于与文件保存路径拼接形成文件访问地址）"`
	IsRes      uint                   `json:"isRes,omitempty" dc:"上传是否有响应数据：0否 1是"`
}

/*--------获取签名 结束--------*/

/*--------获取Sts Token 开始--------*/
type UploadStsReq struct {
	g.Meta `path:"/sts" method:"get,post" tags:"上传" sm:"获取Sts Token(App端直传用)"`
	Type   string `json:"type" v:"" dc:"类型(暂时没用)"`
}

type UploadStsRes struct {
	/*--------App端的SDK需设置一个地址来获取Sts Token，且必须按要求以下字段 开始--------*/
	StatusCode      int    `json:"StatusCode,omitempty" dc:"状态码"`
	AccessKeyId     string `json:"AccessKeyId,omitempty" dc:"阿里云OSS-AccessKeyId"`
	AccessKeySecret string `json:"AccessKeySecret,omitempty" dc:"阿里云OSS-AccessKeySecret"`
	SecurityToken   string `json:"SecurityToken,omitempty" dc:"阿里云OSS-SecurityToken"`
	Expiration      string `json:"Expiration,omitempty" dc:"Expiration"`
	/*--------App端的SDK需设置一个地址来获取Sts Token，且必须按要求以下字段 结束--------*/

	/*--------App端实际上传时需用到的字段，但必须验证权限后才能拿到 开始--------*/
	Endpoint         string `json:"endpoint,omitempty" dc:"阿里云OSS-endpoint"`
	Bucket           string `json:"bucket,omitempty" dc:"阿里云OSS-bucket"`
	Dir              string `json:"dir,omitempty" dc:"上传文件目录"`
	CallbackUrl      string `json:"callbackUrl,omitempty" dc:"回调地址"`
	CallbackBody     string `json:"callbackBody,omitempty" dc:"回调参数"`
	CallbackBodyType string `json:"callbackBodyType,omitempty" dc:"回调方式"`
	/*--------App端实际上传时需用到的字段，但必须验证权限后才能拿到 结束--------*/
}

/*--------获取Sts Token 结束--------*/

/*--------回调 开始--------*/
type UploadNotifyReq struct {
	g.Meta `path:"/notify" method:"get,post" tags:"上传" sm:"回调"`
}

type UploadNotifyRes struct {
	Url string `json:"url" dc:"地址"`
}

/*--------回调 结束--------*/
