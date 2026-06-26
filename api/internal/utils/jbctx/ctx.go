package jbctx

import (
	"api/internal/consts"
	daoAuth "api/internal/dao/auth"
	"context"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/net/ghttp"
)

// 设置场景信息
func SetSceneInfo(r *ghttp.Request, info gdb.Record) {
	r.SetCtxVar(consts.CTX_SCENE_INFO_NAME, info)
}

// 获取场景信息
func GetSceneInfo(ctx context.Context) (info gdb.Record) {
	info, _ = ctx.Value(consts.CTX_SCENE_INFO_NAME).(gdb.Record)
	return
}

// 获取场景ID
func GetSceneId(ctx context.Context) *gvar.Var {
	return GetSceneInfo(ctx)[daoAuth.Scene.Columns().SceneId]
}

// 设置登录身份信息
func SetLoginInfo(r *ghttp.Request, info gdb.Record) {
	r.SetCtxVar(consts.CTX_LOGIN_INFO_NAME, info)
}

// 获取登录身份信息
func GetLoginInfo(ctx context.Context) (info gdb.Record) {
	info, _ = ctx.Value(consts.CTX_LOGIN_INFO_NAME).(gdb.Record)
	return
}

// 获取登录ID
func GetLoginId(ctx context.Context) *gvar.Var {
	return GetLoginInfo(ctx)[consts.CTX_LOGIN_ID_NAME]
}
