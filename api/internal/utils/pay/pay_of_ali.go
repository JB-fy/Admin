package pay

import (
	"context"
	"errors"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/smartwalle/alipay/v3"
)

type PayOfAli struct {
	AppId      string `json:"appId"`
	PrivateKey string `json:"privateKey"`
	PublicKey  string `json:"publicKey"`
	OpAppId    string `json:"opAppId"`
	NotifyUrl  string `json:"notifyUrl"`
}

func NewPayOfAli(ctx context.Context, config map[string]any) *PayOfAli {
	payObj := &PayOfAli{}
	gconv.Struct(config, payObj)
	if payObj.AppId == `` || payObj.PrivateKey == `` || payObj.PublicKey == `` || payObj.NotifyUrl == `` {
		panic(`缺少配置：支付-支付宝`)
	}
	return payObj
}

func (payThis *PayOfAli) App(ctx context.Context, payReqData PayReqData) (payResData PayResData, err error) {
	client, err := alipay.New(payThis.AppId, payThis.PrivateKey, true)
	if err != nil {
		return
	}

	param := alipay.TradeAppPay{
		Trade: alipay.Trade{
			Subject:     payReqData.Desc,
			OutTradeNo:  payReqData.OrderNo,
			TotalAmount: gconv.String(payReqData.Amount),
			ProductCode: `QUICK_MSECURITY_PAY`,
			NotifyURL:   payThis.NotifyUrl,
		},
	}
	result, err := client.TradeAppPay(param)
	if err != nil {
		return
	}

	payResData.PayStr = result
	return
}

func (payThis *PayOfAli) H5(ctx context.Context, payReqData PayReqData) (payResData PayResData, err error) {
	client, err := alipay.New(payThis.AppId, payThis.PrivateKey, true)
	if err != nil {
		return
	}

	param := alipay.TradeWapPay{
		Trade: alipay.Trade{
			Subject:     payReqData.Desc,
			OutTradeNo:  payReqData.OrderNo,
			TotalAmount: gconv.String(payReqData.Amount),
			ProductCode: `QUICK_WAP_WAY`,
			NotifyURL:   payThis.NotifyUrl,
		},
	}
	if payReqData.ReturnUrl != `` {
		param.ReturnURL = payReqData.ReturnUrl
	}
	result, err := client.TradeWapPay(param)
	if err != nil {
		return
	}

	payResData.PayStr = result.String()
	return
}

func (payThis *PayOfAli) QRCode(ctx context.Context, payReqData PayReqData) (payResData PayResData, err error) {
	client, err := alipay.New(payThis.AppId, payThis.PrivateKey, true)
	if err != nil {
		return
	}

	param := alipay.TradePreCreate{
		Trade: alipay.Trade{
			Subject:     payReqData.Desc,
			OutTradeNo:  payReqData.OrderNo,
			TotalAmount: gconv.String(payReqData.Amount),
			ProductCode: `FACE_TO_FACE_PAYMENT`,
			NotifyURL:   payThis.NotifyUrl,
		},
	}
	result, err := client.TradePreCreate(param)
	if err != nil {
		return
	}
	if result.Code != alipay.CodeSuccess {
		err = errors.New(result.Msg)
		return
	}

	payResData.PayStr = result.QRCode
	return
}

func (payThis *PayOfAli) Jsapi(ctx context.Context, payReqData PayReqData) (payResData PayResData, err error) {
	client, err := alipay.New(payThis.AppId, payThis.PrivateKey, true)
	if err != nil {
		return
	}
	if payThis.OpAppId == `` {
		err = errors.New(`缺少插件配置：支付-小程序AppID`)
		return
	}

	param := alipay.TradeCreate{
		Trade: alipay.Trade{
			Subject:     payReqData.Desc,
			OutTradeNo:  payReqData.OrderNo,
			TotalAmount: gconv.String(payReqData.Amount),
			ProductCode: `JSAPI_PAY`,
			NotifyURL:   payThis.NotifyUrl,
		},
		// BuyerId:     ``, //买家支付宝用户ID（未来将被废弃）。BuyerId和BuyerOpenId二选一
		BuyerOpenId: payReqData.Openid, //买家支付宝用户OpenId（推荐）。BuyerId和BuyerOpenId二选一
		OpAppId:     payThis.OpAppId,   //小程序应用ID
	}

	result, err := client.TradeCreate(param)
	if err != nil {
		return
	}
	if result.Code != alipay.CodeSuccess {
		err = errors.New(result.Msg)
		return
	}

	payResData.PayStr = result.TradeNo
	return
}

func (payThis *PayOfAli) Notify(ctx context.Context, r *ghttp.Request) (notifyInfo NotifyInfo, err error) {
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
	notifyInfo.ThirdOrderNo = notifyData.TradeNo
	return
}

func (payThis *PayOfAli) NotifyRes(ctx context.Context, r *ghttp.Request, failMsg string) {
	resData := `success` //success:	成功；fail：失败
	if failMsg != `` {
		resData = `fail`
	}
	r.Response.Write(resData)
}
