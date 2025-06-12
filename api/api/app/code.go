package api

import (
	"github.com/gogf/gf/v2/frame/g"
)

/*--------发送验证码 开始--------*/
type CodeSendReq struct {
	g.Meta `path:"/send" method:"post" tags:"APP/验证码" sm:"发送"`
	Scene  uint   `json:"scene" v:"required|in:0,1,2,3,4,5,10,11,12,13,14,15" dc:"场景：0登录(手机) 1注册(手机) 2密码找回(手机) 3密码修改(手机) 4绑定(手机) 5解绑(手机) 10登录(邮箱) 11注册(邮箱) 12密码找回(邮箱) 13密码修改(邮箱) 14绑定(邮箱) 15解绑(邮箱)"`
	To     string `json:"to" v:"required-if:Scene,0,Scene,1,Scene,2,Scene,4,Scene,10,Scene,11,Scene,12,Scene,14" dc:"手机/邮箱。scene=0,1,2,4,10,11,12,14时必须"`
}

/*--------发送验证码 结束--------*/
