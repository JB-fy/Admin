package cache

import (
	"api/internal/consts"
	"context"
	"fmt"

	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
)

// sceneId 场景ID。注意：在同一权限场景下，存在互相覆盖BUG时，须自定义sceneId规避
//
// loginName 手机/邮箱/账号
var Salt = salt{redis: g.Redis()}

type salt struct{ redis *gredis.Redis }

func (cacheThis *salt) cache() *gredis.Redis {
	return cacheThis.redis
}

func (cacheThis *salt) key(sceneId string, loginName string) string {
	return fmt.Sprintf(consts.CACHE_SALT, sceneId, loginName)
}

func (cacheThis *salt) Set(ctx context.Context, sceneId string, loginName string, value string, ttl int64) (err error) {
	err = cacheThis.cache().SetEX(ctx, cacheThis.key(sceneId, loginName), value, ttl)
	return
}

func (cacheThis *salt) Get(ctx context.Context, sceneId string, loginName string) (value string, err error) {
	valueTmp, err := cacheThis.cache().Get(ctx, cacheThis.key(sceneId, loginName))
	if err != nil {
		return
	}
	value = valueTmp.String()
	return
}
