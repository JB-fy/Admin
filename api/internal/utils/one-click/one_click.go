package one_click

import (
	daoPlatform "api/internal/dao/platform"
	"context"
)

func NewOneClickOfWxByPfCfg(ctx context.Context) *OneClickOfWx {
	config, _ := daoPlatform.Config.CtxDaoModel(ctx).Filter(daoPlatform.Config.Columns().ConfigKey, `oneClickOfWx`).Value(daoPlatform.Config.Columns().ConfigValue)
	return NewOneClickOfWx(ctx, config.Map())
}

func NewOneClickOfYidunByPfCfg(ctx context.Context, configOpt ...map[string]any) *OneClickOfYidun {
	config, _ := daoPlatform.Config.CtxDaoModel(ctx).Filter(daoPlatform.Config.Columns().ConfigKey, `oneClickOfYidun`).Value(daoPlatform.Config.Columns().ConfigValue)
	return NewOneClickOfYidun(ctx, config.Map())
}
