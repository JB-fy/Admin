package model

type Handler interface {
	Push(param PushParam) (err error)
	Tag(param TagParam) (err error)
}
