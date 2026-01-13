package consumer

import (
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
	Ctx          context.Context
	Config       *model.Config
	ConsumerInfo *model.ConsumerInfo
	SaramaConfig *sarama.Config
}

func (handlerThis *GroupHandlerOfTemplate) Setup(session sarama.ConsumerGroupSession) (err error) {
	g.Log(`kafka`).Info(handlerThis.Ctx, fmt.Sprintf(`消费者(分组:%s，组ID:%s，主题:%s)初始化`, handlerThis.Config.Group, handlerThis.ConsumerInfo.GroupId, gconv.String(handlerThis.ConsumerInfo.TopicArr)))
	return
}

func (handlerThis *GroupHandlerOfTemplate) Cleanup(session sarama.ConsumerGroupSession) (err error) {
	g.Log(`kafka`).Info(handlerThis.Ctx, fmt.Sprintf(`消费者(分组:%s，组ID:%s，主题:%s)关闭`, handlerThis.Config.Group, handlerThis.ConsumerInfo.GroupId, gconv.String(handlerThis.ConsumerInfo.TopicArr)))
	// time.Sleep(3 * time.Second /* + time.Duration(3-time.Now().Second()%3)*time.Second */)
	// syscall.Kill(syscall.Getpid(), syscall.SIGTERM) //消费者组中断时，直接关闭进程，触发服务重启
	return
}

func (handlerThis *GroupHandlerOfTemplate) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) (err error) {
	for msg := range claim.Messages() {
		// handlerThis.handle(ctx, msg)执行时间超过handlerThis.SaramaConfig.Consumer.Group.Session.Timeout配置时，会造成kafka消费组假死（不消费消息，但可接收消息）
		timeout := handlerThis.SaramaConfig.Consumer.Group.Session.Timeout
		if timeout <= 0 {
			timeout = 10 * time.Second
		}
		if timeout > time.Second {
			timeout = timeout - 500*time.Millisecond
		}
		ctx, cancel := context.WithTimeout(handlerThis.Ctx, timeout)
		defer cancel()
		ch := make(chan struct{}, 1)
		go func() {
			handlerThis.handle(ctx, msg)
			close(ch)
		}()
		select {
		case <-ch:
		case <-ctx.Done():
		}
		// session.MarkMessage(msg, ``) // 标记消息为已处理
		// session.Commit()             // 马上提交到kafka
	}
	return
}

func (handlerThis *GroupHandlerOfTemplate) handle(ctx context.Context, msg *sarama.ConsumerMessage) (err error) {
	defer func() { //防止panic导致消费者断开
		if rec := recover(); rec != nil {
			err = errors.New(`panic错误：` + gconv.String(rec))
		}
	}()
	// 业务处理
	return
}
