package upload

import (
	"api/internal/utils"
	"context"
	"sort"
	"time"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
)

type UploadOfLocal struct {
	Ctx           context.Context
	Url           string `json:"uploadOfLocalUrl"`
	SignKey       string `json:"uploadOfLocalSignKey"`
	FileSaveDir   string `json:"uploadOfLocalFileSaveDir"`
	FileUrlPrefix string `json:"uploadOfLocalFileUrlPrefix"`
}

func NewUploadOfLocal(ctx context.Context, config map[string]interface{}) *UploadOfLocal {
	uploadOfLocalObj := UploadOfLocal{
		Ctx: ctx,
	}
	gconv.Struct(config, &uploadOfLocalObj)
	return &uploadOfLocalObj
}

// 本地上传
func (uploadThis *UploadOfLocal) Upload() (uploadInfo map[string]interface{}, err error) {
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

// 获取签名（H5直传用）
func (uploadThis *UploadOfLocal) Sign(param UploadParam) (signInfo map[string]interface{}, err error) {
	signInfo = map[string]interface{}{
		`uploadUrl`: uploadThis.Url,
		// `uploadData`:  map[string]interface{}{},
		`host`:   uploadThis.FileUrlPrefix,
		`dir`:    param.Dir,
		`expire`: param.Expire,
		`isRes`:  1,
	}

	uploadData := map[string]interface{}{
		`dir`:     param.Dir,
		`expire`:  param.Expire,
		`minSize`: param.MinSize,
		`maxSize`: param.MaxSize,
		`rand`:    grand.S(8),
	}
	uploadData[`sign`] = uploadThis.CreateSign(uploadData)

	signInfo[`uploadData`] = uploadData
	return
}

// 获取配置信息（APP直传前调用）
func (uploadThis *UploadOfLocal) Config(param UploadParam) (config map[string]interface{}, err error) {
	return
}

// 获取Sts Token（APP直传用）
func (uploadThis *UploadOfLocal) Sts(param UploadParam) (stsInfo map[string]interface{}, err error) {
	return
}

// 回调
func (uploadThis *UploadOfLocal) Notify() (notifyInfo map[string]interface{}, err error) {
	return
}

// 生成签名
func (uploadThis *UploadOfLocal) CreateSign(signData map[string]interface{}) (sign string) {
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
