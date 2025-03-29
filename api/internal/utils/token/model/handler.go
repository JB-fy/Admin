package model

type Handler interface {
	Create(loginId string, opt ...map[string]any) (token string, err error)
	Parse(token string) (tokenInfo TokenInfo, err error)
}
