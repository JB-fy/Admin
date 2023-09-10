package upload

import (
	daoPlatform "api/internal/dao/platform"
	"api/internal/utils"
	"api/internal/utils/upload/internal"
	"context"
	"fmt"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
)

type Local struct{}

func (*Local) Sign(ctx context.Context, uploadFileType string) (signInfo map[string]interface{}, err error) {
	config, _ := daoPlatform.Config.Get(ctx, []string{`localUploadUrl`, `localUploadSignKey`, `localUploadFileUrlPrefix`})
	upload := internal.NewLocal(ctx, config)

	type Option struct {
		Dir     string //上传的文件目录
		Expire  int64  //签名有效时间戳。单位：秒
		MinSize int64  //限制上传的文件大小。单位：字节
		MaxSize int64  //限制上传的文件大小。单位：字节。需要同时设置配置文件api/manifest/config/config.yaml中的server.clientMaxBodySize字段
	}
	option := Option{
		Dir:     fmt.Sprintf(`common/%s/`, gtime.Now().Format(`Ymd`)),
		Expire:  time.Now().Unix() + 15*60,
		MinSize: 0,
		MaxSize: 100 * 1024 * 1024,
	}

	signInfo = map[string]interface{}{
		`uploadUrl`: upload.Url,
		// `uploadData`:  map[string]interface{}{},
		`host`:   upload.FileUrlPrefix,
		`dir`:    option.Dir,
		`expire`: option.Expire,
		`isRes`:  1,
	}

	uploadData := map[string]interface{}{
		`dir`:     option.Dir,
		`expire`:  option.Expire,
		`minSize`: option.MinSize,
		`maxSize`: option.MaxSize,
		`rand`:    grand.S(8),
	}
	uploadData[`sign`] = upload.CreateSign(uploadData)

	signInfo[`uploadData`] = uploadData
	return
}

func (*Local) Sts(ctx context.Context, uploadFileType string) (stsInfo map[string]interface{}, err error) {
	return
}

func (*Local) Notify(ctx context.Context) (notifyInfo map[string]interface{}, err error) {
	return
}

func (*Local) Upload(ctx context.Context) (uploadInfo map[string]interface{}, err error) {
	config, _ := daoPlatform.Config.Get(ctx, []string{`localUploadUrl`, `localUploadSignKey`, `localUploadFileUrlPrefix`})
	upload := internal.NewLocal(ctx, config)

	r := g.RequestFromCtx(ctx)
	dir := r.Get(`dir`).String()
	expire := r.Get(`expire`).Int64()
	minSize := r.Get(`minSize`).Int64()
	maxSize := r.Get(`maxSize`).Int64()
	rand := r.Get(`rand`).String()
	key := r.Get(`key`).String()
	sign := r.Get(`sign`).String()

	if time.Now().Unix() > expire {
		err = utils.NewErrorCode(ctx, 79999999, `签名过期`)
		return
	}
	signData := map[string]interface{}{
		`dir`:     dir,
		`expire`:  expire,
		`minSize`: minSize,
		`maxSize`: maxSize,
		`rand`:    rand,
	}
	if sign != upload.CreateSign(signData) {
		err = utils.NewErrorCode(ctx, 79999999, `签名错误`)
		return
	}

	file := r.GetUploadFile(`file`)
	if minSize > 0 && minSize > file.Size {
		err = utils.NewErrorCode(ctx, 79999999, `文件不能小于`+gconv.String(minSize/(1024*1024))+`MB`)
		return
	}
	if maxSize > 0 && maxSize < file.Size {
		err = utils.NewErrorCode(ctx, 79999999, `文件不能大于`+gconv.String(maxSize/(1024*1024))+`MB`)
		return
	}

	isRand := true
	if key != `` {
		isRand = false
		file.Filename = gstr.Replace(key, dir, ``) //修改保存文件名
	}
	filename, err := file.Save(`../public/`+dir, isRand)
	if err != nil {
		return
	}

	uploadInfo = map[string]interface{}{}
	uploadInfo[`url`] = upload.FileUrlPrefix + `/` + dir + filename
	return
}
