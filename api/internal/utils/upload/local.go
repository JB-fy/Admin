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
)

type Local struct{}

func (*Local) Sign(ctx context.Context, uploadFileType string) (signInfo map[string]interface{}, err error) {
	config, _ := daoPlatform.Config.Get(ctx, []string{`localUploadApi`, `localSignKey`})
	upload := internal.NewLocal(ctx, config)

	dir := fmt.Sprintf(`common/%s/`, gtime.Now().Format(`Ymd`))
	expire := time.Now().Unix() + 15*60
	signInfo = map[string]interface{}{
		// `uploadUrl`: upload.Host + upload.UploadApi,
		`uploadUrl`: upload.Host + `/upload/upload`,
		// `uploadData`:  map[string]interface{}{},
		`host`:   upload.Host,
		`dir`:    dir,
		`expire`: expire,
		`isRes`:  1,
	}

	uploadData := map[string]interface{}{
		`dir`:    dir,
		`expire`: expire,
		`time`:   time.Now().UnixMilli(),
	}
	uploadData[`sign`] = upload.CreateSign(uploadData)

	signInfo[`uploadData`] = uploadData
	return
}

func (*Local) Sts(ctx context.Context, uploadFileType string) (stsInfo map[string]interface{}, err error) {
	return
}

func (*Local) Notify(ctx context.Context) (notifyInfo map[string]interface{}, err error) {
	config, _ := daoPlatform.Config.Get(ctx, []string{`localUploadApi`, `localSignKey`})
	upload := internal.NewLocal(ctx, config)

	r := g.RequestFromCtx(ctx)
	dir := r.PostFormValue(`dir`)
	expire := r.PostFormValue(`expire`)
	time := r.PostFormValue(`time`)
	key := r.PostFormValue(`key`)
	sign := r.PostFormValue(`sign`)

	signData := map[string]interface{}{
		`dir`:    dir,
		`expire`: expire,
		`time`:   time,
	}
	if sign != upload.CreateSign(signData) {
		err = utils.NewErrorCode(ctx, 79999999, `回调签名错误`)
		return
	}

	file := r.GetUploadFile(`file`)
	isRandom := true
	if key != `` {
		isRandom = false
		file.Filename = gstr.Replace(key, dir, ``) //修改保存文件名
	}
	filename, err := file.Save(`../`+dir, isRandom)
	if err != nil {
		return
	}

	notifyInfo = map[string]interface{}{}
	notifyInfo[`url`] = upload.Host + `/` + filename
	return
}
