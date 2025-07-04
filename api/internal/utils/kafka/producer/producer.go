package producer

import (
	"api/internal/utils/kafka/internal"
	"api/internal/utils/kafka/model"
	"context"
	"errors"

	"github.com/IBM/sarama"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	producerMap = map[string]any{}
)

func Add(ctx context.Context, group string, configMap map[string]any) {
	config := model.GetConfig(group, configMap)
	producerConfig := model.CreateProducerConfig(config)
	var producerTmp any
	var err error
	if config.ProducerType == `sync` {
		producerTmp, err = internal.InitSyncProducer(ctx, producerConfig, config)
	} else {
		producerTmp, err = internal.InitAsyncProducer(ctx, producerConfig, config)
	}
	if err != nil {
		panic(`生产者(分组:` + config.Group + `)连接失败：` + err.Error())
	}
	producerMap[config.Group] = producerTmp
	g.Log(`kafka`).Info(ctx, `生产者(分组:`+config.Group+`)连接成功`)
}

// 同步生产者才有返回值
func SendMessage(ctx context.Context, topic string, value []byte, groupOpt ...string) (partition int32, offset int64, err error) {
	group := `default`
	if len(groupOpt) > 0 && groupOpt[0] != `` {
		group = groupOpt[0]
	}
	producerTmp, ok := producerMap[group]
	if !ok {
		err = errors.New(`生产者(分组:` + group + `)不存在`)
		g.Log(`kafka`).Error(ctx, `发送消息失败`, err)
		return
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
			g.Log(`kafka`).Error(ctx, `生产者(分组:`+group+`)同步发送消息失败`, err)
		}
	}
	return
}
