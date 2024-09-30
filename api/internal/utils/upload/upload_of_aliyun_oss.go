package upload

import (
	"api/internal/utils/common"
	"context"
	"crypto"
	"crypto/hmac"
	"crypto/md5"
	"crypto/rsa"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"hash"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	sts20150401 "github.com/alibabacloud-go/sts-20150401/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

type UploadOfAliyunOss struct {
	Ctx             context.Context
	Host            string `json:"host"`
	Bucket          string `json:"bucket"`
	AccessKeyId     string `json:"accessKeyId"`
	AccessKeySecret string `json:"accessKeySecret"`
	Endpoint        string `json:"endpoint"`
	RoleArn         string `json:"roleArn"`
	CallbackUrl     string `json:"callbackUrl"`
}

func NewUploadOfAliyunOss(ctx context.Context, config map[string]any) *UploadOfAliyunOss {
	uploadObj := UploadOfAliyunOss{Ctx: ctx}
	gconv.Struct(config, &uploadObj)
	if uploadObj.Host == `` || uploadObj.Bucket == `` || uploadObj.AccessKeyId == `` || uploadObj.AccessKeySecret == `` || uploadObj.CallbackUrl == `` || uploadObj.Endpoint == `` || uploadObj.RoleArn == `` {
		panic(`缺少配置：上传-阿里云OSS`)
	}
	return &uploadObj
}

type UploadOfAliyunOssCallback struct {
	Url      string `json:"url"`      //回调地址	utils.GetRequestUrl(ctx, 0) + `/upload/notify`
	Body     string `json:"body"`     //回调参数	`filename=${object}&size=${size}&mime_type=${mimeType}&height=${imageInfo.height}&width=${imageInfo.width}`
	BodyType string `json:"bodyType"` //回调方式	`application/x-www-form-urlencoded`
}

// 本地上传
func (uploadThis *UploadOfAliyunOss) Upload(r *ghttp.Request) (notifyInfo NotifyInfo, err error) {
	return
}

// 获取签名（H5直传用）
func (uploadThis *UploadOfAliyunOss) Sign(param UploadParam) (signInfo SignInfo, err error) {
	bucketHost := uploadThis.GetBucketHost()

	signInfo = SignInfo{
		UploadUrl: bucketHost,
		Host:      bucketHost,
		Dir:       param.Dir,
		Expire:    gconv.Uint(param.Expire),
		IsRes:     0,
	}

	policyBase64 := uploadThis.CreatePolicyBase64(param)
	uploadData := map[string]any{
		`OSSAccessKeyId`:        uploadThis.AccessKeyId,
		`policy`:                string(policyBase64),
		`signature`:             uploadThis.CreateSign(policyBase64),
		`success_action_status`: `200`, //让服务端返回200,不然，默认会返回204
	}
	//是否回调
	if uploadThis.CallbackUrl != `` {
		callback := UploadOfAliyunOssCallback{
			Url:      uploadThis.CallbackUrl,
			Body:     `filename=${object}&size=${size}&mime_type=${mimeType}&height=${imageInfo.height}&width=${imageInfo.width}`,
			BodyType: `application/x-www-form-urlencoded`,
		}
		uploadData[`callback`] = uploadThis.CreateCallbackStr(callback)
		signInfo.IsRes = 1
	}

	signInfo.UploadData = uploadData
	return
}

// 获取配置信息（APP直传前调用）
func (uploadThis *UploadOfAliyunOss) Config(param UploadParam) (config map[string]any, err error) {
	config = map[string]any{
		`endpoint`: uploadThis.Host,
		`bucket`:   uploadThis.Bucket,
		`dir`:      param.Dir,
	}
	//是否回调
	if uploadThis.CallbackUrl != `` {
		config[`callbackUrl`] = uploadThis.CallbackUrl
		config[`callbackBody`] = `filename=${object}&size=${size}&mime_type=${mimeType}&height=${imageInfo.height}&width=${imageInfo.width}`
		config[`callbackBodyType`] = `application/x-www-form-urlencoded`
	}
	return
}

// 获取Sts Token（APP直传用）
func (uploadThis *UploadOfAliyunOss) Sts(param UploadParam) (stsInfo map[string]any, err error) {
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
	stsInfo, err = common.CreateStsToken(uploadThis.Ctx, config, assumeRoleRequest)
	return
}

// 回调
func (uploadThis *UploadOfAliyunOss) Notify(r *ghttp.Request) (notifyInfo NotifyInfo, err error) {
	filename := r.Get(`filename`).String()
	notifyInfo.Width = r.Get(`width`).Uint()
	notifyInfo.Height = r.Get(`height`).Uint()
	notifyInfo.Size = r.Get(`size`).Uint()
	notifyInfo.MimeType = r.Get(`mime_type`).String()

	// 获取OSS的签名
	strAuthorizationBase64 := r.Header.Get(`authorization`)
	if strAuthorizationBase64 == `` {
		err = errors.New(`签名不能为空`)
		return
	}
	byteAuthorization, err := base64.StdEncoding.DecodeString(strAuthorizationBase64)
	if err != nil {
		return
	}

	// 获取OSS的公钥
	publicKeyURLBase64 := r.Header.Get(`x-oss-pub-key-url`)
	if publicKeyURLBase64 == `` {
		err = errors.New(`公钥URL不能为空`)
		return
	}
	publicKeyURL, _ := base64.StdEncoding.DecodeString(publicKeyURLBase64)
	responsePublicKeyURL, err := http.Get(string(publicKeyURL))
	if err != nil {
		return
	}
	publicKeyByte, err := io.ReadAll(responsePublicKeyURL.Body)
	if err != nil {
		return
	}
	defer responsePublicKeyURL.Body.Close()
	publicKey, err := common.ParsePublicKey(string(publicKeyByte))
	if err != nil {
		return
	}

	//以设置的回调地址uploadThis.CallbackUrl为准。原因：r.URL.Path可能不是实际对外的回调地址，如/upload/notify/xx在nginx可能被增加了/api前缀，变成/api/upload/notify/xx
	// parsedURL := r.URL
	parsedURL, err := url.Parse(uploadThis.CallbackUrl)
	if err != nil {
		return
	}
	strURLPathDecode, err := uploadThis.unescapePath(parsedURL.Path, encodePathSegment)
	if err != nil {
		return
	}

	callbackBodyStr := r.GetBodyString()
	strAuth := ``
	if r.URL.RawQuery == `` {
		strAuth = fmt.Sprintf("%s\n%s", strURLPathDecode, callbackBodyStr)
	} else {
		strAuth = fmt.Sprintf("%s?%s\n%s", strURLPathDecode, r.URL.RawQuery, callbackBodyStr)
	}
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(strAuth))
	byteMD5 := md5Ctx.Sum(nil)

	// 验证签名
	err = rsa.VerifyPKCS1v15(publicKey.(*rsa.PublicKey), crypto.MD5, byteMD5, byteAuthorization)
	if err != nil {
		return
	}

	notifyInfo.Url = uploadThis.GetBucketHost() + `/` + filename
	//有时文件信息放地址后面，一起保存在数据库中会更好。比如：苹果手机做瀑布流时需要知道图片宽高，这时就能直接从地址中解析获取
	urlQueryArr := []string{}
	if notifyInfo.Width > 0 {
		urlQueryArr = append(urlQueryArr, `w=`+gconv.String(notifyInfo.Width))
	}
	if notifyInfo.Height > 0 {
		urlQueryArr = append(urlQueryArr, `h=`+gconv.String(notifyInfo.Height))
	}
	if notifyInfo.Size > 0 {
		urlQueryArr = append(urlQueryArr, `s=`+gconv.String(notifyInfo.Size))
	}
	/* if notifyInfo.MimeType != `` {
		urlQueryArr = append(urlQueryArr, `m=`+notifyInfo.MimeType)
	} */
	if len(urlQueryArr) > 0 {
		notifyInfo.Url += `?` + gstr.Join(urlQueryArr, `&`)
	}
	return
}

