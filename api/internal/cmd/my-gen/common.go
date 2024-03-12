package my_gen

import (
	"fmt"
	"os/exec"

	"github.com/fatih/color"
)

type myGenDataHandleMethod = uint

const (
	ReturnNil      myGenDataHandleMethod = 0  //默认返回空
	ReturnType     myGenDataHandleMethod = 1  //返回根据字段数据类型解析的数据
	ReturnTypeName myGenDataHandleMethod = 2  //返回根据字段命名类型解析的数据
	ReturnUnion    myGenDataHandleMethod = 10 //返回两种类型解析的数据
)

type myGenDataHandler struct {
	Method       myGenDataHandleMethod //根据该字段返回解析的数据
	DataType     []string              //根据字段数据类型解析的数据
	DataTypeName []string              //根据字段命名类型解析的数据
}

func (myGenDataHandlerThis *myGenDataHandler) getData() []string {
	switch myGenDataHandlerThis.Method {
	case ReturnType:
		return myGenDataHandlerThis.DataType
	case ReturnTypeName:
		return myGenDataHandlerThis.DataTypeName
	case ReturnUnion:
		return append(myGenDataHandlerThis.DataType, myGenDataHandlerThis.DataTypeName...)
	default:
		return nil
	}
}

// 执行命令
func command(title string, isOut bool, dir string, name string, arg ...string) {
	command := exec.Command(name, arg...)
	if dir != `` {
		command.Dir = dir
	}
	fmt.Println()
	fmt.Println(color.GreenString(`================` + title + ` 开始================`))
	fmt.Println(`执行命令：` + command.String())
	stdout, _ := command.StdoutPipe()
	command.Start()
	if isOut {
		buf := make([]byte, 1024)
		for {
			n, err := stdout.Read(buf)
			if err != nil {
				break
			}
			if n > 0 {
				fmt.Print(string(buf[:n]))
			}
		}
	} else {
		fmt.Println(`请稍等，命令正在执行中...`)
	}
	command.Wait()
	fmt.Println(color.GreenString(`================` + title + ` 结束================`))
}
