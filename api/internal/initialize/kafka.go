package initialize

import (
	"api/internal/utils/kafka"
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

func initKafka(ctx context.Context) {
	for group, config := range g.Cfg().MustGet(ctx, `kafka`).Map() {
		configMap := gconv.Map(config)
		kafka.AddProducer(ctx, group, configMap)
		kafka.AddConsumer(ctx, group, configMap)
	}
}
