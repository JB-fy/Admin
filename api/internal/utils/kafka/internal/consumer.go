package internal

import (
	"api/internal/utils/kafka/model"
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/IBM/sarama"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

func InitConsumer(ctx context.Context, consumerConfig *model.ConsumerConfig, runFunc func(msg *sarama.ConsumerMessage)) (consumer sarama.Consumer, err error) {
	consumer, err = sarama.NewConsumer(consumerConfig.Hosts, consumerConfig.SaramaConfig)
	if err != nil {
		err = fmt.Errorf(`消费者(分组:%s,主题:%s)连接错误:%w`, consumerConfig.Group, consumerConfig.TopicArr[0], err)
		return
	}
	go func() {
		defer consumer.Close()
		partitions, err := consumer.Partitions(consumerConfig.TopicArr[0])
		if err != nil {
			err = fmt.Errorf(`消费者(分组:%s,主题:%s)获取分区错误:%w`, consumerConfig.Group, consumerConfig.TopicArr[0], err)
			g.Log(`kafka`).Error(ctx, err)
			return
		}
		var wg sync.WaitGroup
		for _, partition := range partitions { // 每个分区创建消费者
			consumePartition, err := consumer.ConsumePartition(consumerConfig.TopicArr[0], partition, sarama.OffsetNewest)
			if err != nil {
				err = fmt.Errorf(`消费者(分组:%s,主题:%s,分区:%d)创建错误:%w`, consumerConfig.Group, consumerConfig.TopicArr[0], partition, err)
				g.Log(`kafka`).Error(ctx, err)
				continue
			}
			wg.Go(func() {
				defer consumePartition.AsyncClose()
				for msg := range consumePartition.Messages() {
					runFunc(msg)
				}
			})
		}
		wg.Wait()
	}()
	return
}

func InitConsumerGroup(ctx context.Context, consumerConfig *model.ConsumerConfig, consumerGroupHandler sarama.ConsumerGroupHandler) (consumer sarama.ConsumerGroup, err error) {
	consumer, err = sarama.NewConsumerGroup(consumerConfig.Hosts, consumerConfig.GroupId, consumerConfig.SaramaConfig)
	if err != nil {
		err = fmt.Errorf(`消费者(分组:%s,组ID:%s)连接错误:%w`, consumerConfig.Group, consumerConfig.GroupId, err)
		return
	}

	go func() {
		defer consumer.Close()
		for {
			if err := consumer.Consume(ctx, consumerConfig.TopicArr, consumerGroupHandler); err != nil {
				err = fmt.Errorf(`消费者(分组:%s,组ID:%s,主题:%s)创建错误:%w`, consumerConfig.Group, consumerConfig.GroupId, gconv.String(consumerConfig.TopicArr), err)
				g.Log(`kafka`).Error(ctx, err)
			}
			time.Sleep(3 * time.Second)
		}
	}()
	return
}
