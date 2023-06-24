package api

import "github.com/gogf/gf/v2/frame/g"

/*--------获取签名 开始--------*/
type UploadSignReq struct {
	g.Meta `path:"/sign" method:"post" tags:"上传" sm:"获取签名(web端直传用)"`
	Type   string `json:"type"  v:"" dc:"类型(暂时没用)"`
}

type UploadSignRes struct {
	Accessid  string `json:"accessid" dc:"阿里云存储-AccessId"`
	Dir       string `json:"dir" dc:"上传文件目录"`
	Expire    uint   `json:"expire" dc:"过期时间"`
	Host      string `json:"host" dc:"阿里云存储-域名"`
	Policy    string `json:"policy" dc:"上传策略policy"`
	Signature string `json:"signature" dc:"签名"`
}

/*--------获取签名 结束--------*/

/*--------获取Sts Token 开始--------*/
type UploadStsReq struct {
	g.Meta `path:"/sts" method:"post" tags:"上传" sm:"获取Sts Token(App端直传用)"`
	Type   string `json:"type"  v:"" dc:"类型(暂时没用)"`
}

type UploadStsRes struct {
	/*--------阿里云官方文档要求返回的字段 开始--------*/
	StatusCode      int    `json:"StatusCode" dc:"状态码"`
	AccessKeyId     string `json:"AccessKeyId" dc:"阿里云存储-AccessKeyId"`
	AccessKeySecret string `json:"AccessKeySecret" dc:"阿里云存储-AccessKeySecret"`
	SecurityToken   string `json:"SecurityToken" dc:"阿里云存储-SecurityToken"`
	Expiration      string `json:"Expiration" dc:"Expiration"`
	/*--------阿里云官方文档要求返回的字段 结束--------*/
	Endpoint         string `json:"endpoint" dc:"阿里云存储-endpoint"`
	Bucket           string `json:"bucket" dc:"阿里云存储-bucket"`
	Dir              string `json:"dir" dc:"上传文件目录"`
	CallbackUrl      string `json:"callbackUrl" dc:"回调地址"`
	CallbackBody     string `json:"callbackBody" dc:"回调地址"`
	CallbackBodyType string `json:"callbackBodyType" dc:"回调地址"`
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
