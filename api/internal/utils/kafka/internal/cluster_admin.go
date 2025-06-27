package internal

import (
	"api/internal/utils/kafka/model"
	"context"

	"github.com/IBM/sarama"
)

func InitClusterAdmin(ctx context.Context, saramaConfig *sarama.Config, config *model.Config) (clusterAdmin sarama.ClusterAdmin, err error) {
	clusterAdmin, err = sarama.NewClusterAdmin(config.Hosts, saramaConfig)
	// defer clusterAdmin.Close()
	return
}
