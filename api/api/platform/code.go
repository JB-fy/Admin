package api

import (
	"github.com/gogf/gf/v2/frame/g"
)

/*--------发送验证码 开始--------*/
type CodeSendReq struct {
	g.Meta `path:"/send" method:"post" tags:"平台后台/验证码" sm:"发送"`
	Scene  uint   `json:"scene"  v:"required|in:4,14" dc:"场景：4绑定(手机) 14绑定(邮箱)"`
	Phone  string `json:"phone"  v:"required-if:Scene,4|phone" dc:"手机。scene=4时必须"`
	Email  string `json:"email"  v:"required-if:Scene,14|email" dc:"邮箱。scene=14时必须"`
}

/*--------发送验证码 结束--------*/
