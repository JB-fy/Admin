package main

import (
	_ "demo/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"demo/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
