package email

import (
	daoPlatform "api/internal/dao/platform"
	"context"
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
		configTmp, _ := daoPlatform.Config.Get(ctx, []string{`emailCodeSubject`, `emailCodeTemplate`, `emailOfCommon`})
		config := configTmp[`emailOfCommon`].Map()
		config[`codeSubject`] = configTmp[`emailCodeSubject`]
		config[`codeTemplate`] = configTmp[`emailCodeTemplate`]
		return NewEmailOfCommon(ctx, config)
	}
}
