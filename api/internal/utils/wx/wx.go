package wx

import (
	daoPlatform "api/internal/dao/platform"
	"context"
)

func NewWxGzhByPfCfg(ctx context.Context) *WxGzh {
	config, _ := daoPlatform.Config.Get(ctx, []string{`wxGzhHost`, `wxGzhAppId`, `wxGzhSecret`, `wxGzhToken`, `wxGzhEncodingAESKey`})
	return NewWxGzh(ctx, config.Map())
}
