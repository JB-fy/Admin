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

// 本地上传
func (controllerThis *Upload) Upload(ctx context.Context, req *api.UploadUploadReq) (res *api.UploadUploadRes, err error) {
	notifyInfo, err := upload.NewUpload(ctx).Upload()
	if err != nil {
		return
	}
	utils.HttpWriteJson(ctx, notifyInfo, 0, ``)
	return
}

// 获取签名（H5直传用）
func (controllerThis *Upload) Sign(ctx context.Context, req *api.UploadSignReq) (res *api.UploadSignRes, err error) {
	signInfo, err := upload.NewUpload(ctx).Sign(upload.CreateUploadOption(req.UploadType))
	if err != nil {
		return
	}
	utils.HttpWriteJson(ctx, signInfo, 0, ``)
	return
}

// 获取配置信息（APP直传前调用，后期也可用在其它地方）
func (controllerThis *Upload) Config(ctx context.Context, req *api.UploadConfigReq) (res *api.UploadConfigRes, err error) {
	config, err := upload.NewUpload(ctx).Config(upload.CreateUploadOption(req.UploadType))
	if err != nil {
		return
	}
	utils.HttpWriteJson(ctx, config, 0, ``)
	return
}

// 获取Sts Token（APP直传用）
func (controllerThis *Upload) Sts(ctx context.Context, req *api.UploadStsReq) (res *api.UploadStsRes, err error) {
	stsInfo, err := upload.NewUpload(ctx).Sts(upload.CreateUploadOption(req.UploadType))
	if err != nil {
		return
	}
	g.RequestFromCtx(ctx).Response.WriteJson(stsInfo)
	return
}

// 回调
func (controllerThis *Upload) Notify(ctx context.Context, req *api.UploadNotifyReq) (res *api.UploadNotifyRes, err error) {
	notifyInfo, err := upload.NewUpload(ctx).Notify()
	if err != nil {
		return
	}
	utils.HttpWriteJson(ctx, notifyInfo, 0, ``)
	return
}
