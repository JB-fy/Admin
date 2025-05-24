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
	Method  string `json:"method"`
	Key     string `json:"key"`
	KeyName string `json:"key_name"` //不为空时，会在签名的字符串后面追加密钥名称和密钥
	KeySep  string `json:"key_sep"`  //多个字段之间的拼接符
	ValSep  string `json:"val_sep"`  //单个字段名与值的拼接符
}

func NewSign(ctx context.Context, config map[string]any) model.Sign {
	obj := &Sign{}
	gconv.Struct(config, obj)
	if !slices.Contains([]string{`md5`, `hmac-md5`, `hmac-sha1`, `hmac-sha256`, `hmac-sha512`}, obj.Method) || obj.Key == `` {
		panic(`缺少配置：sign-通用`)
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
	for index := range keyArr[:keyArrLen-1] {
		buf.WriteString(keyArr[index])
		buf.WriteString(signThis.ValSep)
		buf.WriteString(gconv.String(data[keyArr[index]]))
		/* if tmp := gvar.New(data[keyArr[index]]); tmp.IsMap() || tmp.IsSlice() {
			buf.Write(gjson.MustEncode(data[keyArr[index]]))
		} else {
			buf.WriteString(gconv.String(data[keyArr[index]]))
		} */
		buf.WriteString(signThis.KeySep)
	}
	buf.WriteString(keyArr[keyArrLen-1])
	buf.WriteString(signThis.ValSep)
	buf.WriteString(gconv.String(data[keyArr[keyArrLen-1]]))
	/* if tmp := gvar.New(data[keyArr[keyArrLen-1]]); tmp.IsMap() || tmp.IsSlice() {
		buf.Write(gjson.MustEncode(data[keyArr[keyArrLen-1]]))
	} else {
		buf.WriteString(gconv.String(data[keyArr[keyArrLen-1]]))
	} */
	if signThis.KeyName != `` {
		buf.WriteString(signThis.KeySep)
		buf.WriteString(signThis.KeyName)
		buf.WriteString(signThis.ValSep)
		buf.WriteString(signThis.Key)
	}

	var h hash.Hash
	switch signThis.Method {
	case `md5`:
		if signThis.KeyName == `` {
			buf.WriteString(signThis.Key)
		}
		h = md5.New()
	case `hmac-md5`:
		h = hmac.New(md5.New, []byte(signThis.Key))
	case `hmac-sha1`:
		h = hmac.New(sha1.New, []byte(signThis.Key))
	case `hmac-sha256`:
		h = hmac.New(sha256.New, []byte(signThis.Key))
	case `hmac-sha512`:
		h = hmac.New(sha512.New, []byte(signThis.Key))
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
