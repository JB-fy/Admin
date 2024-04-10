package initialize

import (
	"github.com/gogf/gf/v2/os/gctx"
)

// 初始化可能有顺序要求，故统一到这里执行初始化函数
func Index() {
	ctx := gctx.New()

	initGenv(ctx)  // 环境变量设置
	initI18n(ctx)  // 多语言设置
	initGtime(ctx) // 时区设置

	initDb(ctx) // 数据库设置

	initGvalid(ctx) // 自定义校验规则注册

	initCron(ctx) // 定时器设置
}
