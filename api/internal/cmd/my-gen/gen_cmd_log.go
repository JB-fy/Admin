package my_gen

import (
	"context"
	"strings"

	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

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
	contentArr = append(contentArr, tpl.Option.CmdLog.RelId...)
	contentArr = append(contentArr, tpl.Option.CmdLog.Extend...)
	contentArr = append(contentArr, tpl.Option.CmdLog.OtherRel...)
	content := strings.Join(myGenCommandArr, ` `) + gstr.Join(contentArr, `
    `)
	if tpl.Option.CmdLog.Content != `` {
		if tpl.Option.CmdLog.Last != `` {
			content = gstr.Replace(tpl.Option.CmdLog.Content, tpl.Option.CmdLog.Last, content)
		} else {
			content = tpl.Option.CmdLog.Content + "\r\n" + content
		}
	}
	gfile.PutContents(tpl.Option.CmdLog.File, content)
}
