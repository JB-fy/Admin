package model

type Handler interface {
	Sts() (stsInfo map[string]any, err error)
}
