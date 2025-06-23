package consumer

import (
	"api/internal/utils/kafka/internal"
	"api/internal/utils/kafka/model"
	"context"

	"github.com/IBM/sarama"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	handlerMapOfGroupId = map[string]func(ctx context.Context, config *model.Config, consumerInfo *model.ConsumerInfo) sarama.ConsumerGroupHandler{
		/* `template`: func(ctx context.Context, config *model.Config, consumerInfo *model.ConsumerInfo) sarama.ConsumerGroupHandler {
			return &kafka.GroupHandlerOfTemplate{Ctx: ctx, Config: config, ConsumerInfo: consumerInfo}
		}, */
	}
	handlerMapOfTopic = map[string]func(ctx context.Context, config *model.Config, consumerInfo *model.ConsumerInfo) func(msg *sarama.ConsumerMessage){
		// `template`: kafka.TopicHandlerOfTemplate,
	}
)

func Add(ctx context.Context, group string, configMap map[string]any) {
	config := model.GetConfig(group, configMap)
	consumerConfig := model.CreateConsumerConfig(config)
	var err error
	for _, consumerInfo := range config.ConsumerList {
		if consumerInfo.GroupId == `` {
			if _, ok := handlerMapOfTopic[consumerInfo.TopicArr[0]]; !ok {
				panic(`消费者(分组:` + config.Group + `，主题:` + consumerInfo.TopicArr[0] + `)缺少处理器，请实现！`)
			}
			_, err = internal.InitConsumer(ctx, consumerConfig, config, consumerInfo, handlerMapOfTopic[consumerInfo.TopicArr[0]](ctx, config, consumerInfo))
		} else {
			if _, ok := handlerMapOfGroupId[consumerInfo.GroupId]; !ok {
				panic(`消费者(分组:` + config.Group + `，组ID:` + consumerInfo.GroupId + `)缺少处理器，请实现！`)
			}
			_, err = internal.InitConsumerGroup(ctx, consumerConfig, config, consumerInfo, handlerMapOfGroupId[consumerInfo.GroupId](ctx, config, consumerInfo))
		}
		if err != nil {
			panic(`消费者(分组:` + config.Group + `)连接失败：` + err.Error())
		}
	}
	g.Log(`kafka`).Info(ctx, `消费者(分组:`+config.Group+`)连接成功`)
}
