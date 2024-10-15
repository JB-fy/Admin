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

	err := r.GetError()
	var code gcode.Code
	if err == nil {
		if r.Response.BufferLength() > 0 {
			return
		}
		if r.Response.Status > 0 && r.Response.Status != http.StatusOK {
			return
			/* switch r.Response.Status {
			case http.StatusNotFound:
				code = gcode.CodeNotFound
			case http.StatusForbidden:
				code = gcode.CodeNotAuthorized
			default:
				// code = gcode.CodeUnknown
				code = utils.NewCode(r.GetCtx(), 19999997, http.StatusText(r.Response.Status), r.GetHandlerResponse())
			}
			// It creates error as it can be retrieved by other middlewares.
			// err = gerror.NewCode(code, code.Message())
			err = gerror.NewCode(code)
			r.SetError(err) */
		}
		code = utils.NewCode(r.GetCtx(), 0, ``, r.GetHandlerResponse())
	} else {
		code = gerror.Code(err)
		if code != gcode.CodeInternalPanic && r.Response.BufferLength() > 0 {
			return
		}
		switch code {
		case gcode.CodeNil:
			code = utils.NewCode(r.GetCtx(), 99999999, err.Error())
		case gcode.CodeValidationFailed:
			code = utils.NewCode(r.GetCtx(), 89999999, err.Error())
		case gcode.CodeDbOperationError:
			match, _ := gregex.MatchString(`Error 1062.*: Duplicate.*for key '(?:[^\.]*\.)?([^']*)'$`, err.Error()) //mysql
			// match, _ := gregex.MatchString(`pq: duplicate key.*constraint "([^"]*)"$`, err.Error()) //pgsql
			if len(match) > 0 {
				code = utils.NewCode(r.GetCtx(), 29991062, ``, g.Map{`i18nValues`: []any{match[1]}})
			} else {
				msg := ``
				if g.Cfg().MustGet(r.GetCtx(), `dev`).Bool() { //开发环境抛出sql错误语句
					msg = err.Error()
				}
				code = utils.NewCode(r.GetCtx(), 29999999, msg)
			}
		case gcode.CodeInternalPanic:
			r.Response.WriteHeader(http.StatusOK)
			r.Response.ClearBuffer()
			code = utils.NewCode(r.GetCtx(), 19999998, err.Error())
		}
	}

	r.Response.WriteJson(map[string]any{
		`code`: code.Code(),
		`msg`:  code.Message(),
		`data`: code.Detail(),
	})
}
