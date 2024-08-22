package controller

import (
	"api/api"
	daoAuth "api/internal/dao/auth"
	daoPay "api/internal/dao/pay"
	daoUsers "api/internal/dao/users"
	"api/internal/utils"
	"api/internal/utils/pay"
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
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

	/**--------订单验证和设置支付数据 开始--------**/
	orderFilter := g.Map{}
	orderFilter[daoPay.Order.Columns().PayStatus] = 0
	if req.OrderId > 0 {
		orderFilter[daoPay.Order.Columns().OrderId] = req.OrderId
	}
	if req.OrderNo != `` {
		orderFilter[daoPay.Order.Columns().OrderNo] = req.OrderNo
	}

	var payReqData pay.PayReqData
	switch utils.GetCtxSceneInfo(ctx)[daoAuth.Scene.Columns().SceneCode].String() {
	case `app`:
		loginInfo := utils.GetCtxLoginInfo(ctx)
		if loginInfo.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39994000, ``)
			return
		}
		if channelInfo[daoPay.Channel.Columns().PayMethod].Uint() == 3 { //小程序支付
			switch payInfo[daoPay.Pay.Columns().PayType].Uint() {
			case 0: //支付宝
				// payReqData.Openid = loginInfo[daoUsers.Users.Columns().AliOpenid].String()
			case 1: //微信
				payReqData.Openid = loginInfo[daoUsers.Users.Columns().WxOpenid].String()
			}
		}
		// orderFilter[daoPay.Order.Columns().UserId] = loginInfo[`login_id`]
	default:
		err = utils.NewErrorCode(ctx, 39999998, ``)
		return
	}

	orderInfo, _ := daoPay.Order.CtxDaoModel(ctx).Filters(orderFilter).One()
	if orderInfo.IsEmpty() {
		err = utils.NewErrorCode(ctx, 30013000, ``)
		return
	}

	payReqData.OrderNo = orderInfo[daoPay.Order.Columns().OrderNo].String()
	payReqData.Amount = orderInfo[daoPay.Order.Columns().Amount].Float64()
	payReqData.Desc = `描述`
	/**--------订单验证和设置支付数据 结束--------**/

	payObj := pay.NewPay(ctx, payInfo)
	var payResData pay.PayResData
	switch channelInfo[daoPay.Channel.Columns().PayMethod].Uint() {
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

	daoPay.Order.CtxDaoModel(ctx).Filters(g.Map{
		daoPay.Order.Columns().OrderId:   orderInfo[daoPay.Order.Columns().OrderId],
		daoPay.Order.Columns().PayStatus: 0,
	}).Data(g.Map{
		daoPay.Order.Columns().PayId:     channelInfo[daoPay.Channel.Columns().PayId],
		daoPay.Order.Columns().ChannelId: channelInfo[daoPay.Channel.Columns().ChannelId],
		daoPay.Order.Columns().PayType:   payInfo[daoPay.Pay.Columns().PayType],
		daoPay.Order.Columns().PayRate:   payInfo[daoPay.Pay.Columns().PayRate], //以订单选择支付通道时的费率为准
	}).Update()

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

	orderInfo, _ := daoPay.Order.CtxDaoModel(ctx).Filters(g.Map{
		daoPay.Order.Columns().PayId:     req.PayId,
		daoPay.Order.Columns().OrderNo:   notifyInfo.OrderNo,
		daoPay.Order.Columns().Amount:    notifyInfo.Amount,
		daoPay.Order.Columns().PayStatus: 0,
	}).One()
	if orderInfo.IsEmpty() {
		err = utils.NewErrorCode(ctx, 30013000, ``)
		payObj.NotifyRes(r, err.Error())
		return
	}

	// 订单回调处理
	payOrderHandler := daoPay.Order.CtxDaoModel(ctx)
	err = payOrderHandler.Transaction(func(ctx context.Context, tx gdb.TX) (err error) {
		row, err := payOrderHandler.CloneNew().TX(tx).Filters(g.Map{
			daoPay.Order.Columns().OrderId:   orderInfo[daoPay.Order.Columns().OrderId],
			daoPay.Order.Columns().PayStatus: 0, //防并发
		}).Data(g.Map{
			daoPay.Order.Columns().ThirdOrderNo: notifyInfo.OrderNoOfThird,
			daoPay.Order.Columns().PayStatus:    1,
			daoPay.Order.Columns().PayTime:      gtime.Now(),
			// daoPay.Order.Columns().PayRate:      payInfo[daoPay.Pay.Columns().PayRate], //以订单回调时的费率为准
		}).UpdateAndGetAffected()
		if err != nil {
			return
		}
		if row == 0 {
			err = utils.NewErrorCode(ctx, 30019000, ``)
			return
		}

		// 支付成功后处理关联订单
		orderRelList, _ := daoPay.OrderRel.CtxDaoModel(ctx).TX(tx).Filter(daoPay.OrderRel.Columns().OrderId, orderInfo[daoPay.Order.Columns().OrderId]).All()
		for _, v := range orderRelList {
			switch v[daoPay.OrderRel.Columns().RelOrderType].Uint() { // 根据订单类型找到对应的订单表做处理
			// case 0:
			default:
				err = utils.NewErrorCode(ctx, 30013001, ``)
				return
			}
		}

		// 累积支付数据
		daoPay.Pay.CtxDaoModel(ctx).Filter(daoPay.Pay.Columns().PayId, orderInfo[daoPay.Order.Columns().PayId]).HookUpdate(g.Map{
			daoPay.Pay.Columns().TotalAmount: gdb.Raw(daoPay.Pay.Columns().TotalAmount + ` + ` + orderInfo[daoPay.Order.Columns().Amount].String()),
			daoPay.Pay.Columns().Balance:     gdb.Raw(daoPay.Pay.Columns().Balance + ` + ` + gconv.String(orderInfo[daoPay.Order.Columns().Amount].Float64()*(1-orderInfo[daoPay.Order.Columns().PayRate].Float64()))), //以订单选择支付通道时的费率为准
			// daoPay.Pay.Columns().Balance:     gdb.Raw(daoPay.Pay.Columns().Balance + ` + ` + gconv.String(orderInfo[daoPay.Order.Columns().Amount].Float64()*(1-payInfo[daoPay.Pay.Columns().PayRate].Float64()))), //以订单回调时的费率为准
		}).Update()
		daoPay.Channel.CtxDaoModel(ctx).Filter(daoPay.Channel.Columns().ChannelId, orderInfo[daoPay.Order.Columns().ChannelId]).HookUpdate(g.Map{
			daoPay.Channel.Columns().TotalAmount: gdb.Raw(daoPay.Channel.Columns().TotalAmount + ` + ` + orderInfo[daoPay.Order.Columns().Amount].String()),
		}).Update()
		return
	})
	if err != nil {
		payObj.NotifyRes(r, err.Error())
		return
	}

	payObj.NotifyRes(r, ``)
	return
}
