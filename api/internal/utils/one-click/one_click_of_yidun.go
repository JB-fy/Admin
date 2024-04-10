package one_click

import (
	"context"
	"errors"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	mobileverify "github.com/yidun/yidun-golang-sdk/yidun/service/mobileverify"
)

type OneClickOfYidun struct {
	Ctx        context.Context
	SecretId   string `json:"oneClickOfYidunSecretId"`
	SecretKey  string `json:"oneClickOfYidunSecretKey"`
	BusinessId string `json:"oneClickOfYidunBusinessId"`
}

func NewOneClickOfYidun(ctx context.Context, config map[string]interface{}) *OneClickOfYidun {
	oneClickOfYidunObj := OneClickOfYidun{
		Ctx: ctx,
	}
	gconv.Struct(config, &oneClickOfYidunObj)
	return &oneClickOfYidunObj
}

func (oneClickThis *OneClickOfYidun) Check(token string, accessToken string) (phone string, err error) {
	client := mobileverify.NewMobileNumberClientWithAccessKey(oneClickThis.SecretId, oneClickThis.SecretKey)

	req := mobileverify.NewMobileNumberGetRequest(oneClickThis.BusinessId)
	req.SetToken(token).SetAccessToken(accessToken)

	res, err := client.GetMobileNumber(req)
	if err != nil {
		g.Log().Error(oneClickThis.Ctx, `易盾一键登录接口错误：`, err)
		return
	}
	if res.GetCode() != 200 {
		err = errors.New(res.GetMsg())
		return
	}
	phone = *res.Data.Phone
	return
}