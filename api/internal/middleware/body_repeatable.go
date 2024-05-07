package middleware

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

/* // GoFrame2.6版本POST表单请求(application/x-www-form-urlencoded 或 multipart/form-data)进入控制器后，r.GetBody()，r.GetBodyString()和r.Body都读取不到数据
body, err := io.ReadAll(r.Body)
r.Body.Close()
if err != nil {
	return
}
bodyStr := string(body)

// 启用该中间件r.GetBody()和r.GetBodyString()一定能读取到数据
// 设置repeatableRead为true，r.Body才能读取到数据 */
//
func BodyRepeatable(repeatableRead bool) func(r *ghttp.Request) {
	return func(r *ghttp.Request) {
		r.MakeBodyRepeatableRead(repeatableRead)
		r.Middleware.Next()
	}
}
