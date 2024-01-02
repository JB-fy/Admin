package middleware

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

// 2.6版本进入控制器后，r.GetBody()和r.GetBodyString()读取不到数据，r.Body也一样
/* body, err := io.ReadAll(r.Body)
r.Body.Close()
if err != nil {
	return
}
bodyStr := string(body) */

// repeatableRead
// fasle时	r.GetBody()和r.GetBodyString()能读取到数据
// true时	r.GetBody()和r.GetBodyString()能读取到数据，r.Body也能
func BodyRepeatable(repeatableRead bool) func(r *ghttp.Request) {
	return func(r *ghttp.Request) {
		r.MakeBodyRepeatableRead(repeatableRead)
		r.Middleware.Next()
	}
}
