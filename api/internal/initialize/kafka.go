package initialize

import (
	"api/internal/consts"
	"api/internal/utils"
	cluster_admin "api/internal/utils/kafka/cluster-admin"
	"api/internal/utils/kafka/consumer"
	"api/internal/utils/kafka/producer"
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/genv"
	"github.com/gogf/gf/v2/util/gconv"
)

func initKafka(ctx context.Context) {
	for group, config := range g.Cfg().MustGet(ctx, `kafka`).Map() {
		configMap := gconv.Map(config)
		if utils.IsDev(ctx) || g.Cfg().MustGet(ctx, `masterServerNetworkIp`).String() == genv.Get(consts.ENV_SERVER_NETWORK_IP).String() {
			cluster_admin.Add(ctx, group, configMap)
		}
		producer.Add(ctx, group, configMap)
		consumer.Add(ctx, group, configMap)
	}
}
