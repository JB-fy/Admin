package controller

import (
	"api/api"
	daoAuth "api/internal/dao/auth"
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

var (
	payList = []api.PayListItem{
		{
			PayMethod: 0,
			PayName:   `支付宝APP支付`,
			PayIcon:   `http://JB.Admin.com/xxxx.png`,
		},
		{
			PayMethod: 1,
			PayName:   `支付宝H5支付`,
			PayIcon:   `http://JB.Admin.com/xxxx.png`,
		},
		{
			PayMethod: 2,
			PayName:   `支付宝JSAPI支付`,
			PayIcon:   `http://JB.Admin.com/xxxx.png`,
		},
		{
			PayMethod: 10,
			PayName:   `微信APP支付`,
			PayIcon:   `http://JB.Admin.com/xxxx.png`,
		},
		{
			PayMethod: 11,
			PayName:   `微信H5支付`,
			PayIcon:   `http://JB.Admin.com/xxxx.png`,
		},
		{
			PayMethod: 12,
			PayName:   `微信JSAPI支付`,
			PayIcon:   `http://JB.Admin.com/xxxx.png`,
		},
	}
)

// 列表
func (controllerThis *Pay) List(ctx context.Context, req *api.PayListReq) (res *api.PayListRes, err error) {
	res = &api.PayListRes{List: []api.PayListItem{}}
	sceneInfo := utils.GetCtxSceneInfo(ctx)
	sceneCode := sceneInfo[daoAuth.Scene.Columns().SceneCode].String()
	switch sceneCode {
	case `app`:
		res.List = append(res.List, payList[0], payList[3])
	default:
		err = utils.NewErrorCode(ctx, 39999998, ``)
		return
	}
	return
}

// 支付
func (controllerThis *Pay) Pay(ctx context.Context, req *api.PayPayReq) (res *api.PayPayRes, err error) {
	sceneInfo := utils.GetCtxSceneInfo(ctx)
	sceneCode := sceneInfo[daoAuth.Scene.Columns().SceneCode].String()
	var payData pay.PayData
	switch sceneCode {
	case `app`:
		loginInfo := utils.GetCtxLoginInfo(ctx)
		if loginInfo.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39994000, ``)
			return
		}
		//订单查询
		/* orderInfo, _ := dao.NewDaoHandler(ctx, &daoXxxx.Order).Filter(g.Map{
			daoXxxx.Order.Columns().OrderNo:   req.OrderNo,
			daoXxxx.Order.Columns().UserId:    loginInfo[`loginId`],
			daoXxxx.Order.Columns().PayStatus: 0,
		}).GetModel().One()
		if orderInfo.IsEmpty() {
			err = utils.NewErrorCode(ctx, 29999998, ``)
			return
		}

		payData.OrderNo = orderInfo[daoXxxx.Order.Columns().OrderNo].String()
		payData.Amount = orderInfo[daoXxxx.Order.Columns().Price].Float64()
		payData.Desc = `订单描述` */
	default:
		err = utils.NewErrorCode(ctx, 39999998, ``)
		return
	}

	var payInfo pay.PayInfo
	switch req.PayMethod {
	case 0: //APP支付(支付宝)
		payInfo, err = pay.NewPay(ctx, `payOfAli`).App(payData)
	case 1: //H5支付(支付宝)
		payInfo, err = pay.NewPay(ctx, `payOfAli`).H5(payData)
	/* case 2: //JSAPI支付(支付宝)
	payData.OpenId = ``
	payInfo, err = pay.NewPay(ctx, `payOfAli`).Jsapi(payData) */
	case 10: //APP支付(微信)
		payInfo, err = pay.NewPay(ctx, `payOfWx`).App(payData)
	case 11: //H5支付(微信)
		payData.ClientIp = g.RequestFromCtx(ctx).GetClientIp()
		payInfo, err = pay.NewPay(ctx, `payOfWx`).H5(payData)
	/* case 12: //JSAPI支付(微信)
	payData.OpenId = ``
	payInfo, err = pay.NewPay(ctx, `payOfWx`).Jsapi(payData) */
	default:
		err = utils.NewErrorCode(ctx, 30010000, ``)
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
			err = utils.NewErrorCode(ctx, 30010001, ``)
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
