package cron

import (
	daoLog "api/internal/dao/log"
	daoPlatform "api/internal/dao/platform"
	"api/internal/utils"
	"context"
	"errors"
	"fmt"

	"github.com/gogf/gf/v2/os/gcron"
)

func InitCron(ctx context.Context) {
	columns := daoPlatform.Cron.Columns()
	cronList, _ := daoPlatform.Cron.ParseDbCtx(ctx).Where(columns.IsStop, 0).All()
	for _, cron := range cronList {
		CronStart(ctx, cron[columns.CronPattern].String(), cron[columns.CronCode].String())
	}
}

// 开启定时任务
func CronStart(ctx context.Context, pattern string, code string) (err error) {
	var JobList = map[string]func(ctx context.Context){
		`LogHttpPartition`: LogHttpPartition,
		`Test`:             Test,
	}
	_, ok := JobList[code]
	if !ok {
		err = errors.New(`定时器方法不存在`)
		return
	}
	gcron.AddSingleton(ctx, pattern, JobList[code], code)
	return
}

// Http日志表每周新增分区
func LogHttpPartition(ctx context.Context) {
	utils.DbTablePartition(ctx, daoLog.Http.Group(), daoLog.Http.Table(), 7, 24*60*60, daoLog.Http.Columns().CreatedAt)
}

// 测试
func Test(ctx context.Context) {
	fmt.Println(1111)
}
