package wx

import (
	"api/internal/utils/pay/model"
	"context"
	"errors"
	"net/url"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/verifiers"
	"github.com/wechatpay-apiv3/wechatpay-go/core/downloader"
	"github.com/wechatpay-apiv3/wechatpay-go/core/notify"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/app"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/h5"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/jsapi"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/native"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
)

type Pay struct {
	Ctx        context.Context
	AppId      string `json:"appId"`
	Mchid      string `json:"mchid"`
	SerialNo   string `json:"serialNo"`
	APIv3Key   string `json:"apiV3Key"`
	PrivateKey string `json:"privateKey"`
	NotifyUrl  string `json:"notifyUrl"`
}

func NewPay(ctx context.Context, config map[string]any) model.Pay {
	obj := &Pay{}
	gconv.Struct(config, obj)
	if obj.AppId == `` || obj.Mchid == `` || obj.SerialNo == `` || obj.APIv3Key == `` || obj.PrivateKey == `` || obj.NotifyUrl == `` {
		panic(`缺少配置：支付-微信`)
	}
	return obj
}

func (payThis *Pay) App(ctx context.Context, payReq model.PayReq) (payRes model.PayRes, err error) {
	privateKey, err := utils.LoadPrivateKey(payThis.PrivateKey)
	if err != nil {
		return
	}
	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	client, err := core.NewClient(ctx, option.WithWechatPayAutoAuthCipher(payThis.Mchid, payThis.SerialNo, privateKey, payThis.APIv3Key))
	if err != nil {
		return
	}

	// 发送请求
	svc := app.AppApiService{Client: client}
	resp, result, err := svc.Prepay(ctx,
		app.PrepayRequest{
			Appid:       core.String(payThis.AppId),
			Mchid:       core.String(payThis.Mchid),
			Description: core.String(payReq.Desc),
			OutTradeNo:  core.String(payReq.OrderNo),
			NotifyUrl:   core.String(payThis.NotifyUrl),
			Amount: &app.Amount{
				// Total: core.Int64(gconv.Int64(math.Ceil(payReq.Amount * 100))),
				Total: core.Int64(gconv.Int64(payReq.Amount * 100)),
			},
		},
	)
	if err != nil {
		return
	}
	if result.Response.StatusCode != 200 {
		err = errors.New(`响应错误`)
		return
	}
	payRes.PayStr = *resp.PrepayId
	return
}

func (payThis *Pay) H5(ctx context.Context, payReq model.PayReq) (payRes model.PayRes, err error) {
	privateKey, err := utils.LoadPrivateKey(payThis.PrivateKey)
	if err != nil {
		return
	}
	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	client, err := core.NewClient(ctx, option.WithWechatPayAutoAuthCipher(payThis.Mchid, payThis.SerialNo, privateKey, payThis.APIv3Key))
	if err != nil {
		return
	}

	if payReq.ClientIp == `` {
		payReq.ClientIp = `127.0.0.1`
	}
	/* if payReq.Device == `` {
		payReq.Device = DeviceUnknown
	} */
	// 发送请求
	svc := h5.H5ApiService{Client: client}
	resp, result, err := svc.Prepay(ctx,
		h5.PrepayRequest{
			Appid:       core.String(payThis.AppId),
			Mchid:       core.String(payThis.Mchid),
			Description: core.String(payReq.Desc),
			OutTradeNo:  core.String(payReq.OrderNo),
			NotifyUrl:   core.String(payThis.NotifyUrl),
			Amount: &h5.Amount{
				// Total: core.Int64(gconv.Int64(math.Ceil(payReq.Amount * 100))),
				Total: core.Int64(gconv.Int64(payReq.Amount * 100)),
			},
			SceneInfo: &h5.SceneInfo{
				PayerClientIp: core.String(payReq.ClientIp),
				H5Info: &h5.H5Info{
					// Type: core.String(string(payReq.Device)),
					Type: core.String(`H5`),
				},
			},
		},
	)
	if err != nil {
		return
	}
	if result.Response.StatusCode != 200 {
		err = errors.New(`响应错误`)
		return
	}
	payRes.PayStr = *resp.H5Url
	if payReq.ReturnUrl != `` {
		payRes.PayStr = payRes.PayStr + `&redirect_url=` + url.QueryEscape(payReq.ReturnUrl)
	}
	return
}

