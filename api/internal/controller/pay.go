package controller

import (
	"api/api"
	daoAuth "api/internal/dao/auth"
	daoPay "api/internal/dao/pay"
	daoUsers "api/internal/dao/users"
	"api/internal/utils"
	"api/internal/utils/pay"
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gutil"
)

type Pay struct{}

func NewPay() *Pay {
	return &Pay{}
}

// 列表
func (controllerThis *Pay) List(ctx context.Context, req *api.PayListReq) (res *api.PayListRes, err error) {
	/* sceneInfo := utils.GetCtxSceneInfo(ctx)
	sceneCode := sceneInfo[daoAuth.Scene.Columns().SceneCode].String()
	switch sceneCode {
	case `app`:
	default:
		err = utils.NewErrorCode(ctx, 39999998, ``)
		return
	} */

	list, err := daoPay.Pay.CtxDaoModel(ctx).Filter(daoPay.Scene.Columns().PayScene, req.PayScene).OrderDesc(daoPay.Pay.Columns().Sort).ListPri()
	if err != nil {
		return
	}

	res = &api.PayListRes{List: []api.PayInfo{}}
	list.Structs(&res.List)
	return
}

// 支付
func (controllerThis *Pay) Pay(ctx context.Context, req *api.PayPayReq) (res *api.PayPayRes, err error) {
	/**--------确定支付数据 开始--------**/
	var payReqData pay.PayReqData
	sceneInfo := utils.GetCtxSceneInfo(ctx)
	sceneCode := sceneInfo[daoAuth.Scene.Columns().SceneCode].String()
	switch sceneCode {
	case `app`:
		loginInfo := utils.GetCtxLoginInfo(ctx)
		if loginInfo.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39994000, ``)
			return
		}
		switch req.PayScene {
		case 10, 11: //微信小程序	微信公众号
			payReqData.Openid = loginInfo[daoUsers.Users.Columns().WxOpenid].String()
		case 20: //支付宝小程序
			// payReqData.Openid = loginInfo[daoUsers.Users.Columns().AliOpenid].String()
		}
	default:
		err = utils.NewErrorCode(ctx, 39999998, ``)
		return
	}

	/* //订单查询
	orderInfo, _ := daoXxxx.Order.CtxDaoModel(ctx).Filters(g.Map{
		daoXxxx.Order.Columns().OrderNo:   req.OrderNo,
		daoXxxx.Order.Columns().UserId:    loginInfo[`login_id`],
		daoXxxx.Order.Columns().PayStatus: 0,
	}).One()
	if orderInfo.IsEmpty() {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}

	payReqData.OrderNo = orderInfo[daoXxxx.Order.Columns().OrderNo].String()
	payReqData.Amount = orderInfo[daoXxxx.Order.Columns().Price].Float64()
	payReqData.Desc = `订单描述` */
	/**--------确定支付数据 开始--------**/

	payInfo, _ := daoPay.Pay.CtxDaoModel(ctx).Filter(daoPay.Pay.Columns().PayId, req.PayId).One()
	if payInfo.IsEmpty() {
		err = utils.NewErrorCode(ctx, 30010000, ``)
		return
	}
	payObj := pay.NewPay(ctx, payInfo)

	var payResData pay.PayResData
	switch req.PayScene {
	case 0: //APP
		payResData, err = payObj.App(payReqData)
	case 1: //H5
		payResData, err = payObj.H5(payReqData)
	case 2: //扫码
		payResData, err = payObj.QRCode(payReqData)
	case 10, 11, 20: //微信小程序	微信公众号	支付宝小程序
		payResData, err = payObj.Jsapi(payReqData)
	}
	if err != nil {
		return
	}

	res = &api.PayPayRes{
		PayStr: payResData.PayStr,
	}
	return
}

// 回调
func (controllerThis *Pay) Notify(ctx context.Context, req *api.PayNotifyReq) (res *api.CommonNoDataRes, err error) {
	payInfo, _ := daoPay.Pay.CtxDaoModel(ctx).Filter(daoPay.Pay.Columns().PayId, req.PayId).One()
	if payInfo.IsEmpty() {
		err = utils.NewErrorCode(ctx, 30010000, ``)
		return
	}
	payObj := pay.NewPay(ctx, payInfo)

	r := g.RequestFromCtx(ctx)
	notifyInfo, err := payObj.Notify(r)
	if err != nil {
		payObj.NotifyRes(r, err.Error())
		return
	}

	// 订单回调处理
	gutil.Dump(notifyInfo)
	/* payAddAmountFunc := func(ctx context.Context, orderAmount float64, payRate float64) { //payRate以订单支付时的费率为准
		daoPay.Pay.CtxDaoModel(ctx).Filter(daoPay.Pay.Columns().PayId, req.PayId).HookUpdate(g.Map{
			daoPay.Pay.Columns().TotalAmount: gdb.Raw(daoPay.Pay.Columns().TotalAmount + ` + ` + gconv.String(orderAmount)),
			daoPay.Pay.Columns().Balance:     gdb.Raw(daoPay.Pay.Columns().Balance + ` + ` + gconv.String(orderAmount*(1-payRate))),
		}).Update()
	}
	xxxxOrderHandler := daoXxxx.Order.CtxDaoModel(ctx)
	err = xxxxOrderHandler.Transaction(func(ctx context.Context, tx gdb.TX) (err error) {
		row, err := tx.Model(xxxxOrderHandler.DbTable).Where(g.Map{
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
			err = utils.NewErrorCode(ctx, 30010001, ``)
			return
		}

		// 支付成功后处理逻辑

		// 支付数据累积
		payRate, _ := tx.Model(xxxxOrderHandler.DbTable).Where(daoXxxx.Order.Columns().OrderNo, notifyInfo.OrderNo).Value(daoXxxx.Order.Columns().PayRate)
		payAddAmountFunc(ctx, notifyInfo.Amount, payRate.Float64())
		return
	})
	if err != nil {
		payObj.NotifyRes(r, err.Error())
		return
	} */

	payObj.NotifyRes(r, ``)
	return
}
