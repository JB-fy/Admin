package push

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

type PushOfTx struct {
	Ctx       context.Context
	Host      string `json:"host"`
	AccessID  uint32 `json:"accessID"`
	SecretKey string `json:"secretKey"`
}

func NewPushOfTx(ctx context.Context, config map[string]any) *PushOfTx {
	pushObj := PushOfTx{Ctx: ctx}
	gconv.Struct(config, &pushObj)
	if pushObj.Host == `` || pushObj.AccessID == 0 || pushObj.SecretKey == `` {
		panic(`缺少插件配置：推送-腾讯移动推送`)
	}
	return &pushObj
}

func (pushThis *PushOfTx) Push(param PushParam) (err error) {
	reqData := g.Map{}
	reqData[`environment`] = `product`
	if param.IsDev {
		reqData[`environment`] = `dev`
	}

	switch param.MessageType {
	case 0:
		reqData[`message_type`] = `notify`
	case 1:
		reqData[`message_type`] = `message`
	}

	switch param.Audience {
	case 0:
		reqData[`audience_type`] = `all`
	case 1:
		lenOfTokenList := len(param.TokenList)
		if lenOfTokenList > 1000 {
			err = errors.New(`token不能超过1000个`)
			return
		}
		reqData[`audience_type`] = `token_list`
		if lenOfTokenList == 1 {
			reqData[`audience_type`] = `token`
		}
		reqData[`token_list`] = param.TokenList
	case 2:
		reqData[`audience_type`] = `tag`
		reqData[`tag_rules`] = param.TagRules
	}

	message := g.Map{
		`title`:   param.Title,
		`content`: param.Content,
	}
	switch param.DeviceType {
	case 0: //安卓
		message[`android`] = g.Map{
			`custom_content`: gjson.MustEncodeString(param.CustomContent),
		}
	case 1, 2: //IOS //MacOS
		message[`ios`] = g.Map{
			`aps`: g.Map{
				`alert`:           g.Map{},
				`mutable-content`: 1,
			},
			`custom_content`: gjson.MustEncodeString(param.CustomContent),
		}
	}
	reqData[`message`] = message

	reqDataJson := gjson.MustEncodeString(reqData)
	res, err := pushThis.NewHttpClient(reqDataJson).Post(pushThis.Ctx, pushThis.Host+`/v3/push/app`, reqDataJson)
	if err != nil {
		return
	}
	defer res.Close()
	resStr := res.ReadAllString()
	resData := gjson.New(resStr)
	if !resData.Contains(`ret_code`) {
		err = errors.New(resStr)
		return
	}
	if resData.Get(`ret_code`).Int() != 0 {
		err = errors.New(resData.Get(`err_msg`).String())
		return
	}
	return
}

func (pushThis *PushOfTx) TagHandle(param TagParam) (err error) {
	lenOfTagList := len(param.TagList)
	lenOfTokenList := len(param.TokenList)
	if lenOfTagList > 1 && lenOfTokenList > 1 {
		err = errors.New(`不支持多tag多token同时操作`)
		return
	}

	reqData := g.Map{}
	switch param.OperatorType {
	case 0: //增加
		reqData[`tag_list`] = param.TagList
		reqData[`token_list`] = param.TokenList
		if lenOfTagList == 1 && lenOfTokenList == 1 {
			reqData[`operator_type`] = 1
		} else if lenOfTagList == 1 {
			reqData[`operator_type`] = 7
		} else if lenOfTokenList == 1 {
			reqData[`operator_type`] = 3
		}
	case 1: //删除
		reqData[`tag_list`] = param.TagList
		reqData[`token_list`] = param.TokenList
		if lenOfTagList == 1 && lenOfTokenList == 1 {
			reqData[`operator_type`] = 2
		} else if lenOfTagList == 1 {
			reqData[`operator_type`] = 8
		} else if lenOfTokenList == 1 {
			reqData[`operator_type`] = 4
		}
	}

	reqDataJson := gjson.MustEncodeString(reqData)
	res, err := pushThis.NewHttpClient(reqDataJson).Post(pushThis.Ctx, pushThis.Host+`/v3/device/tag`, reqDataJson)
	if err != nil {
		return
	}
	defer res.Close()
	resStr := res.ReadAllString()
	resData := gjson.New(resStr)
	if !resData.Contains(`ret_code`) {
		err = errors.New(resStr)
		return
	}
	if resData.Get(`ret_code`).Int() != 0 {
		err = errors.New(resData.Get(`err_msg`).String())
		return
	}
	return
}

func (pushThis *PushOfTx) NewHttpClient(reqDataJson string) (client *gclient.Client) {
	/* // Basic Auth 认证
	client = g.Client().SetHeaderMap(g.MapStrStr{
		`Content-Type`:  `application/json`,
		`Authorization`: `Basic ` + base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf(`%d:%s`, pushThis.AccessID, pushThis.SecretKey))),
	}) */

	//签名认证（推荐）
	timeStamp := gtime.Now().Unix()
	client = g.Client().SetHeaderMap(g.MapStrStr{
		`Content-Type`: `application/json`,
		`AccessId`:     gconv.String(pushThis.AccessID),
		`TimeStamp`:    gconv.String(timeStamp),
		`Sign`:         pushThis.CreateSign(timeStamp, reqDataJson),
	})
	return
}

func (pushThis *PushOfTx) CreateSign(timeStamp int64, reqDataJson string) (sign string) {
	h := hmac.New(sha256.New, []byte(pushThis.SecretKey))
	h.Write([]byte(fmt.Sprintf(`%d%d%s`, timeStamp, pushThis.AccessID, reqDataJson)))
	sha := hex.EncodeToString(h.Sum(nil))
	sign = base64.StdEncoding.EncodeToString([]byte(sha))
	return
}
