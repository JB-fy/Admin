package initialize

import (
	"api/internal/consts"
	"api/internal/utils"
	"api/internal/utils/kafka/admin"
	"api/internal/utils/kafka/consumer"
	"api/internal/utils/kafka/model"
	"api/internal/utils/kafka/producer"
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/genv"
	"github.com/gogf/gf/v2/util/gconv"
)

func initKafka(ctx context.Context) {
	for group, config := range g.Cfg().MustGet(ctx, `kafka`).Map() {
		config := model.GetConfig(group, gconv.Map(config))
		if utils.IsDev(ctx) || g.Cfg().MustGet(ctx, `masterServerNetworkIpArr.0`).String() == genv.Get(consts.ENV_SERVER_NETWORK_IP).String() {
			admin.Add(ctx, config)
		}
		producer.Add(ctx, config)
		consumer.Add(ctx, config)
	}
}
