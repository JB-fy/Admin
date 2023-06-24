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

	var (
		err  = r.GetError()
		res  = r.GetHandlerResponse()
		code = gerror.Code(err)
	)
	if err != nil {
		switch code {
		case gcode.CodeNil:
			code = gcode.CodeInternalError
		case gcode.CodeValidationFailed:
			code = gcode.New(89999999, ``, nil)
		}
		msg := err.Error()
		r.Response.WriteJson(map[string]interface{}{
			`code`: code.Code(),
			`msg`:  msg,
			`data`: res,
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
	} else {
		code = gcode.CodeOK
		msg = g.I18n().T(r.GetCtx(), `code.0`)
	} */

	r.Response.WriteJson(map[string]interface{}{
		`code`: 0,
		`msg`:  g.I18n().T(r.GetCtx(), `code.0`),
		`data`: res,
	})
}
