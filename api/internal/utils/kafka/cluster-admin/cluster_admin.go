package cluster_admin

import (
	"api/internal/utils/kafka/internal"
	"api/internal/utils/kafka/model"
	"context"
	"fmt"

	"github.com/IBM/sarama"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	clusterAdminMap = map[string]sarama.ClusterAdmin{}
)

func Add(ctx context.Context, config *model.Config) {
	adminConfig := &model.AdminConfig{
		CommonConfig: &config.CommonConfig,
		SaramaConfig: model.CreateAdminConfig(config),
	}
	clusterAdmin, err := internal.InitClusterAdmin(ctx, adminConfig)
	if err != nil {
		panic(fmt.Errorf(`管理员(分组:%s)连接错误:%w`, adminConfig.Group, err))
	}
	clusterAdminMap[adminConfig.Group] = clusterAdmin
	g.Log(`kafka`).Info(ctx, fmt.Errorf(`管理员(分组:%s)连接成功`, adminConfig.Group))

	for _, topicInfo := range config.TopicList { // 创建主题
		if topicInfo.Name == `` {
			continue
		}
		err = clusterAdmin.CreateTopic(topicInfo.Name, &sarama.TopicDetail{NumPartitions: topicInfo.PartNum, ReplicationFactor: topicInfo.ReplNum}, false)
		if err == nil {
			continue
		}
		if topicErr, ok := err.(*sarama.TopicError); ok && topicErr.Err == sarama.ErrTopicAlreadyExists {
			topicMetaList, _ := clusterAdmin.DescribeTopics([]string{topicInfo.Name})
			if topicInfo.PartNum <= int32(len(topicMetaList[0].Partitions)) {
				continue
			}
			err = clusterAdmin.CreatePartitions(topicInfo.Name, topicInfo.PartNum, nil, false) // 增加分区数量
			if err == nil {
				continue
			}
		}
		panic(fmt.Errorf(`主题(分组:%s,主题:%s)创建错误:%w`, adminConfig.Group, topicInfo.Name, err))
	}
}
