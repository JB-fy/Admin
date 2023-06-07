package controller

import (
	"api/internal/utils"

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
		// /**--------参数处理 开始--------**/
		// var param *api.UploadSignReq
		// err := r.Parse(&param)
		// if err != nil {
		// 	r.Response.Writeln(err.Error())
		// 	return
		// }
		// /**--------参数处理 结束--------**/

		// token, err := service.Upload().Sign(r.Context(), sceneCode, param.Type)
		// if err != nil {
		// 	utils.HttpFailJson(r, err)
		// 	return
		// }
		// utils.HttpSuccessJson(r, map[string]interface{}{"token": token}, 0)

		/* $option = [];
		   $type = $this->request->input('type');
		   switch ($type) {
		       case 'type1':
		           $option = [
		               'callbackEnable' => true, //是否回调服务器
		               'expireTime' => 5 * 60, //签名有效时间
		               'dir' => 'common/' . date('Y/m/d/His') . mt_rand(1000, 9999) . '_',    //上传的文件前缀
		               'minSize' => 0,    //限制上传的文件大小。单位：字节
		               'maxSize' => 100 * 1024 * 1024,    //限制上传的文件大小。单位：字节
		           ];
		           break;
		   }
		   $upload = make('upload');
		   $upload->createSign($option); */
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
