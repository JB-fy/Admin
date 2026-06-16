package consumer

import (
	"api/internal/utils/kafka/internal"
	"api/internal/utils/kafka/model"
	"context"
	"fmt"

	"github.com/IBM/sarama"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	handlerMapOfGroupId = map[string]map[string]func(ctx context.Context, consumerConfig *model.ConsumerConfig) sarama.ConsumerGroupHandler{
		`default`: {
			/* `template`: func(ctx context.Context, consumerConfig *model.ConsumerConfig) sarama.ConsumerGroupHandler {
				return &GroupHandlerOfTemplate{Ctx: ctx, ConsumerConfig: consumerConfig}
			}, */
		},
	}
	handlerMapOfTopic = map[string]map[string]func(ctx context.Context, consumerConfig *model.ConsumerConfig) func(msg *sarama.ConsumerMessage){
		`default`: {
			// `template`: TopicHandlerOfTemplate,
		},
	}
)

func Add(ctx context.Context, config *model.Config) {
	for _, consumerConfig := range config.ConsumerList {
		consumerConfig.CommonConfig = &config.CommonConfig
		consumerConfig.SaramaConfig = model.CreateConsumerConfig(config, &consumerConfig)
		if consumerConfig.GroupId == `` {
			handler, ok := handlerMapOfTopic[consumerConfig.Group][consumerConfig.TopicArr[0]]
			if !ok {
				panic(fmt.Sprintf(`消费者(分组:%s,主题:%s)缺少处理器，请实现！`, consumerConfig.Group, consumerConfig.TopicArr[0]))
			}
			if _, err := internal.InitConsumer(ctx, &consumerConfig, handler(ctx, &consumerConfig)); err != nil {
				panic(err)
			}
			g.Log(`kafka`).Info(ctx, fmt.Sprintf(`消费者(分组:%s,主题:%s)连接成功`, consumerConfig.Group, consumerConfig.TopicArr[0]))
		} else {
			handler, ok := handlerMapOfGroupId[consumerConfig.Group][consumerConfig.GroupId]
			if !ok {
				panic(fmt.Sprintf(`消费者(分组:%s,组ID:%s)缺少处理器，请实现！`, consumerConfig.Group, consumerConfig.GroupId))
			}
			for range consumerConfig.Number {
				if _, err := internal.InitConsumerGroup(ctx, &consumerConfig, handler(ctx, &consumerConfig)); err != nil {
					panic(err)
				}
			}
			g.Log(`kafka`).Info(ctx, fmt.Sprintf(`消费者(分组:%s,组ID:%s)连接成功`, consumerConfig.Group, consumerConfig.GroupId))
		}
	}
}
