package initialize

import (
	"github.com/gogf/gf/v2/os/gctx"
)

// 初始化可能有顺序要求，故统一到这里执行初始化函数
func init() {
	ctx := gctx.New()

	initOfGenv(ctx)   // 环境变量设置
	initOfCron(ctx)   // 定时器设置
	initOfGtime(ctx)  // 时区设置
	initOfGvalid(ctx) // 自定义校验规则注册
	initOfI18n(ctx)   // 多语言设置
}
