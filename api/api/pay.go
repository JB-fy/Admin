package api

import (
	"github.com/gogf/gf/v2/frame/g"
)

/*--------回调 开始--------*/
type PayNotifyReq struct {
	g.Meta `path:"/notify" method:"get,post" tags:"支付" sm:"回调"`
}

/*--------回调 结束--------*/
