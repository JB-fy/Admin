package push

import (
	"context"
	"errors"

	tpns "git.code.tencent.com/tpns/tpns-server-sdk/gosdk"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/util/gconv"
)

type TxTpns struct {
	Ctx       context.Context
	Host      string `json:"txTpnsHost"`
	AccessID  uint32 `json:"txTpnsAccessID"`
	SecretKey string `json:"txTpnsSecretKey"`
}

func NewTxTpns(ctx context.Context, config map[string]interface{}) *TxTpns {
	txTpnsObj := TxTpns{
		Ctx: ctx,
	}
	gconv.Struct(config, &txTpnsObj)
	return &txTpnsObj
}

func (pushThis *TxTpns) Send(option PushOption) (err error) {
	req := tpns.NewRequest(
		tpns.WithTitle(option.Title),
		tpns.WithContent(option.Content),
	)

	switch option.DeviceType {
	case 0: //安卓
		tpns.WithAndroidMessage(&tpns.AndroidMessage{
			CustomContent: gjson.MustEncodeString(option.CustomContent),
		})(req)
	case 1, 2: //IOS //MacOS
		var aps = tpns.DefaultIOSAps()
		aps.Alert = `大派对`
		tpns.WithIOSMessage(&tpns.IOSMessage{
			Aps:    aps,
			Custom: gjson.MustEncodeString(option.CustomContent),
		})(req)
	}

	req.Environment = tpns.Product
	if option.IsDev {
		req.Environment = tpns.Develop
	}

	switch option.Audience {
	case 0:
		req.Audience = tpns.AudienceAll
	case 1:
		req.Audience = tpns.AudienceToken
		req.TokenList = option.TokenList
	case 2:
		req.Audience = tpns.AudienceTokenList
		req.TokenList = option.TokenList
		if len(option.TokenList) > 1000 {
			err = errors.New(`token不能超过1000个`)
			return
		}
	}

	switch option.MessageType {
	case 0:
		req.MessageType = tpns.Notify
	case 1:
		req.MessageType = tpns.Message
	}
	client := tpns.NewClient(pushThis.Host, pushThis.AccessID, pushThis.SecretKey)
	resp, err := client.Do(req)
	if resp.RetCode != 0 {
		err = errors.New(resp.ErrMsg)
		return
	}
	return
}
