package pay

import (
	daoPay "api/internal/dao/pay"
	"api/internal/utils"
	"api/internal/utils/pay/model"
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/net/ghttp"
)

type Handler struct {
	Ctx context.Context
	pay model.Pay
}

func NewHandler(ctx context.Context, payInfo gdb.Record) model.Handler {
	handlerObj := &Handler{Ctx: ctx}
	config := payInfo[daoPay.Pay.Columns().PayConfig].Map()
	config[`notify_url`] = utils.GetRequestUrl(ctx, 0) + `/pay/notify/` + payInfo[daoPay.Pay.Columns().PayId].String()
	handlerObj.pay = NewPay(ctx, payInfo[daoPay.Pay.Columns().PayType].Uint(), config)
	return handlerObj
}

func (handlerThis *Handler) App(payReq model.PayReq) (payRes model.PayRes, err error) {
	return handlerThis.pay.App(handlerThis.Ctx, payReq)
}

func (handlerThis *Handler) H5(payReq model.PayReq) (payRes model.PayRes, err error) {
	return handlerThis.pay.H5(handlerThis.Ctx, payReq)
}

func (handlerThis *Handler) QRCode(payReq model.PayReq) (payRes model.PayRes, err error) {
	return handlerThis.pay.QRCode(handlerThis.Ctx, payReq)
}

func (handlerThis *Handler) Jsapi(payReq model.PayReq) (payRes model.PayRes, err error) {
	return handlerThis.pay.Jsapi(handlerThis.Ctx, payReq)
}

func (handlerThis *Handler) Notify(r *ghttp.Request) (notifyInfo model.NotifyInfo, err error) {
	return handlerThis.pay.Notify(handlerThis.Ctx, r)
}

func (handlerThis *Handler) NotifyRes(r *ghttp.Request, failMsg string) {
	handlerThis.pay.NotifyRes(handlerThis.Ctx, r, failMsg)
}
