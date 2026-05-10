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
	MaxSize         int            `json:"max_size"`          //最大字节数
	RatioArr        []string       `json:"ratio_arr"`         //宽高比：1:1, 3:4。比例正确时，不修改宽高
	EncodeFormatArr []string       `json:"encode_format_arr"` //需要转换的格式：video/jpeg, image/webp
	TargetFormat    imaging.Format `json:"target_format"`     //当EncodeFormatArr不为空，且需要格式转换时才有用，用于指定转换后的目标格式，默认：imaging.JPEG
	IsError         bool           `json:"is_error"`          //报错：0否 1是
}

func ImgHandle(imgBytesOfRaw []byte, imgOption ImgOption) (imgBytes []byte, err error) {
	imgBytes = imgBytesOfRaw
	imgType := http.DetectContentType(imgBytes[:min(512, len(imgBytes))])
	format, errTmp := imaging.FormatFromExtension(strings.Replace(imgType, `image/`, ``, 1))
	if errTmp != nil { //errors.Is(errTmp, imaging.ErrUnsupportedFormat)
		format = imaging.JPEG //imaging不支持的格式如（webp格式），默认转jpeg格式
	}
	isHandle := len(imgOption.EncodeFormatArr) > 0 && slices.Contains(imgOption.EncodeFormatArr, imgType)
	if isHandle {
		if imgOption.IsError {
			err = fmt.Errorf(`图片格式不支持：%s`, imgType)
			return
		}
		format = imgOption.TargetFormat
	} else if (imgOption.Width == 0 || imgOption.Height == 0) && imgOption.MaxSize == 0 {
		return
	}
	imgObj, err := ImgDecode(imgBytes)
	if err != nil {
		return
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
	if isHandle {
		// buf.Bytes()返回的字节后续还要使用。无法用连接池
		// buf := BytesBufferPoolGet()
		// defer BytesBufferPoolPut(buf)
		buf := bytes.NewBuffer(nil)
		err = imaging.Encode(buf, imgObj, format)
		if err != nil {
			return
		}
		imgBytes = buf.Bytes()
	}
	if imgOption.MaxSize > 0 && len(imgBytes) > imgOption.MaxSize {
		quality := 95
		qualityStep := 10
		qualityLimit := 50
		sigma := 0.0
		sigmaStep := 0.5
		sigmaLimit := 2.0
		var opts []imaging.EncodeOption
		buf := bytes.NewBuffer(nil)
		for len(imgBytes) > imgOption.MaxSize {
			if imgOption.IsError {
				err = fmt.Errorf(`图片大小不符合要求：最大%d`, imgOption.MaxSize)
				return
			}
			switch format {
			case imaging.JPEG:
				quality -= qualityStep
				if quality < qualityLimit {
					err = fmt.Errorf(`图片压缩后大小不符合要求：最大%d,图片质量%d。继续压缩无意义，会变得模糊`, imgOption.MaxSize, quality)
					return
				}
				opts = []imaging.EncodeOption{imaging.JPEGQuality(quality)}
			// case imaging.PNG:
			// 	opts = []imaging.EncodeOption{imaging.PNGCompressionLevel(png.BestCompression)} //实测大小没变化
			// case imaging.GIF:
			// case imaging.TIFF:
			// case imaging.BMP:
			default: //统一使用：高斯模糊
				sigma += sigmaStep
				if sigmaStep > sigmaLimit {
					err = fmt.Errorf(`图片压缩后大小不符合要求：最大%d,高斯模糊%f。继续压缩无意义，会变得模糊`, imgOption.MaxSize, sigma)
					return
				}
				imgObj = imaging.Blur(imgObj, 2)
			}
			buf.Reset()
			err = imaging.Encode(buf, imgObj, format, opts...)
			if err != nil {
				return
			}
			imgBytes = buf.Bytes()
		}
	}
	return
} */
