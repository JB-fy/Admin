package one_click

import (
	daoPlatform "api/internal/dao/platform"
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

func NewOneClickOfYidun(ctx context.Context, configOpt ...map[string]interface{}) *OneClickOfYidun {
	var config map[string]interface{}
	if len(configOpt) > 0 && len(configOpt[0]) > 0 {
		config = configOpt[0]
	} else {
		configTmp, _ := daoPlatform.Config.Get(ctx, []string{`oneClickOfYidunSecretId`, `oneClickOfYidunSecretKey`, `oneClickOfYidunBusinessId`})
		config = configTmp.Map()
	}

	oneClickOfYidunObj := OneClickOfYidun{Ctx: ctx}
	gconv.Struct(config, &oneClickOfYidunObj)
	return &oneClickOfYidunObj
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
