package one_click

import (
	"context"
	"errors"

	"github.com/gogf/gf/v2/util/gconv"
	mobileverify "github.com/yidun/yidun-golang-sdk/yidun/service/mobileverify"
)

type OneClickOfYidun struct {
	Ctx        context.Context
	SecretId   string `json:"oneClickOfYidunSecretId"`
	SecretKey  string `json:"oneClickOfYidunSecretKey"`
	BusinessId string `json:"oneClickOfYidunBusinessId"`
}

func NewOneClickOfYidun(ctx context.Context, config map[string]any) *OneClickOfYidun {
	oneClickObj := OneClickOfYidun{Ctx: ctx}
	gconv.Struct(config, &oneClickObj)
	if oneClickObj.SecretId == `` || oneClickObj.SecretKey == `` || oneClickObj.BusinessId == `` {
		panic(`缺少插件配置：一键登录-易盾`)
	}
	return &oneClickObj
}

func (oneClickThis *OneClickOfYidun) Check(token string, accessToken string) (phone string, err error) {
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
