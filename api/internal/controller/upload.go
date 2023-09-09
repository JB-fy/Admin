package controller

import (
	"api/api"
	"api/internal/utils"
	"api/internal/utils/upload"
	"context"

	"github.com/gogf/gf/v2/frame/g"
)

type Upload struct{}

func NewUpload() *Upload {
	return &Upload{}
}

// 获取签名(web端直传用)
func (controllerThis *Upload) Sign(ctx context.Context, req *api.UploadSignReq) (res *api.UploadSignRes, err error) {
	signInfo, err := upload.NewUpload(ctx).Sign(ctx, req.Type)
	if err != nil {
		return
	}
	utils.HttpWriteJson(ctx, signInfo, 0, ``)
	return
}

// 获取Sts Token(App端直传用)
func (controllerThis *Upload) Sts(ctx context.Context, req *api.UploadStsReq) (res *api.UploadStsRes, err error) {
	stsInfo, err := upload.NewUpload(ctx).Sts(ctx, req.Type)
	if err != nil {
		return
	}
	request := g.RequestFromCtx(ctx)
	if request.URL.Path == `/upload/sts` {
		request.Response.WriteJson(stsInfo)
		return
	}
	utils.HttpWriteJson(ctx, stsInfo, 0, ``)
	return
}

// 回调
func (controllerThis *Upload) Notify(ctx context.Context, req *api.UploadNotifyReq) (res *api.UploadNotifyRes, err error) {
	notifyInfo, err := upload.NewUpload(ctx).Notify(ctx)
	if err != nil {
		return
	}
	utils.HttpWriteJson(ctx, notifyInfo, 0, ``)
	return
}

// 本地上传
func (controllerThis *Upload) Upload(ctx context.Context, req *api.UploadUploadReq) (res *api.UploadNotifyRes, err error) {
	notifyInfo, err := upload.NewUpload(ctx).Notify(ctx)
	if err != nil {
		return
	}
	utils.HttpWriteJson(ctx, notifyInfo, 0, ``)
	return
}
