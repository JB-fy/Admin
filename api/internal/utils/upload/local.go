package upload

import (
	"api/internal/utils"
	"context"
	"fmt"
	"sort"
	"time"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
)

type Local struct {
	Ctx           context.Context
	Url           string `json:"localUploadUrl"`
	SignKey       string `json:"localUploadSignKey"`
	FileSaveDir   string `json:"localUploadFileSaveDir"`
	FileUrlPrefix string `json:"localUploadFileUrlPrefix"`
}

func NewLocal(ctx context.Context, config map[string]interface{}) *Local {
	localObj := Local{
		Ctx: ctx,
	}
	gconv.Struct(config, &localObj)
	return &localObj
}

func (uploadThis *Local) Upload() (uploadInfo map[string]interface{}, err error) {
	r := g.RequestFromCtx(uploadThis.Ctx)
	dir := r.Get(`dir`).String()
	expire := r.Get(`expire`).Int64()
	minSize := r.Get(`minSize`).Int64()
	maxSize := r.Get(`maxSize`).Int64()
	rand := r.Get(`rand`).String()
	key := r.Get(`key`).String()
	sign := r.Get(`sign`).String()

	if time.Now().Unix() > expire {
		err = utils.NewErrorCode(uploadThis.Ctx, 79999999, `签名过期`)
		return
	}
	signData := map[string]interface{}{
		`dir`:     dir,
		`expire`:  expire,
		`minSize`: minSize,
		`maxSize`: maxSize,
		`rand`:    rand,
	}
	if sign != uploadThis.CreateSign(signData) {
		err = utils.NewErrorCode(uploadThis.Ctx, 79999999, `签名错误`)
		return
	}

	file := r.GetUploadFile(`file`)
	if minSize > 0 && minSize > file.Size {
		err = utils.NewErrorCode(uploadThis.Ctx, 79999999, `文件不能小于`+gconv.String(minSize/(1024*1024))+`MB`)
		return
	}
	if maxSize > 0 && maxSize < file.Size {
		err = utils.NewErrorCode(uploadThis.Ctx, 79999999, `文件不能大于`+gconv.String(maxSize/(1024*1024))+`MB`)
		return
	}

	// isRand := true
	if key != `` {
		// isRand = false
		file.Filename = gstr.Replace(key, dir, ``)
	} else {
		file.Filename = dir + gconv.String(time.Now().UnixMilli()) + `_` + gconv.String(grand.N(10000000, 99999999)) + gfile.Ext(file.Filename)
	}
	filename, err := file.Save(uploadThis.FileSaveDir + dir /* , isRand */)
	if err != nil {
		return
	}

	uploadInfo = map[string]interface{}{}
	uploadInfo[`url`] = uploadThis.FileUrlPrefix + `/` + dir + filename
	return
}

func (uploadThis *Local) Sign(uploadFileType string) (signInfo map[string]interface{}, err error) {
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
		`uploadUrl`: uploadThis.Url,
		// `uploadData`:  map[string]interface{}{},
		`host`:   uploadThis.FileUrlPrefix,
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
	uploadData[`sign`] = uploadThis.CreateSign(uploadData)

	signInfo[`uploadData`] = uploadData
	return
}

func (uploadThis *Local) Sts(uploadFileType string) (stsInfo map[string]interface{}, err error) {
	return
}

func (uploadThis *Local) Notify() (notifyInfo map[string]interface{}, err error) {
	return
}

// 生成签名
func (uploadThis *Local) CreateSign(signData map[string]interface{}) (sign string) {
	keyArr := []string{}
	for k := range signData {
		keyArr = append(keyArr, k)
	}
	sort.Strings(keyArr)
	str := ``
	for _, k := range keyArr {
		str += k + `=` + gconv.String(signData[k]) + `&`
	}
	str += `key=` + uploadThis.SignKey
	sign = gmd5.MustEncryptString(str)
	return
}
