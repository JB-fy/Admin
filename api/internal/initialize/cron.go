package initialize

// 定时器设置
import (
	"context"
)

func InitCron(ctx context.Context) {
	// myCronThis := myCron{}

	// myCronThis.LogHttpPartition(ctx) //先执行一次请求日志分区
	// gcron.AddSingleton(ctx, `0 0 3 * * 1`, myCronThis.LogHttpPartition, `LogHttpPartition`)
}

type myCron struct{}

// Http日志表每周新增分区
func (myCron) LogHttpPartition(ctx context.Context) {
	//utils.DbTablePartition(ctx, daoLog.Http.Group(), daoLog.Http.Table(), 7, 24*60*60, daoLog.Http.Columns().CreatedAt)
}
