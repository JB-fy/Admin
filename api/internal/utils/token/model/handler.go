package model

type Handler interface {
	Create(loginId string) (token string, err error)
	Parse(token string) (tokenInfo TokenInfo, err error)
}
