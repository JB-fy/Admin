package utils

import (
	"api/internal/consts"
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
)

func NewErrorCode(ctx context.Context, code int, msg string, data ...map[string]interface{}) error {
	dataTmp := map[string]interface{}{}
	if len(data) > 0 && data[0] != nil {
		dataTmp = data[0]
	}
	if msg == "" {
		switch code {
		case 29991063:
			msg = g.I18n().Tf(ctx, gconv.String(code), dataTmp["uniqueField"])
			delete(dataTmp, "uniqueField")
		case 89999996:
			msg = g.I18n().Tf(ctx, gconv.String(code), gconv.String(dataTmp["paramField"]))
			delete(dataTmp, "paramField")
		default:
			msg = g.I18n().Tf(ctx, gconv.String(code))
		}
	}
	return gerror.NewCode(gcode.New(code, "", dataTmp), msg)
}

func HttpFailJson(r *ghttp.Request, err error) {
	resData := map[string]interface{}{
		"code": 99999999,
		"msg":  err.Error(),
		"data": g.I18n().Tf(r.GetCtx(), "99999999"),
	}
	/* _, ok := err.(*gerror.Error)
	if ok { */
	code := gerror.Code(err)
	if code.Code() > 0 {
		resData["code"] = code.Code()
		resData["data"] = code.Detail()
	}
	r.Response.WriteJsonExit(resData)
}

func HttpSuccessJson(r *ghttp.Request, data map[string]interface{}, code int, msg ...string) {
	resData := map[string]interface{}{
		"code": code,
		"msg":  "",
		"data": data,
	}
	if len(msg) == 0 || msg[0] == "" {
		resData["msg"] = g.I18n().Tf(r.GetCtx(), gconv.String(code))
	} else {
		resData["msg"] = msg[0]
	}
	r.Response.WriteJsonExit(resData)
}

func SetCtxSceneInfo(r *ghttp.Request, info gdb.Record) {
	r.SetCtxVar(consts.ConstCtxSceneInfoName, info)
}

func GetCtxSceneInfo(ctx context.Context) gdb.Record {
	return ctx.Value(consts.ConstCtxSceneInfoName).(gdb.Record)
}

func GetCtxSceneCode(ctx context.Context) string {
	return GetCtxSceneInfo(ctx)["sceneCode"].String()
}

func SetCtxLoginInfo(r *ghttp.Request, info gdb.Record) {
	r.SetCtxVar(consts.ConstCtxLoginInfoName, info)
}

func GetCtxLoginInfo(ctx context.Context) gdb.Record {
	return ctx.Value(consts.ConstCtxLoginInfoName).(gdb.Record)
}
