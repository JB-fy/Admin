package pay

import (
	"context"
	"errors"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/verifiers"
	"github.com/wechatpay-apiv3/wechatpay-go/core/downloader"
	"github.com/wechatpay-apiv3/wechatpay-go/core/notify"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/app"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
)

type PayOfWx struct {
	Ctx      context.Context
	AppId    string `json:"payOfWxAppId"`
	Mchid    string `json:"payOfWxMchid"`
	SerialNo string `json:"payOfWxSerialNo"`
	APIv3Key string `json:"payOfWxApiV3Key"`
	CertPath string `json:"payOfWxCertPath"`
}

func NewPayOfWx(ctx context.Context, config map[string]interface{}) *PayOfWx {
	payOfWxObj := PayOfWx{
		Ctx: ctx,
	}
	gconv.Struct(config, &payOfWxObj)
	return &payOfWxObj
}

func (payThis *PayOfWx) App(payData PayData) (orderInfo PayInfo, err error) {
	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	privateKey, err := utils.LoadPrivateKeyWithPath(payThis.CertPath)
	if err != nil {
		// log.Fatal(`load merchant private key error`)
		return
	}
	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(payThis.Mchid, payThis.SerialNo, privateKey, payThis.APIv3Key),
	}
	client, err := core.NewClient(payThis.Ctx, opts...)
	if err != nil {
		// log.Fatalf(`new wechat pay client err:%s`, err)
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
			NotifyUrl:   core.String(payData.NotifyUrl),
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
		err = errors.New(`支付宝APP支付接口错误`)
		return
	}
	orderInfo.PayStr = *resp.PrepayId
	return
}

func (payThis *PayOfWx) Notify() (notifyInfo NotifyInfo, err error) {
	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	privateKey, err := utils.LoadPrivateKeyWithPath(payThis.CertPath)
	if err != nil {
		// log.Fatal(`load merchant private key error`)
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
	request := g.RequestFromCtx(payThis.Ctx).Request
	_ /* notifyReq */, err = handler.ParseNotifyRequest(payThis.Ctx, request, transaction)
	// 如果验签未通过，或者解密失败
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

func (payThis *PayOfWx) NotifyRes(failMsg string) {
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
	g.RequestFromCtx(payThis.Ctx).Response.WriteJson(resData)
}
