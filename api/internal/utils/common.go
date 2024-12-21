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
	r.SetCtxVar(consts.ConstCtxSceneInfoName, info)
}

// 获取场景信息
func GetCtxSceneInfo(ctx context.Context) (info gdb.Record) {
	info, _ = ctx.Value(consts.ConstCtxSceneInfoName).(gdb.Record)
	return
}

// 设置登录身份信息
func SetCtxLoginInfo(r *ghttp.Request, info gdb.Record) {
	r.SetCtxVar(consts.ConstCtxLoginInfoName, info)
}

// 获取登录身份信息
func GetCtxLoginInfo(ctx context.Context) (info gdb.Record) {
	info, _ = ctx.Value(consts.ConstCtxLoginInfoName).(gdb.Record)
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
	case 3: //http(s)://本地IP或网络IP:端口
		if IsDev(ctx) {
			url = gstr.Replace(r.GetUrl(), r.Host+r.URL.String(), g.Cfg().MustGetWithEnv(ctx, consts.SERVER_LOCAL_IP).String()+ctx.Value(http.ServerContextKey).(*http.Server).Addr)
		} else {
			url = gstr.Replace(r.GetUrl(), r.Host+r.URL.String(), g.Cfg().MustGetWithEnv(ctx, consts.SERVER_NETWORK_IP).String()+ctx.Value(http.ServerContextKey).(*http.Server).Addr)
		}
	}
	return
}

// 获取文件内容（通用）
func GetFileBytes(ctx context.Context, fileUrl string, serverOpt ...string) (fileBytes []byte, err error) {
	hostIp := g.Cfg().MustGetWithEnv(ctx, consts.SERVER_NETWORK_IP).String()
	if IsDev(ctx) {
		hostIp = g.Cfg().MustGetWithEnv(ctx, consts.SERVER_LOCAL_IP).String()
	}
	if hostIp != `` && gstr.Pos(fileUrl, hostIp) != -1 {
		return GetFileBytesByLocal(ctx, fileUrl, serverOpt...)
	}
	return GetFileBytesByRemote(ctx, fileUrl)
}

// 获取文件内容（本地文件）
func GetFileBytesByLocal(ctx context.Context, fileUrl string, serverOpt ...string) (fileBytes []byte, err error) {
	serverRoot := `server`
	if len(serverOpt) > 0 && serverOpt[0] != `` {
		serverRoot = serverOpt[0]
	}
	serverRoot += `.serverRoot`

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
func Tree(list g.List, id uint, priKey string, pidKey string) (tree g.List) {
	for k, v := range list {
		if gconv.Uint(v[pidKey]) == id {
			v[`children`] = Tree(list[(k+1):], gconv.Uint(v[priKey]), priKey, pidKey)
			tree = append(tree, v)
		}
	}
	return
}
