package main

import (
	_ "api/internal/packed"

	_ "api/internal/logic"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"

	initialize "api/internal/initialize"

	"api/internal/cmd"

	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	initialize.InitMy()
	cmd.Main.AddCommand(&cmd.MyGen)
	cmd.Main.AddCommand(&cmd.Http)
	cmd.Main.Run(gctx.New())
}
