package email

import (
	daoPlatform "api/internal/dao/platform"
	"context"
	"errors"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

type Handler struct {
	Ctx   context.Context
	Email Email
}

func NewHandler(ctx context.Context, emailTypeOpt ...string) *Handler {
	handlerObj := &Handler{Ctx: ctx}

	emailType := ``
	if len(emailTypeOpt) > 0 {
		emailType = emailTypeOpt[0]
	} else {
		emailType, _ = daoPlatform.Config.CtxDaoModel(ctx).Filter(daoPlatform.Config.Columns().ConfigKey, `emailType`).ValueStr(daoPlatform.Config.Columns().ConfigValue)
	}

	var config g.Map
	switch emailType {
	// case `emailOfCommon`:
	default:
		config, _ = daoPlatform.Config.CtxDaoModel(ctx).Filter(daoPlatform.Config.Columns().ConfigKey, `emailOfCommon`).ValueMap(daoPlatform.Config.Columns().ConfigValue)
	}

	config[`emailType`] = emailType
	handlerObj.Email = NewEmail(config)
	return handlerObj
}

type CodeTemplate struct {
	Subject  string `json:"subject"`
	Template string `json:"template"`
}

func (handlerThis *Handler) SendCode(toEmail string, code string) (err error) {
	config, _ := daoPlatform.Config.Get(handlerThis.Ctx, `emailCode`)
	codeTemplate := &CodeTemplate{}
	gconv.Struct(config[`emailCode`].Map(), codeTemplate)
	if codeTemplate.Subject == `` || codeTemplate.Template == `` {
		err = errors.New(`缺少配置：邮箱-验证码模板`)
		return
	}

	messageArr := []string{
		`From: ` + handlerThis.Email.GetFromEmail(),
		`To: ` + toEmail,
		`Subject: ` + codeTemplate.Subject,
		gstr.Replace(codeTemplate.Template, `{code}`, code),
	}
	return handlerThis.SendEmail(gstr.Join(messageArr, "\r\n"), toEmail)
}

func (handlerThis *Handler) SendEmail(message string, toEmailArr ...string) (err error) {
	return handlerThis.Email.SendEmail(handlerThis.Ctx, message, toEmailArr...)
}
