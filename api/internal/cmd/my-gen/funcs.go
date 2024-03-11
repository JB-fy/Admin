package my_gen

import (
	"fmt"
	"os/exec"

	"github.com/fatih/color"
)

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
