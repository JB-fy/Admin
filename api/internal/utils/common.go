/* common.go与funcs.go的区别：
common.go：基于当前框架封装的常用函数（与框架耦合）
funcs.go：基于golang封装的常用函数（不与框架耦合） */

package utils

import (
	"api/internal/consts"
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

// 是否开发环境
func IsDev(ctx context.Context) bool {
	return g.Cfg().MustGet(ctx, `dev`).Bool()
}

// 生成错误码
func NewCode(ctx context.Context, code int, msg string, dataOpt ...any) gcode.Code {
	var data any
	if len(dataOpt) > 0 {
		data = dataOpt[0]
	}
	if msg == `` {
		msg = g.I18n().T(ctx, `code.`+gconv.String(code))
		if dataTmp, ok := data.(map[string]any); ok {
			if i18nValues, ok := dataTmp[`i18nValues`]; ok {
				msg = fmt.Sprintf(msg, gconv.SliceAny(i18nValues)...)
				delete(dataTmp, `i18nValues`)
			}
		}
	}
	return gcode.New(code, msg, data)
}

// 生成错误
func NewErrorCode(ctx context.Context, code int, msg string, dataOpt ...any) error {
	codeObj := NewCode(ctx, code, msg, dataOpt...)
	return gerror.NewCode(codeObj /* , codeObj.Message() */)
}

// Http返回json
func HttpWriteJson(ctx context.Context, data map[string]any, code int, msg string) {
	resData := map[string]any{
		`code`: code,
		`msg`:  msg,
		`data`: data,
	}
	if msg == `` {
		resData[`msg`] = g.I18n().T(ctx, `code.`+gconv.String(code))
	}
	g.RequestFromCtx(ctx).Response.WriteJson(resData)
}

// 设置场景信息
func SetCtxSceneInfo(r *ghttp.Request, info gdb.Record) {
	r.SetCtxVar(consts.CTX_SCENE_INFO_NAME, info)
}

// 获取场景信息
func GetCtxSceneInfo(ctx context.Context) (info gdb.Record) {
	info, _ = ctx.Value(consts.CTX_SCENE_INFO_NAME).(gdb.Record)
	return
}

// 设置登录身份信息
func SetCtxLoginInfo(r *ghttp.Request, info gdb.Record) {
	r.SetCtxVar(consts.CTX_LOGIN_INFO_NAME, info)
}

// 获取登录身份信息
func GetCtxLoginInfo(ctx context.Context) (info gdb.Record) {
	info, _ = ctx.Value(consts.CTX_LOGIN_INFO_NAME).(gdb.Record)
	return
}

// 获取当前请求Url
func GetRequestUrl(ctx context.Context, flag int) (url string) {
	r := g.RequestFromCtx(ctx)
	switch flag {
	case 0: //http(s)://www.xxxx.com
		url = gstr.Replace(r.GetUrl(), r.URL.String(), ``)
	case 1: //http(s)://www.xxxx.com/test
		url = gstr.Replace(r.GetUrl(), r.URL.String(), ``) + r.URL.Path
	case 2: //http(s)://www.xxxx.com/test?a=1&b=2
		url = r.GetUrl()
	case 10, 20: //http(s)://外网IP:端口	//http(s)://内网IP:端口
		url = r.GetUrl()
		ip := g.Cfg().MustGetWithEnv(ctx, consts.LOCAL_SERVER_NETWORK_IP).String()
		if flag == 20 {
			ip = g.Cfg().MustGetWithEnv(ctx, consts.LOCAL_SERVER_LOCAL_IP).String()
		}
		addr := ctx.Value(http.ServerContextKey).(*http.Server).Addr
		if gstr.Pos(url, `https`) == 0 {
			if portOfHttps := r.Server.GetListenedHTTPSPort(); portOfHttps != -1 {
				addr = `:` + gconv.String(portOfHttps)
			} else {
				url = gstr.Replace(url, `https`, `http`, 1)
			}
		}
		url = gstr.Replace(url, r.Host+r.URL.String(), ip+addr)
	}
	return
}

// 获取文件内容（通用）
func GetFileBytes(ctx context.Context, fileUrl string) (fileBytes []byte, err error) {
	for _, ip := range []string{g.Cfg().MustGetWithEnv(ctx, consts.LOCAL_SERVER_NETWORK_IP).String(), g.Cfg().MustGetWithEnv(ctx, consts.LOCAL_SERVER_LOCAL_IP).String()} {
		if ip != `` && gstr.Pos(fileUrl, ip) != -1 {
			return GetFileBytesByLocal(ctx, fileUrl)
		}
	}
	return GetFileBytesByRemote(ctx, fileUrl)
}

// 获取文件内容（本地文件）
func GetFileBytesByLocal(ctx context.Context, fileUrl string, serverNameOpt ...string) (fileBytes []byte, err error) {
	serverRoot := `server.serverRoot`
	if len(serverNameOpt) > 0 && serverNameOpt[0] != `` {
		serverRoot = `server.` + serverNameOpt[0] + `.serverRoot`
	} else if r := g.RequestFromCtx(ctx); r != nil {
		if serverName := r.Server.GetName(); serverName != ghttp.DefaultServerName {
			serverRoot = `server.` + serverName + `.serverRoot`
		}
	}
	urlObj, err := url.Parse(fileUrl)
	file := g.Cfg().MustGet(ctx, serverRoot).String() + urlObj.Path
	fileBytes = gfile.GetBytes(file)
	return
}

var getFileClient = g.Client()

// 获取文件内容（远程文件）
func GetFileBytesByRemote(ctx context.Context, fileUrl string) (fileBytes []byte, err error) {
	res, err := getFileClient.Get(ctx, fileUrl)
	if err != nil {
		return
	}
	defer res.Close()

	fileBytes = res.ReadAll()
	return
}

// 列表转树状
func Tree(list g.List, id any, priKey string, pidKey string) (tree g.List) {
	idStr := gconv.String(id)
	for k, v := range list {
		if gconv.String(v[pidKey]) == idStr {
			v[`children`] = Tree(list[(k+1):], v[priKey], priKey, pidKey)
			tree = append(tree, v)
		}
	}
	return
}
