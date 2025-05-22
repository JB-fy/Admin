/* common.go与funcs.go的区别：
common.go：基于当前框架封装的常用函数（与框架耦合）
funcs.go：基于golang封装的常用函数（不与框架耦合） */

package utils

/* import (
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	_ "golang.org/x/image/bmp"
	_ "golang.org/x/image/tiff"
	_ "golang.org/x/image/webp"

	"bytes"
	"image"
	"net/http"
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

func ImgEncode(imgBytes []byte, imgFormat imaging.Format, opts ...imaging.EncodeOption) ([]byte, error) {
	reader := BytesReaderPoolGet(imgBytes)
	defer BytesReaderPoolPut(reader)
	img, err := ImgDecode(imgBytes)
	if err != nil {
		return nil, err
	}
	buf := BytesBufferPoolGet()
	defer BytesBufferPoolPut(buf)
	err = imaging.Encode(buf, img, imgFormat, opts...)
	return buf.Bytes(), err
}

func ImgResize(imgBytes []byte) ([]byte, error) {
	imgFormat, err := imaging.FormatFromExtension(strings.Replace(http.DetectContentType(imgBytes[:min(512, len(imgBytes))]), `image/`, ``, 1))
	if err != nil {
		return nil, err
	}
	img, _, err := image.Decode(bytes.NewReader(imgBytes))
	if err != nil {
		return nil, err
	}
	buf := BytesBufferPoolGet()
	defer BytesBufferPoolPut(buf)
	err = imaging.Encode(buf, imaging.Resize(img, 800, 800, imaging.Lanczos), imgFormat)
	return buf.Bytes(), err
} */
