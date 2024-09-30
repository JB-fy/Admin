package upload

import (
	daoUpload "api/internal/dao/upload"
	"api/internal/utils"
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

func NewHandler(ctx context.Context, scene string, uploadId uint) *Handler {
	handlerObj := Handler{
		Ctx:      ctx,
		Scene:    scene,
		UploadId: uploadId,
	}
	handlerObj.initUpload()
	return &handlerObj
}

type Handler struct {
	Ctx      context.Context
	Scene    string //上传场景。default默认。根据自身需求扩展，用于确定上传通道和上传参数
	UploadId uint   //上传ID
	upload   Upload
}

func (handlerThis *Handler) initUpload() {
	uploadFilter := g.Map{}
	if handlerThis.UploadId > 0 {
		uploadFilter[daoUpload.Upload.Columns().UploadId] = handlerThis.UploadId
	} else {
		uploadFilter[daoUpload.Upload.Columns().IsDefault] = 1
	}
	uploadInfo, _ := daoUpload.Upload.CtxDaoModel(handlerThis.Ctx).Filters(uploadFilter).One()

	config := uploadInfo[daoUpload.Upload.Columns().UploadConfig].Map()
	config[`upload_type`] = uploadInfo[daoUpload.Upload.Columns().UploadType]
	config[`uploadId`] = uploadInfo[daoUpload.Upload.Columns().UploadId]
	if gconv.Bool(config[`isNotify`]) {
		config[`callbackUrl`] = utils.GetRequestUrl(handlerThis.Ctx, 0) + `/upload/notify/` + uploadInfo[daoUpload.Upload.Columns().UploadId].String()
	}

	handlerThis.upload = NewUpload(handlerThis.Ctx, config)
}

func (handlerThis *Handler) createUploadParam() (param UploadParam) {
	switch handlerThis.Scene {
	default:
		param = UploadParam{
			Dir:        `upload/` + gtime.Now().Format(`Ymd`) + `/`,
			Expire:     gtime.Now().Unix() + 15*60,
			ExpireTime: 15 * 60,
			MinSize:    0,
			MaxSize:    1024 * 1024 * 1024,
		}
	}
	/* sceneInfo := utils.GetCtxSceneInfo(handlerThis.Ctx)
	sceneCode := sceneInfo[daoAuth.Scene.Columns().SceneCode].String()
	loginInfo := utils.GetCtxLoginInfo(handlerThis.Ctx)
	loginId := loginInfo[`login_id`]
	switch sceneCode {
	case `platform`:
	case `org`:
		orgId := loginInfo[daoOrg.Admin.Columns().OrgId]
	case `app`:
	default:
	} */
	return
}

func (handlerThis *Handler) Upload(r *ghttp.Request) (notifyInfo NotifyInfo, err error) {
	return handlerThis.upload.Upload(r)
}

func (handlerThis *Handler) Sign() (signInfo SignInfo, err error) {
	return handlerThis.upload.Sign(handlerThis.createUploadParam())
}

func (handlerThis *Handler) Config() (config map[string]any, err error) {
	return handlerThis.upload.Config(handlerThis.createUploadParam())
}

func (handlerThis *Handler) Sts() (stsInfo map[string]any, err error) {
	return handlerThis.upload.Sts(handlerThis.createUploadParam())
}

func (handlerThis *Handler) Notify(r *ghttp.Request) (notifyInfo NotifyInfo, err error) {
	return handlerThis.upload.Notify(r)
}
