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
	Url     string `json:"url"`
	Appcode string `json:"appcode"`
	client  *gclient.Client
}

func NewIdCardOfAliyun(ctx context.Context, config map[string]any) *IdCardOfAliyun {
	idCardObj := &IdCardOfAliyun{}
	gconv.Struct(config, idCardObj)
	if idCardObj.Url == `` || idCardObj.Appcode == `` {
		panic(`缺少插件配置：实名认证-阿里云`)
	}
	idCardObj.client = g.Client().SetHeaderMap(g.MapStrStr{`Authorization`: `APPCODE ` + idCardObj.Appcode})
	return idCardObj
}

func (idCardThis *IdCardOfAliyun) Auth(ctx context.Context, idCardName string, idCardNo string) (idCardInfo IdCardInfo, err error) {
	res, err := idCardThis.client.Get(idCardThis.Ctx, idCardThis.Url, g.Map{
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
	if gender, ok := map[string]uint{`男`: 1, `女`: 2}[gconv.String(idCardInfoMap[`sex`])]; ok {
		idCardInfo.Gender = gender
	}
	idCardInfo.Address = gconv.String(idCardInfoMap[`address`])
	// idCardInfo.Birthday = gtime.NewFromStr(gconv.String(idCardInfoMap[`birthday`]))
	idCardInfo.Birthday = gconv.String(idCardInfoMap[`birthday`])
	return
}
