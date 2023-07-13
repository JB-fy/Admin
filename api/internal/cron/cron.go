package cron

import (
	"context"
)

func InitCron(ctx context.Context) {
	//LogHttpPartition(ctx) //先执行一次请求日志分区

	//gcron.AddSingleton(ctx, `0 0 3 * * 1`, LogHttpPartition, `LogHttpPartition`)
}

// Http日志表每周新增分区
func LogHttpPartition(ctx context.Context) {
	//utils.DbTablePartition(ctx, daoLog.Http.Group(), daoLog.Http.Table(), 7, 24*60*60, daoLog.Http.Columns().CreatedAt)
}
