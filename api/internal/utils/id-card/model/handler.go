package model

type Handler interface {
	Auth(idCardName string, idCardNo string) (idCardInfo IdCardInfo, err error)
}
