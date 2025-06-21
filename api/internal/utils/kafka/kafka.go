package kafka

import (
	"api/internal/utils/kafka/consumer"
	"api/internal/utils/kafka/handler"
	"api/internal/utils/kafka/model"
	"api/internal/utils/kafka/producer"
	"context"
	"errors"

	"github.com/IBM/sarama"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	producerMap = map[string]any{}
)

func AddProducer(ctx context.Context, group string, configMap map[string]any) {
	config := model.GetConfig(group, configMap)
	producerConfig := model.CreateProducerConfig(config)
	var producerTmp any
	var err error
	if config.ProducerType == `sync` {
		producerTmp, err = producer.InitSyncProducer(ctx, producerConfig, config)
	} else {
		producerTmp, err = producer.InitAsyncProducer(ctx, producerConfig, config)
	}
	if err != nil {
		panic(`生产者(分组:` + config.Group + `)连接失败：` + err.Error())
	}
	producerMap[config.Group] = producerTmp
	g.Log(`kafka`).Info(ctx, `生产者(分组:`+config.Group+`)连接成功`)
}

// 同步生产者才有返回值
func SendMessage(ctx context.Context, topic, value string, groupOpt ...string) (partition int32, offset int64, err error) {
	group := `default`
	if len(groupOpt) > 0 && groupOpt[0] != `` {
		group = groupOpt[0]
	}
	producerTmp, ok := producerMap[group]
	if !ok {
		err = errors.New(`生产者(分组:` + group + `)不存在`)
		return
	}
	switch producer := producerTmp.(type) {
	case sarama.AsyncProducer:
		producer.Input() <- &sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.StringEncoder(value),
		}
	case sarama.SyncProducer:
		partition, offset, err = producer.SendMessage(&sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.StringEncoder(value),
		})
		if err != nil {
			g.Log(`kafka`).Error(ctx, `生产者(分组:`+group+`)同步发送消息失败`, err)
		}
	}
	return
}

var (
	handlerMapOfTopic = map[string]func(ctx context.Context, config *model.Config, consumerInfo *model.ConsumerInfo) func(msg *sarama.ConsumerMessage){
		`test`: handler.TopicHandlerOfTemplate,
	}
	handlerMapOfGroupId = map[string]func(ctx context.Context, config *model.Config, consumerInfo *model.ConsumerInfo) sarama.ConsumerGroupHandler{
		`test`: func(ctx context.Context, config *model.Config, consumerInfo *model.ConsumerInfo) sarama.ConsumerGroupHandler {
			return &handler.GroupHandlerOfTemplate{Ctx: ctx, Config: config, ConsumerInfo: consumerInfo}
		},
	}
)

func AddConsumer(ctx context.Context, group string, configMap map[string]any) {
	config := model.GetConfig(group, configMap)
	consumerConfig := model.CreateConsumerConfig(config)
	var err error
	for _, consumerInfo := range config.ConsumerList {
		if consumerInfo.GroupId == `` {
			if _, ok := handlerMapOfTopic[consumerInfo.TopicArr[0]]; !ok {
				panic(`消费者(分组:` + config.Group + `，主题:` + consumerInfo.TopicArr[0] + `)缺少处理器，请实现！`)
			}
			_, err = consumer.InitConsumer(ctx, consumerConfig, config, consumerInfo, handlerMapOfTopic[consumerInfo.TopicArr[0]](ctx, config, consumerInfo))
		} else {
			if _, ok := handlerMapOfGroupId[consumerInfo.GroupId]; !ok {
				panic(`消费者(分组:` + config.Group + `，组ID:` + consumerInfo.GroupId + `)缺少处理器，请实现！`)
			}
			_, err = consumer.InitConsumerGroup(ctx, consumerConfig, config, consumerInfo, handlerMapOfGroupId[consumerInfo.GroupId](ctx, config, consumerInfo))
		}
		if err != nil {
			panic(`消费者(分组:` + config.Group + `)连接失败：` + err.Error())
		}
	}
	g.Log(`kafka`).Info(ctx, `消费者(分组:`+config.Group+`)连接成功`)
}
