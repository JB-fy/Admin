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
	Ctx       context.Context
	SmtpHost  string `json:"emailOfCommonSmtpHost"`
	SmtpPort  string `json:"emailOfCommonSmtpPort"`
	FromEmail string `json:"emailOfCommonFromEmail"`
	Password  string `json:"emailOfCommonPassword"` //注意：这里是QQ的授权码，不是密码
}

func NewEmailOfCommon(ctx context.Context, configOpt ...map[string]any) *EmailOfCommon {
	var config map[string]any
	if len(configOpt) > 0 && len(configOpt[0]) > 0 {
		config = configOpt[0]
	} else {
		configTmp, _ := daoPlatform.Config.Get(ctx, []string{`emailOfCommonSmtpHost`, `emailOfCommonSmtpPort`, `emailOfCommonFromEmail`, `emailOfCommonPassword`})
		config = configTmp.Map()
	}
	config = g.Map{
		`emailOfCommonSmtpHost`:  `smtp.qq.com`,
		`emailOfCommonSmtpPort`:  `465`,
		`emailOfCommonFromEmail`: `274456806@qq.com`,
		`emailOfCommonPassword`:  `nsiuuffaemvpbjei`,
	}

	emailOfCommonObj := EmailOfCommon{Ctx: ctx}
	gconv.Struct(config, &emailOfCommonObj)
	return &emailOfCommonObj
}

func (emailThis *EmailOfCommon) SendCode(toEmail string, code string) (err error) {
	message := `From: ` + emailThis.FromEmail + "\r\n" +
		`To: ` + toEmail + "\r\n" +
		`Subject: 您的邮箱验证码` + "\r\n\r\n" +
		`验证码：` + code + `
说明：
1. 请在验证码输入框中输入上面的验证码，以完成您的邮箱验证。
2. 验证码在发送后的5分钟内有效。如果验证码过期，请重新请求一个新的验证码。
3. 出于安全考虑，请不要将此验证码分享给任何人。` + "\r\n"
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
