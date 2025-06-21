package model

import (
	"github.com/gogf/gf/v2/util/gconv"
)

type Config struct {
	Group    string
	Hosts    []string `json:"hosts"`
	UserName string   `json:"userName"`
	Password string   `json:"password"`
	Keyspace string   `json:"keyspace"`
	DcName   string   `json:"dcName"`
	DcList   []struct {
		DcName  string `json:"dcName"`
		ReplNum uint8  `json:"replNum"`
	} `json:"dcList"`
	ProtoVersion *int    `json:"protoVersion"`
	Consistency  *uint16 `json:"consistency"`
	NumConns     *int    `json:"mumConns"`
	Debug        bool    `json:"debug"`
}

func GetConfig(group string, configMap map[string]any) (config *Config) {
	config = &Config{Group: group}
	gconv.Struct(configMap, config)
	return
}