func (payThis *Pay) QRCode(ctx context.Context, payReq model.PayReq) (payRes model.PayRes, err error) {
	privateKey, err := utils.LoadPrivateKey(payThis.PrivateKey)
	if err != nil {
		return
	}
	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	client, err := core.NewClient(ctx, option.WithWechatPayAutoAuthCipher(payThis.Mchid, payThis.SerialNo, privateKey, payThis.APIv3Key))
	if err != nil {
		return
	}

	// 发送请求
	svc := native.NativeApiService{Client: client}
	resp, result, err := svc.Prepay(ctx,
		native.PrepayRequest{
			Appid:       core.String(payThis.AppId),
			Mchid:       core.String(payThis.Mchid),
			Description: core.String(payReq.Desc),
			OutTradeNo:  core.String(payReq.OrderNo),
			NotifyUrl:   core.String(payThis.NotifyUrl),
			Amount: &native.Amount{
				// Total: core.Int64(gconv.Int64(math.Ceil(payReq.Amount * 100))),
				Total: core.Int64(gconv.Int64(payReq.Amount * 100)),
			},
		},
	)
	if err != nil {
		return
	}
	if result.Response.StatusCode != 200 {
		err = errors.New(`响应错误`)
		return
	}
	payRes.PayStr = *resp.CodeUrl
	return
}

func (payThis *Pay) Jsapi(ctx context.Context, payReq model.PayReq) (payRes model.PayRes, err error) {
	privateKey, err := utils.LoadPrivateKey(payThis.PrivateKey)
	if err != nil {
		return
	}
	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	client, err := core.NewClient(ctx, option.WithWechatPayAutoAuthCipher(payThis.Mchid, payThis.SerialNo, privateKey, payThis.APIv3Key))
	if err != nil {
		return
	}

	// 发送请求
	svc := jsapi.JsapiApiService{Client: client}
	resp, result, err := svc.Prepay(ctx,
		jsapi.PrepayRequest{
			Appid:       core.String(payThis.AppId),
			Mchid:       core.String(payThis.Mchid),
			Description: core.String(payReq.Desc),
			OutTradeNo:  core.String(payReq.OrderNo),
			NotifyUrl:   core.String(payThis.NotifyUrl),
			Amount: &jsapi.Amount{
				// Total: core.Int64(gconv.Int64(math.Ceil(payReq.Amount * 100))),
				Total: core.Int64(gconv.Int64(payReq.Amount * 100)),
			},
			Payer: &jsapi.Payer{
				Openid: core.String(payReq.Openid),
			},
		},
	)
	if err != nil {
		return
	}
	if result.Response.StatusCode != 200 {
		err = errors.New(`响应错误`)
		return
	}
	payRes.PayStr = *resp.PrepayId
	return
}

func (payThis *Pay) Notify(ctx context.Context, r *ghttp.Request) (notifyInfo model.NotifyInfo, err error) {
	privateKey, err := utils.LoadPrivateKey(payThis.PrivateKey)
	if err != nil {
		return
	}

	// 1. 使用 `RegisterDownloaderWithPrivateKey` 注册下载器
	err = downloader.MgrInstance().RegisterDownloaderWithPrivateKey(ctx, privateKey, payThis.SerialNo, payThis.Mchid, payThis.APIv3Key)
	if err != nil {
		return
	}
	// 2. 获取商户号对应的微信支付平台证书访问器
	certificateVisitor := downloader.MgrInstance().GetCertificateVisitor(payThis.Mchid)
	// 3. 使用证书访问器初始化 `notify.Handler`
	handler := notify.NewNotifyHandler(payThis.APIv3Key, verifiers.NewSHA256WithRSAVerifier(certificateVisitor))

	transaction := new(payments.Transaction)
	_ /* notifyReq */, err = handler.ParseNotifyRequest(ctx, r.Request, transaction)
	if err != nil {
		return
	}
	/* if notifyReq.EventType != `TRANSACTION.SUCCESS` {
		return
	} */

	notifyInfo.Amount = gconv.Float64(transaction.Amount.Total) / 100
	notifyInfo.OrderNo = *transaction.OutTradeNo
	notifyInfo.ThirdOrderNo = *transaction.TransactionId
	return
}

func (payThis *Pay) NotifyRes(ctx context.Context, r *ghttp.Request, failMsg string) {
	resData := map[string]string{
		`code`:    `SUCCESS`, //错误码，SUCCESS为清算机构接收成功，其他错误码为失败。
		`message`: ``,        //返回信息，如非空，为错误原因。
	}
	if failMsg != `` {
		resData = map[string]string{
			`code`:    `FAIL`,
			`message`: failMsg,
		}
	}
	r.Response.WriteJson(resData)
}
