package token

import (
	"api/internal/cache"
	"api/internal/utils"
	"context"

	"github.com/gogf/gf/v2/util/gconv"
)

func NewHandler(ctx context.Context, config map[string]any, sceneId string) *Handler {
	handlerObj := Handler{
		Ctx:     ctx,
		Token:   NewToken(ctx, config),
		SceneId: sceneId,
	}
	gconv.Struct(config, &handlerObj)
	return &handlerObj
}

type Handler struct {
	Ctx        context.Context
	Token      Token
	SceneId    string //场景ID。缓存使用，注意：在同一权限场景下，存在互相覆盖BUG时，须自定义sceneId规避
	ActiveTime int64  `json:"active_time"` //失活时间。大于0生效，防止长时间无操作（人离开）时，被他人趁机而入（一段秒数内Token未使用，判定失活）
	IsIP       bool   `json:"is_ip"`       //验证IP。开启后，可防止Token被盗用（验证使用Token时的IP与生成Token时的IP是否一致）
	IsUnique   bool   `json:"is_unique"`   //Token唯一。开启后，可限制用户多地、多设备登录（同时只会有一个Token有效，生成新Token时，旧Token失效）
}

func (handlerThis *Handler) Create(tokenInfo TokenInfo) (token string, err error) {
	token, err = handlerThis.Token.Create(tokenInfo)
	if err != nil {
		return
	}

	if handlerThis.ActiveTime > 0 {
		cache.NewTokenActive(handlerThis.Ctx, handlerThis.SceneId, tokenInfo.LoginId).Set(handlerThis.ActiveTime)
	}

	if handlerThis.IsUnique {
		cache.NewTokenIsUnique(handlerThis.Ctx, handlerThis.SceneId, tokenInfo.LoginId).Set(token, handlerThis.Token.GetExpireTime())
	}
	return
}

// 不验证IP时，ip传空
func (handlerThis *Handler) Parse(token string, ip string) (tokenInfo TokenInfo, err error) {
	tokenInfo, err = handlerThis.Token.Parse(token)
	if err != nil {
		err = utils.NewErrorCode(handlerThis.Ctx, 39994001, err.Error())
		return
	}
	if handlerThis.IsIP && ip != tokenInfo.IP {
		err = utils.NewErrorCode(handlerThis.Ctx, 39994001, ``) //直接使用39994001错误码！不报出具体原因，防止被攻击（攻击者知道原因会做针对处理）
		return
	}

	if handlerThis.ActiveTime > 0 {
		cacheTokenActive := cache.NewTokenActive(handlerThis.Ctx, handlerThis.SceneId, tokenInfo.LoginId)
		if isExists, _ := cacheTokenActive.Get(); isExists == 0 {
			err = utils.NewErrorCode(handlerThis.Ctx, 39994002, ``)
			return
		}
		cacheTokenActive.Set(handlerThis.ActiveTime)
	}

	if handlerThis.IsUnique {
		if checkToken, _ := cache.NewTokenIsUnique(handlerThis.Ctx, handlerThis.SceneId, tokenInfo.LoginId).Get(); checkToken != token {
			err = utils.NewErrorCode(handlerThis.Ctx, 39994003, ``)
			return
		}
	}
	return
}
