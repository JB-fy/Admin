package id_card

import (
	"context"
	"errors"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/util/gconv"
)

type AliyunIdCard struct {
	Ctx     context.Context
	Host    string `json:"aliyunIdCardHost"`
	Path    string `json:"aliyunIdCardPath"`
	Appcode string `json:"aliyunIdCardAppcode"`
}

func NewAliyunIdCard(ctx context.Context, config map[string]interface{}) *AliyunIdCard {
	aliyunIdCardObj := AliyunIdCard{
		Ctx: ctx,
	}
	gconv.Struct(config, &aliyunIdCardObj)
	return &aliyunIdCardObj
}

func (idCardThis *AliyunIdCard) Auth(idCardName string, idCardNo string) (idCardInfo map[string]interface{}, err error) {
	// res, err := idCardThis.CreateClient().Get(idCardThis.Ctx, idCardThis.Host+idCardThis.Path+`?cardno=`+idCardNo+`&name=`+url.QueryEscape(idCardName))
	res, err := idCardThis.CreateClient().Get(idCardThis.Ctx, idCardThis.Host+idCardThis.Path, g.Map{
		`cardno`: idCardNo,
		`name`:   idCardName,
	})
	if err != nil {
		return
	}
	defer res.Close()
	resData := gjson.New(res.ReadAllString())
	if resData.Get(`resp.code`).Int() != 0 {
		err = errors.New(resData.Get(`resp.desc`).String())
		return
	}
	idCardInfo = resData.Get(`data`).Map()
	return
}

func (idCardThis *AliyunIdCard) CreateClient() (client *gclient.Client) {
	client = g.Client().SetHeaderMap(g.MapStrStr{`Authorization`: `APPCODE ` + idCardThis.Appcode})
	return
}
