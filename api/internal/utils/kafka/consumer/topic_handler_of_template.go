package consumer

import (
	"api/internal/utils/kafka/model"
	"context"

	"github.com/IBM/sarama"
)

func TopicHandlerOfTemplate(ctx context.Context, config *model.Config, consumerInfo *model.ConsumerInfo) func(msg *sarama.ConsumerMessage) {
	return func(msg *sarama.ConsumerMessage) {
		// 对应主题的处理逻辑
	}
}
