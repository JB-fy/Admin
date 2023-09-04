package middleware

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

func Cross(r *ghttp.Request) {
	//也可以在nginx中直接设置全站跨域
	/* location / {
	    add_header Access-Control-Allow-Credentials true;
	    #add_header Access-Control-Allow-Origin $http_origin;
	    #add_header Access-Control-Allow-Origin 'http://www.xxxx.com';
	    add_header Access-Control-Allow-Origin *;
	    #add_header Access-Control-Allow-Methods 'GET, POST, PUT, DELETE, PATCH, OPTIONS';
	    add_header Access-Control-Allow-Methods *;
	    #add_header Access-Control-Allow-Headers 'X-Requested-With, Content-Type, Accept, Origin, Authorization';
	    add_header Access-Control-Allow-Headers *;
	    if ($request_method = 'OPTIONS') {
	        return 204;
	    }
	} */

	/* corsOptions := r.Response.DefaultCORSOptions()
	corsOptions.AllowHeaders = `*`
	corsOptions.AllowMethods = `*`
	corsOptions.AllowDomain = []string{`xxxx.com`}
	if !r.Response.CORSAllowedOrigin(corsOptions) {
		r.Response.WriteStatus(http.StatusForbidden)
		return
	}
	r.Response.CORS(corsOptions) */
	r.Response.CORSDefault()
	r.Middleware.Next()
}
