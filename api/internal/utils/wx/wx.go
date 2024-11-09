package wx

import (
	daoPlatform "api/internal/dao/platform"
	"context"
)

func NewWxGzhByPfCfg(ctx context.Context) *WxGzh {
	config, _ := daoPlatform.Config.CtxDaoModel(ctx).Filter(daoPlatform.Config.Columns().ConfigKey, `wxGzh`).ValueMap(daoPlatform.Config.Columns().ConfigValue)
	return NewWxGzh(ctx, config)
}
