package controller

import (
	"api/api"
	daoAuth "api/internal/dao/auth"
	daoPay "api/internal/dao/pay"
	daoUsers "api/internal/dao/users"
	"api/internal/utils"
	"api/internal/utils/jbctx"
	"api/internal/utils/pay"
	payModel "api/internal/utils/pay/model"
	"context"
	"strconv"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
)

type Pay struct{}

func NewPay() *Pay {
	return &Pay{}
}

// 列表
func (controllerThis *Pay) List(ctx context.Context, req *api.PayChannelListReq) (res *api.PayChannelListRes, err error) {
	paySceneInfo, _ := daoPay.Scene.CacheGetInfo(ctx, req.SceneId)
	if paySceneInfo.IsEmpty() {
		err = utils.NewErrorCode(ctx, 30011000, ``)
		return
	}
	if paySceneInfo[daoPay.Scene.Columns().IsStop].Uint8() == 0 {
		err = utils.NewErrorCode(ctx, 30011001, ``)
		return
	}

	/* sceneInfo := jbctx.GetSceneInfo(ctx)
	sceneId := sceneInfo[daoAuth.Scene.Columns().SceneId].String()
	switch sceneId {
	case `app`:
	default:
		err = utils.NewErrorCode(ctx, 39999998, ``)
		return
	} */

	list, err := daoPay.Channel.CacheGetList(ctx, req.SceneId)
	if err != nil {
		return
	}

	res = &api.PayChannelListRes{List: []api.PayChannelInfo{}}
	gconv.Structs(list.List(), &res.List)
	return
}

// 新增
func (controllerThis *Pay) Create(ctx context.Context, req *api.PayCreateReq) (res *api.PayCreateRes, err error) {
	/**--------参数处理 开始--------**/
	data := g.Map{}
	data[daoPay.Order.Columns().OrderNo] = strconv.FormatInt(gtime.Now().UnixNano(), 36) + grand.S(4)
	data[daoPay.Order.Columns().OrderType] = req.OrderType
	data[daoPay.Order.Columns().OrderIp] = g.RequestFromCtx(ctx).GetClientIp()
	switch jbctx.GetSceneInfo(ctx)[daoAuth.Scene.Columns().SceneId].String() {
	case `app`:
		loginInfo := jbctx.GetLoginInfo(ctx)
		if loginInfo.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39994000, ``)
			return
		}
		data[daoPay.Order.Columns().RelId] = loginInfo[`login_id`]
		switch *req.OrderType {
		case 0:
			data[daoPay.Order.Columns().Amount] = req.Param0.Amount
			// data[daoPay.Order.Columns().ExtData] = g.Map{}
		default:
			err = utils.NewErrorCode(ctx, 30013001, ``)
			return
		}
	default:
		err = utils.NewErrorCode(ctx, 39999998, ``)
		return
	}
	/**--------参数处理 结束--------**/

	id, err := daoPay.Order.CtxDaoModel(ctx).Data(data).InsertAndGetId()
	if err != nil {
		return
	}
	data[daoPay.Order.Columns().OrderId] = id

	res = &api.PayCreateRes{}
	gconv.Struct(data, &res.Info)
	return
}

// 支付
func (controllerThis *Pay) Pay(ctx context.Context, req *api.PayPayReq) (res *api.PayPayRes, err error) {
	channelInfo, _ := daoPay.Channel.CacheGetInfo(ctx, req.ChannelId)
	if channelInfo.IsEmpty() {
		err = utils.NewErrorCode(ctx, 30012000, ``)
		return
	}
	if channelInfo[daoPay.Channel.Columns().IsStop].Uint8() == 1 {
		err = utils.NewErrorCode(ctx, 30012001, ``)
		return
	}
	payInfo, _ := daoPay.Pay.CacheGetInfo(ctx, channelInfo[daoPay.Channel.Columns().PayId].Uint())
	if payInfo.IsEmpty() {
		err = utils.NewErrorCode(ctx, 30010000, ``)
		return
	}

	/**--------订单验证和设置支付数据 开始--------**/
	orderFilter := g.Map{}
	orderFilter[daoPay.Order.Columns().PayStatus] = 0
	if req.OrderId > 0 {
		orderFilter[daoPay.Order.Columns().OrderId] = req.OrderId
	} else /* if req.OrderNo != `` */ {
		orderFilter[daoPay.Order.Columns().OrderNo] = req.OrderNo
	}

	var payReq payModel.PayReq
	switch jbctx.GetSceneInfo(ctx)[daoAuth.Scene.Columns().SceneId].String() {
	case `app`:
		loginInfo := jbctx.GetLoginInfo(ctx)
		if loginInfo.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39994000, ``)
			return
		}
		/* if !slices.Contains([]uint8{0}, orderInfo[daoPay.Order.Columns().OrderType].Uint8()) {
			err = utils.NewErrorCode(ctx, 30013001, ``)
			return
		}
		if orderInfo[daoPay.Order.Columns().RelId].Uint() != loginInfo[`login_id`].Uint() {
			err = utils.NewErrorCode(ctx, 30013000, ``)
			return
		} */
		orderFilter[daoPay.Order.Columns().RelId] = loginInfo[`login_id`]
		orderFilter[daoPay.Order.Columns().OrderType] = []uint8{0}
		if channelInfo[daoPay.Channel.Columns().PayMethod].Uint8() == 3 { //小程序支付
			switch payInfo[daoPay.Pay.Columns().PayType].Uint8() {
			case 0: //支付宝
				// payReq.Openid = loginInfo[daoUsers.Users.Columns().AliOpenid].String()
			case 1: //微信
				payReq.Openid = loginInfo[daoUsers.Users.Columns().WxOpenid].String()
			}
		}
	default:
		err = utils.NewErrorCode(ctx, 39999998, ``)
		return
	}

	orderInfo, _ := daoPay.Order.CtxDaoModel(ctx).Filters(orderFilter).One()
	if orderInfo.IsEmpty() {
		err = utils.NewErrorCode(ctx, 30013000, ``)
		return
	}

	payReq.OrderNo = orderInfo[daoPay.Order.Columns().OrderNo].String()
	payReq.Amount = orderInfo[daoPay.Order.Columns().Amount].Float64()
	payReq.Desc = `描述`
	/* switch orderInfo[daoPay.Order.Columns().RelOrderType].Uint8() { // 根据订单类型确认是否设置不同描述
	case 0:
		payReq.Desc = `默认订单`
	} */
	/**--------订单验证和设置支付数据 结束--------**/

	payObj := pay.NewHandler(ctx, payInfo)
	var payRes payModel.PayRes
	switch channelInfo[daoPay.Channel.Columns().PayMethod].Uint8() {
	case 0: //APP支付
		payRes, err = payObj.App(payReq)
	case 1: //H5支付
		payRes, err = payObj.H5(payReq)
	case 2: //扫码支付
		payRes, err = payObj.QRCode(payReq)
	case 3: //小程序支付
		payRes, err = payObj.Jsapi(payReq)
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
		PayStr: payRes.PayStr,
	}
	return
}

