package initialize

import (
	daoPlatform "api/internal/dao/platform"
	"context"

	"github.com/gogf/gf/v2/util/gutil"
)

// 初始化可能有顺序要求，故统一到这里执行初始化函数
func Entry(ctx context.Context) {
	initGenv(ctx)  // 环境变量设置
	initI18n(ctx)  // 多语言设置
	initGtime(ctx) // 时区设置

	initDb(ctx) // 数据库设置

	initGvalid(ctx) // 自定义校验规则注册

	initCron(ctx)  // 定时任务设置
	initTimer(ctx) // 定时器设置

	gutil.Dump(daoPlatform.Admin.CtxDaoModel(ctx).SetIdArr(1).HookUpdateOne(daoPlatform.AdminPrivacy.Columns().Password, `123456`).Update())

}
