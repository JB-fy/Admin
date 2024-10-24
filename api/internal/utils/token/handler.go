package token

import (
	"api/internal/cache"
	"api/internal/utils"
	"context"

	"github.com/gogf/gf/v2/util/gconv"
)

func NewHandler(ctx context.Context, config map[string]any, sceneCode string) *Handler {
	handlerObj := Handler{
		Ctx:       ctx,
		Token:     NewToken(ctx, config),
		SceneCode: sceneCode,
	}
	gconv.Struct(config, &handlerObj)
	return &handlerObj
}

type Handler struct {
	Ctx        context.Context
	Token      Token
	SceneCode  string //场景标识。缓存使用，注意：在同一权限场景下，存在互相覆盖BUG时，须自定义sceneCode规避
	ActiveTime int64  `json:"active_time"` //失活时间。大于0生效，即当Token在一段秒数内未使用，判定失活
	IsUnique   bool   `json:"is_unique"`   //Token唯一。开启后，可限制用户多地，多设备登录，因同时只会有一个Token有效（新Token生成时，旧Token会失效）
}

func (handlerThis *Handler) Create(tokenInfo TokenInfo) (token string, err error) {
	token, err = handlerThis.Token.Create(tokenInfo)
	if err != nil {
		return
	}

	if handlerThis.ActiveTime > 0 {
		cache.NewTokenActive(handlerThis.Ctx, handlerThis.SceneCode, tokenInfo.LoginId).Set(handlerThis.ActiveTime)
	}

	if handlerThis.IsUnique {
		cache.NewTokenIsUnique(handlerThis.Ctx, handlerThis.SceneCode, tokenInfo.LoginId).Set(token, handlerThis.Token.GetExpireTime())
	}
	return
}

func (handlerThis *Handler) Parse(token string) (tokenInfo TokenInfo, err error) {
	tokenInfo, err = handlerThis.Token.Parse(token)
	if err != nil {
		err = utils.NewErrorCode(handlerThis.Ctx, 39994001, err.Error())
		return
	}

	if handlerThis.ActiveTime > 0 {
		cacheTokenActive := cache.NewTokenActive(handlerThis.Ctx, handlerThis.SceneCode, tokenInfo.LoginId)
		if isExists, _ := cacheTokenActive.Get(); isExists == 0 {
			err = utils.NewErrorCode(handlerThis.Ctx, 39994002, ``)
			return
		}
		cacheTokenActive.Set(handlerThis.ActiveTime)
	}

	if handlerThis.IsUnique {
		if checkToken, _ := cache.NewTokenIsUnique(handlerThis.Ctx, handlerThis.SceneCode, tokenInfo.LoginId).Get(); checkToken != token {
			err = utils.NewErrorCode(handlerThis.Ctx, 39994003, ``)
			return
		}
	}
	return
}
