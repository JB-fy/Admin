/* common.go与funcs.go的区别：
common.go：基于当前框架封装的常用函数（与框架耦合）
funcs.go：基于golang封装的常用函数（不与框架耦合） */

package utils

import (
	"api/internal/consts"
	"context"
	"fmt"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
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
	if len(dataOpt) > 0 && dataOpt[0] != nil {
		data = dataOpt[0]
	}
	if msg == `` {
		msg = g.I18n().T(ctx, `code.`+gconv.String(code))
		if dataTmp, ok := data.(map[string]any); ok {
			if _, ok := dataTmp[`i18nValues`]; ok {
				msg = fmt.Sprintf(msg, gconv.SliceAny(dataTmp[`i18nValues`])...)
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
	}
	return
}

// 列表转树状
func Tree(list g.List, id uint, priKey string, pidKey string) (tree g.List) {
	tree = g.List{}
	for k, v := range list {
		if gconv.Uint(v[pidKey]) == id {
			v[`children`] = Tree(list[(k+1):], gconv.Uint(v[priKey]), priKey, pidKey)
			tree = append(tree, v)
		}
	}
	return
}
