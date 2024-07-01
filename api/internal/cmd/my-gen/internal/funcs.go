package internal

import (
	"fmt"
	"math"
	"os/exec"

	"github.com/fatih/color"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

// 执行命令
func Command(title string, isOut bool, dir string, name string, arg ...string) {
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

// 获取Handle.PasswordMap的Key（以Password为主）
func GetHandlePasswordMapKey(passwordOrsalt string) (passwordMapKey string) {
	passwordOrsalt = gstr.Replace(gstr.CaseCamel(passwordOrsalt), `Salt`, `Password`, 1) //替换salt
	passwordOrsalt = gstr.Replace(passwordOrsalt, `Passwd`, `Password`, 1)               //替换passwd
	passwordMapKey = gstr.CaseCamelLower(passwordOrsalt)
	return
}

// status字段注释解析
func GetStatusList(comment string, isStr bool) (statusList [][2]string) {
	var tmp [][]string
	if isStr {
		tmp, _ = gregex.MatchAllString(`([A-Za-z0-9]+)[-=:：]?([^\s,，.。;；]+)`, comment)
	} else {
		// tmp, _ = gregex.MatchAllString(`(-?\d+)[-=:：]?([^\d\s,，.。;；]+)`, comment)
		tmp, _ = gregex.MatchAllString(`(-?\d+)[-=:：]?([^\s,，.。;；]+)`, comment)
	}

	if len(tmp) == 0 {
		statusList = [][2]string{{`0`, `请设置表字段注释后，再生成代码`}}
		return
	}
	statusList = make([][2]string, len(tmp))
	for k, v := range tmp {
		statusList[k] = [2]string{v[1], v[2]}
	}
	return
}

// 获取显示长度。汉字个数 + (其它字符个数 / 2) 后的值
func GetShowLen(str string) int {
	len := len(str)
	lenRune := gstr.LenRune(str)
	countHan := (len - lenRune) / 2
	countOther := gconv.Int(math.Ceil(float64(len-countHan*3) / 2))
	return countHan + countOther
}
