package upload

import (
	"api/internal/consts"
	daoUpload "api/internal/dao/upload"
	"api/internal/utils"
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

type Handler struct {
	Ctx      context.Context
	Scene    string //上传场景。default默认。根据自身需求扩展，用于确定上传通道和上传参数
	UploadId uint   //上传ID
	upload   Upload
}

func NewHandler(ctx context.Context, scene string, uploadId uint) *Handler {
	handlerObj := &Handler{
		Ctx:      ctx,
		Scene:    scene,
		UploadId: uploadId,
	}

	uploadInfo, _ := daoUpload.Upload.CacheGetInfo(handlerObj.Ctx, handlerObj.UploadId)

	config := uploadInfo[daoUpload.Upload.Columns().UploadConfig].Map()
	config[`uploadId`] = uploadInfo[daoUpload.Upload.Columns().UploadId]
	if gconv.Bool(config[`isNotify`]) {
		config[`callbackUrl`] = utils.GetRequestUrl(handlerObj.Ctx, 0) + `/upload/notify/` + uploadInfo[daoUpload.Upload.Columns().UploadId].String()
	}
	switch uploadInfo[daoUpload.Upload.Columns().UploadType].Uint() {
	case 1: //阿里云OSS
	// case 0: //本地
	default:
		handleUrl := func(strRaw string) (str string) {
			str = strRaw
			if gstr.Pos(str, `http`) != 0 {
				currentUrl := utils.GetRequestUrl(ctx, 0)
				for _, v := range []string{`0.0.0.0`, `127.0.0.1`} {
					if gstr.Pos(currentUrl, v) != -1 {
						if utils.IsDev(ctx) {
							currentUrl = gstr.Replace(currentUrl, v, g.Cfg().MustGetWithEnv(ctx, consts.SERVER_LOCAL_IP).String(), 1)
						} else {
							currentUrl = gstr.Replace(currentUrl, v, g.Cfg().MustGetWithEnv(ctx, consts.SERVER_NETWORK_IP).String(), 1)
						}
						break
					}
				}
				if str != `` && gstr.Pos(str, `/`) != 0 {
					str = `/` + str
				}
				str = currentUrl + str
			}
			return
		}
		config[`url`] = handleUrl(gconv.String(config[`url`]))
		config[`fileUrlPrefix`] = handleUrl(gconv.String(config[`fileUrlPrefix`]))
	}

	config[`uploadType`] = uploadInfo[daoUpload.Upload.Columns().UploadType]
	handlerObj.upload = NewUpload(config)
	return handlerObj
}

func (handlerThis *Handler) Upload(r *ghttp.Request) (notifyInfo NotifyInfo, err error) {
	return handlerThis.upload.Upload(handlerThis.Ctx, r)
}

func (handlerThis *Handler) Sign() (signInfo SignInfo, err error) {
	return handlerThis.upload.Sign(handlerThis.Ctx, handlerThis.createUploadParam())
}

func (handlerThis *Handler) Config() (config map[string]any, err error) {
	return handlerThis.upload.Config(handlerThis.Ctx, handlerThis.createUploadParam())
}

func (handlerThis *Handler) Sts() (stsInfo map[string]any, err error) {
	return handlerThis.upload.Sts(handlerThis.Ctx, handlerThis.createUploadParam())
}

func (handlerThis *Handler) Notify(r *ghttp.Request) (notifyInfo NotifyInfo, err error) {
	return handlerThis.upload.Notify(handlerThis.Ctx, r)
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
	sceneId := sceneInfo[daoAuth.Scene.Columns().SceneId].String()
	loginInfo := utils.GetCtxLoginInfo(handlerThis.Ctx)
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
