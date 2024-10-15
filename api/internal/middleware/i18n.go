package middleware

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/gogf/gf/v2/net/ghttp"
)

func I18n(r *ghttp.Request) {
	language := r.Header.Get(`Language`)
	if !(language == `` || language == g.Cfg().MustGet(r.GetCtx(), `i18n.language`).String()) {
		ctx := gi18n.WithLanguage(r.GetCtx(), language)
		r.SetCtx(ctx)
	}
	r.Middleware.Next()
}
