package corn

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

func InitCorn(ctx context.Context) {
	columns := daoPlatform.Corn.Columns()
	cornList, _ := daoPlatform.Corn.ParseDbCtx(ctx).Where(columns.IsStop, 0).All()
	for _, corn := range cornList {
		code := corn[columns.CornCode].String()
		pattern := corn[columns.CornPattern].String()
		gcron.AddSingleton(ctx, pattern, JobList[code], code)
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
