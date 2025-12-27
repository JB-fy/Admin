package api

import (
	"github.com/gogf/gf/v2/frame/g"
)

/*--------公众号回调 开始--------*/
type WxGzhNotifyReq struct {
	g.Meta `path:"/gzh-notify" method:"get,post" tags:"微信" sm:"公众号回调"`
	CommonHeaderReq
}

/*--------公众号回调 结束--------*/
