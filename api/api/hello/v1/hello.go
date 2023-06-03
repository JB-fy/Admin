package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type Req struct {
	g.Meta `path:"/hello" tags:"Hello" method:"get" summary:"You first hello api"`
}
type Res struct {
	g.Meta   `mime:"text/html" example:"string"`
	UserName string `p:"username"  v:"required|length:4,30#请输入账号|账号长度为:{min}到:{max}位"`
}

type TestReq struct {
	UserName string `p:"username"  v:"required|length:4,30#请输入账号|账号长度为:{min}到:{max}位"`
}
