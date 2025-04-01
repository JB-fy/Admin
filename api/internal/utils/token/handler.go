package token

import (
	"api/internal/cache"
	daoAuth "api/internal/dao/auth"
	"api/internal/utils"
	"api/internal/utils/token/model"
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type Handler struct {
	Ctx        context.Context
	token      model.Token
	SceneId    string //场景ID。缓存使用，注意：在同一权限场景下，存在互相覆盖BUG时，须自定义sceneId规避
	ActiveTime int64  `json:"active_time"` //失活时间。大于0生效，防止长时间无操作（人离开）时，被他人趁机而入（一段秒数内Token未使用，判定失活）
	IsIP       bool   `json:"is_ip"`       //验证IP。开启后，可防止Token被盗用（验证使用Token时的IP与生成Token时的IP是否一致）
	IsUnique   bool   `json:"is_unique"`   //Token唯一。开启后，可限制用户多地、多设备登录（同时只会有一个Token有效，生成新Token时，旧Token失效）
}

func NewHandler(ctx context.Context /* , sceneIdOpt ...string */) model.Handler {
	handlerObj := &Handler{Ctx: ctx}
	sceneInfo := utils.GetCtxSceneInfo(ctx)
	/* if len(sceneIdOpt) > 0 {
		sceneInfo, _ = daoAuth.Scene.GetInfoFromCache(ctx, sceneIdOpt[0])
	} */
	config, _ := sceneInfo[daoAuth.Scene.Columns().SceneConfig].Map()[`token_config`].(g.Map)
	gconv.Struct(config, handlerObj)
	handlerObj.SceneId = sceneInfo[daoAuth.Scene.Columns().SceneId].String()
	handlerObj.token = NewToken(ctx, gconv.Uint(config[`token_type`]), config)
	return handlerObj
}

func (handlerThis *Handler) Create(loginId string, extData map[string]any) (token string, err error) {
	tokenInfo := model.TokenInfo{LoginId: loginId, ExtData: extData}
	if handlerThis.IsIP {
		tokenInfo.IP = g.RequestFromCtx(handlerThis.Ctx).GetClientIp()
	}

	token, err = handlerThis.token.Create(handlerThis.Ctx, tokenInfo)
	if err != nil {
		return
	}

	if handlerThis.ActiveTime > 0 {
		cache.TokenActive.Set(handlerThis.Ctx, handlerThis.SceneId, tokenInfo.LoginId, handlerThis.ActiveTime)
	}

	if handlerThis.IsUnique {
		cache.TokenIsUnique.Set(handlerThis.Ctx, handlerThis.SceneId, tokenInfo.LoginId, token, handlerThis.token.GetExpireTime())
	}
	return
}

func (handlerThis *Handler) Parse(token string) (tokenInfo model.TokenInfo, err error) {
	tokenInfo, err = handlerThis.token.Parse(handlerThis.Ctx, token)
	if err != nil {
		err = utils.NewErrorCode(handlerThis.Ctx, 39994001, err.Error())
		return
	}
	if handlerThis.IsIP && tokenInfo.IP != g.RequestFromCtx(handlerThis.Ctx).GetClientIp() {
		err = utils.NewErrorCode(handlerThis.Ctx, 39994001, ``) //直接使用39994001错误码！不报出具体原因，防止被攻击（攻击者知道原因会做针对处理）
		return
	}

	if handlerThis.ActiveTime > 0 {
		if isSet, _ := cache.TokenActive.Reset(handlerThis.Ctx, handlerThis.SceneId, tokenInfo.LoginId, handlerThis.ActiveTime); !isSet {
			err = utils.NewErrorCode(handlerThis.Ctx, 39994002, ``)
			return
		}
	}

	if handlerThis.IsUnique {
		if checkToken, _ := cache.TokenIsUnique.Get(handlerThis.Ctx, handlerThis.SceneId, tokenInfo.LoginId); checkToken != token {
			err = utils.NewErrorCode(handlerThis.Ctx, 39994003, ``)
			return
		}
	}
	return
}
