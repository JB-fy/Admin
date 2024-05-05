package controller

import (
	"api/api"
	"context"
)

type Wx struct{}

func NewWx() *Wx {
	return &Wx{}
}

// 公众号回调
func (controllerThis *Wx) GzhNotify(ctx context.Context, req *api.WxGzhNotifyReq) (res *api.CommonNoDataRes, err error) {
	// wxGzhObj := wx.NewWxGzh(ctx)
	/* notifyInfo, err := wxGzhObj.Notify()
	if err != nil {
		payObj.NotifyRes(err.Error())
		return
	}

	payObj.NotifyRes(``) */
	return
}
