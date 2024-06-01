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
		}
		r.Response.WriteJson(map[string]any{
			`code`: code.Code(),
			`msg`:  code.Message(),
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

	r.Response.WriteJson(map[string]any{
		`code`: 0,
		`msg`:  g.I18n().T(r.GetCtx(), `code.0`),
		`data`: r.GetHandlerResponse(),
	})
}
