package cron

import (
	daoLog "api/internal/dao/log"
	daoPlatform "api/internal/dao/platform"
	"api/internal/utils"
	"context"
	"fmt"

	"github.com/gogf/gf/v2/os/gcron"
)

var JobList = map[string]func(ctx context.Context){
	`LogHttpPartition`: LogHttpPartition,
	`Test`:             Test,
}

func InitCron(ctx context.Context) {
	columns := daoPlatform.Cron.Columns()
	cronList, _ := daoPlatform.Cron.ParseDbCtx(ctx).Where(columns.IsStop, 0).All()
	for _, cron := range cronList {
		code := cron[columns.CronCode].String()
		pattern := cron[columns.CronPattern].String()
		_, ok := JobList[code]
		if ok {
			gcron.AddSingleton(ctx, pattern, JobList[code], code)
		}
	}
}

// Http日志表每周新增分区
func LogHttpPartition(ctx context.Context) {
	utils.DbTablePartition(ctx, daoLog.Http.Group(), daoLog.Http.Table(), 7, 24*60*60, daoLog.Http.Columns().CreatedAt)
}

// 测试
func Test(ctx context.Context) {
	fmt.Println(1111)
}