// 回调
func (controllerThis *Pay) Notify(ctx context.Context, req *api.PayNotifyReq) (res *api.CommonNoDataRes, err error) {
	payInfo, _ := daoPay.Pay.CacheGetInfo(ctx, req.PayId)
	if payInfo.IsEmpty() {
		err = utils.NewErrorCode(ctx, 30010000, ``)
		return
	}
	payObj := pay.NewHandler(ctx, payInfo)

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
	payOrderDaoModel := daoPay.Order.CtxDaoModel(ctx)
	err = payOrderDaoModel.Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		row, _ := payOrderDaoModel.ResetNew().TX(tx).Filters(g.Map{
			daoPay.Order.Columns().OrderId:   orderInfo[daoPay.Order.Columns().OrderId],
			daoPay.Order.Columns().PayStatus: 0, //防并发
		}).Data(g.Map{
			daoPay.Order.Columns().ThirdOrderNo: notifyInfo.ThirdOrderNo,
			daoPay.Order.Columns().PayStatus:    1,
			daoPay.Order.Columns().PayAt:        gtime.Now(),
			// daoPay.Order.Columns().PayRate:      payInfo[daoPay.Pay.Columns().PayRate], //以订单回调时的费率为准
		}).UpdateAndGetAffected()
		if row == 0 {
			err = utils.NewErrorCode(ctx, 30019000, ``)
			return
		}

		/**--------处理关联订单 开始--------**/
		switch orderInfo[daoPay.Order.Columns().OrderType].Uint8() { // 根据订单类型处理
		// case 0: //默认
		default:
			err = utils.NewErrorCode(ctx, 30013001, ``)
			return
		}
		/**--------处理关联订单 结束--------**/

		// 累积支付数据
		daoPay.Pay.CtxDaoModel(ctx).SetIdArr(orderInfo[daoPay.Order.Columns().PayId]).HookUpdate(g.Map{
			daoPay.Pay.Columns().TotalAmount: gdb.Raw(daoPay.Pay.Columns().TotalAmount + ` + ` + orderInfo[daoPay.Order.Columns().Amount].String()),
			daoPay.Pay.Columns().Balance:     gdb.Raw(daoPay.Pay.Columns().Balance + ` + ` + gconv.String(orderInfo[daoPay.Order.Columns().Amount].Float64()*(1-orderInfo[daoPay.Order.Columns().PayRate].Float64()))), //以订单选择支付通道时的费率为准
			// daoPay.Pay.Columns().Balance:     gdb.Raw(daoPay.Pay.Columns().Balance + ` + ` + gconv.String(orderInfo[daoPay.Order.Columns().Amount].Float64()*(1-payInfo[daoPay.Pay.Columns().PayRate].Float64()))), //以订单回调时的费率为准
		}).Update()
		daoPay.Channel.CtxDaoModel(ctx).SetIdArr(orderInfo[daoPay.Order.Columns().ChannelId]).HookUpdateOne(daoPay.Channel.Columns().TotalAmount, gdb.Raw(daoPay.Channel.Columns().TotalAmount+` + `+orderInfo[daoPay.Order.Columns().Amount].String())).Update()
		return
	})
	if err != nil {
		payObj.NotifyRes(r, err.Error())
		return
	}

	payObj.NotifyRes(r, ``)
	return
}
