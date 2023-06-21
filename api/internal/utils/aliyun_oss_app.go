package utils

import (
	"context"
	"crypto/hmac"
	"crypto/sha1"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
)

const (
	StsSignVersion = "1.0"
	StsAPIVersion  = "2015-04-01"
	StsHost        = "https://sts.aliyuncs.com/"
	TimeFormat     = "2006-01-02T15:04:05Z"
	RespBodyFormat = "JSON"
	PercentEncode  = "%2F"
	HTTPGet        = "GET"
)

type AliyunOssApp struct {
	Ctx             context.Context
	AccessKeyId     string `json:"aliyunOssAccessKeyId"`     // LTAI5t9jGNGpb9hhtV8M8q2x
	AccessKeySecret string `json:"aliyunOssAccessKeySecret"` // vhfbJ2QAZsFoTZ6m5XF0qwikqWeR0x
	RoleArn         string `json:"aliyunOssRoleArn"`         // acs:ram::1359390739767110:role/aliyunosstokengeneratorrole
}

type AliyunOssAppStsOption struct {
	SessionName string `c:"sessionName"` //可自定义
	CallbackUrl string `c:"callbackUrl"` //是否回调服务器。空字符串不回调
	ExpireTime  int    `c:"expireTime"`  //签名有效时间。单位：秒
	Policy      string `json:"policy"`   //写入权限：{"Statement": [{"Action": ["oss:PutObject","oss:ListParts","oss:AbortMultipartUpload"],"Effect": "Allow","Resource": ["acs:oss:*:*:$BUCKET_NAME/$OBJECT_PREFIX*"]}],"Version": "1"}。读取权限：{"Statement": [{"Action": ["oss:GetObject"],"Effect": "Allow","Resource": ["acs:oss:*:*:$BUCKET_NAME/$OBJECT_PREFIX*"]}],"Version": "1"}
}

func NewAliyunOssApp(ctx context.Context, config map[string]interface{}) *AliyunOssApp {
	aliyunOssAppObj := AliyunOssApp{
		Ctx: ctx,
	}
	gconv.Struct(config, &aliyunOssAppObj)
	return &aliyunOssAppObj
}

func (aliyunOssAppThis *AliyunOssApp) GetStsToken(option AliyunOssAppStsOption) (stsInfo map[string]interface{}, err error) {
	url, err := aliyunOssAppThis.GenerateSignedURL(option)
	if err != nil {
		err = NewErrorCode(aliyunOssAppThis.Ctx, 40000004, err.Error())
		return
	}

	body, status, err := aliyunOssAppThis.SendRequest(url)
	if err != nil {
		err = NewErrorCode(aliyunOssAppThis.Ctx, 40000005, err.Error())
		return
	}
	if status != http.StatusOK {
		err = NewErrorCode(aliyunOssAppThis.Ctx, 40000005, ``)
		return
	}

	type Credentials struct {
		AccessKeyId     string `json:"AccessKeyId"`
		AccessKeySecret string `json:"AccessKeySecret"`
		Expiration      string `json:"Expiration"`
		SecurityToken   string `json:"SecurityToken"`
	}
	type AssumedRoleUser struct {
		Arn           string `json:"Arn"`
		AssumedRoleId string `json:"AssumedRoleId"`
	}
	type StsResponse struct {
		Credentials     Credentials     `json:"Credentials"`
		AssumedRoleUser AssumedRoleUser `json:"AssumedRoleUser"`
		RequestId       string          `json:"RequestId"`
	}
	resp := StsResponse{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		err = NewErrorCode(aliyunOssAppThis.Ctx, 40000005, ``)
		return
	}
	stsInfo = map[string]interface{}{
		`StatusCode`:      http.StatusOK,
		`AccessKeyId`:     resp.Credentials.AccessKeyId,
		`AccessKeySecret`: resp.Credentials.AccessKeySecret,
		`SecurityToken`:   resp.Credentials.SecurityToken,
		`Expiration`:      resp.Credentials.Expiration,
	}
	return
}

func (aliyunOssAppThis *AliyunOssApp) GenerateSignedURL(option AliyunOssAppStsOption) (string, error) {
	rand.Seed(time.Now().UnixNano())
	queryStr := "SignatureVersion=" + StsSignVersion
	queryStr += "&Format=" + RespBodyFormat
	queryStr += "&Timestamp=" + url.QueryEscape(time.Now().UTC().Format(TimeFormat))
	queryStr += "&RoleArn=" + url.QueryEscape(aliyunOssAppThis.RoleArn)
	queryStr += "&RoleSessionName=" + option.SessionName
	queryStr += "&AccessKeyId=" + aliyunOssAppThis.AccessKeyId
	queryStr += "&SignatureMethod=HMAC-SHA1"
	queryStr += "&Version=" + StsAPIVersion
	queryStr += "&Action=AssumeRole"
	queryStr += "&SignatureNonce=" + grand.Str(`ABCDEFGHIJKLMNOPQRSTUVWXYZ`, 10)
	queryStr += "&DurationSeconds=" + gconv.String(option.ExpireTime)
	if option.Policy != "" {
		queryStr += "&Policy=" + url.QueryEscape(option.Policy)
	}

	queryParams, err := url.ParseQuery(queryStr)
	if err != nil {
		return "", err
	}
	sortUrl := strings.Replace(queryParams.Encode(), "+", "%20", -1)
	strToSign := HTTPGet + "&" + PercentEncode + "&" + url.QueryEscape(sortUrl)

	hashSign := hmac.New(sha1.New, []byte(aliyunOssAppThis.AccessKeySecret+"&"))
	hashSign.Write([]byte(strToSign))
	signature := base64.StdEncoding.EncodeToString(hashSign.Sum(nil))

	assumeURL := StsHost + "?" + queryStr + "&Signature=" + url.QueryEscape(signature)
	return assumeURL, nil
}

func (aliyunOssAppThis *AliyunOssApp) SendRequest(url string) ([]byte, int, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	resp, err := client.Get(url)
	if err != nil {
		return nil, -1, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	return body, resp.StatusCode, err
}
