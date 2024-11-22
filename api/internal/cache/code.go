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
// to 手机/邮箱
//
// scene 场景
var Code = code{redis: g.Redis()}

type code struct{ redis *gredis.Redis }

// 可在该方法增加参数根据参数做分库逻辑处理
func (cacheThis *code) cache() *gredis.Redis {
	// return g.Redis(`其它redis组`)
	return cacheThis.redis
}

func (cacheThis *code) key(sceneId string, to string, scene uint) string {
	return fmt.Sprintf(consts.CacheCodeFormat, sceneId, to, scene)
}

func (cacheThis *code) Set(ctx context.Context, sceneId string, to string, scene uint, value string, ttl int64) (err error) {
	err = cacheThis.cache().SetEX(ctx, cacheThis.key(sceneId, to, scene), value, ttl)
	return
}

func (cacheThis *code) Get(ctx context.Context, sceneId string, to string, scene uint) (value string, err error) {
	valueTmp, err := cacheThis.cache().Get(ctx, cacheThis.key(sceneId, to, scene))
	if err != nil {
		return
	}
	value = valueTmp.String()
	return
}
