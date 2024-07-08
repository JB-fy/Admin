package upload

import (
	"context"
	"errors"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"net/http"
	"os"
	"sort"
	"time"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/net/ghttp"
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

func NewUploadOfLocal(ctx context.Context, config map[string]any) *UploadOfLocal {
	uploadObj := UploadOfLocal{Ctx: ctx}
	gconv.Struct(config, &uploadObj)
	if uploadObj.Url == `` || uploadObj.SignKey == `` || uploadObj.FileSaveDir == `` || uploadObj.FileUrlPrefix == `` {
		panic(`缺少配置：上传-本地`)
	}
	return &uploadObj
}

// 本地上传
func (uploadThis *UploadOfLocal) Upload(r *ghttp.Request) (notifyInfo NotifyInfo, err error) {
	dir := r.Get(`dir`).String()
	expire := r.Get(`expire`).Int64()
	minSize := r.Get(`min_size`).Int64()
	maxSize := r.Get(`max_size`).Int64()
	rand := r.Get(`rand`).String()
	key := r.Get(`key`).String()
	sign := r.Get(`sign`).String()

	if time.Now().Unix() > expire {
		err = errors.New(`签名过期`)
		return
	}
	signData := map[string]any{
		`dir`:      dir,
		`expire`:   expire,
		`min_size`: minSize,
		`max_size`: maxSize,
		`rand`:     rand,
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

	fileTmp, err := os.Open(uploadThis.FileSaveDir + dir + filename)
	if err == nil {
		defer fileTmp.Close()
		//获取图片宽高
		img, _, errTmp := image.Decode(fileTmp)
		if errTmp == nil {
			notifyInfo.Width = gconv.Uint(img.Bounds().Dx())
			notifyInfo.Height = gconv.Uint(img.Bounds().Dy())
		}
		//获取文件的MIME类型
		buffer := make([]byte, 512)
		_, errTmp = fileTmp.ReadAt(buffer, 0)
		if errTmp == nil {
			notifyInfo.MimeType = http.DetectContentType(buffer)
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

	uploadData := map[string]any{
		`dir`:      param.Dir,
		`expire`:   param.Expire,
		`min_size`: param.MinSize,
		`max_size`: param.MaxSize,
		`rand`:     grand.S(8),
	}
	uploadData[`sign`] = uploadThis.CreateSign(uploadData)

	signInfo.UploadData = uploadData
	return
}

// 获取配置信息（APP直传前调用）
func (uploadThis *UploadOfLocal) Config(param UploadParam) (config map[string]any, err error) {
	return
}

// 获取Sts Token（APP直传用）
func (uploadThis *UploadOfLocal) Sts(param UploadParam) (stsInfo map[string]any, err error) {
	return
}

// 回调
func (uploadThis *UploadOfLocal) Notify(r *ghttp.Request) (notifyInfo NotifyInfo, err error) {
	return
}

// 生成签名
func (uploadThis *UploadOfLocal) CreateSign(data map[string]any) (sign string) {
	keyArr := []string{}
	for k := range data {
		keyArr = append(keyArr, k)
	}
	sort.Strings(keyArr)
	strArr := []string{}
	for _, k := range keyArr {
		tmp := gvar.New(data[k])
		if tmp.IsMap() || tmp.IsSlice() {
			strArr = append(strArr, k+`=`+gjson.MustEncodeString(data[k]))
		} else {
			strArr = append(strArr, k+`=`+gconv.String(data[k]))
		}
	}
	strArr = append(strArr, `key=`+uploadThis.SignKey)
	sign = gmd5.MustEncryptString(gstr.Join(strArr, `&`))
	return
}
