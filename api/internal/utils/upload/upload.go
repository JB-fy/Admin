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

type UploadParam struct {
	Dir        string //上传的文件目录
	Expire     int64  //签名有效时间戳。单位：秒
	ExpireTime int64  //签名有效时间。单位：秒
	MinSize    int    //限制上传的文件大小。单位：字节
	MaxSize    int    //限制上传的文件大小。单位：字节。本地上传（uploadOfLocal.go）需要同时设置配置文件api/manifest/config/config.yaml中的server.clientMaxBodySize字段
}

type SignInfo struct {
	UploadUrl  string         //上传地址
	UploadData map[string]any //上传数据
	Host       string         //站点域名（当上传无响应数据，前端组件用于与上传目录拼接形成文件访问地址）
	Dir        string         //上传目录
	Expire     uint           //过期时间。单位：秒
	IsRes      uint           //是否有响应信息。0否 1是
}

type NotifyInfo struct {
	Url      string //地址
	Width    uint   //宽度
	Height   uint   //高度
	Size     uint   //大小。单位：比特
	MimeType string //文件类型
}

type Upload interface {
	Upload(r *ghttp.Request) (notifyInfo NotifyInfo, err error)  // 本地上传
	Sign(param UploadParam) (signInfo SignInfo, err error)       // 获取签名（H5直传用）
	Config(param UploadParam) (config map[string]any, err error) // 获取配置信息（APP直传前调用）
	Sts(param UploadParam) (stsInfo map[string]any, err error)   // 获取Sts Token（APP直传用）
	Notify(r *ghttp.Request) (notifyInfo NotifyInfo, err error)  // 回调
}

// scene	上传场景。default默认。根据自身需求扩展，用于确定上传通道和上传参数
// uploadId	上传ID
func NewUpload(ctx context.Context, scene string, uploadId uint) Upload {
	uploadFilter := g.Map{}
	if uploadId > 0 {
		uploadFilter[daoUpload.Upload.Columns().UploadId] = uploadId
	} else {
		switch scene {
		default:
			uploadFilter[daoUpload.Upload.Columns().IsDefault] = 1
		}
	}
	uploadInfo, _ := daoUpload.Upload.CtxDaoModel(ctx).Filters(uploadFilter).One()

	config := uploadInfo[daoUpload.Upload.Columns().UploadConfig].Map()
	switch uploadInfo[daoUpload.Upload.Columns().UploadType].Uint() {
	case 1: //阿里云OSS
		if gconv.Bool(config[`isNotify`]) {
			config[`callbackUrl`] = utils.GetRequestUrl(ctx, 0) + `/upload/notify/` + uploadInfo[daoUpload.Upload.Columns().UploadId].String()
		}
		return NewUploadOfAliyunOss(ctx, config)
	// case 0: //本地
	default:
		config[`uploadId`] = uploadInfo[daoUpload.Upload.Columns().UploadId]
		return NewUploadOfLocal(ctx, config)
	}
}

func CreateUploadParam(ctx context.Context, scene string) (param UploadParam) {
	/* sceneInfo := utils.GetCtxSceneInfo(ctx)
	sceneCode := sceneInfo[daoAuth.Scene.Columns().SceneCode].String()
	loginInfo := utils.GetCtxLoginInfo(ctx)
	loginId := loginInfo[`login_id`]
	switch sceneCode {
	case `platform`:
	case `org`:
		orgId := loginInfo[daoOrg.Admin.Columns().OrgId]
	case `app`:
	default:
	} */
	switch scene {
	default:
		param = UploadParam{
			Dir:        `upload/` + gtime.Now().Format(`Ymd`) + `/`,
			Expire:     gtime.Now().Unix() + 15*60,
			ExpireTime: 15 * 60,
			MinSize:    0,
			MaxSize:    1024 * 1024 * 1024,
		}
	}
	return
}
