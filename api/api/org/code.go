package api

import (
	"github.com/gogf/gf/v2/frame/g"
)

/*--------发送验证码 开始--------*/
type CodeSendReq struct {
	g.Meta `path:"/send" method:"post" tags:"机构后台/验证码" sm:"发送"`
	Scene  uint   `json:"scene"  v:"required|in:4,14" dc:"场景：4绑定(手机) 14绑定(邮箱)"`
	To     string `json:"to"  v:"required-if:Scene,4,Scene,14" dc:"手机/邮箱。scene=4,14时必须"`
}

/*--------发送验证码 结束--------*/
