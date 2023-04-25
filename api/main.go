package main

import (
	_ "api/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"api/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
