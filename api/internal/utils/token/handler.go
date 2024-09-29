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
	/* if len(sceneCode) > 0 {
		handlerObj.SceneCode = sceneCode[0]
	} */
	gconv.Struct(config, &handlerObj)
	return &handlerObj
}

type Handler struct {
	Ctx       context.Context
	Token     Token
	SceneCode string //场景标识。缓存使用，注意：在同一权限场景下，存在互相覆盖BUG时，须自定义sceneCode规避
	IsUnique  bool   `json:"isUnique"`
}

func (handlerThis *Handler) Create(tokenInfo TokenInfo) (token string, err error) {
	token, err = handlerThis.Token.Create(tokenInfo)
	if err != nil {
		return
	}

	// 限制多地登录，多设备登录
	if !handlerThis.IsUnique {
		cache.NewTokenUnique(handlerThis.Ctx, handlerThis.SceneCode, tokenInfo.LoginId).Set(token, handlerThis.Token.GetExpireTime())
	}
	return
}

func (handlerThis *Handler) Parse(token string) (tokenInfo TokenInfo, err error) {
	tokenInfo, err = handlerThis.Token.Parse(token)
	if err != nil {
		err = utils.NewErrorCode(handlerThis.Ctx, 39994001, err.Error())
		return
	}

	// 限制多地登录，多设备登录
	if !handlerThis.IsUnique {
		checkToken, _ := cache.NewTokenUnique(handlerThis.Ctx, handlerThis.SceneCode, tokenInfo.LoginId).Get()
		if checkToken != token {
			err = utils.NewErrorCode(handlerThis.Ctx, 39994002, ``)
			return
		}
	}
	return
}
