package email

import (
	daoPlatform "api/internal/dao/platform"
	"api/internal/utils/email/model"
	"context"
	"errors"

	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

type Handler struct {
	Ctx   context.Context
	email model.Email
}

func NewHandler(ctx context.Context, emailTypeOpt ...string) model.Handler {
	handlerObj := &Handler{Ctx: ctx}
	emailType := ``
	if len(emailTypeOpt) > 0 {
		emailType = emailTypeOpt[0]
	} else {
		emailType = daoPlatform.Config.GetOne(ctx, `email_type`).String()
	}
	if _, ok := emailFuncMap[emailType]; !ok {
		emailType = emailTypeDef
	}
	config := daoPlatform.Config.GetOne(ctx, emailType).Map()
	handlerObj.email = NewEmail(ctx, emailType, config)
	return handlerObj
}

func (handlerThis *Handler) SendEmail(message string, toEmailArr ...string) (err error) {
	return handlerThis.email.SendEmail(handlerThis.Ctx, message, toEmailArr...)
}

func (handlerThis *Handler) SendCode(toEmail string, code string) (err error) {
	codeData := daoPlatform.Config.GetOne(handlerThis.Ctx, `email_code`).Map()
	subject := gconv.String(codeData[`subject`])
	template := gconv.String(codeData[`template`])
	if subject == `` || template == `` {
		err = errors.New(`缺少配置：邮箱-验证码模板`)
		return
	}

	messageArr := []string{
		`From: ` + handlerThis.email.GetFromEmail(),
		`To: ` + toEmail,
		`Subject: ` + subject,
		gstr.Replace(template, `{code}`, code),
	}
	return handlerThis.SendEmail(gstr.Join(messageArr, "\r\n"), toEmail)
}
