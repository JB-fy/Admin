package upload

import (
	"api/internal/utils"
	"api/internal/utils/common"
	"context"
	"crypto"
	"crypto/hmac"
	"crypto/md5"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"hash"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	sts20150401 "github.com/alibabacloud-go/sts-20150401/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

type AliyunOss struct {
	Ctx             context.Context
	Host            string `json:"aliyunOssHost"`
	Bucket          string `json:"aliyunOssBucket"`
	AccessKeyId     string `json:"aliyunOssAccessKeyId"`
	AccessKeySecret string `json:"aliyunOssAccessKeySecret"`
	CallbackUrl     string `json:"aliyunOssCallbackUrl"`
	Endpoint        string `json:"aliyunOssEndpoint"`
	RoleArn         string `json:"aliyunOssRoleArn"`
}

func NewAliyunOss(ctx context.Context, config map[string]interface{}) *AliyunOss {
	aliyunOssObj := AliyunOss{
		Ctx: ctx,
	}
	gconv.Struct(config, &aliyunOssObj)
	return &aliyunOssObj
}

type AliyunOssCallback struct {
	Url      string `json:"url"`      //回调地址	utils.GetRequestUrl(ctx, 0) + `/upload/notify`
	Body     string `json:"body"`     //回调参数	`filename=${object}&size=${size}&mimeType=${mimeType}&height=${imageInfo.height}&width=${imageInfo.width}`
	BodyType string `json:"bodyType"` //回调方式	`application/x-www-form-urlencoded`
}

// 本地上传
func (uploadThis *AliyunOss) Upload() (uploadInfo map[string]interface{}, err error) {
	return
}

// 获取签名（H5直传用）
func (uploadThis *AliyunOss) Sign(param UploadParam) (signInfo map[string]interface{}, err error) {
	bucketHost := uploadThis.GetBucketHost()

	signInfo = map[string]interface{}{
		`uploadUrl`: bucketHost,
		// `uploadData`:  map[string]interface{}{},
		`host`:     bucketHost,
		`dir`:      param.Dir,
		`expire`:   param.Expire,
		`isRes`:    0,
		`endpoint`: uploadThis.Host,
		`bucket`:   uploadThis.Bucket,
	}

	policyBase64 := uploadThis.CreatePolicyBase64(param)
	uploadData := map[string]interface{}{
		`OSSAccessKeyId`:        uploadThis.AccessKeyId,
		`policy`:                string(policyBase64),
		`signature`:             uploadThis.CreateSign(policyBase64),
		`success_action_status`: `200`, //让服务端返回200,不然，默认会返回204
	}
	//是否回调
	if uploadThis.CallbackUrl != `` {
		callback := AliyunOssCallback{
			Url:      uploadThis.CallbackUrl,
			Body:     `filename=${object}&size=${size}&mimeType=${mimeType}&height=${imageInfo.height}&width=${imageInfo.width}`,
			BodyType: `application/x-www-form-urlencoded`,
		}
		uploadData[`callback`] = uploadThis.CreateCallbackStr(callback)
		signInfo[`isRes`] = 1
		signInfo[`callbackUrl`] = uploadThis.CallbackUrl
		signInfo[`callbackBody`] = `filename=${object}&size=${size}&mimeType=${mimeType}&height=${imageInfo.height}&width=${imageInfo.width}`
		signInfo[`callbackBodyType`] = `application/x-www-form-urlencoded`
	}

	signInfo[`uploadData`] = uploadData
	return
}

// 获取配置信息（APP直传前调用）
func (uploadThis *AliyunOss) Config(param UploadParam) (config map[string]interface{}, err error) {
	config = map[string]interface{}{
		`endpoint`: uploadThis.Host,
		`bucket`:   uploadThis.Bucket,
		`dir`:      param.Dir,
	}
	//是否回调
	if uploadThis.CallbackUrl != `` {
		config[`callbackUrl`] = uploadThis.CallbackUrl
		config[`callbackBody`] = `filename=${object}&size=${size}&mimeType=${mimeType}&height=${imageInfo.height}&width=${imageInfo.width}`
		config[`callbackBodyType`] = `application/x-www-form-urlencoded`
	}
	return
}

// 获取Sts Token（APP直传用）
func (uploadThis *AliyunOss) Sts(param UploadParam) (stsInfo map[string]interface{}, err error) {
	config := &openapi.Config{
		AccessKeyId:     tea.String(uploadThis.AccessKeyId),
		AccessKeySecret: tea.String(uploadThis.AccessKeySecret),
		Endpoint:        tea.String(uploadThis.Endpoint),
	}
	assumeRoleRequest := &sts20150401.AssumeRoleRequest{
		DurationSeconds: tea.Int64(param.ExpireTime),
		Policy:          tea.String(`{"Statement": [{"Action": ["oss:PutObject","oss:ListParts","oss:AbortMultipartUpload"],"Effect": "Allow","Resource": ["acs:oss:*:*:` + uploadThis.Bucket + `/` + param.Dir + `*"]}],"Version": "1"}`),
		RoleArn:         tea.String(uploadThis.RoleArn),
		RoleSessionName: tea.String(`sts_token_to_oss`),
	}
	stsInfo, _ = common.CreateStsToken(uploadThis.Ctx, config, assumeRoleRequest)
	return
}

// 回调
func (uploadThis *AliyunOss) Notify() (notifyInfo map[string]interface{}, err error) {
	r := g.RequestFromCtx(uploadThis.Ctx)
	filename := r.Get(`filename`).String()
	size := r.Get(`size`).String()
	// mimeType := r.Get(`mimeType`).String()
	width := r.Get(`width`).String()
	height := r.Get(`height`).String()

	// 1.获取OSS的签名header和公钥url header
	strAuthorizationBase64 := r.Header.Get(`authorization`)
	if strAuthorizationBase64 == `` {
		err = utils.NewErrorCode(uploadThis.Ctx, 79990000, err.Error())
		return
	}
	publicKeyURLBase64 := r.Header.Get(`x-oss-pub-key-url`)
	if publicKeyURLBase64 == `` {
		err = utils.NewErrorCode(uploadThis.Ctx, 79990001, ``)
		return
	}

	// 2.获取OSS的签名
	byteAuthorization, _ := base64.StdEncoding.DecodeString(strAuthorizationBase64)

	// 3.获取公钥
	publicKeyURL, _ := base64.StdEncoding.DecodeString(publicKeyURLBase64)
	responsePublicKeyURL, err := http.Get(string(publicKeyURL))
	if err != nil {
		err = utils.NewErrorCode(uploadThis.Ctx, 79990002, err.Error())
		return
	}
	bytePublicKey, err := ioutil.ReadAll(responsePublicKeyURL.Body)
	if err != nil {
		err = utils.NewErrorCode(uploadThis.Ctx, 79990002, err.Error())
		return
	}
	defer responsePublicKeyURL.Body.Close()

	// 4.获取回调body
	bodyContent, err := ioutil.ReadAll(r.Body)
	r.Body.Close()
	if err != nil {
		err = utils.NewErrorCode(uploadThis.Ctx, 79990003, err.Error())
		return
	}
	strCallbackBody := string(bodyContent)

	//以设置的回调地址uploadThis.CallbackUrl为准。原因：r.URL.Path可能不是实际对外的回调地址，如upload/notify在nginx可能被增加了api/前缀，变成api/upload/notify
	parsedURL, _ := url.Parse(uploadThis.CallbackUrl)
	// strURLPathDecode, err := uploadThis.unescapePath(r.URL.Path, encodePathSegment)
	strURLPathDecode, err := uploadThis.unescapePath(parsedURL.Path, encodePathSegment)
	if err != nil {
		err = utils.NewErrorCode(uploadThis.Ctx, 79990003, err.Error())
		return
	}

	strAuth := ``
	if r.URL.RawQuery == `` {
		strAuth = fmt.Sprintf("%s\n%s", strURLPathDecode, strCallbackBody)
	} else {
		strAuth = fmt.Sprintf("%s?%s\n%s", strURLPathDecode, r.URL.RawQuery, strCallbackBody)
	}

	md5Ctx := md5.New()
	md5Ctx.Write([]byte(strAuth))
	byteMD5 := md5Ctx.Sum(nil)

	// 5.拼接待签名字符串
	pubBlock, _ := pem.Decode(bytePublicKey)
	if pubBlock == nil {
		err = utils.NewErrorCode(uploadThis.Ctx, 79990003, ``)
		return
	}
	pubInterface, err := x509.ParsePKIXPublicKey(pubBlock.Bytes)
	if (pubInterface == nil) || (err != nil) {
		err = utils.NewErrorCode(uploadThis.Ctx, 79990003, err.Error())
		return
	}
	pub := pubInterface.(*rsa.PublicKey)

	// 6.验证签名
	err = rsa.VerifyPKCS1v15(pub, crypto.MD5, byteMD5, byteAuthorization)
	if err != nil {
		err = utils.NewErrorCode(uploadThis.Ctx, 79990003, err.Error())
		return
	}

	notifyInfo = map[string]interface{}{}
	notifyInfo[`url`] = uploadThis.GetBucketHost() + `/` + filename + `?w=` + width + `&h=` + height + `&s=` + size /* + `&m=` + mimeType */
	return
}

// 生成签名（web前端直传用）
func (uploadThis *AliyunOss) CreateSign(policyBase64 string) (sign string) {
	h := hmac.New(func() hash.Hash { return sha1.New() }, []byte(uploadThis.AccessKeySecret))
	io.WriteString(h, policyBase64)
	signBase64 := base64.StdEncoding.EncodeToString(h.Sum(nil))
	sign = string(signBase64)
	return
}

// 生成PolicyBase64（web前端直传用）
func (uploadThis *AliyunOss) CreatePolicyBase64(param UploadParam) (policyBase64 string) {
	policyMap := map[string]interface{}{
		`expiration`: uploadThis.GetGmtIso8601(param.Expire),
		`conditions`: [][]interface{}{
			{`content-length-range`, param.MinSize, param.MaxSize},
			{`starts-with`, `$key`, param.Dir},
		},
	}
	policyStr, _ := json.Marshal(policyMap)
	policyBase64 = base64.StdEncoding.EncodeToString(policyStr)
	// policy = string(policy)
	return
}

// 生成回调字符串（web前端直传用）
func (uploadThis *AliyunOss) CreateCallbackStr(callback AliyunOssCallback) string {
	callbackParam := map[string]interface{}{
		`callbackUrl`:      callback.Url,
		`callbackBody`:     callback.Body,
		`callbackBodyType`: callback.BodyType,
	}
	callbackStr, _ := json.Marshal(callbackParam)
	callbackBase64 := base64.StdEncoding.EncodeToString(callbackStr)
	return string(callbackBase64)
}

// 获取bucketHost
func (uploadThis *AliyunOss) GetBucketHost() string {
	scheme := `https://`
	if gstr.Pos(uploadThis.Host, `https://`) == -1 {
		scheme = `http://`
	}
	return gstr.Replace(uploadThis.Host, scheme, scheme+uploadThis.Bucket+`.`, 1)
}

func (uploadThis *AliyunOss) GetGmtIso8601(expireEnd int64) string {
	var tokenExpire = time.Unix(expireEnd, 0).UTC().Format(`2006-01-02T15:04:05Z`)
	return tokenExpire
}

type encoding int

const (
	encodePath encoding = 1 + iota
	encodePathSegment
	encodeHost
	encodeZone
	encodeUserPassword
	encodeQueryComponent
	encodeFragment
)

// unescapePath : unescapes a string; the mode specifies, which section of the URL string is being unescaped.
func (uploadThis *AliyunOss) unescapePath(s string, mode encoding) (string, error) {
	// Count %, check that they're well-formed.
	mode = encodePathSegment
	n := 0
	hasPlus := false
	for i := 0; i < len(s); {
		switch s[i] {
		case '%':
			n++
			if i+2 >= len(s) || !uploadThis.ishex(s[i+1]) || !uploadThis.ishex(s[i+2]) {
				s = s[i:]
				if len(s) > 3 {
					s = s[:3]
				}
				return "", errors.New("invalid URL escape " + strconv.Quote(string(s)))
			}
			// Per https://tools.ietf.org/html/rfc3986#page-21
			// in the host component %-encoding can only be used
			// for non-ASCII bytes.
			// But https://tools.ietf.org/html/rfc6874#section-2
			// introduces %25 being allowed to escape a percent sign
			// in IPv6 scoped-address literals. Yay.
			if mode == encodeHost && uploadThis.unhex(s[i+1]) < 8 && s[i:i+3] != "%25" {
				return "", errors.New("invalid URL escape " + strconv.Quote(string(s[i:i+3])))
			}
			if mode == encodeZone {
				// RFC 6874 says basically "anything goes" for zone identifiers
				// and that even non-ASCII can be redundantly escaped,
				// but it seems prudent to restrict %-escaped bytes here to those
				// that are valid host name bytes in their unescaped form.
				// That is, you can use escaping in the zone identifier but not
				// to introduce bytes you couldn't just write directly.
				// But Windows puts spaces here! Yay.
				v := uploadThis.unhex(s[i+1])<<4 | uploadThis.unhex(s[i+2])
				if s[i:i+3] != "%25" && v != ' ' && uploadThis.shouldEscape(v, encodeHost) {
					return "", errors.New("invalid URL escape " + strconv.Quote(string(s[i:i+3])))
				}
			}
			i += 3
		case '+':
			hasPlus = mode == encodeQueryComponent
			i++
		default:
			if (mode == encodeHost || mode == encodeZone) && s[i] < 0x80 && uploadThis.shouldEscape(s[i], mode) {
				return "", errors.New("invalid character " + strconv.Quote(string(s[i:i+1])) + " in host name")
			}
			i++
		}
	}

	if n == 0 && !hasPlus {
		return s, nil
	}

	t := make([]byte, len(s)-2*n)
	j := 0
	for i := 0; i < len(s); {
		switch s[i] {
		case '%':
			t[j] = uploadThis.unhex(s[i+1])<<4 | uploadThis.unhex(s[i+2])
			j++
			i += 3
		case '+':
			if mode == encodeQueryComponent {
				t[j] = ' '
			} else {
				t[j] = '+'
			}
			j++
			i++
		default:
			t[j] = s[i]
			j++
			i++
		}
	}
	return string(t), nil
}

// Please be informed that for now shouldEscape does not check all
// reserved characters correctly. See golang.org/issue/5684.
func (uploadThis *AliyunOss) shouldEscape(c byte, mode encoding) bool {
	// §2.3 Unreserved characters (alphanum)
	if 'A' <= c && c <= 'Z' || 'a' <= c && c <= 'z' || '0' <= c && c <= '9' {
		return false
	}

	if mode == encodeHost || mode == encodeZone {
		// §3.2.2 Host allows
		//	sub-delims = "!" / "$" / "&" / "'" / "(" / ")" / "*" / "+" / "," / ";" / "="
		// as part of reg-name.
		// We add : because we include :port as part of host.
		// We add [ ] because we include [ipv6]:port as part of host.
		// We add < > because they're the only characters left that
		// we could possibly allow, and Parse will reject them if we
		// escape them (because hosts can't use %-encoding for
		// ASCII bytes).
		switch c {
		case '!', '$', '&', '\'', '(', ')', '*', '+', ',', ';', '=', ':', '[', ']', '<', '>', '"':
			return false
		}
	}

	switch c {
	case '-', '_', '.', '~': // §2.3 Unreserved characters (mark)
		return false

	case '$', '&', '+', ',', '/', ':', ';', '=', '?', '@': // §2.2 Reserved characters (reserved)
		// Different sections of the URL allow a few of
		// the reserved characters to appear unescaped.
		switch mode {
		case encodePath: // §3.3
			// The RFC allows : @ & = + $ but saves / ; , for assigning
			// meaning to individual path segments. This package
			// only manipulates the path as a whole, so we allow those
			// last three as well. That leaves only ? to escape.
			return c == '?'

		case encodePathSegment: // §3.3
			// The RFC allows : @ & = + $ but saves / ; , for assigning
			// meaning to individual path segments.
			return c == '/' || c == ';' || c == ',' || c == '?'

		case encodeUserPassword: // §3.2.1
			// The RFC allows ';', ':', '&', '=', '+', '$', and ',' in
			// userinfo, so we must escape only '@', '/', and '?'.
			// The parsing of userinfo treats ':' as special so we must escape
			// that too.
			return c == '@' || c == '/' || c == '?' || c == ':'

		case encodeQueryComponent: // §3.4
			// The RFC reserves (so we must escape) everything.
			return true

		case encodeFragment: // §4.1
			// The RFC text is silent but the grammar allows
			// everything, so escape nothing.
			return false
		}
	}

	// Everything else must be escaped.
	return true
}

func (uploadThis *AliyunOss) ishex(c byte) bool {
	switch {
	case '0' <= c && c <= '9':
		return true
	case 'a' <= c && c <= 'f':
		return true
	case 'A' <= c && c <= 'F':
		return true
	}
	return false
}

func (uploadThis *AliyunOss) unhex(c byte) byte {
	switch {
	case '0' <= c && c <= '9':
		return c - '0'
	case 'a' <= c && c <= 'f':
		return c - 'a' + 10
	case 'A' <= c && c <= 'F':
		return c - 'A' + 10
	}
	return 0
}
