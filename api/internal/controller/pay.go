package controller

import (
	"api/api"
	"api/internal/utils"
	"api/internal/utils/pay"
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gutil"
)

type Pay struct{}

func NewPay() *Pay {
	return &Pay{}
}

// 支付
func (controllerThis *Pay) Pay(ctx context.Context, req *api.PayPayReq) (res *api.PayPayRes, err error) {
	loginInfo := utils.GetCtxLoginInfo(ctx)
	if loginInfo.IsEmpty() {
		err = utils.NewErrorCode(ctx, 39994000, ``)
		return
	}
	var orderInfo gdb.Record
	//订单查询
	/* orderInfo, _ = dao.NewDaoHandler(ctx, &daoXxxx.Order).Filter(g.Map{
		daoXxxx.Order.Columns().OrderNo:   req.OrderNo,
		daoXxxx.Order.Columns().UserId:    loginInfo[`loginId`],
		daoXxxx.Order.Columns().PayStatus: 0,
	}).GetModel().One()
	if orderInfo.IsEmpty() {
		err = utils.NewErrorCode(ctx, 99999999, `订单不存在`)
		return
	} */

	payData := pay.PayData{
		OrderNo: orderInfo[`orderNo`].String(),
		Amount:  orderInfo[`amount`].Float64(),
		Desc:    `订单描述`,
	}
	var payInfo pay.PayInfo
	switch req.PayMethod {
	case 0: //APP支付(支付宝)
		payInfo, err = pay.NewPay(ctx, `payOfAli`).App(payData)
	case 1: //H5支付(支付宝)
		payInfo, err = pay.NewPay(ctx, `payOfAli`).H5(payData)
	case 2: //JSAPI支付(支付宝)
		/* payData.OpenId = ``
		payInfo, err = pay.NewPay(ctx, `payOfAli`).Jsapi(payData) */
		err = utils.NewErrorCode(ctx, 99999999, `暂不支持，无法获取用户openId`)
	case 10: //APP支付(微信)
		payInfo, err = pay.NewPay(ctx, `payOfWx`).App(payData)
	case 11: //H5支付(微信)
		payData.ClientIp = g.RequestFromCtx(ctx).GetClientIp()
		payInfo, err = pay.NewPay(ctx, `payOfWx`).H5(payData)
	case 12: //JSAPI支付(微信)
		/* payData.OpenId = ``
		payInfo, err = pay.NewPay(ctx, `payOfWx`).Jsapi(payData) */
		err = utils.NewErrorCode(ctx, 99999999, `暂不支持，无法获取用户openId`)
	}
	if err != nil {
		return
	}

	res = &api.PayPayRes{
		PayStr: payInfo.PayStr,
	}
	return
}

// 回调
func (controllerThis *Pay) Notify(ctx context.Context, req *api.PayNotifyReq) (res *api.CommonNoDataRes, err error) {
	payObj := pay.NewPay(ctx, req.PayType)
	notifyInfo, err := payObj.Notify()
	if err != nil {
		payObj.NotifyRes(err.Error())
		return
	}
	//订单回调处理
	gutil.Dump(notifyInfo)
	/* err = daoXxxx.Order.ParseDbCtx(ctx).Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		row, err := tx.Model(daoXxxx.Order.ParseDbTable(ctx)).Where(g.Map{
			daoXxxx.Order.Columns().OrderNo:   notifyInfo.OrderNo,
			daoXxxx.Order.Columns().Price:     notifyInfo.Amount,
			daoXxxx.Order.Columns().PayStatus: 0,
		}).Data(g.Map{
			daoXxxx.Order.Columns().OrderNoOfThird: notifyInfo.OrderNoOfThird,
			daoXxxx.Order.Columns().PayStatus:      1,
			daoXxxx.Order.Columns().PayTime:        gtime.Now(),
		}).UpdateAndGetAffected()
		if err != nil {
			return
		}
		if row == 0 {
			err = utils.NewErrorCode(ctx, 99999999, `请勿重复通知`)
			return
		}

		// 支付成功后处理逻辑
		return
	})
	if err != nil {
		payObj.NotifyRes(err.Error())
		return
	} */

	payObj.NotifyRes(``)
	return
}
