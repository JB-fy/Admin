package sub

import (
	"api/internal/utils"
	"context"
	"strings"
	"time"

	"github.com/gogf/gf/v2/util/gutil"
	"github.com/redis/go-redis/v9"
)

var (
	channelArrMap = map[string][]string{}
	subHandlerMap = map[string]map[string]func(ctx context.Context, msg *redis.Message) (err error){
		`default`: {
			`__keyevent@0__:expired`: func(ctx context.Context, msg *redis.Message) (err error) {
				gutil.Dump(msg)
				return
			},
		},
	}
)

func Add(ctx context.Context, group string, db redis.UniversalClient, channelArr ...string) {
	channelArrMap[group] = append(channelArrMap[group], channelArr...)
	channelSet := map[string]struct{}{}
	buf := utils.BytesBufferPoolGet()
	defer utils.BytesBufferPoolPut(buf)
	for _, channel := range channelArrMap[group] {
		if _, ok := channelSet[channel]; ok {
			continue
		}
		channelSet[channel] = struct{}{}
		if strings.HasSuffix(channel, `__:expired`) {
			buf.WriteString(`Ex`)
		}
	}
	if notifyKeyspaceEvents := buf.String(); notifyKeyspaceEvents != `` {
		if err := db.ConfigSet(ctx, `notify-keyspace-events`, notifyKeyspaceEvents).Err(); err != nil {
			panic(err)
		}
	}
	for _, channel := range channelArr {
		subHandler, ok := subHandlerMap[group][channel]
		if !ok {
			panic(`订阅(分组:` + group + `，通道:` + channel + `)缺少处理器，请实现！`)
		}
		go func() {
			for {
				subscribe := db.Subscribe(ctx, channel)
				if _, err := subscribe.Receive(ctx); err != nil {
					subscribe.Close()
					return
				}
				ch := subscribe.Channel()
				for msg := range ch {
					subHandler(ctx, msg)
				}
				subscribe.Close()
				time.Sleep(3 * time.Second)
			}
		}()
	}
}
