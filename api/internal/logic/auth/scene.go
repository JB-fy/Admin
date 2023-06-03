package logic

import (
	"api/internal/service"
	"context"
	"fmt"
)

type sScene struct{}

func NewScene() *sScene {
	return &sScene{}
}

func init() {
	service.RegisterScene(NewScene())
}

func (logic *sScene) List(ctx context.Context) {
	fmt.Println("Scene1")
}
