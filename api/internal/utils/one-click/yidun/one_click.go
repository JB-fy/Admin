package yidun

import (
	"context"
	"errors"

	"github.com/gogf/gf/v2/util/gconv"
	mobileverify "github.com/yidun/yidun-golang-sdk/yidun/service/mobileverify"
)

type OneClick struct {
	SecretId   string `json:"secretId"`
	SecretKey  string `json:"secretKey"`
	BusinessId string `json:"businessId"`
}

func NewOneClick(ctx context.Context, config map[string]any) *OneClick {
	obj := &OneClick{}
	gconv.Struct(config, obj)
	if obj.SecretId == `` || obj.SecretKey == `` || obj.BusinessId == `` {
		panic(`缺少插件配置：一键登录-易盾`)
	}
	return obj
}

func (oneClickThis *OneClick) Check(ctx context.Context, token string, accessToken string) (phone string, err error) {
	client := mobileverify.NewMobileNumberClientWithAccessKey(oneClickThis.SecretId, oneClickThis.SecretKey)

	req := mobileverify.NewMobileNumberGetRequest(oneClickThis.BusinessId)
	req.SetToken(token).SetAccessToken(accessToken)

	res, err := client.GetMobileNumber(req)
	if err != nil {
		return
	}
	if res.GetCode() != 200 {
		err = errors.New(res.GetMsg())
		return
	}
	phone = *res.Data.Phone
	return
}
