package aliyun_oss

type UploadCallback struct {
	Url      string `json:"url"`      //回调地址	utils.GetRequestUrl(ctx, 0) + `/upload/notify`
	Body     string `json:"body"`     //回调参数	`filename=${object}&size=${size}&mime_type=${mimeType}&height=${imageInfo.height}&width=${imageInfo.width}`
	BodyType string `json:"bodyType"` //回调方式	`application/x-www-form-urlencoded`
}
