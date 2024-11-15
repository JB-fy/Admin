package pay

import (
	daoPay "api/internal/dao/pay"
	"api/internal/utils"
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/net/ghttp"
)

type Handler struct {
	Ctx context.Context
	Pay Pay
}

func NewHandler(ctx context.Context, payInfo gdb.Record) *Handler {
	handlerObj := &Handler{Ctx: ctx}

	config := payInfo[daoPay.Pay.Columns().PayConfig].Map()
	config[`notifyUrl`] = utils.GetRequestUrl(ctx, 0) + `/pay/notify/` + payInfo[daoPay.Pay.Columns().PayId].String()

	config[`payType`] = payInfo[daoPay.Pay.Columns().PayType]
	handlerObj.Pay = NewPay(config)
	return handlerObj
}

func (handlerThis *Handler) App(payReqData PayReqData) (payResData PayResData, err error) {
	return handlerThis.Pay.App(handlerThis.Ctx, payReqData)
}

func (handlerThis *Handler) H5(payReqData PayReqData) (payResData PayResData, err error) {
	return handlerThis.Pay.H5(handlerThis.Ctx, payReqData)
}

func (handlerThis *Handler) QRCode(payReqData PayReqData) (payResData PayResData, err error) {
	return handlerThis.Pay.QRCode(handlerThis.Ctx, payReqData)
}

func (handlerThis *Handler) Jsapi(payReqData PayReqData) (payResData PayResData, err error) {
	return handlerThis.Pay.Jsapi(handlerThis.Ctx, payReqData)
}

func (handlerThis *Handler) Notify(r *ghttp.Request) (notifyInfo NotifyInfo, err error) {
	return handlerThis.Pay.Notify(handlerThis.Ctx, r)
}

func (handlerThis *Handler) NotifyRes(r *ghttp.Request, failMsg string) {
	handlerThis.Pay.NotifyRes(handlerThis.Ctx, r, failMsg)
}
