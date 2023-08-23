package middleware

import (
	"api/internal/utils"
	"net/http"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gregex"
)

func HandlerResponse(r *ghttp.Request) {
	r.Middleware.Next()

	if r.Response.BufferLength() > 0 {
		return
	}

	err := r.GetError()
	if err != nil {
		code := gerror.Code(err)
		msg := err.Error()
		switch code {
		case gcode.CodeNil:
			code = gcode.CodeInternalError
		case gcode.CodeValidationFailed:
			code = gcode.New(89999999, ``, nil)
		case gcode.CodeDbOperationError:
			match, _ := gregex.MatchString(`1062.*Duplicate.*for key '(?:[^\.]*\.)?([^']*)'$`, msg)
			if len(match) > 0 {
				code = gcode.New(29991062, ``, nil)
				msg = g.I18n().Tf(r.GetCtx(), `code.29991062`, match[1])
			} else {
				code = gcode.New(29999999, ``, nil)
				if !utils.IsDev(r.GetCtx()) {
					msg = g.I18n().T(r.GetCtx(), `code.29999999`)
				}
			}
		}
		r.Response.WriteJson(map[string]interface{}{
			`code`: code.Code(),
			`msg`:  msg,
			`data`: code.Detail(),
		})
		return
	}

	if r.Response.Status > 0 && r.Response.Status != http.StatusOK {
		/* msg := http.StatusText(r.Response.Status)
		code := gcode.CodeUnknown
		switch r.Response.Status {
		case http.StatusNotFound:
			code = gcode.CodeNotFound
		case http.StatusForbidden:
			code = gcode.CodeNotAuthorized
		}
		err = gerror.NewCode(code, msg)
		r.SetError(err) */
		r.Response.WriteStatus(r.Response.Status)
		return
	}

	r.Response.WriteJson(map[string]interface{}{
		`code`: 0,
		`msg`:  g.I18n().T(r.GetCtx(), `code.0`),
		`data`: r.GetHandlerResponse(),
	})
}
