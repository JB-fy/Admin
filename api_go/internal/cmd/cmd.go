package cmd

import (
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  `main`,
		Brief: `通过这个命令启动`,
	}

	Http = gcmd.Command{
		Name:  `http`,
		Usage: `http`,
		Brief: `http服务`,
		Func:  HttpFunc,
	}

	MyGen = gcmd.Command{
		Name:  `myGen`,
		Usage: `myGen`,
		Brief: `代码自动生成`,
		Func:  MyGenFunc,
	}
)
