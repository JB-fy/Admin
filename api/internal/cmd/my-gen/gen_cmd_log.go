package my_gen

import (
	"context"
	"strings"

	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

type cmdLog struct {
	File     string   //文件路径
	Content  string   //日志
	Last     string   //上一次日志
	RelId    []string //id后缀字段关联表日志
	Extend   []string //扩展表日志
	OtherRel []string //其它关联表日志
}

func createCmdLog(_ context.Context, option *myGenOption) (log *cmdLog) {
	saveFileName := option.SceneId
	if !(option.IsApi || option.IsView) {
		saveFileName = `gen_dao`
	}
	log = &cmdLog{}
	log.File = gfile.SelfDir() + `/internal/cmd/my-gen/log/` + saveFileName + `.log`
	if gfile.IsFile(log.File) {
		log.Content = gfile.GetContents(log.File)
		myGenCommandArr := []string{
			`./main`,
			`myGen`,
			`-dbGroup=` + option.DbGroup,
			`-dbTable=` + option.DbTable,
		}
		match, _ := gregex.MatchString(`(`+gregex.Quote(strings.Join(myGenCommandArr, ` `)+` `)+`[\s\S]*?)(((\r|\n)\./main)|$)`, log.Content)
		if len(match) > 0 {
			log.Last = match[1]
		}
	}
	return
}

func genCmdLog(_ context.Context, tpl *myGenTpl) {
	myGenCommandArr := []string{
		`./main`,
		`myGen`,
		`-dbGroup=` + tpl.Option.DbGroup,
		`-dbTable=` + tpl.Option.DbTable,
		`-removePrefixCommon=` + tpl.Option.RemovePrefixCommon,
		`-removePrefixAlone=` + tpl.Option.RemovePrefixAlone,
		`-cacheType=` + gconv.String(tpl.Option.CacheType),
	}
	if tpl.Option.CacheType != 0 {
		myGenCommandArr = append(myGenCommandArr, `-cacheTime=`+tpl.Option.CacheTime)
	}
	myGenCommandArr = append(myGenCommandArr, `-isApi=`+gconv.String(gconv.Uint(tpl.Option.IsApi)))
	if tpl.Option.IsApi {
		myGenCommandArr = append(myGenCommandArr,
			`-isResetLogic=`+gconv.String(gconv.Uint(tpl.Option.IsResetLogic)),
			`-isAuthAction=`+gconv.String(gconv.Uint(tpl.Option.IsAuthAction)),
			`-commonName=`+tpl.Option.CommonName,
			`-loginRelId=`+tpl.Option.LoginRelId,
			`-loginIdStr="`+tpl.Option.LoginIdStr+`"`,
			`-filterIsStop=`+gconv.String(gconv.Uint(tpl.Option.FilterIsStop)))
	}
	myGenCommandArr = append(myGenCommandArr, `-isView=`+gconv.String(gconv.Uint(tpl.Option.IsView)))
	if tpl.Option.IsApi || tpl.Option.IsView {
		myGenCommandArr = append(myGenCommandArr,
			`-sceneId=`+tpl.Option.SceneId,
			`-isList=`+gconv.String(gconv.Uint(tpl.Option.IsList)),
			`-isCount=`+gconv.String(gconv.Uint(tpl.Option.IsCount)),
			`-isInfo=`+gconv.String(gconv.Uint(tpl.Option.IsInfo)),
			`-isCreate=`+gconv.String(gconv.Uint(tpl.Option.IsCreate)),
			`-isUpdate=`+gconv.String(gconv.Uint(tpl.Option.IsUpdate)),
			`-isDelete=`+gconv.String(gconv.Uint(tpl.Option.IsDelete)))
	}

	contentArr := []string{``}
	contentArr = append(contentArr, tpl.CmdLog.RelId...)
	contentArr = append(contentArr, tpl.CmdLog.Extend...)
	contentArr = append(contentArr, tpl.CmdLog.OtherRel...)
	content := strings.Join(myGenCommandArr, ` `) + gstr.Join(contentArr, `
    `)
	if tpl.CmdLog.Content != `` {
		if tpl.CmdLog.Last != `` {
			content = gstr.Replace(tpl.CmdLog.Content, tpl.CmdLog.Last, content)
		} else {
			content = tpl.CmdLog.Content + "\r\n" + content
		}
	}
	gfile.PutContents(tpl.CmdLog.File, content)
}
