package cluster_admin

import (
	"api/internal/utils/kafka/internal"
	"api/internal/utils/kafka/model"
	"context"

	"github.com/IBM/sarama"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	clusterAdminMap = map[string]sarama.ClusterAdmin{}
)

func Add(ctx context.Context, group string, configMap map[string]any) {
	config := model.GetConfig(group, configMap)
	clusterAdminConfig := model.CreateClusterAdmin(config)
	clusterAdmin, err := internal.InitClusterAdmin(ctx, clusterAdminConfig, config)
	if err != nil {
		panic(`管理员(分组:` + config.Group + `)连接失败：` + err.Error())
	}
	clusterAdminMap[config.Group] = clusterAdmin
	g.Log(`kafka`).Info(ctx, `管理员(分组:`+config.Group+`)连接成功`)

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
		panic(`主题(分组:` + config.Group + `，主题:` + topicInfo.Name + `)创建失败：` + err.Error())
	}
}
