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
// loginName 手机/邮箱/账号
var Salt = salt{}

type salt struct{}

func (cacheThis *salt) cache() redis.UniversalClient {
	return jbredis.DB()
}

func (cacheThis *salt) key(sceneId string, loginName string) string {
	return fmt.Sprintf(consts.CACHE_SALT, sceneId, loginName)
}

func (cacheThis *salt) Set(ctx context.Context, sceneId string, loginName string, value string, ttl time.Duration) error {
	return cacheThis.cache().SetEx(ctx, cacheThis.key(sceneId, loginName), value, ttl).Err()
}

func (cacheThis *salt) Get(ctx context.Context, sceneId string, loginName string) (string, error) {
	return cacheThis.cache().Get(ctx, cacheThis.key(sceneId, loginName)).Result()
}
