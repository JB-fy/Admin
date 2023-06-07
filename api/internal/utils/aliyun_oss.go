package utils

import (
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
	"strconv"
	"time"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

type AliyunOss struct {
	Ctx             context.Context
	AccessKeyId     string `c:"aliyunOssAccessKeyId"`     //LTAI5tHx81H64BRJA971DPZF,	LTAI5tSjYikt3bX33riHezmk
	AccessKeySecret string `c:"aliyunOssAccessKeySecret"` //nJyNpTtUuIgZqx21FF4G2zi0WHOn51,	k4uRZU6flv73yz1j4LJu9VY5eNlHas
	Host            string `c:"aliyunOssHost"`            //http://oss-cn-hongkong.aliyuncs.com,	https://oss-cn-hangzhou.aliyuncs.com
	Bucket          string `c:"aliyunOssBucket"`          //4724382110,	gamemt
	//BucketHost      string //http://4724382110.oss-cn-hongkong.aliyuncs.com
}

type AliyunOssSignOption struct {
	CallbackUrl string `c:"callbackUrl"` //是否回调服务器。空字符串不回调
	ExpireTime  int    `c:"expireTime"`  //签名有效时间。单位：秒
	Dir         string `c:"dir"`         //上传的文件前缀
	MinSize     int    `c:"MinSize"`     //限制上传的文件大小。单位：字节
	MaxSize     int    `c:"maxSize"`     //限制上传的文件大小。单位：字节
}

func NewAliyunOss(ctx context.Context, config map[string]interface{}) *AliyunOss {
	aliyunOssObj := AliyunOss{
		Ctx: ctx,
	}
	gconv.Struct(config, &aliyunOssObj)
	return &aliyunOssObj
}

// 创建签名（web前端直传用）
func (aliyunOssThis *AliyunOss) CreateSign(option AliyunOssSignOption) (signInfo map[string]interface{}, err error) {
	expireEnd := time.Now().Unix() + int64(option.ExpireTime)
	signInfo = map[string]interface{}{
		"accessid": aliyunOssThis.AccessKeyId,
		"host":     aliyunOssThis.GetBucketHost(),
		"dir":      option.Dir,
		"expire":   expireEnd,
	}

	if option.CallbackUrl != "" {
		callbackParam := map[string]interface{}{
			"callbackUrl":      option.CallbackUrl,
			"callbackBody":     "filename=${object}&size=${size}&mimeType=${mimeType}&height=${imageInfo.height}&width=${imageInfo.width}",
			"callbackBodyType": "application/x-www-form-urlencoded",
		}
		callbackStr, _ := json.Marshal(callbackParam)
		callbackBase64 := base64.StdEncoding.EncodeToString(callbackStr)
		signInfo["callback"] = string(callbackBase64)
	}

	policy := map[string]interface{}{
		"expiration": aliyunOssThis.GetGmtIso8601(expireEnd),
		"conditions": [][]interface{}{
			{"content-length-range", option.MinSize, option.MaxSize},
			{"starts-with", "$key", option.Dir},
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
func (aliyunOssThis *AliyunOss) Notify(r *ghttp.Request) (err error) {
	// 1.获取OSS的签名header和公钥url header
	strAuthorizationBase64 := r.Header.Get("authorization")
	if strAuthorizationBase64 == "" {
		err = NewErrorCode(aliyunOssThis.Ctx, 40000000, err.Error())
		return
	}
	publicKeyURLBase64 := r.Header.Get("x-oss-pub-key-url")
	if publicKeyURLBase64 == "" {
		err = NewErrorCode(aliyunOssThis.Ctx, 40000001, "")
		return
	}

	// 2.获取OSS的签名
	byteAuthorization, _ := base64.StdEncoding.DecodeString(strAuthorizationBase64)

	// 3.获取公钥
	publicKeyURL, _ := base64.StdEncoding.DecodeString(publicKeyURLBase64)
	responsePublicKeyURL, err := http.Get(string(publicKeyURL))
	if err != nil {
		err = NewErrorCode(aliyunOssThis.Ctx, 40000002, err.Error())
		return
	}
	bytePublicKey, err := ioutil.ReadAll(responsePublicKeyURL.Body)
	if err != nil {
		err = NewErrorCode(aliyunOssThis.Ctx, 40000002, err.Error())
		return
	}
	defer responsePublicKeyURL.Body.Close()

	// 4.获取回调body
	bodyContent, err := ioutil.ReadAll(r.Body)
	r.Body.Close()
	if err != nil {
		err = NewErrorCode(aliyunOssThis.Ctx, 40000003, err.Error())
		return
	}
	strCallbackBody := string(bodyContent)
	strURLPathDecode, errUnescape := aliyunOssThis.unescapePath(r.URL.Path, encodePathSegment) //url.PathUnescape(r.URL.Path) for Golang v1.8.2+
	if errUnescape != nil {
		err = NewErrorCode(aliyunOssThis.Ctx, 40000003, err.Error())
		return
	}

	strAuth := ""
	if r.URL.RawQuery == "" {
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
		err = NewErrorCode(aliyunOssThis.Ctx, 40000003, "")
		return
	}
	pubInterface, err := x509.ParsePKIXPublicKey(pubBlock.Bytes)
	if (pubInterface == nil) || (err != nil) {
		err = NewErrorCode(aliyunOssThis.Ctx, 40000003, err.Error())
		return
	}
	pub := pubInterface.(*rsa.PublicKey)

	// 6.验证签名
	errorVerifyPKCS1v15 := rsa.VerifyPKCS1v15(pub, crypto.MD5, byteMD5, byteAuthorization)
	if errorVerifyPKCS1v15 != nil {
		err = NewErrorCode(aliyunOssThis.Ctx, 40000003, errorVerifyPKCS1v15.Error())
		return
	}
	return
}

// 获取bucketHost
func (aliyunOssThis *AliyunOss) GetBucketHost() string {
	scheme := "https://"
	if gstr.Pos(aliyunOssThis.Host, "https://") == -1 {
		scheme = "http://"
	}
	return gstr.Replace(aliyunOssThis.Host, scheme, scheme+aliyunOssThis.Bucket+".", 1)
}

func (aliyunOssThis *AliyunOss) GetGmtIso8601(expireEnd int64) string {
	var tokenExpire = time.Unix(expireEnd, 0).UTC().Format("2006-01-02T15:04:05Z")
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
func (aliyunOssThis *AliyunOss) unescapePath(s string, mode encoding) (string, error) {
	// Count %, check that they're well-formed.
	mode = encodePathSegment
	n := 0
	hasPlus := false
	for i := 0; i < len(s); {
		switch s[i] {
		case '%':
			n++
			if i+2 >= len(s) || !aliyunOssThis.ishex(s[i+1]) || !aliyunOssThis.ishex(s[i+2]) {
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
			if mode == encodeHost && aliyunOssThis.unhex(s[i+1]) < 8 && s[i:i+3] != "%25" {
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
				v := aliyunOssThis.unhex(s[i+1])<<4 | aliyunOssThis.unhex(s[i+2])
				if s[i:i+3] != "%25" && v != ' ' && aliyunOssThis.shouldEscape(v, encodeHost) {
					return "", errors.New("invalid URL escape " + strconv.Quote(string(s[i:i+3])))
				}
			}
			i += 3
		case '+':
			hasPlus = mode == encodeQueryComponent
			i++
		default:
			if (mode == encodeHost || mode == encodeZone) && s[i] < 0x80 && aliyunOssThis.shouldEscape(s[i], mode) {
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
			t[j] = aliyunOssThis.unhex(s[i+1])<<4 | aliyunOssThis.unhex(s[i+2])
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
func (aliyunOssThis *AliyunOss) shouldEscape(c byte, mode encoding) bool {
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

func (aliyunOssThis *AliyunOss) ishex(c byte) bool {
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

func (aliyunOssThis *AliyunOss) unhex(c byte) byte {
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
