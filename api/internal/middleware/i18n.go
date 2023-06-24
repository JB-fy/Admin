package middleware

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/gogf/gf/v2/net/ghttp"
)

func I18n(r *ghttp.Request) {
	/* //建议放服务启动位置执行。省去重复执行
	g.I18n().SetPath(g.Cfg().MustGet(r.GetCtx(), `i18n.path`).String())         //设置资源目录
	g.I18n().SetLanguage(g.Cfg().MustGet(r.GetCtx(), `i18n.language`).String()) //设置默认为中文（原默认为英文en） */
	language := r.Header.Get(`Language`)
	if !(language == `` || language == g.Cfg().MustGet(r.GetCtx(), `i18n.language`).String()) {
		ctx := gi18n.WithLanguage(r.GetCtx(), language)
		r.SetCtx(ctx)
	}
	r.Middleware.Next()
}
