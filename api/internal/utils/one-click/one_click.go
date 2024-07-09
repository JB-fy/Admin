package one_click

import (
	daoPlatform "api/internal/dao/platform"
	"context"
)

func NewOneClickOfWxByPfCfg(ctx context.Context) *OneClickOfWx {
	configTmp, _ := daoPlatform.Config.Get(ctx, []string{`oneClickOfWx`})
	config := configTmp[`oneClickOfWx`].Map()
	return NewOneClickOfWx(ctx, config)
}

func NewOneClickOfYidunByPfCfg(ctx context.Context, configOpt ...map[string]any) *OneClickOfYidun {
	configTmp, _ := daoPlatform.Config.Get(ctx, []string{`oneClickOfYidun`})
	config := configTmp[`oneClickOfYidun`].Map()
	return NewOneClickOfYidun(ctx, config)
}
