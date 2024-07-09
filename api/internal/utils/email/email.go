package email

import (
	daoPlatform "api/internal/dao/platform"
	"context"

	"github.com/gogf/gf/v2/container/gmap"
)

type Email interface {
	SendCode(toEmail string, code string) (err error)
	SendEmail(toEmailArr []string, message string) (err error)
}

func NewEmail(ctx context.Context, emailTypeOpt ...string) Email {
	emailType := ``
	if len(emailTypeOpt) > 0 {
		emailType = emailTypeOpt[0]
	} else {
		emailType, _ = daoPlatform.Config.CtxDaoModel(ctx).Filter(daoPlatform.Config.Columns().ConfigKey, `emailType`).ValueStr(daoPlatform.Config.Columns().ConfigValue)
	}

	switch emailType {
	// case `emailOfCommon`:
	default:
		configTmp, _ := daoPlatform.Config.Get(ctx, []string{`emailCode`, `emailOfCommon`})
		config := gmap.NewStrAnyMapFrom(configTmp[`emailCode`].Map())
		config.Merge(gmap.NewStrAnyMapFrom(configTmp[`emailOfCommon`].Map()))
		return NewEmailOfCommon(ctx, config.Map())
	}
}
