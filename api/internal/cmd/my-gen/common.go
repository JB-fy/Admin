package my_gen

import (
	"fmt"
	"os/exec"

	"github.com/fatih/color"
)

type myGenTableType = uint
type myGenFieldTypePrimary = string
type myGenFieldType = int
type myGenFieldTypeName = string
type myGenDataHandleMethod = uint

const (
	TableTypeExtendOne  myGenTableType = 1  //扩展表（一对一）
	TableTypeExtendMany myGenTableType = 2  //扩展表（一对多）
	TableTypeMiddleOne  myGenTableType = 11 //中间表（一对一）
	TableTypeMiddleMany myGenTableType = 12 //中间表（一对多）

	//用于结构体中，需从1开始，否则结构体会默认0
	TypeInt       myGenFieldType = iota + 1 // `int等类型`
	TypeIntU                                // `int等类型（unsigned）`
	TypeFloat                               // `float等类型`
	TypeFloatU                              // `float等类型（unsigned）`
	TypeVarchar                             // `varchar类型`
	TypeChar                                // `char类型`
	TypeText                                // `text类型`
	TypeJson                                // `json类型`
	TypeTimestamp                           // `timestamp类型`
	TypeDatetime                            // `datetime类型`
	TypeDate                                // `date类型`

	TypePrimary            myGenFieldTypePrimary = `独立主键`
	TypePrimaryAutoInc     myGenFieldTypePrimary = `独立主键（自增）`
	TypePrimaryMany        myGenFieldTypePrimary = `联合主键`
	TypePrimaryManyAutoInc myGenFieldTypePrimary = `联合主键（自增）`

	TypeNameDeleted        myGenFieldTypeName = `软删除字段`
	TypeNameUpdated        myGenFieldTypeName = `更新时间字段`
	TypeNameCreated        myGenFieldTypeName = `创建时间字段`
	TypeNamePid            myGenFieldTypeName = `命名：pid；	类型：int等类型；`
	TypeNameLevel          myGenFieldTypeName = `命名：level，且pid,level,idPath|id_path同时存在时（才）有效；	类型：int等类型；`
	TypeNameIdPath         myGenFieldTypeName = `命名：idPath|id_path，且pid,level,idPath|id_path同时存在时（才）有效；	类型：varchar或text；`
	TypeNameSort           myGenFieldTypeName = `命名：sort，且pid,level,idPath|id_path,sort同时存在时（才）有效；	类型：int等类型；`
	TypeNamePasswordSuffix myGenFieldTypeName = `命名：password,passwd后缀；		类型：char(32)；`
	TypeNameSaltSuffix     myGenFieldTypeName = `命名：salt后缀，且对应的password,passwd后缀存在时（才）有效；	类型：char；`
	TypeNameNameSuffix     myGenFieldTypeName = `命名：name,title后缀；	类型：varchar；`
	TypeNameCodeSuffix     myGenFieldTypeName = `命名：code后缀；	类型：varchar；`
	TypeNameAccountSuffix  myGenFieldTypeName = `命名：account后缀；	类型：varchar；`
	TypeNamePhoneSuffix    myGenFieldTypeName = `命名：phone,mobile后缀；	类型：varchar；`
	TypeNameEmailSuffix    myGenFieldTypeName = `命名：email后缀；	类型：varchar；`
	TypeNameUrlSuffix      myGenFieldTypeName = `命名：url,link后缀；	类型：varchar；`
	TypeNameIpSuffix       myGenFieldTypeName = `命名：IP后缀；	类型：varchar；`
	TypeNameIdSuffix       myGenFieldTypeName = `命名：id后缀；	类型：int等类型；`
	TypeNameSortSuffix     myGenFieldTypeName = `命名：sort,weight等后缀；	类型：int等类型；`
	TypeNameStatusSuffix   myGenFieldTypeName = `命名：status,type,method,pos,position,gender等后缀；	类型：int等类型或varchar或char；	注释：多状态之间用[\s,，;；]等字符分隔。示例（状态：0待处理 1已处理 2驳回 yes是 no否）`
	TypeNameIsPrefix       myGenFieldTypeName = `命名：is_前缀；		类型：int等类型；注释：多状态之间用[\s,，;；]等字符分隔。示例（停用：0否 1是）`
	TypeNameStartPrefix    myGenFieldTypeName = `命名：start_前缀；	类型：timestamp或datetime或date；`
	TypeNameEndPrefix      myGenFieldTypeName = `命名：end_前缀；	类型：timestamp或datetime或date；`
	TypeNameRemarkSuffix   myGenFieldTypeName = `命名：remark,desc,msg,message,intro,content后缀；	类型：varchar或text；前端对应组件：varchar文本输入框，text富文本编辑器`
	TypeNameImageSuffix    myGenFieldTypeName = `命名：icon,cover,avatar,img,img_list,imgList,img_arr,imgArr,image,image_list,imageList,image_arr,imageArr等后缀；	类型：单图片varchar，多图片json或text`
	TypeNameVideoSuffix    myGenFieldTypeName = `命名：video,video_list,videoList,video_arr,videoArr等后缀；		类型：单视频varchar，多视频json或text`
	TypeNameArrSuffix      myGenFieldTypeName = `命名：list,arr等后缀；	类型：json或text；`

	ReturnEmpty    myGenDataHandleMethod = 0  //默认返回空
	ReturnType     myGenDataHandleMethod = 1  //返回根据字段数据类型解析的数据
	ReturnTypeName myGenDataHandleMethod = 2  //返回根据字段命名类型解析的数据
	ReturnUnion    myGenDataHandleMethod = 10 //返回两种类型解析的数据
)

type myGenDataSliceHandler struct {
	Method       myGenDataHandleMethod //根据该字段返回解析的数据
	DataType     []string              //根据字段数据类型解析的数据
	DataTypeName []string              //根据字段命名类型解析的数据
}

func (myGenDataSliceHandlerThis *myGenDataSliceHandler) getData() []string {
	switch myGenDataSliceHandlerThis.Method {
	case ReturnType:
		return myGenDataSliceHandlerThis.DataType
	case ReturnTypeName:
		return myGenDataSliceHandlerThis.DataTypeName
	case ReturnUnion:
		return append(myGenDataSliceHandlerThis.DataType, myGenDataSliceHandlerThis.DataTypeName...)
	default:
		return nil
	}
}

type myGenDataStrHandler struct {
	Method       myGenDataHandleMethod //根据该字段返回解析的数据
	DataType     string                //根据字段数据类型解析的数据
	DataTypeName string                //根据字段命名类型解析的数据
}

func (myGenDataStrHandlerThis *myGenDataStrHandler) getData() string {
	switch myGenDataStrHandlerThis.Method {
	case ReturnType:
		return myGenDataStrHandlerThis.DataType
	case ReturnTypeName:
		return myGenDataStrHandlerThis.DataTypeName
	case ReturnUnion:
		return myGenDataStrHandlerThis.DataType + myGenDataStrHandlerThis.DataTypeName
	default:
		return ``
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
