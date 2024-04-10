package one_click

import (
	daoPlatform "api/internal/dao/platform"
	"context"
)

type OneClick interface {
	Check(token string, accessToken string) (phone string, err error)
}

func NewOneClick(ctx context.Context, oneClickTypeOpt ...string) OneClick {
	oneClickType := ``
	if len(oneClickTypeOpt) > 0 {
		oneClickType = oneClickTypeOpt[0]
	} else {
		oneClickTypeVar, _ := daoPlatform.Config.CtxDaoModel(ctx).Filter(daoPlatform.Config.Columns().ConfigKey, `oneClickType`).Value(daoPlatform.Config.Columns().ConfigValue)
		oneClickType = oneClickTypeVar.String()
	}

	switch oneClickType {
	// case `oneClickOfYidun`:
	default:
		config, _ := daoPlatform.Config.Get(ctx, []string{`oneClickOfYidunSecretId`, `oneClickOfYidunSecretKey`, `oneClickOfYidunBusinessId`})
		return NewOneClickOfYidun(ctx, config.Map())
	}
}
