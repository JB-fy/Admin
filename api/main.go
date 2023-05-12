package main

import (
	_ "api/internal/packed"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	"github.com/gogf/gf/v2/os/gctx"

	"api/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
