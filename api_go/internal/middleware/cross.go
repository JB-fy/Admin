package middleware

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

func Cross(r *ghttp.Request) {
	/* corsOptions := r.Response.DefaultCORSOptions()
	corsOptions.AllowHeaders = "*"
	corsOptions.AllowMethods = "*"
	corsOptions.AllowDomain = []string{"xxxx.com"}
	if !r.Response.CORSAllowedOrigin(corsOptions) {
		r.Response.WriteStatus(http.StatusForbidden)
		return
	}
	r.Response.CORS(corsOptions) */
	r.Response.CORSDefault()
	r.Middleware.Next()
}
