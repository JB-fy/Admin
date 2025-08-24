package jbredis

import (
	"context"
	"net"
	"slices"
	"time"

	"github.com/gogf/gf/v2/util/gconv"
	"github.com/redis/go-redis/v9"
)

var redisMap = map[string]redis.UniversalClient{}

func AddDB(ctx context.Context, group string, configMap map[string]any) {
	config := &redis.UniversalOptions{}
	gconv.Struct(configMap, config)
	if dialer := gconv.Strings(configMap[`dialer`]); len(dialer) > 0 {
		config.Dialer = func(ctx context.Context, network, addr string) (net.Conn, error) {
			if index := slices.Index(dialer, addr); index > 0 && index < len(config.Addrs) {
				addr = config.Addrs[index]
			}
			return net.Dial(network, addr)
		}
	}
	if config.MaxIdleConns == 0 {
		config.MaxIdleConns = 10
	}
	if config.MaxActiveConns == 0 {
		config.MaxActiveConns = 100
	}
	if config.ConnMaxLifetime == 0 {
		config.ConnMaxLifetime = 30 * time.Second
	}
	if config.ConnMaxIdleTime == 0 {
		config.ConnMaxIdleTime = 10 * time.Second
	}
	if config.PoolTimeout == 0 {
		config.PoolTimeout = 10 * time.Second
	}
	if config.WriteTimeout == 0 {
		config.WriteTimeout = -1
	}
	if config.ReadTimeout == 0 {
		config.ReadTimeout = -1
	}
	redisMap[group] = redis.NewUniversalClient(config)
}

func DB(opt ...string) (redis redis.UniversalClient) {
	group := `default`
	if len(opt) > 0 && opt[0] != `` {
		group = opt[0]
	}
	redis, ok := redisMap[group]
	if !ok {
		panic(`redis数据库连接(分组:` + group + `)不存在`)
	}
	return redis
}
