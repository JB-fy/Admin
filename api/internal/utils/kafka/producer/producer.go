package producer

import (
	"api/internal/utils/kafka/internal"
	"api/internal/utils/kafka/model"
	"context"
	"fmt"

	"github.com/IBM/sarama"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	producerMap = map[string]map[string]any{}
)

func Add(ctx context.Context, config *model.Config) {
	for _, producerConfig := range config.ProducerList {
		producerConfig.CommonConfig = &config.CommonConfig
		producerConfig.SaramaConfig = model.CreateProducerConfig(config, &producerConfig)
		var producerTmp any
		var err error
		if producerConfig.IsSync {
			producerTmp, err = internal.InitSyncProducer(ctx, &producerConfig)
		} else {
			producerTmp, err = internal.InitAsyncProducer(ctx, &producerConfig)
		}
		if err != nil {
			panic(fmt.Errorf(`生产者(分组:%s,主题:%s)连接错误:%w`, producerConfig.Group, producerConfig.Topic, err))
		}
		if producerMap[producerConfig.Group] == nil {
			producerMap[producerConfig.Group] = map[string]any{}
		}
		producerMap[producerConfig.Group][producerConfig.Topic] = producerTmp
		g.Log(`kafka`).Info(ctx, fmt.Sprintf(`生产者(分组:%s,主题:%s)连接成功`, producerConfig.Group, producerConfig.Topic))
	}
}

// 同步生产者才有返回值
func SendMessage(ctx context.Context, topic string, value []byte, groupOpt ...string) (partition int32, offset int64, err error) {
	group := `default`
	if len(groupOpt) > 0 && groupOpt[0] != `` {
		group = groupOpt[0]
	}
	producerTmp, ok := producerMap[group][topic]
	if !ok {
		producerTmp, ok = producerMap[group][``] //共用生产者
		if !ok {
			err = fmt.Errorf(`生产者(分组:%s,主题:%s)不存在`, group, topic)
			g.Log(`kafka`).Error(ctx, err)
			return
		}
	}
	switch producer := producerTmp.(type) {
	case sarama.AsyncProducer:
		producer.Input() <- &sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.ByteEncoder(value),
		}
	case sarama.SyncProducer:
		partition, offset, err = producer.SendMessage(&sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.ByteEncoder(value),
		})
		if err != nil {
			err = fmt.Errorf(`生产者(分组:%s,主题:%s)同步发送消息错误:%w`, group, topic, err)
			g.Log(`kafka`).Error(ctx, err)
			return
		}
	default:
		err = fmt.Errorf(`生产者(分组:%s,主题:%s)类型错误`, group, topic)
		g.Log(`kafka`).Error(ctx, err)
		return
	}
	return
}
