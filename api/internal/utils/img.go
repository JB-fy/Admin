package utils

/* import (
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	_ "golang.org/x/image/bmp"
	_ "golang.org/x/image/tiff"
	_ "golang.org/x/image/webp"

	"image"
	"net/http"
	"slices"
	"strings"

	"github.com/disintegration/imaging"
)

func ImgDecode(imgBytes []byte, opts ...imaging.DecodeOption) (img image.Image, err error) {
	reader := BytesReaderPoolGet(imgBytes)
	defer BytesReaderPoolPut(reader)
	// 必须import _ "image/gif"导入要处理的图片类型的包，否则会报错：image: unknown format
	// img, _, err = image.Decode(reader)
	img, err = imaging.Decode(reader, opts...)
	return
}

func ImgEncode(img image.Image, format imaging.Format, opts ...imaging.EncodeOption) ([]byte, error) {
	buf := BytesBufferPoolGet()
	defer BytesBufferPoolPut(buf)
	err := imaging.Encode(buf, img, format, opts...)
	return buf.Bytes(), err
}

func ImgEncodeBytes(imgBytes []byte, format imaging.Format, opts ...imaging.EncodeOption) ([]byte, error) {
	img, err := ImgDecode(imgBytes)
	if err != nil {
		return nil, err
	}
	return ImgEncode(img, format, opts...)
}

func ImgFill(imgBytes []byte, width int, height int, anchor imaging.Anchor, filter imaging.ResampleFilter) ([]byte, error) {
	img, err := ImgDecode(imgBytes)
	if err != nil {
		return nil, err
	}
	format := imaging.JPEG //imaging不支持的格式如（webp格式），默认转jpeg格式
	if imgType := http.DetectContentType(imgBytes[:min(512, len(imgBytes))]); !slices.Contains([]string{`image/webp`}, imgType) {
		format, err = imaging.FormatFromExtension(strings.Replace(imgType, `image/`, ``, 1))
		if err != nil {
			return nil, err
		}
	}
	return ImgEncode(img, format)
} */
