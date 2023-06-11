package main

import (
	_ "api/internal/packed"

	_ "api/internal/logic"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"

	"github.com/gogf/gf/v2/os/gctx"

	"api/internal/cmd"
)

func main() {
	cmd.MyGen.Run(gctx.New())
	//cmd.Http.Run(gctx.New())

	// cmd.Main.AddCommand(&cmd.MyGen)
	// cmd.Main.AddCommand(&cmd.Http)
	// cmd.Main.Run(gctx.New())
}
