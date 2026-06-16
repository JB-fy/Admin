package internal

import (
	"api/internal/utils/kafka/model"
	"context"
	"fmt"

	"github.com/IBM/sarama"
	"github.com/gogf/gf/v2/frame/g"
)

func InitSyncProducer(ctx context.Context, producerConfig *model.ProducerConfig) (producer sarama.SyncProducer, err error) {
	producer, err = sarama.NewSyncProducer(producerConfig.Hosts, producerConfig.SaramaConfig)
	// defer producer.Close()
	return
}

func InitAsyncProducer(ctx context.Context, producerConfig *model.ProducerConfig) (producer sarama.AsyncProducer, err error) {
	producer, err = sarama.NewAsyncProducer(producerConfig.Hosts, producerConfig.SaramaConfig)
	if err != nil {
		return
	}
	// defer producer.Close()

	// 监听错误的消息
	go func() {
		for err := range producer.Errors() {
			// producer.Input() <- err.Msg
			g.Log(`kafka`).Error(ctx, fmt.Errorf(`生产者(分组:%s,主题:%s)异步发送消息错误:%w`, producerConfig.Group, producerConfig.Topic, err))
		}
	}()
	return
}
