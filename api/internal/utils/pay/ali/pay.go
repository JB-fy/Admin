package ali

import (
	"api/internal/utils/pay/model"
	"context"
	"errors"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/smartwalle/alipay/v3"
)

type Pay struct {
	AppId      string `json:"app_id"`
	PrivateKey string `json:"private_key"`
	PublicKey  string `json:"public_key"`
	OpAppId    string `json:"op_app_id"`
	NotifyUrl  string `json:"notify_url"`
}

func NewPay(ctx context.Context, config map[string]any) model.Pay {
	obj := &Pay{}
	gconv.Struct(config, obj)
	if obj.AppId == `` || obj.PrivateKey == `` || obj.PublicKey == `` || obj.NotifyUrl == `` {
		panic(`缺少配置：支付-支付宝`)
	}
	return obj
}

func (payThis *Pay) App(ctx context.Context, payReq model.PayReq) (payRes model.PayRes, err error) {
	client, err := alipay.New(payThis.AppId, payThis.PrivateKey, true)
	if err != nil {
		return
	}

	param := alipay.TradeAppPay{
		Trade: alipay.Trade{
			Subject:     payReq.Desc,
			OutTradeNo:  payReq.OrderNo,
			TotalAmount: gconv.String(payReq.Amount),
			ProductCode: `QUICK_MSECURITY_PAY`,
			NotifyURL:   payThis.NotifyUrl,
		},
	}
	result, err := client.TradeAppPay(param)
	if err != nil {
		return
	}

	payRes.PayStr = result
	return
}

func (payThis *Pay) H5(ctx context.Context, payReq model.PayReq) (payRes model.PayRes, err error) {
	client, err := alipay.New(payThis.AppId, payThis.PrivateKey, true)
	if err != nil {
		return
	}

	param := alipay.TradeWapPay{
		Trade: alipay.Trade{
			Subject:     payReq.Desc,
			OutTradeNo:  payReq.OrderNo,
			TotalAmount: gconv.String(payReq.Amount),
			ProductCode: `QUICK_WAP_WAY`,
			NotifyURL:   payThis.NotifyUrl,
		},
	}
	if payReq.ReturnUrl != `` {
		param.ReturnURL = payReq.ReturnUrl
	}
	result, err := client.TradeWapPay(param)
	if err != nil {
		return
	}

	payRes.PayStr = result.String()
	return
}

func (payThis *Pay) QRCode(ctx context.Context, payReq model.PayReq) (payRes model.PayRes, err error) {
	client, err := alipay.New(payThis.AppId, payThis.PrivateKey, true)
	if err != nil {
		return
	}

	param := alipay.TradePreCreate{
		Trade: alipay.Trade{
			Subject:     payReq.Desc,
			OutTradeNo:  payReq.OrderNo,
			TotalAmount: gconv.String(payReq.Amount),
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

	payRes.PayStr = result.QRCode
	return
}

func (payThis *Pay) Jsapi(ctx context.Context, payReq model.PayReq) (payRes model.PayRes, err error) {
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
			Subject:     payReq.Desc,
			OutTradeNo:  payReq.OrderNo,
			TotalAmount: gconv.String(payReq.Amount),
			ProductCode: `JSAPI_PAY`,
			NotifyURL:   payThis.NotifyUrl,
		},
		// BuyerId:     ``, //买家支付宝用户ID（未来将被废弃）。BuyerId和BuyerOpenId二选一
		BuyerOpenId: payReq.Openid,   //买家支付宝用户OpenId（推荐）。BuyerId和BuyerOpenId二选一
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

	payRes.PayStr = result.TradeNo
	return
}

func (payThis *Pay) Notify(ctx context.Context, r *ghttp.Request) (notifyInfo model.NotifyInfo, err error) {
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

func (payThis *Pay) NotifyRes(ctx context.Context, r *ghttp.Request, failMsg string) {
	resData := `success` //success:	成功；fail：失败
	if failMsg != `` {
		resData = `fail`
	}
	r.Response.Write(resData)
}
