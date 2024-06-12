package email

import (
	daoPlatform "api/internal/dao/platform"
	"context"
	"crypto/tls"
	"net/smtp"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type EmailOfCommon struct {
	Ctx      context.Context
	SmtpHost string `json:"emailOfCommonSmtpHost"`
	SmtpPort string `json:"emailOfCommonSmtpPort"`
	Email    string `json:"emailOfCommonEmail"`
	Password string `json:"emailOfCommonPassword"` //注意：这里是QQ的授权码，不是密码
}

func NewEmailOfCommon(ctx context.Context, configOpt ...map[string]any) *EmailOfCommon {
	var config map[string]any
	if len(configOpt) > 0 && len(configOpt[0]) > 0 {
		config = configOpt[0]
	} else {
		configTmp, _ := daoPlatform.Config.Get(ctx, []string{`emailOfCommonSmtpHost`, `emailOfCommonSmtpPort`, `emailOfCommonEmail`, `emailOfCommonPassword`})
		config = configTmp.Map()
	}
	config = g.Map{
		`emailOfCommonSmtpHost`: `smtp.qq.com`,
		`emailOfCommonSmtpPort`: `465`,
		`emailOfCommonEmail`:    `274456806@qq.com`,
		`emailOfCommonPassword`: `nsiuuffaemvpbjei`,
	}

	emailOfCommonObj := EmailOfCommon{Ctx: ctx}
	gconv.Struct(config, &emailOfCommonObj)
	return &emailOfCommonObj
}

func (emailThis *EmailOfCommon) SendCode(toEmail string, code string) (err error) {
	message := "To: " + toEmail + "\r\n" +
		"Subject: " + code + "\r\n" +
		"\r\n" +
		"This is the body of the email sent from QQ Mail using Go.\r\n"
	err = emailThis.SendEmail([]string{toEmail}, message)
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
	auth := smtp.PlainAuth(``, emailThis.Email, emailThis.Password, emailThis.SmtpHost)
	err = client.Auth(auth)
	if err != nil {
		return
	}

	// 发送邮件
	if err = client.Mail(emailThis.Email); err != nil {
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
	if err != nil {
		return
	}
	return
}
