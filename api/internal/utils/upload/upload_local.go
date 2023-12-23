package upload

import (
	"context"
	"errors"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
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
func (uploadThis *UploadOfLocal) Upload() (notifyInfo NotifyInfo, err error) {
	r := g.RequestFromCtx(uploadThis.Ctx)
	dir := r.Get(`dir`).String()
	expire := r.Get(`expire`).Int64()
	minSize := r.Get(`minSize`).Int64()
	maxSize := r.Get(`maxSize`).Int64()
	rand := r.Get(`rand`).String()
	key := r.Get(`key`).String()
	sign := r.Get(`sign`).String()

	if time.Now().Unix() > expire {
		err = errors.New(`签名过期`)
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
		err = errors.New(`签名错误`)
		return
	}

	file := r.GetUploadFile(`file`)
	if minSize > 0 && minSize > file.Size {
		err = errors.New(`文件不能小于` + gconv.String(minSize/(1024*1024)) + `MB`)
		return
	}
	if maxSize > 0 && maxSize < file.Size {
		err = errors.New(`文件不能大于` + gconv.String(maxSize/(1024*1024)) + `MB`)
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

	//获取图片宽高
	fileTmp, err := os.Open(uploadThis.FileSaveDir + dir + filename)
	if err == nil {
		defer fileTmp.Close()
		img, _, err := image.Decode(fileTmp)
		if err == nil {
			notifyInfo.Width = gconv.Uint(img.Bounds().Dx())
			notifyInfo.Height = gconv.Uint(img.Bounds().Dy())
		}
	}
	notifyInfo.Size = gconv.Uint(file.Size)

	notifyInfo.Url = uploadThis.FileUrlPrefix + `/` + dir + filename
	//有时文件信息放地址后面，一起保存在数据库中会更好。比如：苹果手机做瀑布流时需要知道图片宽高，这时就能直接从地址中解析获取
	urlQueryArr := []string{}
	if notifyInfo.Width > 0 {
		urlQueryArr = append(urlQueryArr, `w=`+gconv.String(notifyInfo.Width))
	}
	if notifyInfo.Height > 0 {
		urlQueryArr = append(urlQueryArr, `h=`+gconv.String(notifyInfo.Height))
	}
	if notifyInfo.Size > 0 {
		urlQueryArr = append(urlQueryArr, `s=`+gconv.String(notifyInfo.Size))
	}
	/* if notifyInfo.MimeType != `` {
		urlQueryArr = append(urlQueryArr, `m=`+notifyInfo.MimeType)
	} */
	if len(urlQueryArr) > 0 {
		notifyInfo.Url += `?` + gstr.Join(urlQueryArr, `&`)
	}
	return
}

// 获取签名（H5直传用）
func (uploadThis *UploadOfLocal) Sign(param UploadParam) (signInfo SignInfo, err error) {
	signInfo = SignInfo{
		UploadUrl: uploadThis.Url,
		Host:      uploadThis.FileUrlPrefix,
		Dir:       param.Dir,
		Expire:    gconv.Uint(param.Expire),
		IsRes:     1,
	}

	uploadData := map[string]interface{}{
		`dir`:     param.Dir,
		`expire`:  param.Expire,
		`minSize`: param.MinSize,
		`maxSize`: param.MaxSize,
		`rand`:    grand.S(8),
	}
	uploadData[`sign`] = uploadThis.CreateSign(uploadData)

	signInfo.UploadData = uploadData
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
func (uploadThis *UploadOfLocal) Notify() (notifyInfo NotifyInfo, err error) {
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
