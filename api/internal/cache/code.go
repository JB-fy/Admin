package cache

import (
	"api/internal/consts"
	"api/internal/utils/jbredis"
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

// sceneId 场景ID。注意：在同一权限场景下，存在互相覆盖BUG时，须自定义sceneId规避
//
// to 手机/邮箱
//
// scene 场景
var Code = code{}

type code struct{}

// 可在该方法增加参数根据参数做分库逻辑处理
func (cacheThis *code) cache() redis.UniversalClient {
	// return jbredis.DB(`其它redis组`)
	return jbredis.DB()
}

func (cacheThis *code) key(sceneId string, to string, toType, scene uint8) string {
	return fmt.Sprintf(consts.CACHE_CODE, sceneId, to, toType, scene)
}

func (cacheThis *code) Set(ctx context.Context, sceneId string, to string, toType, scene uint8, value string, ttl time.Duration) error {
	return cacheThis.cache().SetEx(ctx, cacheThis.key(sceneId, to, toType, scene), value, ttl).Err()
}

func (cacheThis *code) Get(ctx context.Context, sceneId string, to string, toType, scene uint8) (string, error) {
	return cacheThis.cache().Get(ctx, cacheThis.key(sceneId, to, toType, scene)).Result()
}
