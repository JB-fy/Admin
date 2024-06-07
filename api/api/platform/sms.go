package api

import (
	"github.com/gogf/gf/v2/frame/g"
)

/*--------发送验证码 开始--------*/
type SmsSendReq struct {
	g.Meta   `path:"/send" method:"post" tags:"平台后台/短信" sm:"发送验证码"`
	UseScene int    `json:"use_scene"  v:"required|in:4" dc:"使用场景：4绑定手机"`
	Phone    string `json:"phone"  v:"required|phone" dc:"手机"`
}

/*--------发送验证码 结束--------*/
