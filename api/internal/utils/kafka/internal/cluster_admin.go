package internal

import (
	"api/internal/utils/kafka/model"
	"context"

	"github.com/IBM/sarama"
)

func InitClusterAdmin(ctx context.Context, adminConfig *model.AdminConfig) (clusterAdmin sarama.ClusterAdmin, err error) {
	clusterAdmin, err = sarama.NewClusterAdmin(adminConfig.Hosts, adminConfig.SaramaConfig)
	// defer clusterAdmin.Close()
	return
}
