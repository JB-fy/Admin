package model

import (
	"github.com/IBM/sarama"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/xdg-go/scram"
)

type Config struct {
	Group        string
	Hosts        []string        `json:"hosts"`
	UserName     string          `json:"userName"`
	Password     string          `json:"password"`
	SaslType     string          `json:"saslType"`
	ProducerType string          `json:"producerType"`
	ConsumerList []*ConsumerInfo `json:"consumerList"`
}

type ConsumerInfo struct {
	GroupId  string   `json:"groupId"`
	Number   int      `json:"number"`
	TopicArr []string `json:"topicArr"`
}

func GetConfig(group string, configMap map[string]any) (config *Config) {
	config = &Config{Group: group}
	gconv.Struct(configMap, config)
	return
}

func CreateProducerConfig(config *Config) (saramaConfig *sarama.Config) {
	saramaConfig = createSaramaConfig(config)
	if config.ProducerType == `sync` { //同步生成者必须设置
		saramaConfig.Producer.Return.Successes = true
	}
	return
}

func CreateConsumerConfig(config *Config) (saramaConfig *sarama.Config) {
	saramaConfig = createSaramaConfig(config)
	saramaConfig.Consumer.Return.Errors = true
	return
}

func createSaramaConfig(config *Config) (saramaConfig *sarama.Config) {
	saramaConfig = sarama.NewConfig()
	if config.UserName != `` && config.Password != `` {
		saramaConfig.Net.SASL.Enable = true
		saramaConfig.Net.SASL.User = config.UserName
		saramaConfig.Net.SASL.Password = config.Password
		saramaConfig.Net.SASL.Mechanism = sarama.SASLMechanism(config.SaslType)
		switch saramaConfig.Net.SASL.Mechanism {
		case sarama.SASLTypeSCRAMSHA512:
			saramaConfig.Net.SASL.SCRAMClientGeneratorFunc = func() sarama.SCRAMClient {
				return &ScramClient{HashGeneratorFcn: scram.SHA512}
			}
		case sarama.SASLTypeSCRAMSHA256:
			saramaConfig.Net.SASL.SCRAMClientGeneratorFunc = func() sarama.SCRAMClient {
				return &ScramClient{HashGeneratorFcn: scram.SHA256}
			}
		}
	}
	return
}
