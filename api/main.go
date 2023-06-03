package main

import (
	_ "api/internal/packed"
	"context"

	_ "api/internal/logic"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/v2/os/gctx"

	"api/internal/cmd"
)

func main() {
	g.I18n().SetPath(g.Cfg().MustGet(context.TODO(), "i18n.path").String())         //设置资源目录
	g.I18n().SetLanguage(g.Cfg().MustGet(context.TODO(), "i18n.language").String()) //设置默认为中文（原默认为英文en）
	cmd.Main.Run(gctx.New())
}
