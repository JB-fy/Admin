package api

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

/*--------上传本地 开始--------*/
type UploadUploadReq struct {
	g.Meta   `path:"/upload" method:"post" tags:"上传" sm:"上传本地"`
	UploadId uint              `json:"upload_id" v:"required|between:1,4294967295" dc:"上传ID"`
	Dir      string            `json:"dir" v:"required" dc:"上传目录"`
	Expire   string            `json:"expire" v:"required" dc:"过期时间"`
	Rand     string            `json:"rand" v:"required" dc:"随机字符串"`
	Sign     string            `json:"sign" v:"required" dc:"签名"`
	Key      string            `json:"key" v:"" dc:"文件名称"`
	File     *ghttp.UploadFile `json:"file" v:"required" dc:"上传文件"`
}

/*--------上传本地 结束--------*/

/*--------获取签名（H5直传用） 开始--------*/
type UploadSignReq struct {
	g.Meta `path:"/sign" method:"post" tags:"上传" sm:"获取签名（H5直传用）"`
	Scene  string `json:"scene" v:"" dc:"上传场景"`
}

type UploadSignRes struct {
	UploadUrl  string         `json:"upload_url,omitempty" dc:"上传地址"`
	UploadData map[string]any `json:"upload_data,omitempty" dc:"上传数据"`
	Host       string         `json:"host,omitempty" dc:"站点域名（当上传无响应信息，前端组件用于与上传目录拼接形成文件访问地址）"`
	Dir        string         `json:"dir,omitempty" dc:"上传目录"`
	Expire     uint           `json:"expire,omitempty" dc:"过期时间"`
	IsRes      uint           `json:"is_res,omitempty" dc:"是否有响应信息。0否 1是"`
}

/*--------获取签名（H5直传用） 结束--------*/

/*--------获取配置信息（APP直传前调用） 开始--------*/
type UploadConfigReq struct {
	g.Meta `path:"/config" method:"post" tags:"上传" sm:"获取配置信息（APP直传前调用）"`
	Scene  string `json:"scene" v:"" dc:"上传场景"`
}

type UploadConfigRes struct {
	Endpoint         string `json:"endpoint,omitempty" dc:"阿里云OSS-endpoint"`
	Bucket           string `json:"bucket,omitempty" dc:"阿里云OSS-bucket"`
	Dir              string `json:"dir,omitempty" dc:"上传文件目录"`
	CallbackUrl      string `json:"callbackUrl,omitempty" dc:"回调地址"`
	CallbackBody     string `json:"callbackBody,omitempty" dc:"回调参数"`
	CallbackBodyType string `json:"callbackBodyType,omitempty" dc:"回调方式"`
}

/*--------获取Sts 获取配置信息（APP直传前调用） 结束--------*/

/*--------获取Sts Token（APP直传用） 开始--------*/
type UploadStsReq struct { //阿里云的APP SDK通过设置地址来获取Sts Token。请求方式必须是GET
	g.Meta `path:"/sts" method:"get" tags:"上传" sm:"获取Sts Token（APP直传用）"`
	Scene  string `json:"scene" v:"" dc:"上传场景"`
}

type UploadStsRes struct {
	StatusCode      int    `json:"StatusCode,omitempty" dc:"状态码"`
	RequestId       string `json:"RequestId,omitempty" dc:"请求ID"`
	AccessKeyId     string `json:"AccessKeyId,omitempty" dc:"阿里云OSS-AccessKeyId"`
	AccessKeySecret string `json:"AccessKeySecret,omitempty" dc:"阿里云OSS-AccessKeySecret"`
	SecurityToken   string `json:"SecurityToken,omitempty" dc:"阿里云OSS-SecurityToken"`
	Expiration      string `json:"Expiration,omitempty" dc:"Expiration"`
}

/*--------获取Sts Token（APP直传用） 结束--------*/

/*--------回调 开始--------*/
type UploadNotifyReq struct {
	g.Meta   `path:"/notify/:upload_id" method:"get,post" tags:"上传" sm:"回调"`
	UploadId uint `json:"upload_id" v:"required|between:1,4294967295" in:"path" dc:"上传ID"`
}

type UploadNotifyRes struct {
	Url      string `json:"url" dc:"地址"`
	Width    uint   `json:"width" dc:"宽度"`
	Height   uint   `json:"height" dc:"高度"`
	Size     uint   `json:"size" dc:"大小。单位：比特"`
	MimeType string `json:"mime_type" dc:"文件类型"`
}

/*--------回调 结束--------*/
