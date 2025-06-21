package initialize

import (
	"api/internal/utils/kafka/consumer"
	"api/internal/utils/kafka/producer"
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

func initKafka(ctx context.Context) {
	for group, config := range g.Cfg().MustGet(ctx, `kafka`).Map() {
		configMap := gconv.Map(config)
		producer.Add(ctx, group, configMap)
		consumer.Add(ctx, group, configMap)
	}
}
