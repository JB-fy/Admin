package middleware

import (
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
)

// DefaultHandlerResponse is the default implementation of HandlerResponse.
type DefaultHandlerResponse struct {
	Code    int         `json:"code"    dc:"Error code"`
	Message string      `json:"message" dc:"Error message"`
	Data    interface{} `json:"data"    dc:"Result data for certain request according API definition"`
}

func Cross(r *ghttp.Request) {
	r.Middleware.Next()

	var (
		err  = r.GetError()
		code = gerror.Code(err)
	)

	r.Response.WriteJson(DefaultHandlerResponse{
		Code:    code.Code(),
		Message: code.Message(),
		Data:    code.Detail(),
	})
}
