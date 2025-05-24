package common

import (
	"api/internal/utils"
	"api/internal/utils/sign/model"
	"context"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"hash"
	"slices"
	"sort"

	"github.com/gogf/gf/v2/util/gconv"
)

type Sign struct {
	SignMethod string `json:"sign_method"`
	SignKey    string `json:"sign_key"`
	SignName   string `json:"sign_name"` //不为空时，在签名字符串后追加&sign_name=sign_key
}

func NewSign(ctx context.Context, config map[string]any) model.Sign {
	obj := &Sign{}
	gconv.Struct(config, obj)
	if !slices.Contains([]string{`md5`, `hmac-md5`, `hmac-sha1`, `hmac-sha256`}, obj.SignMethod) || obj.SignKey == `` {
		panic(`sign-通用`)
	}
	return obj
}

func (signThis *Sign) Create(ctx context.Context, data map[string]any) (sign string) {
	keyArr := make([]string, 0, len(data))
	for key := range data {
		keyArr = append(keyArr, key)
	}
	sort.Strings(keyArr)

	buf := utils.BytesBufferPoolGet()
	defer utils.BytesBufferPoolPut(buf)
	keyArrLen := len(keyArr)
	for _, key := range keyArr[:keyArrLen-1] {
		buf.WriteString(key)
		buf.WriteString(`=`)
		/* if tmp := gvar.New(data[key]); tmp.IsMap() || tmp.IsSlice() {
			buf.Write(gjson.MustEncode(data[key]))
		} else {
			buf.WriteString(gconv.String(data[key]))
		} */
		buf.WriteString(gconv.String(data[key]))
		buf.WriteString(`&`)
	}
	buf.WriteString(keyArr[keyArrLen-1])
	buf.WriteString(`=`)
	/* if tmp := gvar.New(data[keyArr[keyArrLen-1]]); tmp.IsMap() || tmp.IsSlice() {
		buf.Write(gjson.MustEncode(data[keyArr[keyArrLen-1]]))
	} else {
		buf.WriteString(gconv.String(data[keyArr[keyArrLen-1]]))
	} */
	buf.WriteString(gconv.String(data[keyArr[keyArrLen-1]]))
	if signThis.SignName != `` {
		buf.WriteString(`&`)
		buf.WriteString(signThis.SignName)
		buf.WriteString(`=`)
		buf.WriteString(signThis.SignKey)
	}

	var h hash.Hash
	switch signThis.SignMethod {
	case `md5`:
		if signThis.SignName == `` {
			buf.WriteString(signThis.SignKey)
		}
		h = md5.New()
	case `hmac-md5`:
		h = hmac.New(md5.New, []byte(signThis.SignKey))
	case `hmac-sha1`:
		h = hmac.New(sha1.New, []byte(signThis.SignKey))
	case `hmac-sha256`:
		h = hmac.New(sha256.New, []byte(signThis.SignKey))
	case `hmac-sha512`:
		h = hmac.New(sha512.New, []byte(signThis.SignKey))
	}
	h.Write(buf.Bytes())
	sign = hex.EncodeToString(h.Sum(nil)) //fmt.Sprintf("%x", h.Sum(nil))
	return
}

func (signThis *Sign) Verify(ctx context.Context, data map[string]any, sign string) (err error) {
	if sign == `` {
		err = errors.New(`缺少签名`)
		return
	}
	if sign != signThis.Create(ctx, data) {
		err = errors.New(`签名错误`)
		return
	}
	return
}
