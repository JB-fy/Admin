package controller

import (
	"api/api"
	daoAuth "api/internal/dao/auth"
	daoPay "api/internal/dao/pay"
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
	/**--------确定支付场景 开始--------**/
	var payScene uint
	sceneInfo := utils.GetCtxSceneInfo(ctx)
	sceneCode := sceneInfo[daoAuth.Scene.Columns().SceneCode].String()
	switch sceneCode {
	case `app`:
		payScene = 0
	default:
		err = utils.NewErrorCode(ctx, 39999998, ``)
		return
	}
	/**--------确定支付场景 结束--------**/

	list, err := daoPay.Pay.CtxDaoModel(ctx).Filter(daoPay.Scene.Columns().PayScene, payScene).OrderDesc(daoPay.Pay.Columns().Sort).ListPri()
	if err != nil {
		return
	}

	res = &api.PayListRes{List: []api.PayInfo{}}
	list.Structs(&res.List)
	return
}

// 支付
func (controllerThis *Pay) Pay(ctx context.Context, req *api.PayPayReq) (res *api.PayPayRes, err error) {
	/**--------确定支付场景 开始--------**/
	var payScene uint
	sceneInfo := utils.GetCtxSceneInfo(ctx)
	sceneCode := sceneInfo[daoAuth.Scene.Columns().SceneCode].String()
	switch sceneCode {
	case `app`:
		loginInfo := utils.GetCtxLoginInfo(ctx)
		if loginInfo.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39994000, ``)
			return
		}
		payScene = 0
	default:
		err = utils.NewErrorCode(ctx, 39999998, ``)
		return
	}
	/**--------确定支付场景 结束--------**/

	/**--------确定支付数据 开始--------**/
	var payData pay.PayData
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

	payData.OrderNo = orderInfo[daoXxxx.Order.Columns().OrderNo].String()
	payData.Amount = orderInfo[daoXxxx.Order.Columns().Price].Float64()
	payData.Desc = `订单描述` */
	/**--------确定支付数据 开始--------**/

	payObj, err := pay.NewPay(ctx, req.PayId)
	if err != nil {
		return
	}

	var payResult pay.PayInfo
	switch payScene {
	case 0: //APP
		payResult, err = payObj.App(payData)
	case 1: //H5
		payResult, err = payObj.H5(payData)
	case 2: //扫码
		payResult, err = payObj.QRCode(payData)
	case 10, 11, 20: //微信小程序	微信公众号	支付宝小程序
		payData.Openid = ``
		payResult, err = payObj.Jsapi(payData)
	}
	if err != nil {
		return
	}

	res = &api.PayPayRes{
		PayStr: payResult.PayStr,
	}
	return
}

// 回调
func (controllerThis *Pay) Notify(ctx context.Context, req *api.PayNotifyReq) (res *api.CommonNoDataRes, err error) {
	payObj, err := pay.NewPay(ctx, req.PayId)
	if err != nil {
		return
	}

	r := g.RequestFromCtx(ctx)
	notifyInfo, err := payObj.Notify(r)
	if err != nil {
		payObj.NotifyRes(r, err.Error())
		return
	}

	// 订单回调处理
	gutil.Dump(notifyInfo)
	/* xxxxOrderHandler := daoXxxx.Order.CtxDaoModel(ctx)
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
		return
	})
	if err != nil {
		payObj.NotifyRes(r, err.Error())
		return
	} */

	payObj.NotifyRes(r, ``)
	return
}
