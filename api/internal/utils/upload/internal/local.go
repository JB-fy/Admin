package internal

import (
	"context"
	"sort"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/util/gconv"
)

type Local struct {
	Ctx           context.Context
	Url           string `json:"localUploadUrl"`
	SignKey       string `json:"localUploadSignKey"`
	FileUrlPrefix string `json:"localUploadFileUrlPrefix"`
}

func NewLocal(ctx context.Context, config map[string]interface{}) *Local {
	localObj := Local{
		Ctx: ctx,
	}
	gconv.Struct(config, &localObj)
	return &localObj
}

// 生成签名
func (localThis *Local) CreateSign(signData map[string]interface{}) (sign string) {
	keyArr := []string{}
	for k := range signData {
		keyArr = append(keyArr, k)
	}
	sort.Strings(keyArr)
	str := ``
	for _, k := range keyArr {
		str += k + `=` + gconv.String(signData[k]) + `&`
	}
	str += `key=` + localThis.SignKey
	sign = gmd5.MustEncryptString(str)
	return
}
