package model

type Handler interface {
	SendEmail(message string, toEmailArr ...string) (err error)
	SendCode(toEmail string, code string) (err error)
}
