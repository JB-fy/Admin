package pay

import (
	"context"
	"errors"

	"github.com/gogf/gf/v2/util/gconv"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
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
