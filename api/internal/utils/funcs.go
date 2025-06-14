/* common.go与funcs.go的区别：
common.go：基于当前框架封装的常用函数（与框架耦合）
funcs.go：基于golang封装的常用函数（不与框架耦合） */

package utils

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"os/exec"
	"reflect"
	"runtime"
	"strings"
	"sync"

	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
	"golang.org/x/tools/imports"
)

var bytesBufferPool = sync.Pool{
	New: func() any {
		return new(bytes.Buffer)
	},
}

// 返回值必须在使用完成后，才可以放回对象池。注意：因为buf.Bytes()返回的是底层切片引用，故当buf.Bytes()的返回值还再使用时，也不能将buf放回对象池
func BytesBufferPoolGet() *bytes.Buffer {
	buf := bytesBufferPool.Get().(*bytes.Buffer)
	buf.Reset()
	return buf
}

func BytesBufferPoolPut(buf *bytes.Buffer) {
	bytesBufferPool.Put(buf)
}

var bytesReaderPool = sync.Pool{
	New: func() any {
		return bytes.NewReader(nil)
	},
}

// 返回值必须在使用完成后，才可以放回对象池
func BytesReaderPoolGet(b []byte) *bytes.Reader {
	reader := bytesReaderPool.Get().(*bytes.Reader)
	reader.Reset(b)
	return reader
}

func BytesReaderPoolPut(reader *bytes.Reader) {
	bytesReaderPool.Put(reader)
}

var stringsBuilderPool = sync.Pool{
	New: func() any {
		return new(strings.Builder)
	},
}

// 返回值必须在使用完成后，才可以放回对象池
func StringsBuilderPoolGet() *strings.Builder {
	builder := stringsBuilderPool.Get().(*strings.Builder)
	builder.Reset()
	return builder
}

func StringsBuilderPoolPut(builder *strings.Builder) {
	stringsBuilderPool.Put(builder)
}

// 获取服务器外网ip
func GetServerNetworkIp() string {
	for _, v := range []string{`ifconfig.me`, `https://ipinfo.io/ip`, `https://checkip.amazonaws.com`, `https://icanhazip.com`, `https://api.ipify.org`} {
		cmd := exec.Command(`/bin/bash`, `-c`, `curl -s --max-time 3 `+v)
		output, _ := cmd.CombinedOutput()
		if ip := string(output); ip != `` {
			return ip
		}
	}
	panic(`获取外网IP失败`)
}

// 获取服务器内网ip
func GetServerLocalIp() string {
	cmd := exec.Command(`/bin/bash`, `-c`, `hostname -I | awk '{printf "%s", $1}'`)
	output, _ := cmd.CombinedOutput()
	if ip := string(output); ip != `` {
		return ip
	}
	panic(`获取内网IP失败`)
}

// 获取调用该函数的上层第几个函数的函数名
func GetMethodName(skip int) (methodName string) {
	if pc, _, _, ok := runtime.Caller(skip); ok {
		fullName := runtime.FuncForPC(pc).Name()
		parts := strings.Split(fullName, `.`)
		methodName = parts[len(parts)-1]
	}
	return
}

// 文件代码格式化
func FilePutFormat(filePath string, src ...byte) (err error) {
	contentFormat, err := imports.Process(filePath, src, nil)
	if err != nil {
		if src == nil {
			return
		}
		contentFormat = src
	}
	return gfile.PutBytes(filePath, contentFormat)
}

// 逐行读取文件内容。框架gfile.ReadLines()方法在行数据超过默认的缓冲区大小（一般4KB）时，scanner.Scan()这行代码会返回false中断执行
func FileReadLine(filePath string, callback func(line []byte) error) (err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	var line []byte
	var isPrefix bool
	var fullLine []byte
	for {
		line, isPrefix, err = reader.ReadLine()
		if err != nil && err != bufio.ErrBufferFull {
			if err == io.EOF {
				return nil
			}
			return
		}
		fullLine = append(fullLine, line...)
		if isPrefix {
			continue
		}
		if err = callback(fullLine); err != nil {
			return
		}
		fullLine = fullLine[:0]
	}
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
	for i := range len(numStr) {
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
