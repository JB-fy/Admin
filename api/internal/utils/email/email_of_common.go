package email

import (
	daoPlatform "api/internal/dao/platform"
	"context"
	"crypto/tls"
	"errors"
	"net/smtp"

	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

type EmailOfCommon struct {
	Ctx          context.Context
	SmtpHost     string `json:"emailOfCommonSmtpHost"`
	SmtpPort     string `json:"emailOfCommonSmtpPort"`
	FromEmail    string `json:"emailOfCommonFromEmail"`
	Password     string `json:"emailOfCommonPassword"` //QQ邮箱需注意：填QQ邮箱的授权码，而不是密码
	CodeSubject  string `json:"emailCodeSubject"`
	CodeTemplate string `json:"emailCodeTemplate"`
}

func NewEmailOfCommon(ctx context.Context, configOpt ...map[string]any) *EmailOfCommon {
	var config map[string]any
	if len(configOpt) > 0 && len(configOpt[0]) > 0 {
		config = configOpt[0]
	} else {
		configTmp, _ := daoPlatform.Config.Get(ctx, []string{`emailOfCommonSmtpHost`, `emailOfCommonSmtpPort`, `emailOfCommonFromEmail`, `emailOfCommonPassword`, `emailCodeSubject`, `emailCodeTemplate`})
		config = configTmp.Map()
	}

	emailObj := EmailOfCommon{Ctx: ctx}
	gconv.Struct(config, &emailObj)
	if emailObj.SmtpHost == `` || emailObj.SmtpPort == `` || emailObj.FromEmail == `` || emailObj.Password == `` {
		panic(`缺少插件配置：邮箱-通用`)
	}
	/* if emailObj.CodeSubject == `` || emailObj.CodeTemplate == `` {
		panic(`缺少插件配置：邮箱-验证码模板`)
	} */
	return &emailObj
}

func (emailThis *EmailOfCommon) SendCode(toEmail string, code string) (err error) {
	if emailThis.CodeSubject != `` || emailThis.CodeTemplate == `` {
		err = errors.New(`缺少插件配置：邮箱-验证码模板`)
		return
	}
	messageArr := []string{
		`From: ` + emailThis.FromEmail,
		`To: ` + toEmail,
		`Subject: ` + emailThis.CodeSubject,
		gstr.Replace(emailThis.CodeTemplate, `{code}`, code),
	}
	err = emailThis.SendEmail([]string{toEmail}, gstr.Join(messageArr, "\r\n"))
	return
}

func (emailThis *EmailOfCommon) SendEmail(toEmailArr []string, message string) (err error) {
	// 设置TLS配置
	tlsConfig := &tls.Config{ServerName: emailThis.SmtpHost}
	// tlsConfig.InsecureSkipVerify = true // 在开发或测试环境中可以设置为true，但在生产环境中应该验证服务器证书

	// 连接到SMTP服务器
	conn, err := tls.Dial(`tcp`, emailThis.SmtpHost+`:`+emailThis.SmtpPort, tlsConfig)
	if err != nil {
		return
	}
	defer conn.Close()

	// 创建SMTP客户端
	client, err := smtp.NewClient(conn, emailThis.SmtpHost)
	if err != nil {
		return
	}
	defer client.Quit()

	// 设置SMTP的认证信息
	auth := smtp.PlainAuth(``, emailThis.FromEmail, emailThis.Password, emailThis.SmtpHost)
	err = client.Auth(auth)
	if err != nil {
		return
	}

	// 发送邮件
	if err = client.Mail(emailThis.FromEmail); err != nil {
		return
	}
	for _, toEmail := range toEmailArr {
		if err = client.Rcpt(toEmail); err != nil {
			return
		}
	}
	w, err := client.Data()
	if err != nil {
		return
	}
	_, err = w.Write([]byte(message))
	if err != nil {
		return
	}
	err = w.Close()
	return
}
