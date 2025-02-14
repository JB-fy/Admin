package model

type Handler interface {
	SendCode(phone string, code string) (err error)
	SendSms(phoneArr []string, message string, paramOpt ...any) (err error)
}
