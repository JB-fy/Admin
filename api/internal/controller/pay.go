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
func (controllerThis *Pay) List(ctx context.Context, req *api.PayChannelListReq) (res *api.PayChannelListRes, err error) {
	paySceneInfo, _ := daoPay.Scene.CtxDaoModel(ctx).Filter(daoPay.Scene.Columns().SceneId, req.SceneId).One()
	if paySceneInfo.IsEmpty() {
		err = utils.NewErrorCode(ctx, 30011000, ``)
		return
	}
	if paySceneInfo[daoPay.Scene.Columns().IsStop].Uint() == 0 {
		err = utils.NewErrorCode(ctx, 30011001, ``)
		return
	}

	/* sceneInfo := utils.GetCtxSceneInfo(ctx)
	sceneCode := sceneInfo[daoAuth.Scene.Columns().SceneCode].String()
	switch sceneCode {
	case `app`:
	default:
		err = utils.NewErrorCode(ctx, 39999998, ``)
		return
	} */

	list, err := daoPay.Channel.CtxDaoModel(ctx).Filter(daoPay.Channel.Columns().SceneId, req.SceneId).OrderDesc(daoPay.Channel.Columns().Sort).ListPri()
	if err != nil {
		return
	}

	res = &api.PayChannelListRes{List: []api.PayChannelInfo{}}
	list.Structs(&res.List)
	return
}

// 支付
func (controllerThis *Pay) Pay(ctx context.Context, req *api.PayPayReq) (res *api.PayPayRes, err error) {
	channelInfo, _ := daoPay.Channel.CtxDaoModel(ctx).Filter(daoPay.Channel.Columns().ChannelId, req.ChannelId).One()
	if channelInfo.IsEmpty() {
		err = utils.NewErrorCode(ctx, 30012000, ``)
		return
	}
	if channelInfo[daoPay.Channel.Columns().IsStop].Uint() == 0 {
		err = utils.NewErrorCode(ctx, 30012001, ``)
		return
	}
	payInfo, _ := daoPay.Pay.CtxDaoModel(ctx).Filter(daoPay.Pay.Columns().PayId, channelInfo[daoPay.Channel.Columns().PayId]).One()
	if payInfo.IsEmpty() {
		err = utils.NewErrorCode(ctx, 30010000, ``)
		return
	}

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
		if channelInfo[daoPay.Channel.Columns().Method].Uint() == 3 { //小程序支付
			switch payInfo[daoPay.Pay.Columns().PayType].Uint() {
			case 0: //支付宝
				// payReqData.Openid = loginInfo[daoUsers.Users.Columns().AliOpenid].String()
			case 1: //微信
				payReqData.Openid = loginInfo[daoUsers.Users.Columns().WxOpenid].String()
			}
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
	/**--------确定支付数据 结束--------**/

	payObj := pay.NewPay(ctx, payInfo)
	var payResData pay.PayResData
	switch channelInfo[daoPay.Channel.Columns().Method].Uint() {
	case 0: //APP支付
		payResData, err = payObj.App(payReqData)
	case 1: //H5支付
		payResData, err = payObj.H5(payReqData)
	case 2: //扫码支付
		payResData, err = payObj.QRCode(payReqData)
	case 3: //小程序支付
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

	gutil.Dump(notifyInfo)
	/* // 支付数据累积方法。作用：方便不同订单调用
	payAddFunc := func(ctx context.Context, payId uint, channelId uint, orderAmount float64, payRate float64) { //payRate以订单支付时的费率为准
		daoPay.Pay.CtxDaoModel(ctx).Filter(daoPay.Pay.Columns().PayId, payId).HookUpdate(g.Map{
			daoPay.Pay.Columns().TotalAmount: gdb.Raw(daoPay.Pay.Columns().TotalAmount + ` + ` + gconv.String(orderAmount)),
			daoPay.Pay.Columns().Balance:     gdb.Raw(daoPay.Pay.Columns().Balance + ` + ` + gconv.String(orderAmount*(1-payRate))),
		}).Update()
		daoPay.Channel.CtxDaoModel(ctx).Filter(daoPay.Channel.Columns().ChannelId, channelId).HookUpdate(g.Map{
			daoPay.Channel.Columns().TotalAmount: gdb.Raw(daoPay.Channel.Columns().TotalAmount + ` + ` + gconv.String(orderAmount)),
		}).Update()
	}
	// 订单回调处理
	xxxxOrderHandler := daoXxxx.Order.CtxDaoModel(ctx)
	err = xxxxOrderHandler.Transaction(func(ctx context.Context, tx gdb.TX) (err error) {
		row, err := tx.Model(xxxxOrderHandler.DbTable).Where(g.Map{
			daoXxxx.Order.Columns().PayId:     req.PayId,
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
			err = utils.NewErrorCode(ctx, 30019000, ``)
			return
		}

		// 支付成功后处理逻辑

		// 支付数据累积
		orderInfo, _ := tx.Model(xxxxOrderHandler.DbTable).Where(daoXxxx.Order.Columns().OrderNo, notifyInfo.OrderNo).One()
		payAddFunc(ctx, req.PayId, orderInfo[daoXxxx.Order.Columns().ChannelId].Uint(), notifyInfo.Amount, orderInfo[daoXxxx.Order.Columns().PayRate].Float64())
		return
	})
	if err != nil {
		payObj.NotifyRes(r, err.Error())
		return
	} */

	payObj.NotifyRes(r, ``)
	return
}
