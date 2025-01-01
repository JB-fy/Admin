/* common.go与funcs.go的区别：
common.go：基于当前框架封装的常用函数（与框架耦合）
funcs.go：基于golang封装的常用函数（不与框架耦合） */

package utils

import (
	"bytes"
	"crypto/aes"
	"errors"
	"os/exec"
	"reflect"

	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
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

// aes加密（ECB模式，PKCS5补码）
func AesEncrypt(rawStr string, key string) (cipherByte []byte, err error) {
	keyByte := []byte(key)
	block, err := aes.NewCipher(keyByte)
	if err != nil {
		return
	}

	rawStrByte := []byte(rawStr)

	blockSize := block.BlockSize()
	fillLen := blockSize - (len(rawStrByte) % blockSize)
	fillByte := bytes.Repeat([]byte{byte(fillLen)}, fillLen)
	rawStrByte = append(rawStrByte, fillByte...)

	cipherByteTmp := make([]byte, blockSize)
	for i := 0; i < len(rawStrByte); i += blockSize {
		block.Encrypt(cipherByteTmp, rawStrByte[i:i+blockSize])
		cipherByte = append(cipherByte, cipherByteTmp...)
	}
	return
}

// aes解密（ECB模式，PKCS5补码）
func AesDecrypt(cipherByte []byte, key string) (rawStr string, err error) {
	keyByte := []byte(key)
	block, err := aes.NewCipher(keyByte)
	if err != nil {
		return
	}

	blockSize := block.BlockSize()
	if len(cipherByte)%blockSize != 0 {
		err = errors.New(`加密串必须是块大小的整数倍`)
		return
	}

	rawStrByte := []byte{}
	rawStrByteTmp := make([]byte, blockSize)
	for i := 0; i < len(cipherByte); i += blockSize {
		block.Decrypt(rawStrByteTmp, cipherByte[i:i+blockSize])
		rawStrByte = append(rawStrByte, rawStrByteTmp...)
	}

	fillLen := int(rawStrByte[len(rawStrByte)-1])
	rawStr = string(rawStrByte[:(len(rawStrByte) - fillLen)])
	return
}
