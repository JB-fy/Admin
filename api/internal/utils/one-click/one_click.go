package one_click

import (
	daoPlatform "api/internal/dao/platform"
	"context"
)

func NewOneClickOfWxByPfCfg(ctx context.Context) *OneClickOfWx {
	config, _ := daoPlatform.Config.Get(ctx, []string{`oneClickOfWxHost`, `oneClickOfWxAppId`, `oneClickOfWxSecret`})
	return NewOneClickOfWx(ctx, config.Map())
}

func NewOneClickOfYidunByPfCfg(ctx context.Context, configOpt ...map[string]any) *OneClickOfYidun {
	config, _ := daoPlatform.Config.Get(ctx, []string{`oneClickOfYidunSecretId`, `oneClickOfYidunSecretKey`, `oneClickOfYidunBusinessId`})
	return NewOneClickOfYidun(ctx, config.Map())
}
