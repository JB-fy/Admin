package internal

import (
	"api/internal/utils/kafka/model"
	"context"

	"github.com/IBM/sarama"
	"github.com/gogf/gf/v2/frame/g"
)

func InitSyncProducer(ctx context.Context, saramaConfig *sarama.Config, config *model.Config) (producer sarama.SyncProducer, err error) {
	producer, err = sarama.NewSyncProducer(config.Hosts, saramaConfig)
	// defer producer.Close()
	return
}

func InitAsyncProducer(ctx context.Context, saramaConfig *sarama.Config, config *model.Config) (producer sarama.AsyncProducer, err error) {
	producer, err = sarama.NewAsyncProducer(config.Hosts, saramaConfig)
	if err != nil {
		return
	}
	// defer producer.Close()

	// 监听错误的消息
	go func() {
		for err := range producer.Errors() {
			// producer.Input() <- err.Msg
			g.Log(`kafka`).Error(ctx, `生产者(组:`+config.Group+`)异步发送消息失败`, err)
		}
	}()
	return
}
