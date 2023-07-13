package initialize

// 多语言设置
import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
)

func InitI18n(ctx context.Context) {
	g.I18n().SetPath(g.Cfg().MustGet(ctx, `i18n.path`).String())         //设置资源目录
	g.I18n().SetLanguage(g.Cfg().MustGet(ctx, `i18n.language`).String()) //设置默认为中文（原默认为英文en）
}
