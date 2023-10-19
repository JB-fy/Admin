package api

import (
	"github.com/gogf/gf/v2/frame/g"
)

/*--------发送验证码 开始--------*/
type SmsSendReq struct {
	g.Meta   `path:"/send" method:"post" tags:"APP/短信" sm:"发送验证码"`
	UseScene uint   `json:"useScene"  v:"required|in:0,1,2,3,4,5" dc:"使用场景：0登录 1注册 2密码找回 3密码修改 4绑定手机 5解绑手机"`
	Phone    string `json:"phone"  v:"required-unless:UseScene,3,UseScene,5|phone" dc:"手机。useScene=3或5时可不传"`
}

/*--------发送验证码 结束--------*/
