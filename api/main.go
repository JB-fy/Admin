package main

import (
	_ "api/internal/packed"

	_ "api/internal/logic"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"

	// _ "github.com/gogf/gf/contrib/nosql/redis/v2"	// 不再使用框架的redis（缺少dialer，无法远程连接redis集群）

	"api/internal/cmd"
	"api/internal/initialize"

	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	ctx := gctx.GetInitCtx()

	initialize.Entry(ctx)

	cmd.Main.AddCommand(&cmd.MyGen)
	cmd.Main.AddCommand(&cmd.Http)
	cmd.Main.Run(ctx)
}
