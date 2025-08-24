package jbredis

import (
	"context"
	"net"
	"slices"

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
