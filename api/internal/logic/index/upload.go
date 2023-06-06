package logic

import (
	"api/internal/service"
	"context"
)

type sUpload struct{}

func NewUpload() *sUpload {
	return &sUpload{}
}

func init() {
	service.RegisterUpload(NewUpload())
}

// 签名
func (logicThis *sUpload) Sign(ctx context.Context, sceneCode string, account string) (encryptStr string, err error) {

	return
}

// 回调
func (logicThis *sUpload) Notify(ctx context.Context, sceneCode string, account string, password string) (token string, err error) {

	return
}
