package upload

import (
	daoUpload "api/internal/dao/upload"
	"api/internal/utils"
	"api/internal/utils/upload/model"
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

type Handler struct {
	Ctx      context.Context
	Scene    string //上传场景。default默认。根据自身需求扩展，用于确定上传通道和上传参数
	UploadId uint   //上传ID
	upload   model.Upload
}

func NewHandler(ctx context.Context, scene string, uploadId uint) model.Handler {
	handlerObj := &Handler{
		Ctx:      ctx,
		Scene:    scene,
		UploadId: uploadId,
	}
	uploadInfo, _ := daoUpload.Upload.CacheGetInfo(handlerObj.Ctx, handlerObj.UploadId)
	config := uploadInfo[daoUpload.Upload.Columns().UploadConfig].Map()
	config[`upload_id`] = uploadInfo[daoUpload.Upload.Columns().UploadId]
	uploadType := uploadInfo[daoUpload.Upload.Columns().UploadType].Uint()
	if _, ok := uploadFuncMap[uploadType]; !ok {
		uploadType = uploadTypeDef
	}
	switch uploadType {
	case 0:
	case 1:
		if gconv.Bool(config[`is_notify`]) {
			config[`callback_url`] = utils.GetRequestUrl(handlerObj.Ctx, 0) + `/upload/notify/` + uploadInfo[daoUpload.Upload.Columns().UploadId].String()
		}
	}
	handlerObj.upload = NewUpload(ctx, uploadType, config)
	return handlerObj
}

func (handlerThis *Handler) Upload(r *ghttp.Request) (notifyInfo model.NotifyInfo, err error) {
	return handlerThis.upload.Upload(handlerThis.Ctx, r)
}

func (handlerThis *Handler) Sign() (signInfo model.SignInfo, err error) {
	return handlerThis.upload.Sign(handlerThis.Ctx, handlerThis.createUploadParam())
}

func (handlerThis *Handler) Config() (config map[string]any, err error) {
	return handlerThis.upload.Config(handlerThis.Ctx, handlerThis.createUploadParam())
}

func (handlerThis *Handler) Sts() (stsInfo map[string]any, err error) {
	return handlerThis.upload.Sts(handlerThis.Ctx, handlerThis.createUploadParam())
}

func (handlerThis *Handler) Notify(r *ghttp.Request) (notifyInfo model.NotifyInfo, err error) {
	return handlerThis.upload.Notify(handlerThis.Ctx, r)
}

func (handlerThis *Handler) createUploadParam() (param model.UploadParam) {
	switch handlerThis.Scene {
	default:
		param = model.UploadParam{
			Dir:        `upload/` + gtime.Now().Format(`Ymd`) + `/`,
			Expire:     gtime.Now().Unix() + 15*60,
			ExpireTime: 15 * 60,
			MinSize:    0,
			MaxSize:    1024 * 1024 * 1024,
		}
	}
	/* sceneInfo := get_or_set_ctx.GetCtxSceneInfo(handlerThis.Ctx)
	sceneId := sceneInfo[daoAuth.Scene.Columns().SceneId].String()
	loginInfo := get_or_set_ctx.GetCtxLoginInfo(handlerThis.Ctx)
	loginId := loginInfo[`login_id`]
	switch sceneId {
	case `platform`:
	case `org`:
		orgId := loginInfo[daoOrg.Admin.Columns().OrgId]
	case `app`:
	default:
	} */
	return
}
