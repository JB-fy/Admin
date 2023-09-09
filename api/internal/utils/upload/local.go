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

	dir := fmt.Sprintf(`common/%s/`, gtime.Now().Format(`Ymd`))
	expire := time.Now().Unix() + 15*60
	signInfo = map[string]interface{}{
		`uploadUrl`: upload.Url,
		// `uploadData`:  map[string]interface{}{},
		`host`:   upload.FileUrlPrefix,
		`dir`:    dir,
		`expire`: expire,
		`isRes`:  1,
	}

	uploadData := map[string]interface{}{
		`dir`:    dir,
		`expire`: expire,
		`rand`:   grand.S(8),
	}
	uploadData[`sign`] = upload.CreateSign(uploadData)

	signInfo[`uploadData`] = uploadData
	return
}

func (*Local) Sts(ctx context.Context, uploadFileType string) (stsInfo map[string]interface{}, err error) {
	return
}

func (*Local) Notify(ctx context.Context) (notifyInfo map[string]interface{}, err error) {
	config, _ := daoPlatform.Config.Get(ctx, []string{`localUploadUrl`, `localUploadSignKey`, `localUploadFileUrlPrefix`})
	upload := internal.NewLocal(ctx, config)

	r := g.RequestFromCtx(ctx)
	expire := r.PostFormValue(`expire`)
	if time.Now().Unix() > gconv.Int64(expire) {
		err = utils.NewErrorCode(ctx, 79999999, `签名失效`)
		return
	}

	dir := r.PostFormValue(`dir`)
	rand := r.PostFormValue(`rand`)
	key := r.PostFormValue(`key`)
	sign := r.PostFormValue(`sign`)

	signData := map[string]interface{}{
		`dir`:    dir,
		`expire`: expire,
		`rand`:   rand,
	}
	if sign != upload.CreateSign(signData) {
		err = utils.NewErrorCode(ctx, 79999999, `回调签名错误`)
		return
	}

	file := r.GetUploadFile(`file`)
	isRand := true
	if key != `` {
		isRand = false
		file.Filename = gstr.Replace(key, dir, ``) //修改保存文件名
	}
	filename, err := file.Save(`../`+dir, isRand)
	if err != nil {
		return
	}

	notifyInfo = map[string]interface{}{}
	notifyInfo[`url`] = upload.FileUrlPrefix + `/` + filename
	return
}
