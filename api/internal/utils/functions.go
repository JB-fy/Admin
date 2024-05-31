package utils

import (
	"api/internal/consts"
	"context"
	"fmt"
	"os/exec"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"golang.org/x/tools/imports"
)

// 生成错误码
func NewCode(ctx context.Context, code int, msg string, data ...map[string]any) gcode.Code {
	detail := map[string]any{}
	if len(data) > 0 && data[0] != nil {
		detail = data[0]
	}
	if msg == `` {
		msg = g.I18n().T(ctx, `code.`+gconv.String(code))
		if _, ok := detail[`i18nValues`]; ok {
			msg = fmt.Sprintf(msg, gconv.SliceAny(detail[`i18nValues`])...)
			delete(detail, `i18nValues`)
		}
	}
	return gcode.New(code, msg, detail)
}

// 生成错误码
func NewErrorCode(ctx context.Context, code int, msg string, data ...map[string]any) error {
	codeObj := NewCode(ctx, code, msg, data...)
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
func GetCtxSceneInfo(ctx context.Context) gdb.Record {
	tmp := ctx.Value(consts.ConstCtxSceneInfoName)
	if tmp == nil {
		return nil
	}
	return tmp.(gdb.Record)
}

// 设置登录身份信息
func SetCtxLoginInfo(r *ghttp.Request, info gdb.Record) {
	r.SetCtxVar(consts.ConstCtxLoginInfoName, info)
}

// 获取登录身份信息
func GetCtxLoginInfo(ctx context.Context) gdb.Record {
	tmp := ctx.Value(consts.ConstCtxLoginInfoName)
	if tmp == nil {
		return nil
	}
	return tmp.(gdb.Record)
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

// 获取服务器外网ip
func GetServerNetworkIp() string {
	cmd := exec.Command(`/bin/bash`, `-c`, `curl -s ifconfig.me`)
	output, _ := cmd.CombinedOutput()
	return string(output)
}

// 获取服务器内网ip
func GetServerLocalIp() string {
	cmd := exec.Command(`/bin/bash`, `-c`, `hostname -I`)
	output, _ := cmd.CombinedOutput()
	return gstr.Trim(string(output))
}

// go文件代码格式化
func GoFileFmt(filePath string) {
	fmtFuc := func(path, content string) string {
		res, err := imports.Process(path, []byte(content), nil)
		if err != nil {
			return content
		}
		return string(res)
	}
	gfile.ReplaceFileFunc(fmtFuc, filePath)
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
