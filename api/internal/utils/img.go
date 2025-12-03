package utils

/* import (
	"bytes"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"net/http"
	"slices"
	"strings"

	"github.com/disintegration/imaging"
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
	EncodeFormatArr []string       `json:"encode_format_arr"` //需要转换的格式：video/jpeg, image/webp
	TargerFormat    imaging.Format `json:"targer_format"`     //当EncodeFormatArr不为空，且需要格式转换时才有用，用于指定转换后的目标格式，默认：imaging.JPEG
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
		}
	}
	imgObj, err := ImgDecode(imgBytes)
	if err != nil {
		return
	}
	if imgOption.Width > 0 && imgOption.Height > 0 && !(imgObj.Bounds().Dx() == imgOption.Width && imgObj.Bounds().Dy() == imgOption.Height) {
		imgObj = imaging.Fill(imgObj, imgOption.Width, imgOption.Height, imaging.Center, imaging.NearestNeighbor)
	}
	format, errTmp := imaging.FormatFromExtension(strings.Replace(imgType, `image/`, ``, 1))
	if errTmp != nil { //errors.Is(errTmp, imaging.ErrUnsupportedFormat)
		format = imaging.JPEG //imaging不支持的格式如（webp格式），默认转jpeg格式
	}
	if len(imgOption.EncodeFormatArr) > 0 && slices.Contains(imgOption.EncodeFormatArr, imgType) {
		format = imgOption.TargerFormat
	}
	// buf.Bytes()返回的字节后续还要使用。无法用连接池
	// buf := BytesBufferPoolGet()
	// defer BytesBufferPoolPut(buf)
	buf := bytes.NewBuffer(nil)
	err = imaging.Encode(buf, imgObj, format)
	imgBytes = buf.Bytes()
	return
} */