// 生成签名（web前端直传用）
func (uploadThis *UploadOfAliyunOss) CreateSign(policyBase64 string) (sign string) {
	h := hmac.New(func() hash.Hash { return sha1.New() }, []byte(uploadThis.AccessKeySecret))
	io.WriteString(h, policyBase64)
	signBase64 := base64.StdEncoding.EncodeToString(h.Sum(nil))
	sign = string(signBase64)
	return
}

// 生成PolicyBase64（web前端直传用）
func (uploadThis *UploadOfAliyunOss) CreatePolicyBase64(param UploadParam) (policyBase64 string) {
	policyMap := map[string]any{
		`expiration`: uploadThis.GetGmtIso8601(param.Expire),
		`conditions`: [][]any{
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
func (uploadThis *UploadOfAliyunOss) CreateCallbackStr(callback UploadOfAliyunOssCallback) string {
	callbackParam := map[string]any{
		`callbackUrl`:      callback.Url,
		`callbackBody`:     callback.Body,
		`callbackBodyType`: callback.BodyType,
	}
	callbackStr, _ := json.Marshal(callbackParam)
	callbackBase64 := base64.StdEncoding.EncodeToString(callbackStr)
	return string(callbackBase64)
}

// 获取bucketHost
func (uploadThis *UploadOfAliyunOss) GetBucketHost() string {
	scheme := `https://`
	if gstr.Pos(uploadThis.Host, `https://`) == -1 {
		scheme = `http://`
	}
	return gstr.Replace(uploadThis.Host, scheme, scheme+uploadThis.Bucket+`.`, 1)
}

func (uploadThis *UploadOfAliyunOss) GetGmtIso8601(expireEnd int64) string {
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
func (uploadThis *UploadOfAliyunOss) unescapePath(s string, mode encoding) (string, error) {
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
func (uploadThis *UploadOfAliyunOss) shouldEscape(c byte, mode encoding) bool {
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

func (uploadThis *UploadOfAliyunOss) ishex(c byte) bool {
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

func (uploadThis *UploadOfAliyunOss) unhex(c byte) byte {
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
