package main

import (
	_ "api/internal/packed"

	_ "api/internal/logic"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"

	"api/internal/cmd"
	initialize "api/internal/initialize"

	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	ctx := gctx.GetInitCtx()

	initialize.Entry(ctx)

	cmd.Main.AddCommand(&cmd.MyGen)
	cmd.Main.AddCommand(&cmd.Http)
	cmd.Main.Run(ctx)
}
