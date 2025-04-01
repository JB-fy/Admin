package model

type Handler interface {
	Create(loginId string, extData map[string]any) (token string, err error)
	Parse(token string) (tokenInfo TokenInfo, err error)
}
