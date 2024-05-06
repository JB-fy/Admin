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
func (controllerThis *Upload) Upload(ctx context.Context, req *api.UploadUploadReq) (res *api.UploadNotifyRes, err error) {
	notifyInfo, err := upload.NewUpload(ctx).Upload(g.RequestFromCtx(ctx))
	if err != nil {
		return
	}
	res = &api.UploadNotifyRes{
		Url:      notifyInfo.Url,
		Width:    notifyInfo.Width,
		Height:   notifyInfo.Height,
		Size:     notifyInfo.Size,
		MimeType: notifyInfo.MimeType,
	}
	return
}

// 获取签名（H5直传用）
func (controllerThis *Upload) Sign(ctx context.Context, req *api.UploadSignReq) (res *api.UploadSignRes, err error) {
	signInfo, err := upload.NewUpload(ctx).Sign(upload.CreateUploadParam(req.FileType))
	if err != nil {
		return
	}
	res = &api.UploadSignRes{
		UploadUrl:  signInfo.UploadUrl,
		UploadData: signInfo.UploadData,
		Host:       signInfo.Host,
		Dir:        signInfo.Dir,
		Expire:     signInfo.Expire,
		IsRes:      signInfo.IsRes,
	}
	return
}

// 获取配置信息（APP直传前调用）
func (controllerThis *Upload) Config(ctx context.Context, req *api.UploadConfigReq) (res *api.UploadConfigRes, err error) {
	config, err := upload.NewUpload(ctx).Config(upload.CreateUploadParam(req.FileType))
	if err != nil {
		return
	}
	utils.HttpWriteJson(ctx, config, 0, ``)
	return
}

// 获取Sts Token（APP直传用）
func (controllerThis *Upload) Sts(ctx context.Context, req *api.UploadStsReq) (res *api.UploadStsRes, err error) {
	stsInfo, err := upload.NewUpload(ctx).Sts(upload.CreateUploadParam(req.FileType))
	if err != nil {
		return
	}
	g.RequestFromCtx(ctx).Response.WriteJson(stsInfo)
	return
}

// 回调
func (controllerThis *Upload) Notify(ctx context.Context, req *api.UploadNotifyReq) (res *api.UploadNotifyRes, err error) {
	notifyInfo, err := upload.NewUpload(ctx, req.UploadType).Notify(g.RequestFromCtx(ctx))
	if err != nil {
		return
	}
	res = &api.UploadNotifyRes{
		Url:      notifyInfo.Url,
		Width:    notifyInfo.Width,
		Height:   notifyInfo.Height,
		Size:     notifyInfo.Size,
		MimeType: notifyInfo.MimeType,
	}
	return
}
