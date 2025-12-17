package common

import (
	"api/internal/consts"
	"api/internal/utils"
	"api/internal/utils/jbredis"
	"context"
	"fmt"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/genv"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/redis/go-redis/v9"
)

var IsLimit = isLimit{}

type isLimit struct {
	incrScripty string
	decrScripty string
}

func InitIsLimit(ctx context.Context) {
	var err error
	IsLimit.incrScripty, err = IsLimit.cache().ScriptLoad(ctx, `local count = redis.call('INCR', KEYS[1])
if tonumber(ARGV[2]) > 0 and (count == 1 or tonumber(ARGV[3]) == 1) then
	redis.call('PEXPIRE', KEYS[1], ARGV[2])	-- 毫秒
end
if count <= tonumber(ARGV[1]) then
	return count
else
	redis.call('DECR', KEYS[1])
    return 0
end`).Result()
	if err != nil {
		panic(`redis加载脚本错误：` + err.Error())
	}
	IsLimit.decrScripty, err = IsLimit.cache().ScriptLoad(ctx, `local count = redis.call('GET', KEYS[1])
if count and tonumber(count) > 0 then
    redis.call('DECR', KEYS[1])
end`).Result()
	if err != nil {
		panic(`redis加载脚本错误：` + err.Error())
	}

	if utils.IsDev(ctx) || g.Cfg().MustGet(ctx, `masterServerNetworkIp`).String() == genv.Get(consts.ENV_SERVER_NETWORK_IP).String() {
		match := IsLimit.key(`*`)
		var keyArr []string
		var cursor uint64
		for {
			keyArr, cursor, err = IsLimit.cache().Scan(ctx, cursor, match, 5000).Result()
			if err != nil {
				panic(`redis批量查询（` + match + `）错误：` + err.Error())
			}
			if len(keyArr) > 0 {
				/* // 主要删除无过期时间的key（系统重启会被影响）。而一般情况下都会允许短时间内超出限制，故直接全部删除
				var keyArrOfDel []string
				for _, key := range keyArr {
					ttl, _ := IsLimit.cache().PTTL(ctx, key).Result()
					if ttl > 0 {
						keyArrOfDel = append(keyArrOfDel, key)
					}
				}
				err = IsLimit.cache().Del(ctx, keyArrOfDel...).Err() */
				err = IsLimit.cache().Del(ctx, keyArr...).Err()
				if err != nil {
					panic(`redis批量删除（` + match + `）错误：` + err.Error())
				}
			}
			if cursor == 0 {
				break
			}
		}
	}
}

func (cacheThis *isLimit) cache() redis.UniversalClient {
	return jbredis.DB()
}

func (cacheThis *isLimit) key(key string) string {
	return fmt.Sprintf(consts.CACHE_IS_LIMIT, key)
}

func (cacheThis *isLimit) Incr(ctx context.Context, key string, limitNum uint, ttl time.Duration, isRefreshTTLOpt ...bool) (isLimit bool, err error) {
	isRefreshTTL := 0
	if len(isRefreshTTLOpt) > 0 && isRefreshTTLOpt[0] {
		isRefreshTTL = 1
	}
	countTmp, err := jbredis.DB().EvalSha(ctx, cacheThis.incrScripty, []string{cacheThis.key(key)}, []any{limitNum, ttl.Milliseconds(), isRefreshTTL}).Result()
	isLimit = gconv.Int64(countTmp) == 0
	return
}

func (cacheThis *isLimit) Decr(ctx context.Context, key string) (err error) {
	_, err = jbredis.DB().EvalSha(ctx, cacheThis.decrScripty, []string{cacheThis.key(key)}).Result()
	return
}

func (cacheThis *isLimit) keyOfNum(key string) string {
	return fmt.Sprintf(consts.CACHE_IS_LIMIT, key+`:num`)
}

func (cacheThis *isLimit) GetNum(ctx context.Context, key string, limitNumOfDef uint) (limitNum uint, err error) {
	limitNumTmp, err := cacheThis.cache().Get(ctx, cacheThis.keyOfNum(key)).Result()
	limitNum = gconv.Uint(limitNumTmp)
	if limitNum == 0 {
		limitNum = limitNumOfDef
	}
	return
}

func (cacheThis *isLimit) SetNum(ctx context.Context, key string, limitNum uint, ttl time.Duration) (err error) {
	err = cacheThis.cache().SetEx(ctx, cacheThis.keyOfNum(key), limitNum, ttl).Err()
	return
}
