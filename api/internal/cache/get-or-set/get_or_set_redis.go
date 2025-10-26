package get_or_set

import (
	"api/internal/utils/jbredis"
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

var GetOrSetRedis = getOrSetRedis{}

type getOrSetRedis struct{}

func (cacheThis *getOrSetRedis) cache() redis.UniversalClient {
	return jbredis.DB()
}

func (cacheThis *getOrSetRedis) GetOrSet(ctx context.Context, key string, setFunc func() (value any, notExist bool, err error), ttl time.Duration, numLock, numRead uint8, oneTime time.Duration) (value any, notExist bool, err error) {
	value, notExist, err = GetOrSet.GetOrSet(ctx, key, func() (value any, notExist bool, err error) {
		value, notExist, err = setFunc()
		if notExist || err != nil {
			return
		}
		err = cacheThis.cache().SetEx(ctx, key, value, ttl).Err()
		return
	}, func() (value any, notExist bool, err error) {
		value, err = cacheThis.cache().Get(ctx, key).Result()
		notExist = err == redis.Nil
		if notExist {
			err = nil
		}
		return
	}, numLock, numRead, oneTime)
	return
}
