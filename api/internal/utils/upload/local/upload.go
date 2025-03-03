package local

import (
	"api/internal/consts"
	"api/internal/utils"
	"api/internal/utils/upload/model"
	"context"
	"errors"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"net/http"
	"net/url"
	"os"
	"sort"
	"time"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
)

type Upload struct {
	UploadId      uint   `json:"uploadId"`
	SignKey       string `json:"signKey"`
	Url           string `json:"url"`
	FileSaveDir   string `json:"fileSaveDir"`
	FileUrlPrefix string `json:"fileUrlPrefix"`
	IsCluster     uint   `json:"isCluster"`
	IsSameServer  uint   `json:"isSameServer"`
	ServerList    []struct {
		Ip   string `json:"ip"`
		Host string `json:"host"`
	} `json:"serverList"`
}

func NewUpload(ctx context.Context, config map[string]any) model.Upload {
	obj := &Upload{}
	gconv.Struct(config, obj)
	if obj.UploadId == 0 || obj.SignKey == `` {
		panic(`缺少配置：上传-本地`)
	}
	if obj.FileSaveDir == `` {
		obj.FileSaveDir = `../public/`
	}
	if obj.Url == `` {
		obj.Url = utils.GetRequestUrl(ctx, 0) + `/upload/upload`
		if utils.IsDev(ctx) {
			obj.Url = utils.GetRequestUrl(ctx, 20) + `/upload/upload`
		}
	}
	urlObj, _ := url.Parse(obj.Url)
	obj.FileUrlPrefix = urlObj.Scheme + `://` + urlObj.Host
	if !utils.IsDev(ctx) && obj.IsCluster == 1 {
		obj.FileUrlPrefix = utils.GetRequestUrl(ctx, 10)
		serverHostObj, _ := url.Parse(obj.FileUrlPrefix)
		serverIp := g.Cfg().MustGetWithEnv(ctx, consts.LOCAL_SERVER_NETWORK_IP).String()
		for _, v := range obj.ServerList {
			if v.Ip == serverIp {
				serverHostObj, _ = url.Parse(v.Host)
				obj.FileUrlPrefix = serverHostObj.Scheme + `://` + serverHostObj.Host
				break
			}
		}
		if obj.IsSameServer == 1 {
			urlObj.Scheme = serverHostObj.Scheme
			urlObj.Host = serverHostObj.Host
			obj.Url = urlObj.String()
		}
	}
	return obj
}

// 本地上传
func (uploadThis *Upload) Upload(ctx context.Context, r *ghttp.Request) (notifyInfo model.NotifyInfo, err error) {
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
		`upload_id`: uploadThis.UploadId,
		`dir`:       dir,
		`expire`:    expire,
		`min_size`:  minSize,
		`max_size`:  maxSize,
		`rand`:      rand,
	}
	if sign != uploadThis.sign(signData) {
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
func (uploadThis *Upload) Sign(ctx context.Context, param model.UploadParam) (signInfo model.SignInfo, err error) {
	signInfo = model.SignInfo{
		UploadUrl: uploadThis.Url,
		Host:      uploadThis.FileUrlPrefix,
		Dir:       param.Dir,
		Expire:    gconv.Uint(param.Expire),
		IsRes:     1,
	}

	uploadData := map[string]any{
		`upload_id`: uploadThis.UploadId,
		`dir`:       param.Dir,
		`expire`:    param.Expire,
		`min_size`:  param.MinSize,
		`max_size`:  param.MaxSize,
		`rand`:      grand.S(8),
	}
	uploadData[`sign`] = uploadThis.sign(uploadData)

	signInfo.UploadData = uploadData
	return
}

// 获取配置信息（APP直传前调用）
func (uploadThis *Upload) Config(ctx context.Context, param model.UploadParam) (config map[string]any, err error) {
	return
}

// 获取Sts Token（APP直传用）
func (uploadThis *Upload) Sts(ctx context.Context, param model.UploadParam) (stsInfo map[string]any, err error) {
	return
}

// 回调
func (uploadThis *Upload) Notify(ctx context.Context, r *ghttp.Request) (notifyInfo model.NotifyInfo, err error) {
	return
}

// 生成签名
func (uploadThis *Upload) sign(data map[string]any) (sign string) {
	keyArr := make([]string, 0, len(data))
	for key := range data {
		keyArr = append(keyArr, key)
	}
	sort.Strings(keyArr)

	buf := utils.BytesBufferPoolGet()
	defer utils.BytesBufferPoolPut(buf)
	for _, key := range keyArr {
		buf.WriteString(key)
		buf.WriteString(`=`)
		if tmp := gvar.New(data[key]); tmp.IsMap() || tmp.IsSlice() {
			buf.Write(gjson.MustEncode(data[key]))
		} else {
			buf.WriteString(gconv.String(data[key]))
		}
		buf.WriteString(`&`)
	}
	buf.WriteString(`signSecret=`)
	buf.WriteString(uploadThis.SignKey)

	sign = gmd5.MustEncryptBytes(buf.Bytes())
	return
}
