package consumer

import (
	"api/internal/utils/kafka/model"
	"context"
	"fmt"

	"github.com/IBM/sarama"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

func InitConsumer(ctx context.Context, saramaConfig *sarama.Config, config *model.Config, consumerInfo *model.ConsumerInfo, runFunc func(msg *sarama.ConsumerMessage)) (consumer sarama.Consumer, err error) {
	consumer, err = sarama.NewConsumer(config.Hosts, saramaConfig)
	if err != nil {
		return
	}
	// defer consumer.Close() //长期跑，不准关闭

	partitions, err := consumer.Partitions(consumerInfo.TopicArr[0]) // 获取主题的分区列表
	if err != nil {
		return
	}
	// consumePartitionArr := []sarama.PartitionConsumer{}
	// var consumePartition sarama.PartitionConsumer
	for _, partition := range partitions { // 每个分区创建消费者
		consumePartition, err := consumer.ConsumePartition(consumerInfo.TopicArr[0], partition, sarama.OffsetNewest)
		if err != nil {
			g.Log(`kafka`).Error(ctx, fmt.Sprintf(`消费者(分组:%s，主题:%s，分区:%d)创建失败`, config.Group, consumerInfo.TopicArr[0], partition), err)
			continue
		}
		// consumePartitionArr = append(consumePartitionArr, consumePartition)
		// defer consumePartition.AsyncClose()	//长期跑，不准关闭

		go func() {
			for msg := range consumePartition.Messages() {
				runFunc(msg)
			}
		}()
	}
	return
}

func InitConsumerGroup(ctx context.Context, saramaConfig *sarama.Config, config *model.Config, consumerInfo *model.ConsumerInfo, consumerGroupHandler sarama.ConsumerGroupHandler) (consumer sarama.ConsumerGroup, err error) {
	consumer, err = sarama.NewConsumerGroup(config.Hosts, consumerInfo.GroupId, saramaConfig)
	if err != nil {
		return
	}
	// defer consumer.Close() //长期跑，不准关闭

	go func() {
		if err := consumer.Consume(ctx, consumerInfo.TopicArr, consumerGroupHandler); err != nil {
			g.Log(`kafka`).Error(ctx, fmt.Sprintf(`消费者(分组:%s，组ID:%s，主题:%s)创建失败`, config.Group, consumerInfo.GroupId, gconv.String(consumerInfo.TopicArr)), err)
			return
		}
	}()
	return
}
