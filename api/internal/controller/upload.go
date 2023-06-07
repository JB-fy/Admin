package controller

import (
	"api/api"
	"api/internal/service"
	"api/internal/utils"
	"fmt"

	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/grand"
	"github.com/gogf/gf/v2/net/ghttp"
)

type Upload struct{}

func NewUpload() *Upload {
	return &Upload{}
}

// 获取签名
func (c *Upload) Sign(r *ghttp.Request) {
	sceneCode := utils.GetCtxSceneCode(r.GetCtx())
	switch sceneCode {
	case "platformAdmin":
		/**--------参数处理 开始--------**/
		var param *api.UploadSignReq
		err := r.Parse(&param)
		if err != nil {
			r.Response.Writeln(err.Error())
			return
		}
		/**--------参数处理 结束--------**/

		option := map[string]interface{}{}
		switch param.UploadType {
		default:
			option = map[string]interface{}{
				"callbackUrl": "",                                                                                   //是否回调服务器。空字符串不回调
				"expireTime":  15 * 60,                                                                              //签名有效时间。单位：秒
				"dir":         fmt.Sprintf("common/%s_%d_", gtime.Now().Format("Y-m-d H:i:s"), grand.N(1000, 9999)), //上传的文件前缀
				"minSize":     0,                                                                                    //限制上传的文件大小。单位：字节
				"maxSize":     100 * 1024 * 1024,                                                                    //限制上传的文件大小。单位：字节
			}
		}

		filter := map[string]interface{}{
			"configKey": []string{"aliyunOssAccessKeyId", "aliyunOssAccessKeySecret", "aliyunOssHost", "aliyunOssBucket"},
		}
		config, _ := service.Config().Get(r.Context(), filter)
		upload := utils.NewAliyunOss(r.GetCtx(), config)
		signInfo, _ := upload.CreateSign(option)
		utils.HttpSuccessJson(r, signInfo, 0)
	}
}

// 回调
func (c *Upload) Notify(r *ghttp.Request) {
	// /**
	//  * @var \App\Plugin\Upload\AbstractUpload
	//  */
	//  $upload = make('upload');
	//  $upload->notify();
}
