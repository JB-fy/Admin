package api

import (
	"github.com/gogf/gf/v2/frame/g"
)

/*--------发送验证码 开始--------*/
type EmailSendReq struct {
	g.Meta   `path:"/send" method:"post" tags:"APP/短信" sm:"发送验证码"`
	UseScene int    `json:"use_scene"  v:"required|in:4" dc:"使用场景：4绑定邮箱"`
	Email    string `json:"email"  v:"required|email" dc:"邮箱"`
	// UseScene int    `json:"use_scene"  v:"required|in:0,1,2,3,4,5" dc:"使用场景：0登录 1注册 2密码找回 3密码修改 4绑定邮箱 5解绑邮箱"`
	// Email    string `json:"email"  v:"required-unless:UseScene,3,UseScene,5|email" dc:"邮箱。use_scene=3或5时可不传"`
}

/*--------发送验证码 结束--------*/
