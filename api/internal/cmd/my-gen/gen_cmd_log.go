package my_gen

import (
	"strings"

	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

func initCmdLog(option myGenOption, tpl *myGenTpl) {
	saveFileName := option.SceneId
	if !(option.IsApi || option.IsView) {
		saveFileName = `gen_dao`
	}
	tpl.CmdLog.File = gfile.SelfDir() + `/internal/cmd/my-gen/log/` + saveFileName + `.log`
	if gfile.IsFile(tpl.CmdLog.File) {
		tpl.CmdLog.Content = gfile.GetContents(tpl.CmdLog.File)
		myGenCommandArr := []string{
			`./main`,
			`myGen`,
			`-dbGroup=` + option.DbGroup,
			`-dbTable=` + option.DbTable,
		}
		match, _ := gregex.MatchString(`(`+gregex.Quote(strings.Join(myGenCommandArr, ` `)+` `)+`[\s\S]*?)(((\r|\n)\./main)|$)`, tpl.CmdLog.Content)
		if len(match) > 0 {
			tpl.CmdLog.Last = match[1]
		}
	}
}

func genCmdLog(option myGenOption, tpl *myGenTpl) {
	tableCmdLog := []string{}
	tableCmdLog = append(tableCmdLog, tpl.CmdLog.RelId...)
	tableCmdLog = append(tableCmdLog, tpl.CmdLog.Extend...)
	tableCmdLog = append(tableCmdLog, tpl.CmdLog.OtherRel...)

	myGenCommandArr := []string{
		`./main`,
		`myGen`,
		`-dbGroup=` + option.DbGroup,
		`-dbTable=` + option.DbTable,
		`-removePrefixCommon=` + option.RemovePrefixCommon,
		`-removePrefixAlone=` + option.RemovePrefixAlone,
		`-cacheType=` + gconv.String(option.CacheType),
	}
	if option.CacheType != 0 {
		myGenCommandArr = append(myGenCommandArr, `-cacheTime=`+option.CacheTime)
	}
	myGenCommandArr = append(myGenCommandArr, `-isApi=`+gconv.String(gconv.Uint(option.IsApi)))
	if option.IsApi {
		myGenCommandArr = append(myGenCommandArr,
			`-isResetLogic=`+gconv.String(gconv.Uint(option.IsResetLogic)),
			`-isAuthAction=`+gconv.String(gconv.Uint(option.IsAuthAction)),
			`-commonName=`+option.CommonName,
			`-loginRelId=`+option.LoginRelId,
			`-loginIdStr="`+option.LoginIdStr+`"`,
			`-filterIsStop=`+gconv.String(gconv.Uint(option.FilterIsStop)))
	}
	myGenCommandArr = append(myGenCommandArr, `-isView=`+gconv.String(gconv.Uint(option.IsView)))
	if option.IsApi || option.IsView {
		myGenCommandArr = append(myGenCommandArr,
			`-sceneId=`+option.SceneId,
			`-isList=`+gconv.String(gconv.Uint(option.IsList)),
			`-isCount=`+gconv.String(gconv.Uint(option.IsCount)),
			`-isInfo=`+gconv.String(gconv.Uint(option.IsInfo)),
			`-isCreate=`+gconv.String(gconv.Uint(option.IsCreate)),
			`-isUpdate=`+gconv.String(gconv.Uint(option.IsUpdate)),
			`-isDelete=`+gconv.String(gconv.Uint(option.IsDelete)))
	}
	cmdLog := strings.Join(myGenCommandArr, ` `) + gstr.Join(append([]string{``}, tableCmdLog...), `
    `)
	if tpl.CmdLog.Content != `` {
		if tpl.CmdLog.Last != `` {
			cmdLog = gstr.Replace(tpl.CmdLog.Content, tpl.CmdLog.Last, cmdLog)
		} else {
			cmdLog = tpl.CmdLog.Content + "\r\n" + cmdLog
		}
	}
	gfile.PutContents(tpl.CmdLog.File, cmdLog)
}
