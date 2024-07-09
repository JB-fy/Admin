package email

import (
	"context"
	"crypto/tls"
	"errors"
	"net/smtp"

	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

type EmailOfCommon struct {
	Ctx       context.Context
	SmtpHost  string `json:"smtpHost"`
	SmtpPort  string `json:"smtpPort"`
	FromEmail string `json:"fromEmail"`
	Password  string `json:"password"`
	Code      struct {
		Subject  string `json:"subject"`
		Template string `json:"template"`
	} `json:"code"`
}

func NewEmailOfCommon(ctx context.Context, config map[string]any) *EmailOfCommon {
	emailObj := EmailOfCommon{Ctx: ctx}
	gconv.Struct(config, &emailObj)
	/* if emailObj.Code.Subject == `` || emailObj.Code.Template == `` {
		panic(`缺少插件配置：邮箱-验证码模板`)
	} */
	if emailObj.SmtpHost == `` || emailObj.SmtpPort == `` || emailObj.FromEmail == `` || emailObj.Password == `` {
		panic(`缺少插件配置：邮箱-通用`)
	}
	return &emailObj
}

func (emailThis *EmailOfCommon) SendCode(toEmail string, code string) (err error) {
	if emailThis.Code.Subject == `` || emailThis.Code.Template == `` {
		err = errors.New(`缺少插件配置：邮箱-验证码模板`)
		return
	}
	messageArr := []string{
		`From: ` + emailThis.FromEmail,
		`To: ` + toEmail,
		`Subject: ` + emailThis.Code.Subject,
		gstr.Replace(emailThis.Code.Template, `{code}`, code),
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
