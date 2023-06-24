package middleware

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func HandlerResponse(r *ghttp.Request) {
	r.Middleware.Next()

	if r.Response.BufferLength() > 0 {
		return
	}

	err := r.GetError()
	if err != nil {
		code := gerror.Code(err)
		switch code {
		case gcode.CodeNil:
			code = gcode.CodeInternalError
		case gcode.CodeValidationFailed:
			code = gcode.New(89999999, ``, nil)
		}
		r.Response.WriteJson(map[string]interface{}{
			`code`: code.Code(),
			`msg`:  err.Error(),
			`data`: code.Detail(),
		})
		return
	}
	/* if r.Response.Status > 0 && r.Response.Status != http.StatusOK {
		msg = http.StatusText(r.Response.Status)
		switch r.Response.Status {
		case http.StatusNotFound:
			code = gcode.CodeNotFound
		case http.StatusForbidden:
			code = gcode.CodeNotAuthorized
		default:
			code = gcode.CodeUnknown
		}
		err = gerror.NewCode(code, msg)
		r.SetError(err)
		return
	} */

	r.Response.WriteJson(map[string]interface{}{
		`code`: 0,
		`msg`:  g.I18n().T(r.GetCtx(), `code.0`),
		`data`: r.GetHandlerResponse(),
	})
}
