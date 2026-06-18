package consumer

import (
	cacheCommon "api/internal/cache/common"
	"api/internal/utils/kafka/model"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/IBM/sarama"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type GroupHandlerOfTemplate struct {
	Ctx            context.Context
	ConsumerConfig *model.ConsumerConfig
}

func (handlerThis *GroupHandlerOfTemplate) Setup(session sarama.ConsumerGroupSession) (err error) {
	g.Log(`kafka`).Info(handlerThis.Ctx, fmt.Sprintf(`消费者(分组:%s,组ID:%s,主题:%s)初始化`, handlerThis.ConsumerConfig.Group, handlerThis.ConsumerConfig.GroupId, gconv.String(handlerThis.ConsumerConfig.TopicArr)))
	return
}

func (handlerThis *GroupHandlerOfTemplate) Cleanup(session sarama.ConsumerGroupSession) (err error) {
	err = session.Context().Err()
	if err == nil {
		err = errors.New(`可能是Rebalance`)
	}
	g.Log(`kafka`).Info(handlerThis.Ctx, fmt.Errorf(`消费者(分组:%s,组ID:%s,主题:%s)关闭:%w`, handlerThis.ConsumerConfig.Group, handlerThis.ConsumerConfig.GroupId, gconv.String(handlerThis.ConsumerConfig.TopicArr), err))
	return
}

func (handlerThis *GroupHandlerOfTemplate) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) (err error) {
	chLimit := cacheCommon.IsLimitLocal.GetChan(`kafkaTemplateLimit`, handlerThis.ConsumerConfig.Number+5) //限制超时任务数量
	for msg := range claim.Messages() {
		// handlerThis.handle(ctx, msg)执行时间超过handlerThis.ConsumerConfig.SaramaConfig.Consumer.Group.Session.Timeout配置时，会造成kafka消费组假死（不再消费消息，却还能接收消息）
		isRunTimeout, _, _ := cacheCommon.IsRunTimeout.IsRunTimeout(handlerThis.Ctx, handlerThis.ConsumerConfig.SaramaConfig.Consumer.Group.Session.Timeout-2*time.Second, func(ctx context.Context) (value any, err error) {
			err = cacheCommon.IsLimitLocal.Acquire(handlerThis.Ctx, chLimit, 0)
			if err != nil { //排队超时处理
				g.Log(`kafka`).Debug(handlerThis.Ctx, fmt.Sprintf(`任务(消息:%v)排队超时`, msg))
				return
			}
			defer cacheCommon.IsLimitLocal.Release(handlerThis.Ctx, chLimit)
			err = handlerThis.handle(handlerThis.Ctx, msg)
			return
		})
		if isRunTimeout { //超时处理
			g.Log(`kafka`).Debug(handlerThis.Ctx, fmt.Sprintf(`任务(消息:%v)超时`, msg))
		}
		session.MarkMessage(msg, ``) // 标记消息为已处理
		if !handlerThis.ConsumerConfig.SaramaConfig.Consumer.Offsets.AutoCommit.Enable {
			session.Commit() // 马上提交到kafka
		}
	}
	return
}

func (handlerThis *GroupHandlerOfTemplate) handle(ctx context.Context, msg *sarama.ConsumerMessage) (err error) {
	defer func() { //防止panic导致消费者断开
		if rec := recover(); rec != nil {
			err = fmt.Errorf(`panic错误:%v`, rec)
		}
	}()
	// 业务处理
	return
}
