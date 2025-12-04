package utils

/* import (
	"bytes"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"net/http"
	"slices"
	"strings"

	"github.com/disintegration/imaging"
	"github.com/gogf/gf/v2/util/gconv"
	_ "golang.org/x/image/bmp"
	_ "golang.org/x/image/tiff"
	_ "golang.org/x/image/webp"
)

func ImgDecode(imgBytes []byte, opts ...imaging.DecodeOption) (img image.Image, err error) {
	reader := BytesReaderPoolGet(imgBytes)
	defer BytesReaderPoolPut(reader)
	// 必须import _ "image/gif"导入要处理的图片类型的包，否则会报错：image: unknown format
	// img, _, err = image.Decode(reader)
	img, err = imaging.Decode(reader, opts...)
	return
}

type ImgOption struct {
	Width           int            `json:"width"`
	Height          int            `json:"height"`
	RatioArr        []string       `json:"ratio_arr"`         //宽高比：1:1, 3:4。比例正确时，不修改宽高
	EncodeFormatArr []string       `json:"encode_format_arr"` //需要转换的格式：video/jpeg, image/webp
	TargerFormat    imaging.Format `json:"targer_format"`     //当EncodeFormatArr不为空，且需要格式转换时才有用，用于指定转换后的目标格式，默认：imaging.JPEG
	IsError         bool           `json:"is_error"`          //报错：0否 1是
}

func ImgHandle(imgBytesOfRaw []byte, imgOption ImgOption) (imgBytes []byte, err error) {
	imgBytes = imgBytesOfRaw
	var imgType string
	if imgOption.Width == 0 || imgOption.Height == 0 {
		if len(imgOption.EncodeFormatArr) == 0 {
			return
		}
		imgType = http.DetectContentType(imgBytes[:min(512, len(imgBytes))])
		if !slices.Contains(imgOption.EncodeFormatArr, imgType) {
			return
		} else {
			if imgOption.IsError {
				err = fmt.Errorf(`图片格式不支持：%s`, imgType)
				return
			}
		}
	}
	imgObj, err := ImgDecode(imgBytes)
	if err != nil {
		return
	}
	if imgType == `` {
		imgType = http.DetectContentType(imgBytes[:min(512, len(imgBytes))])
	}
	isHandle := false
	format, errTmp := imaging.FormatFromExtension(strings.Replace(imgType, `image/`, ``, 1))
	if errTmp != nil { //errors.Is(errTmp, imaging.ErrUnsupportedFormat)
		format = imaging.JPEG //imaging不支持的格式如（webp格式），默认转jpeg格式
	}
	if imgOption.Width > 0 && imgOption.Height > 0 && !(imgObj.Bounds().Dx() == imgOption.Width && imgObj.Bounds().Dy() == imgOption.Height) {
		if len(imgOption.RatioArr) == 0 {
			if !(imgObj.Bounds().Dx() == imgOption.Width && imgObj.Bounds().Dy() == imgOption.Height) {
				if imgOption.IsError {
					err = fmt.Errorf(`图片宽高不符合要求：宽%d,高%d`, imgOption.Width, imgOption.Height)
					return
				}
				isHandle = true
			}
		} else if !slices.Contains(imgOption.RatioArr, GetRatio(imgObj.Bounds().Dx(), imgObj.Bounds().Dy())) {
			if imgOption.IsError {
				err = fmt.Errorf(`图片宽高比不符合要求：%s`, gconv.String(imgOption.RatioArr))
				return
			}
			isHandle = true
		}
		if isHandle {
			imgObj = imaging.Fill(imgObj, imgOption.Width, imgOption.Height, imaging.Center, imaging.NearestNeighbor)
		}
	}
	if len(imgOption.EncodeFormatArr) > 0 && slices.Contains(imgOption.EncodeFormatArr, imgType) {
		isHandle = true
		format = imgOption.TargerFormat
	}
	if !isHandle {
		return
	}
	// buf.Bytes()返回的字节后续还要使用。无法用连接池
	// buf := BytesBufferPoolGet()
	// defer BytesBufferPoolPut(buf)
	buf := bytes.NewBuffer(nil)
	err = imaging.Encode(buf, imgObj, format)
	imgBytes = buf.Bytes()
	return
} */
