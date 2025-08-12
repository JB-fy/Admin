package consumer

import (
	"api/internal/utils/kafka/model"
	"context"
	"errors"
	"fmt"
	"syscall"
	"time"

	"github.com/IBM/sarama"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type GroupHandlerOfTemplate struct {
	Ctx          context.Context
	Config       *model.Config
	ConsumerInfo *model.ConsumerInfo
}

func (handlerThis *GroupHandlerOfTemplate) Setup(session sarama.ConsumerGroupSession) (err error) {
	g.Log(`kafka`).Info(handlerThis.Ctx, fmt.Sprintf(`消费者(分组:%s，组ID:%s，主题:%s)初始化`, handlerThis.Config.Group, handlerThis.ConsumerInfo.GroupId, gconv.String(handlerThis.ConsumerInfo.TopicArr)))
	return
}

func (handlerThis *GroupHandlerOfTemplate) Cleanup(session sarama.ConsumerGroupSession) (err error) {
	g.Log(`kafka`).Info(handlerThis.Ctx, fmt.Sprintf(`消费者(分组:%s，组ID:%s，主题:%s)关闭`, handlerThis.Config.Group, handlerThis.ConsumerInfo.GroupId, gconv.String(handlerThis.ConsumerInfo.TopicArr)))
	time.Sleep(3 * time.Second)
	syscall.Kill(syscall.Getpid(), syscall.SIGTERM) //消费者组中断时，直接关闭进程，触发服务重启
	return
}

func (handlerThis *GroupHandlerOfTemplate) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) (err error) {
	for msg := range claim.Messages() {
		handlerThis.handle(msg)
		// session.MarkMessage(msg, ``) // 标记消息为已处理
		// session.Commit()             // 马上提交到kafka
	}
	return
}

func (handlerThis *GroupHandlerOfTemplate) handle(msg *sarama.ConsumerMessage) (err error) {
	defer func() { //防止panic导致消费者断开
		if rec := recover(); rec != nil {
			err = errors.New(`panic错误：` + gconv.String(rec))
		}
	}()
	// 业务处理
	return
}
