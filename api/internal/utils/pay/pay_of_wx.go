package pay

import (
	daoPlatform "api/internal/dao/platform"
	"context"
	"errors"

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
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
)

type PayOfWx struct {
	Ctx        context.Context
	AppId      string `json:"payOfWxAppId"`
	Mchid      string `json:"payOfWxMchid"`
	SerialNo   string `json:"payOfWxSerialNo"`
	APIv3Key   string `json:"payOfWxApiV3Key"`
	PrivateKey string `json:"payOfWxPrivateKey"`
	NotifyUrl  string `json:"payOfWxNotifyUrl"`
}

func NewPayOfWx(ctx context.Context, configOpt ...map[string]any) *PayOfWx {
	var config map[string]any
	if len(configOpt) > 0 && len(configOpt[0]) > 0 {
		config = configOpt[0]
	} else {
		configTmp, _ := daoPlatform.Config.Get(ctx, []string{`payOfWxAppId`, `payOfWxMchid`, `payOfWxSerialNo`, `payOfWxApiV3Key`, `payOfWxPrivateKey`, `payOfWxNotifyUrl`})
		config = configTmp.Map()
	}

	payOfWxObj := PayOfWx{Ctx: ctx}
	gconv.Struct(config, &payOfWxObj)
	return &payOfWxObj
}

func (payThis *PayOfWx) App(payData PayData) (orderInfo PayInfo, err error) {
	privateKey, err := utils.LoadPrivateKey(payThis.PrivateKey)
	if err != nil {
		return
	}
	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(payThis.Mchid, payThis.SerialNo, privateKey, payThis.APIv3Key),
	}
	client, err := core.NewClient(payThis.Ctx, opts...)
	if err != nil {
		return
	}

	// 发送请求
	svc := app.AppApiService{Client: client}
	resp, result, err := svc.Prepay(payThis.Ctx,
		app.PrepayRequest{
			Appid:       core.String(payThis.AppId),
			Mchid:       core.String(payThis.Mchid),
			Description: core.String(payData.Desc),
			OutTradeNo:  core.String(payData.OrderNo),
			NotifyUrl:   core.String(payThis.NotifyUrl),
			Amount: &app.Amount{
				// Total: core.Int64(gconv.Int64(math.Ceil(payData.Amount * 100))),
				Total: core.Int64(gconv.Int64(payData.Amount * 100)),
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
	orderInfo.PayStr = *resp.PrepayId
	return
}

func (payThis *PayOfWx) H5(payData PayData) (orderInfo PayInfo, err error) {
	privateKey, err := utils.LoadPrivateKey(payThis.PrivateKey)
	if err != nil {
		return
	}
	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(payThis.Mchid, payThis.SerialNo, privateKey, payThis.APIv3Key),
	}
	client, err := core.NewClient(payThis.Ctx, opts...)
	if err != nil {
		return
	}

	// 发送请求
	svc := h5.H5ApiService{Client: client}
	if payData.ClientIp == `` {
		payData.ClientIp = `127.0.0.1`
	}
	/* if payData.Device == `` {
		payData.Device = DeviceUnknown
	} */
	resp, result, err := svc.Prepay(payThis.Ctx,
		h5.PrepayRequest{
			Appid:       core.String(payThis.AppId),
			Mchid:       core.String(payThis.Mchid),
			Description: core.String(payData.Desc),
			OutTradeNo:  core.String(payData.OrderNo),
			NotifyUrl:   core.String(payThis.NotifyUrl),
			Amount: &h5.Amount{
				// Total: core.Int64(gconv.Int64(math.Ceil(payData.Amount * 100))),
				Total: core.Int64(gconv.Int64(payData.Amount * 100)),
			},
			SceneInfo: &h5.SceneInfo{
				PayerClientIp: core.String(payData.ClientIp),
				H5Info: &h5.H5Info{
					// Type: core.String(string(payData.Device)),
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
	orderInfo.PayStr = *resp.H5Url
	return
}

func (payThis *PayOfWx) Jsapi(payData PayData) (orderInfo PayInfo, err error) {
	privateKey, err := utils.LoadPrivateKey(payThis.PrivateKey)
	if err != nil {
		return
	}
	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(payThis.Mchid, payThis.SerialNo, privateKey, payThis.APIv3Key),
	}
	client, err := core.NewClient(payThis.Ctx, opts...)
	if err != nil {
		return
	}

	// 发送请求
	svc := jsapi.JsapiApiService{Client: client}
	resp, result, err := svc.Prepay(payThis.Ctx,
		jsapi.PrepayRequest{
			Appid:       core.String(payThis.AppId),
			Mchid:       core.String(payThis.Mchid),
			Description: core.String(payData.Desc),
			OutTradeNo:  core.String(payData.OrderNo),
			NotifyUrl:   core.String(payThis.NotifyUrl),
			Amount: &jsapi.Amount{
				// Total: core.Int64(gconv.Int64(math.Ceil(payData.Amount * 100))),
				Total: core.Int64(gconv.Int64(payData.Amount * 100)),
			},
			Payer: &jsapi.Payer{
				Openid: core.String(payData.Openid),
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
	orderInfo.PayStr = *resp.PrepayId
	return
}

func (payThis *PayOfWx) Notify(r *ghttp.Request) (notifyInfo NotifyInfo, err error) {
	privateKey, err := utils.LoadPrivateKey(payThis.PrivateKey)
	if err != nil {
		return
	}

	// 1. 使用 `RegisterDownloaderWithPrivateKey` 注册下载器
	err = downloader.MgrInstance().RegisterDownloaderWithPrivateKey(payThis.Ctx, privateKey, payThis.SerialNo, payThis.Mchid, payThis.APIv3Key)
	if err != nil {
		return
	}
	// 2. 获取商户号对应的微信支付平台证书访问器
	certificateVisitor := downloader.MgrInstance().GetCertificateVisitor(payThis.Mchid)
	// 3. 使用证书访问器初始化 `notify.Handler`
	handler := notify.NewNotifyHandler(payThis.APIv3Key, verifiers.NewSHA256WithRSAVerifier(certificateVisitor))

	transaction := new(payments.Transaction)
	_ /* notifyReq */, err = handler.ParseNotifyRequest(payThis.Ctx, r.Request, transaction)
	if err != nil {
		return
	}
	/* if notifyReq.EventType != `TRANSACTION.SUCCESS` {
		return
	} */

	notifyInfo.Amount = gconv.Float64(transaction.Amount.Total) / 100
	notifyInfo.OrderNo = *transaction.OutTradeNo
	notifyInfo.OrderNoOfThird = *transaction.TransactionId
	return
}

func (payThis *PayOfWx) NotifyRes(r *ghttp.Request, failMsg string) {
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
