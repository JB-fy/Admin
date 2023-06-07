package utils

import (
	"context"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"hash"
	"io"
	"time"

	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

type AliyunOss struct {
	Ctx             context.Context
	AccessKeyId     string `c:"aliyunOssAccessKeyId"`
	AccessKeySecret string `c:"aliyunOssAccessKeySecret"`
	Host            string `c:"aliyunOssHost"`
	Bucket          string `c:"aliyunOssBucket"`
	//BucketHost      string //http://4724382110.oss-cn-hongkong.aliyuncs.com  //web前端直传地址（内部用getBucketHost方法获取）
}

func NewAliyunOss(ctx context.Context, config map[string]interface{}) *AliyunOss {
	/* return &AliyunOss{
		AccessKeyId:     "LTAI5tHx81H64BRJA971DPZF",            //LTAI5tSjYikt3bX33riHezmk
		AccessKeySecret: "nJyNpTtUuIgZqx21FF4G2zi0WHOn51",      //k4uRZU6flv73yz1j4LJu9VY5eNlHas
		Host:            "http://oss-cn-hongkong.aliyuncs.com", //https://oss-cn-hangzhou.aliyuncs.com
		Bucket:          "4724382110",                          //gamemt
	} */
	aliyunOssObj := AliyunOss{
		Ctx: ctx,
	}
	gconv.Struct(config, &aliyunOssObj)
	return &aliyunOssObj
}

// 创建签名（web前端直传用）
func (aliyunOssThis *AliyunOss) CreateSign(option map[string]interface{}) (signInfo map[string]interface{}, err error) {
	/*--------配置示例 开始--------*/
	/* option := map[string]interface{}{
		"callbackUrl": "",                                                                                   //是否回调服务器。空字符串不回调
		"expireTime":  15 * 60,                                                                              //签名有效时间。单位：秒
		"dir":         fmt.Sprintf("common/%s_%d_", gtime.Now().Format("Y-m-d H:i:s"), grand.N(1000, 9999)), //上传的文件前缀
		"minSize":     0,                                                                                    //限制上传的文件大小。单位：字节
		"maxSize":     100 * 1024 * 1024,                                                                    //限制上传的文件大小。单位：字节
	} */
	/*--------配置示例 结束--------*/
	expireEnd := time.Now().Unix() + gconv.Int64(option["expireTime"])
	signInfo = map[string]interface{}{
		"accessid": aliyunOssThis.AccessKeyId,
		"host":     aliyunOssThis.GetBucketHost(),
		"dir":      option["dir"],
		"expire":   expireEnd,
	}

	if gconv.String(option["callbackUrl"]) != "" {
		callbackParam := map[string]interface{}{
			"callbackUrl":      option["callbackUrl"],
			"callbackBody":     "filename=${object}&size=${size}&mimeType=${mimeType}&height=${imageInfo.height}&width=${imageInfo.width}",
			"callbackBodyType": "application/x-www-form-urlencoded",
		}
		callbackStr, _ := json.Marshal(callbackParam)
		callbackBase64 := base64.StdEncoding.EncodeToString(callbackStr)
		signInfo["callback"] = string(callbackBase64)
	}

	policy := map[string]interface{}{
		"expiration": aliyunOssThis.GetGmtIso8601(expireEnd),
		"conditions": [][]string{
			{"content-length-range", gconv.String(option["minSize"]), gconv.String(option["maxSize"])},
			{"starts-with", "$key", gconv.String(option["dir"])},
		},
	}
	policyStr, _ := json.Marshal(policy)
	policyBase64 := base64.StdEncoding.EncodeToString(policyStr)
	signInfo["policy"] = string(policyBase64)

	h := hmac.New(func() hash.Hash { return sha1.New() }, []byte(aliyunOssThis.AccessKeySecret))
	io.WriteString(h, policyBase64)
	signedStr := base64.StdEncoding.EncodeToString(h.Sum(nil))
	signInfo["signature"] = string(signedStr)
	return
}

// 回调
func (aliyunOssThis *AliyunOss) Notify(tokenString string) (claims *CustomClaims, err error) {

	return
}

// 获取bucketHost
func (aliyunOssThis *AliyunOss) GetBucketHost() string {
	scheme := "https://"
	if gstr.Pos(aliyunOssThis.Host, "https://") == -1 {
		scheme = "http://"
	}
	return gstr.Replace(aliyunOssThis.Host, scheme, scheme+aliyunOssThis.Bucket, 1)
}

func (aliyunOssThis *AliyunOss) GetGmtIso8601(expireEnd int64) string {
	var tokenExpire = time.Unix(expireEnd, 0).UTC().Format("2006-01-02T15:04:05Z")
	return tokenExpire
}
