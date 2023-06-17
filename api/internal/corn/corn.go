package corn

import (
	daoLog "api/internal/dao/log"
	"api/internal/utils"
	"context"

	"github.com/gogf/gf/v2/os/gcron"
)

func InitCorn(ctx context.Context) {
	// 星期一的凌晨3点执行
	gcron.Add(ctx, `0 0 3 * * 1`, func(ctx context.Context) {
		LogRequestPartition(ctx)
	}, `LogRequestPartition`)
}

// 请求日志表每周新增分区
func LogRequestPartition(ctx context.Context) {
	utils.DbTablePartition(ctx, daoLog.Request.Group(), daoLog.Request.Table(), 7, 24*60*60, `createdAt`)
}
