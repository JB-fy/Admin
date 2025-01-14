/* common.go与funcs.go的区别：
common.go：基于当前框架封装的常用函数（与框架耦合）
funcs.go：基于golang封装的常用函数（不与框架耦合） */

package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"errors"
	"os/exec"
	"reflect"

	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"golang.org/x/tools/imports"
)

// 获取服务器外网ip
func GetServerNetworkIp() string {
	cmd := exec.Command(`/bin/bash`, `-c`, `curl -s ifconfig.me`)
	output, _ := cmd.CombinedOutput()
	return string(output)
}

// 获取服务器内网ip
func GetServerLocalIp() string {
	cmd := exec.Command(`/bin/bash`, `-c`, `hostname -I`)
	output, _ := cmd.CombinedOutput()
	return gstr.Trim(string(output))
}

// go文件代码格式化
func GoFileFmt(filePath string) {
	fmtFuc := func(path, content string) string {
		res, err := imports.Process(path, []byte(content), nil)
		if err != nil {
			return content
		}
		return string(res)
	}
	gfile.ReplaceFileFunc(fmtFuc, filePath)
}

// 十进制转其它进制
func DecimalToBase(decimal int, base int) (numStr string /* , err error */) {
	if base < 2 || base > 36 {
		/* err = errors.New(`base必须是2~36进制`)
		return */
		panic(`base必须是2~36进制`)
	}
	if decimal == 0 {
		numStr = `0`
		return
	}
	chars := `0123456789abcdefghijklmnopqrstuvwxyz`
	for decimal > 0 {
		remainder := decimal % base
		numStr = string(chars[remainder]) + numStr
		decimal /= base
	}
	return
}

// 其它进制转十进制
func BaseToDecimal(numStr string, base int) (decimal int /* , err error */) {
	if base < 2 || base > 36 {
		/* err = errors.New(`base必须是2~36进制`)
		return */
		panic(`base必须是2~36进制`)
	}
	if numStr == `` {
		/* err = errors.New(`numStr不是符合base对应进制的字符串`)
		return */
		panic(`numStr不是符合base对应进制的字符串`)
	}
	chars := `0123456789abcdefghijklmnopqrstuvwxyz`
	chars = chars[:base]
	for i := 0; i < len(numStr); i++ {
		// remainder := strings.IndexByte(chars, numStr[i])
		remainder := gstr.PosI(chars, string(numStr[i]))
		if remainder == -1 {
			/* err = errors.New(`numStr不是符合base对应进制的字符串`)
			return */
			panic(`numStr不是符合base对应进制的字符串`)
		}
		decimal = decimal*base + remainder
	}
	return
}

// 从结构体中获取对应字段的值
func GetValueFromStruct(Obj any, name string) (val any) {
	v := reflect.ValueOf(Obj)

	for {
		if v.Kind() != reflect.Ptr {
			break
		}
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return nil
	}

	field := v.FieldByName(name)
	if !field.IsValid() {
		return nil
	}

	val = field.Interface()
	return
}

// PKCS补码（PKCS5，PKCS7通用）
func PKCS5Pad(rawByte []byte, padLen int) []byte {
	rawByteLen := len(rawByte)
	fillLen := padLen - (rawByteLen % padLen)
	fillByte := bytes.Repeat([]byte{byte(fillLen)}, fillLen)
	return append(rawByte, fillByte...)
}

// PKCS补码移除（PKCS5，PKCS7通用）
func PKCS5UnPad(rawByte []byte, padLen int) ([]byte, error) {
	rawByteLen := len(rawByte)
	fillByte := rawByte[rawByteLen-1]
	fillLen := int(fillByte)
	if fillLen > padLen || fillLen == 0 {
		return nil, errors.New(`无效的填充长度：` + gconv.String(fillLen))
	}
	fillPosition := rawByteLen - fillLen
	for _, v := range rawByte[fillPosition:] {
		if v != fillByte {
			return nil, errors.New(`无效的填充位：预期` + gconv.String(fillLen) + `实际` + gconv.String(v))
		}
	}
	return rawByte[:fillPosition], nil
}

// AES加密
func AesEncrypt(rawByte []byte, keyByte []byte, cipherType string, iv ...byte) (cipherByte []byte, err error) {
	block, err := aes.NewCipher(keyByte)
	if err != nil {
		return
	}
	blockSize := block.BlockSize()
	rawByteLen := len(rawByte)
	if rawByteLen%blockSize != 0 {
		err = errors.New(`加密串必须是块大小的整数倍`)
		return
	}
	cipherByte = make([]byte, rawByteLen)
	switch cipherType {
	case `ECB`:
		for i := 0; i < rawByteLen; i += blockSize {
			block.Encrypt(cipherByte[i:i+blockSize], rawByte[i:i+blockSize])
		}
	// case `CBC`:
	default:
		blockMode := cipher.NewCBCEncrypter(block, iv)
		blockMode.CryptBlocks(cipherByte, rawByte)
	}
	return
}

// AES解密
func AesDecrypt(cipherByte []byte, keyByte []byte, cipherType string, iv ...byte) (rawByte []byte, err error) {
	block, err := aes.NewCipher(keyByte)
	if err != nil {
		return
	}
	blockSize := block.BlockSize()
	cipherByteLen := len(cipherByte)
	if cipherByteLen%blockSize != 0 {
		err = errors.New(`解密串必须是块大小的整数倍`)
		return
	}
	rawByte = make([]byte, cipherByteLen)
	switch cipherType {
	case `ECB`:
		for i := 0; i < cipherByteLen; i += blockSize {
			block.Decrypt(rawByte[i:i+blockSize], cipherByte[i:i+blockSize])
		}
	// case `CBC`:
	default:
		blockMode := cipher.NewCBCDecrypter(block, iv)
		blockMode.CryptBlocks(rawByte, cipherByte)
	}
	return
}

/* var imageMimeTypeExtMap = g.MapStrStr{
	`image/x-icon`: `icon`,
	`image/jpeg`:   `jpeg`,
	`image/bmp`:    `bmp`,
	`image/gif`:    `gif`,
	`image/webp`:   `webp`,
	`image/png`:    `png`,
}

// 获取图片后缀
func GetImageExt(imageBytes []byte, defExt string) (imageExt string) {
	imageExt = imageMimeTypeExtMap[http.DetectContentType(imageBytes[:512])]
	if imageExt == `` {
		imageExt = defExt
	}
	return
} */

/* // 转成jpeg图片。目前只支持webp转换
func JpegEncode(imageBytes []byte, quality int, imageTypeArr ...string) ([]byte, error) {
	var err error
	imageType := http.DetectContentType(imageBytes[:512])
	for _, v := range imageTypeArr {
		if imageType != v {
			continue
		}

		var img image.Image
		switch v {
		case `image/webp`:
			img, err = webp.Decode(bytes.NewReader(imageBytes))
		}
		if err != nil {
			return nil, err
		}

		var jpegData bytes.Buffer
		err = jpeg.Encode(&jpegData, img, &jpeg.Options{Quality: quality})
		if err != nil {
			return nil, err
		}
		imageBytes = jpegData.Bytes()
	}
	return imageBytes, err
} */
