package common

import (
	"api/internal/utils/email/model"
	"context"
	"crypto/tls"
	"net/smtp"

	"github.com/gogf/gf/v2/util/gconv"
)

type Email struct {
	SmtpHost  string `json:"smtpHost"`
	SmtpPort  string `json:"smtpPort"`
	FromEmail string `json:"fromEmail"`
	Password  string `json:"password"`
	// client    *smtp.Client //非并发安全的客户端，不能在初始化时生成
	// clientMu  sync.Mutex   //互斥锁。确实需要复用客户端时，必须在发送时上锁
}

func NewEmail(ctx context.Context, config map[string]any) model.Email {
	emailObj := &Email{}
	gconv.Struct(config, emailObj)
	if emailObj.SmtpHost == `` || emailObj.SmtpPort == `` || emailObj.FromEmail == `` || emailObj.Password == `` {
		panic(`缺少插件配置：邮箱-通用`)
	}
	/* // 设置TLS配置
	tlsConfig := &tls.Config{ServerName: emailObj.SmtpHost}
	// tlsConfig.InsecureSkipVerify = true // 在开发环境中可以设置为true，但在生产环境中应该验证服务器证书

	// 连接到SMTP服务器
	conn, err := tls.Dial(`tcp`, emailObj.SmtpHost+`:`+emailObj.SmtpPort, tlsConfig)
	if err != nil {
		panic(`连接到SMTP服务器错误：` + err.Error())
	}
	// defer conn.Close()

	// 创建SMTP客户端
	client, err := smtp.NewClient(conn, emailObj.SmtpHost)
	if err != nil {
		panic(`创建SMTP客户端错误：` + err.Error())
	}
	// defer client.Quit()

	// 设置SMTP的认证信息
	auth := smtp.PlainAuth(``, emailObj.FromEmail, emailObj.Password, emailObj.SmtpHost)
	err = client.Auth(auth)
	if err != nil {
		panic(`SMTP的认证信息错误：` + err.Error())
	}

	// 发送邮件
	err = client.Mail(emailObj.FromEmail)
	if err != nil {
		panic(`设置邮件发送人错误：` + err.Error())
	}
	emailObj.client = client */
	return emailObj
}

func (emailThis *Email) GetFromEmail() string {
	return emailThis.FromEmail
}

func (emailThis *Email) SendEmail(ctx context.Context, message string, toEmailArr ...string) (err error) { // 发送邮件
	/* // 复用客户端时需上锁
	emailThis.clientMu.Lock()
	defer emailThis.clientMu.Unlock()

	for _, toEmail := range toEmailArr {
		if err = emailThis.client.Rcpt(toEmail); err != nil {
			return
		}
	}
	w, err := emailThis.client.Data()
	if err != nil {
		return
	} */

	// 设置TLS配置
	tlsConfig := &tls.Config{ServerName: emailThis.SmtpHost}
	// tlsConfig.InsecureSkipVerify = true // 在开发环境中可以设置为true，但在生产环境中应该验证服务器证书

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

	//设置邮件发送人
	if err = client.Mail(emailThis.FromEmail); err != nil {
		return
	}
	// 发送邮件
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
