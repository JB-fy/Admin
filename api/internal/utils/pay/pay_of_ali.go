package pay

import (
	daoPlatform "api/internal/dao/platform"
	"context"
	"errors"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/smartwalle/alipay/v3"
)

type PayOfAli struct {
	Ctx        context.Context
	AppId      string `json:"payOfAliAppId"`
	PrivateKey string `json:"payOfAliPrivateKey"`
	PublicKey  string `json:"payOfAliPublicKey"`
	NotifyUrl  string `json:"payOfAliNotifyUrl"`
	OpAppId    string `json:"payOfAliOpAppId"`
}

func NewPayOfAli(ctx context.Context, configOpt ...map[string]any) *PayOfAli {
	var config map[string]any
	if len(configOpt) > 0 && len(configOpt[0]) > 0 {
		config = configOpt[0]
	} else {
		configTmp, _ := daoPlatform.Config.Get(ctx, []string{`payOfAliAppId`, `payOfAliPrivateKey`, `payOfAliPublicKey`, `payOfAliNotifyUrl`, `payOfAliOpAppId`})
		config = configTmp.Map()
	}

	payOfAliObj := PayOfAli{Ctx: ctx}
	gconv.Struct(config, &payOfAliObj)
	return &payOfAliObj
}

func (payThis *PayOfAli) App(payData PayData) (orderInfo PayInfo, err error) {
	client, err := alipay.New(payThis.AppId, payThis.PrivateKey, true)
	if err != nil {
		return
	}

	param := alipay.TradeAppPay{
		Trade: alipay.Trade{
			Subject:     payData.Desc,
			OutTradeNo:  payData.OrderNo,
			TotalAmount: gconv.String(payData.Amount),
			ProductCode: `QUICK_MSECURITY_PAY`,
			NotifyURL:   payThis.NotifyUrl,
		},
	}
	result, err := client.TradeAppPay(param)
	if err != nil {
		return
	}

	orderInfo.PayStr = result
	return
}

func (payThis *PayOfAli) H5(payData PayData) (orderInfo PayInfo, err error) {
	client, err := alipay.New(payThis.AppId, payThis.PrivateKey, true)
	if err != nil {
		return
	}

	param := alipay.TradeWapPay{
		Trade: alipay.Trade{
			Subject:     payData.Desc,
			OutTradeNo:  payData.OrderNo,
			TotalAmount: gconv.String(payData.Amount),
			ProductCode: `QUICK_WAP_WAY`,
			NotifyURL:   payThis.NotifyUrl,
		},
	}
	if payData.ReturnUrl != `` {
		param.ReturnURL = payData.ReturnUrl
	}
	result, err := client.TradeWapPay(param)
	if err != nil {
		return
	}

	orderInfo.PayStr = result.String()
	return
}

func (payThis *PayOfAli) Jsapi(payData PayData) (orderInfo PayInfo, err error) {
	client, err := alipay.New(payThis.AppId, payThis.PrivateKey, true)
	if err != nil {
		return
	}

	param := alipay.TradeCreate{
		Trade: alipay.Trade{
			Subject:     payData.Desc,
			OutTradeNo:  payData.OrderNo,
			TotalAmount: gconv.String(payData.Amount),
			ProductCode: `JSAPI_PAY`,
			NotifyURL:   payThis.NotifyUrl,
		},
		// BuyerId:     ``, //买家支付宝用户ID（未来将被废弃）。BuyerId和BuyerOpenId二选一
		BuyerOpenId: payData.Openid,  //买家支付宝用户OpenId（推荐）。BuyerId和BuyerOpenId二选一
		OpAppId:     payThis.OpAppId, //小程序应用ID
	}

	result, err := client.TradeCreate(param)
	if err != nil {
		return
	}
	if result.Code != alipay.CodeSuccess {
		err = errors.New(result.Msg)
		return
	}

	orderInfo.PayStr = result.TradeNo
	return
}

func (payThis *PayOfAli) Notify(r *ghttp.Request) (notifyInfo NotifyInfo, err error) {
	client, err := alipay.New(payThis.AppId, payThis.PrivateKey, true)
	if err != nil {
		return
	}
	err = client.LoadAliPayPublicKey(payThis.PublicKey)
	if err != nil {
		return
	}

	notifyData, err := client.DecodeNotification(r.Form)
	if err != nil {
		return
	}

	notifyInfo.Amount = gconv.Float64(notifyData.TotalAmount)
	notifyInfo.OrderNo = notifyData.OutTradeNo
	notifyInfo.OrderNoOfThird = notifyData.TradeNo
	return
}

func (payThis *PayOfAli) NotifyRes(r *ghttp.Request, failMsg string) {
	resData := `success` //success:	成功；fail：失败
	if failMsg != `` {
		resData = `fail`
	}
	r.Response.Write(resData)
}
