package consumer

import (
	"api/internal/utils/kafka/model"
	"context"
	"fmt"

	"github.com/IBM/sarama"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
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
			_, err = initConsumer(ctx, consumerConfig, config, consumerInfo, handlerMapOfTopic[consumerInfo.TopicArr[0]](ctx, config, consumerInfo))
		} else {
			if _, ok := handlerMapOfGroupId[consumerInfo.GroupId]; !ok {
				panic(`消费者(分组:` + config.Group + `，组ID:` + consumerInfo.GroupId + `)缺少处理器，请实现！`)
			}
			_, err = initConsumerGroup(ctx, consumerConfig, config, consumerInfo, handlerMapOfGroupId[consumerInfo.GroupId](ctx, config, consumerInfo))
		}
		if err != nil {
			panic(`消费者(分组:` + config.Group + `)连接失败：` + err.Error())
		}
	}
	g.Log(`kafka`).Info(ctx, `消费者(分组:`+config.Group+`)连接成功`)
}

func initConsumerGroup(ctx context.Context, saramaConfig *sarama.Config, config *model.Config, consumerInfo *model.ConsumerInfo, consumerGroupHandler sarama.ConsumerGroupHandler) (consumer sarama.ConsumerGroup, err error) {
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

func initConsumer(ctx context.Context, saramaConfig *sarama.Config, config *model.Config, consumerInfo *model.ConsumerInfo, runFunc func(msg *sarama.ConsumerMessage)) (consumer sarama.Consumer, err error) {
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
