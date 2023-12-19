package controller

import (
	"api/api"
	"api/internal/utils/pay"
	"context"

	"github.com/gogf/gf/v2/util/gutil"
)

type Pay struct{}

func NewPay() *Pay {
	return &Pay{}
}

// 回调
func (controllerThis *Pay) Notify(ctx context.Context, req *api.PayNotifyReq) (res *api.CommonNoDataRes, err error) {
	payObj := pay.NewPay(ctx)
	notifyInfo, err := payObj.Notify()
	if err != nil {
		payObj.NotifyRes(err.Error())
		return
	}
	//订单回调处理
	gutil.Dump(notifyInfo)

	payObj.NotifyRes(``)
	return
}
