package model

import (
	"time"

	"github.com/IBM/sarama"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/xdg-go/scram"
)

type Config struct {
	CommonConfig
	TopicList []struct {
		Name    string `json:"name"`
		PartNum int32  `json:"partNum"`
		ReplNum int16  `json:"replNum"`
	} `json:"topicList"`
	ProducerList []ProducerConfig `json:"producerList"`
	ConsumerList []ConsumerConfig `json:"consumerList"`
}

type CommonConfig struct {
	Group    string
	Hosts    []string `json:"hosts"`
	UserName string   `json:"userName"`
	Password string   `json:"password"`
	SaslType string   `json:"saslType"`
}

type AdminConfig struct {
	*CommonConfig
	SaramaConfig *sarama.Config
}

type ProducerConfig struct {
	*CommonConfig
	SaramaConfig *sarama.Config
	Topic        string `json:"topic"`
	IsSync       bool   `json:"isSync"`
	Idempotent   bool   `json:"idempotent"`
}

type ConsumerConfig struct {
	*CommonConfig
	SaramaConfig      *sarama.Config
	GroupId           string        `json:"groupId"`
	TopicArr          []string      `json:"topicArr"`
	Number            uint          `json:"number"`
	AutoCommit        *bool         `json:"autoCommit"`
	MaxWaitTime       time.Duration `json:"maxWaitTime"`
	MaxProcessingTime time.Duration `json:"maxProcessingTime"`
	SessionTimeout    time.Duration `json:"sessionTimeout"`
	HeartbeatInterval time.Duration `json:"heartbeatInterval"`
}

func GetConfig(group string, configMap map[string]any) (config *Config) {
	config = &Config{}
	gconv.Struct(configMap, config)
	config.Group = group
	return
}

func CreateAdminConfig(config *Config) (saramaConfig *sarama.Config) {
	saramaConfig = createSaramaConfig(config)
	return
}

func CreateProducerConfig(config *Config, producerConfig *ProducerConfig) (saramaConfig *sarama.Config) {
	saramaConfig = createSaramaConfig(config)
	if producerConfig.IsSync { //同步生成者必须设置
		saramaConfig.Producer.Return.Successes = true
	}
	saramaConfig.Producer.Idempotent = producerConfig.Idempotent
	if saramaConfig.Producer.Idempotent { //保证消息唯一时必须设置
		saramaConfig.Producer.RequiredAcks = sarama.WaitForAll
		saramaConfig.Net.MaxOpenRequests = 1
	}
	return
}

func CreateConsumerConfig(config *Config, ConsumerConfig *ConsumerConfig) (saramaConfig *sarama.Config) {
	saramaConfig = createSaramaConfig(config)
	saramaConfig.Consumer.Return.Errors = true
	saramaConfig.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategySticky() /* , sarama.NewBalanceStrategyRoundRobin(), sarama.NewBalanceStrategyRange() */}
	if ConsumerConfig.AutoCommit != nil {
		saramaConfig.Consumer.Offsets.AutoCommit.Enable = *ConsumerConfig.AutoCommit
	}
	if ConsumerConfig.MaxWaitTime > 0 {
		saramaConfig.Consumer.MaxWaitTime = ConsumerConfig.MaxWaitTime // 多久拉取一次消息
	}
	if ConsumerConfig.MaxProcessingTime > 0 {
		saramaConfig.Consumer.MaxProcessingTime = ConsumerConfig.MaxProcessingTime // 单次消息处理的最大时间
	}
	if ConsumerConfig.SessionTimeout > 0 {
		saramaConfig.Consumer.Group.Session.Timeout = ConsumerConfig.SessionTimeout
	}
	if ConsumerConfig.HeartbeatInterval > 0 {
		saramaConfig.Consumer.Group.Heartbeat.Interval = ConsumerConfig.HeartbeatInterval
	}
	return
}

func createSaramaConfig(config *Config) (saramaConfig *sarama.Config) {
	saramaConfig = sarama.NewConfig()
	saramaConfig.Version = sarama.V4_0_0_0
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
