package id_card

import (
	"context"
	"errors"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/util/gconv"
)

type IdCardOfAliyun struct {
	Ctx     context.Context
	Host    string `json:"host"`
	Path    string `json:"path"`
	Appcode string `json:"appcode"`
}

func NewIdCardOfAliyun(ctx context.Context, config map[string]any) *IdCardOfAliyun {
	idCardObj := IdCardOfAliyun{Ctx: ctx}
	gconv.Struct(config, &idCardObj)
	if idCardObj.Host == `` || idCardObj.Path == `` || idCardObj.Appcode == `` {
		panic(`缺少插件配置：实名认证-阿里云`)
	}
	return &idCardObj
}

func (idCardThis *IdCardOfAliyun) Auth(idCardName string, idCardNo string) (idCardInfo IdCardInfo, err error) {
	res, err := idCardThis.CreateClient().Get(idCardThis.Ctx, idCardThis.Host+idCardThis.Path, g.Map{
		`cardno`: idCardNo,
		`name`:   idCardName,
	})
	if err != nil {
		return
	}
	defer res.Close()
	resStr := res.ReadAllString()
	resData := gjson.New(resStr)
	if !resData.Contains(`resp.code`) {
		err = errors.New(resStr)
		return
	}
	if resData.Get(`resp.code`).Int() != 0 {
		err = errors.New(resData.Get(`resp.desc`).String())
		return
	}

	/* idCardInfoMap示例：{
		"sex":      "男",
		"address":  "福建省-泉州市-晋江市",
		"birthday": "1986-08-30",
	} */
	idCardInfoMap := resData.Get(`data`).Map()
	switch gconv.String(idCardInfoMap[`sex`]) {
	case `男`:
		idCardInfo.Gender = 1
	case `女`:
		idCardInfo.Gender = 2
	}
	idCardInfo.Address = gconv.String(idCardInfoMap[`address`])
	// idCardInfo.Birthday = gtime.NewFromStr(gconv.String(idCardInfoMap[`birthday`]))
	idCardInfo.Birthday = gconv.String(idCardInfoMap[`birthday`])
	return
}

func (idCardThis *IdCardOfAliyun) CreateClient() (client *gclient.Client) {
	client = g.Client().SetHeaderMap(g.MapStrStr{`Authorization`: `APPCODE ` + idCardThis.Appcode})
	return
}
