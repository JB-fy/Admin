package my_gen

import (
	"api/internal/cmd/my-gen/internal"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
)

type myGenViewI18n struct {
	name   []string
	status []string
	tip    []string
}

type myGenViewI18nField struct {
	name   internal.MyGenDataStrHandler
	status internal.MyGenDataSliceHandler
	tip    internal.MyGenDataStrHandler
}

func (viewI18nThis *myGenViewI18n) Add(viewI18nField myGenViewI18nField, field string) {
	if viewI18nField.name.GetData() != `` {
		viewI18nThis.name = append(viewI18nThis.name, field+`: `+viewI18nField.name.GetData()+`,`)
	}
	if len(viewI18nField.status.GetData()) > 0 {
		viewI18nThis.status = append(viewI18nThis.status, field+`: [`+gstr.Join(append([]string{``}, viewI18nField.status.GetData()...), `
		`)+`
	],`)
	}
	if viewI18nField.tip.GetData() != `` {
		viewI18nThis.tip = append(viewI18nThis.tip, field+`: `+viewI18nField.tip.GetData()+`,`)
	}
}

func (viewI18nThis *myGenViewI18n) Merge(viewI18nOther myGenViewI18n) {
	viewI18nThis.name = append(viewI18nThis.name, viewI18nOther.name...)
	viewI18nThis.status = append(viewI18nThis.status, viewI18nOther.status...)
	viewI18nThis.tip = append(viewI18nThis.tip, viewI18nOther.tip...)
}

func (viewI18nThis *myGenViewI18n) Unique() {
	// viewI18nThis.name = garray.NewStrArrayFrom(viewI18nThis.name).Unique().Slice()
	// viewI18nThis.status = garray.NewStrArrayFrom(viewI18nThis.status).Unique().Slice()
	// viewI18nThis.tip = garray.NewStrArrayFrom(viewI18nThis.tip).Unique().Slice()
}

// 视图模板Query生成
func genViewI18n(option myGenOption, tpl myGenTpl) {
	viewI18n := myGenViewI18n{}
	for _, v := range tpl.FieldListOfDefault {
		viewI18n.Add(getViewI18nField(tpl, v), v.FieldRaw)
	}
	for _, v := range tpl.FieldListOfAfter {
		viewI18n.Add(getViewI18nField(tpl, v), v.FieldRaw)
	}
	for _, v := range tpl.Handle.ExtendTableOneList {
		viewI18n.Merge(getViewI18nExtendMiddleOne(v))
	}
	for _, v := range tpl.Handle.MiddleTableOneList {
		viewI18n.Merge(getViewI18nExtendMiddleOne(v))
	}
	for _, v := range tpl.Handle.ExtendTableManyList {
		viewI18n.Merge(getViewI18nExtendMiddleMany(v))
	}
	for _, v := range tpl.Handle.MiddleTableManyList {
		viewI18n.Merge(getViewI18nExtendMiddleMany(v))
	}
	viewI18n.Unique()

	tplView := `export default {
    name: {` + gstr.Join(append([]string{``}, viewI18n.name...), `
        `) + `
    },
    status: {` + gstr.Join(append([]string{``}, viewI18n.status...), `
        `) + `
    },
    tip: {` + gstr.Join(append([]string{``}, viewI18n.tip...), `
        `) + `
    },
}
`

	saveFile := gfile.SelfDir() + `/../view/` + option.SceneCode + `/src/i18n/language/zh-cn/` + tpl.ModuleDirCaseKebab + `/` + tpl.TableCaseKebab + `.ts`
	gfile.PutContents(saveFile, tplView)
}

func getViewI18nField(tpl myGenTpl, v myGenField) (viewI18nField myGenViewI18nField) {
	viewI18nField.name.Method = internal.ReturnType
	viewI18nField.name.DataType = `'` + v.FieldName + `'`

	/*--------根据字段数据类型处理（注意：这里的代码改动对字段命名类型处理有影响） 开始--------*/
	switch v.FieldType {
	case internal.TypeInt, internal.TypeIntU: // `int等类型` // `int等类型（unsigned）`
	case internal.TypeFloat, internal.TypeFloatU: // `float等类型`  // `float等类型（unsigned）`
	case internal.TypeVarchar, internal.TypeChar: // `varchar类型`	// `char类型`
	case internal.TypeText: // `text类型`
	case internal.TypeJson: // `json类型`
		viewI18nField.tip.Method = internal.ReturnType
		if v.FieldTip != `` {
			viewI18nField.tip.DataType = `'` + v.FieldTip + `'`
		}
	case internal.TypeDatetime, internal.TypeTimestamp: // `datetime类型`	// `timestamp类型`
	case internal.TypeDate: // `date类型`
	case internal.TypeTime: // `time类型`
	}
	/*--------根据字段数据类型处理（注意：这里的代码改动对字段命名类型处理有影响） 结束--------*/

	/*--------根据字段主键类型处理 开始--------*/
	switch v.FieldTypePrimary {
	case internal.TypePrimary: // 独立主键
	case internal.TypePrimaryAutoInc: // 独立主键（自增）
		return myGenViewI18nField{}
	case internal.TypePrimaryMany: // 联合主键
	case internal.TypePrimaryManyAutoInc: // 联合主键（自增）
	}
	/*--------根据字段主键类型处理 结束--------*/

	/*--------根据字段命名类型处理 开始--------*/
	switch v.FieldTypeName {
	case internal.TypeNameDeleted: // 软删除字段
		return myGenViewI18nField{}
	case internal.TypeNameUpdated: // 更新时间字段
		return myGenViewI18nField{}
	case internal.TypeNameCreated: // 创建时间字段
		return myGenViewI18nField{}
	case internal.TypeNamePid: // pid；	类型：int等类型；
		viewI18nField.name.Method = internal.ReturnTypeName
		viewI18nField.name.DataTypeName = `'父级'`
	case internal.TypeNameLevel: // level，且pid,level,idPath|id_path同时存在时（才）有效；	类型：int等类型；
	case internal.TypeNameIdPath: // idPath|id_path，且pid,level,idPath|id_path同时存在时（才）有效；	类型：varchar或text；
	case internal.TypeNamePasswordSuffix: // password,passwd后缀；	类型：char(32)；
	case internal.TypeNameSaltSuffix: // salt后缀，且对应的password,passwd后缀存在时（才）有效；	类型：char；
		return myGenViewI18nField{}
	case internal.TypeNameNameSuffix: // name,title后缀；	类型：varchar；
	case internal.TypeNameCodeSuffix: // code后缀；	类型：varchar；
	case internal.TypeNameAccountSuffix: // account后缀；	类型：varchar；
	case internal.TypeNamePhoneSuffix: // phone,mobile后缀；	类型：varchar；
	case internal.TypeNameEmailSuffix: // email后缀；	类型：varchar；
	case internal.TypeNameUrlSuffix: // url,link后缀；	类型：varchar；
	case internal.TypeNameIpSuffix: // IP后缀；	类型：varchar；
	case internal.TypeNameIdSuffix: // id后缀；	类型：int等类型；
		relIdObj := tpl.Handle.RelIdMap[v.FieldRaw]
		if relIdObj.tpl.Table != `` && !relIdObj.IsRedundName {
			viewI18nField.name.Method = internal.ReturnTypeName
			viewI18nField.name.DataTypeName = `'` + relIdObj.FieldName + `'`
		}
	case internal.TypeNameSortSuffix: // sort,num,number,weight,level,rank等后缀；	类型：int等类型；
		viewI18nField.tip.Method = internal.ReturnTypeName
		viewI18nField.tip.DataTypeName = `'` + v.FieldTip + `'`
	case internal.TypeNameStatusSuffix: // status,type,method,pos,position,gender等后缀；	类型：int等类型或varchar或char；	注释：多状态之间用[\s,，;；]等字符分隔。示例（状态：0待处理 1已处理 2驳回 yes是 no否）
		viewI18nField.status.Method = internal.ReturnTypeName
		if garray.NewIntArrayFrom([]int{internal.TypeVarchar, internal.TypeChar}).Contains(v.FieldType) {
			for _, status := range v.StatusList {
				viewI18nField.status.DataTypeName = append(viewI18nField.status.DataTypeName, `{ value: '`+status[0]+`', label: '`+status[1]+`' },`)
			}
		} else {
			for _, status := range v.StatusList {
				viewI18nField.status.DataTypeName = append(viewI18nField.status.DataTypeName, `{ value: `+status[0]+`, label: '`+status[1]+`' },`)
			}
		}
	case internal.TypeNameIsPrefix: // is_前缀；	类型：int等类型；注释：多状态之间用[\s,，;；]等字符分隔。示例（停用：0否 1是）
	case internal.TypeNameStartPrefix: // start_前缀；	类型：datetime或date或timestamp或time；
	case internal.TypeNameEndPrefix: // end_前缀；	类型：datetime或date或timestamp或time；
	case internal.TypeNameRemarkSuffix: // remark,desc,msg,message,intro,content后缀；	类型：varchar或text；前端对应组件：varchar文本输入框，text富文本编辑器
	case internal.TypeNameImageSuffix, internal.TypeNameVideoSuffix, internal.TypeNameFileSuffix: // icon,cover,avatar,img,img_list,imgList,img_arr,imgArr,image,image_list,imageList,image_arr,imageArr等后缀；	类型：单图片varchar，多图片json或text	// video,video_list,videoList,video_arr,videoArr等后缀；	类型：单视频varchar，多视频json或text	// file,file_list,fileList,file_arr,fileArr等后缀；	类型：单文件varchar，多文件json或text
	case internal.TypeNameArrSuffix: // list,arr等后缀；	类型：json或text；
	}
	/*--------根据字段命名类型处理 结束--------*/
	return
}

func getViewI18nExtendMiddleOne(tplEM handleExtendMiddle) (viewI18n myGenViewI18n) {
	switch tplEM.TableType {
	case internal.TableTypeExtendOne:
		for _, v := range tplEM.FieldList {
			viewI18n.Add(getViewI18nField(tplEM.tpl, v), v.FieldRaw)
		}
	case internal.TableTypeMiddleOne:
		for _, v := range tplEM.FieldListOfIdSuffix {
			viewI18n.Add(getViewI18nField(tplEM.tpl, v), v.FieldRaw)
		}
		for _, v := range tplEM.FieldListOfOther {
			viewI18n.Add(getViewI18nField(tplEM.tpl, v), v.FieldRaw)
		}
	}
	return
}

func getViewI18nExtendMiddleMany(tplEM handleExtendMiddle) (viewI18n myGenViewI18n) {
	if len(tplEM.FieldList) == 1 {
		v := tplEM.FieldList[0]

		viewI18nField := myGenViewI18nField{}
		viewI18nField.name.Method = internal.ReturnType
		viewI18nField.name.DataType = `'` + v.FieldName /*  + `列表` */ + `'`
		if v.FieldTypeName == internal.TypeNameIdSuffix && gstr.ToUpper(gstr.SubStr(v.FieldName, -2)) == `ID` {
			viewI18nField.name.DataType = `'` + gstr.SubStr(v.FieldName, 0, -2) /*  + `列表` */ + `'`
		}

		/*--------根据字段数据类型处理（注意：这里的代码改动对字段命名类型处理有影响） 开始--------*/
		switch v.FieldType {
		case internal.TypeInt, internal.TypeIntU: // `int等类型` // `int等类型（unsigned）`
		case internal.TypeFloat, internal.TypeFloatU: // `float等类型`  // `float等类型（unsigned）`
		case internal.TypeVarchar, internal.TypeChar: // `varchar类型`	// `char类型`
		case internal.TypeText: // `text类型`
		case internal.TypeJson: // `json类型`
			viewI18nField.tip.Method = internal.ReturnType
			if v.FieldTip != `` {
				viewI18nField.tip.DataType = `'` + v.FieldTip + `'`
			}
		case internal.TypeDatetime, internal.TypeTimestamp: // `datetime类型`	// `timestamp类型`
		case internal.TypeDate: // `date类型`
		case internal.TypeTime: // `time类型`
		}
		/*--------根据字段数据类型处理（注意：这里的代码改动对字段命名类型处理有影响） 结束--------*/

		/*--------根据字段命名类型处理 开始--------*/
		switch v.FieldTypeName {
		case internal.TypeNameDeleted, internal.TypeNameUpdated, internal.TypeNameCreated: // 软删除字段 // 更新时间字段 // 创建时间字段
		case internal.TypeNamePid: // pid；	类型：int等类型；
		case internal.TypeNameLevel: // level，且pid,level,idPath|id_path同时存在时（才）有效；	类型：int等类型；
		case internal.TypeNameIdPath: // idPath|id_path，且pid,level,idPath|id_path同时存在时（才）有效；	类型：varchar或text；
		case internal.TypeNamePasswordSuffix: // password,passwd后缀；	类型：char(32)；
		case internal.TypeNameSaltSuffix: // salt后缀，且对应的password,passwd后缀存在时（才）有效；	类型：char；
		case internal.TypeNameNameSuffix: // name,title后缀；	类型：varchar；
		case internal.TypeNameCodeSuffix: // code后缀；	类型：varchar；
		case internal.TypeNameAccountSuffix: // account后缀；	类型：varchar；
		case internal.TypeNamePhoneSuffix: // phone,mobile后缀；	类型：varchar；
		case internal.TypeNameEmailSuffix: // email后缀；	类型：varchar；
		case internal.TypeNameUrlSuffix: // url,link后缀；	类型：varchar；
		case internal.TypeNameIpSuffix: // IP后缀；	类型：varchar；
		case internal.TypeNameIdSuffix: // id后缀；	类型：int等类型；
		case internal.TypeNameSortSuffix: // sort,num,number,weight,level,rank等后缀；	类型：int等类型；
			if v.FieldTip != `` {
				viewI18nField.tip.Method = internal.ReturnTypeName
				viewI18nField.tip.DataTypeName = `'` + v.FieldTip + `'`
			}
		case internal.TypeNameStatusSuffix: // status,type,method,pos,position,gender等后缀；	类型：int等类型或varchar或char；	注释：多状态之间用[\s,，;；]等字符分隔。示例（状态：0待处理 1已处理 2驳回 yes是 no否）
			viewI18nField.status.Method = internal.ReturnTypeName
			if garray.NewIntArrayFrom([]int{internal.TypeVarchar, internal.TypeChar}).Contains(v.FieldType) {
				for _, status := range v.StatusList {
					viewI18nField.status.DataTypeName = append(viewI18nField.status.DataTypeName, `{ value: '`+status[0]+`', label: '`+status[1]+`' },`)
				}
			} else {
				for _, status := range v.StatusList {
					viewI18nField.status.DataTypeName = append(viewI18nField.status.DataTypeName, `{ value: `+status[0]+`, label: '`+status[1]+`' },`)
				}
			}
		case internal.TypeNameIsPrefix: // is_前缀；	类型：int等类型；注释：多状态之间用[\s,，;；]等字符分隔。示例（停用：0否 1是）
		case internal.TypeNameStartPrefix: // start_前缀；	类型：datetime或date或timestamp或time；
		case internal.TypeNameEndPrefix: // end_前缀；	类型：datetime或date或timestamp或time；
		case internal.TypeNameRemarkSuffix: // remark,desc,msg,message,intro,content后缀；	类型：varchar或text；前端对应组件：varchar文本输入框，text富文本编辑器
		case internal.TypeNameImageSuffix, internal.TypeNameVideoSuffix, internal.TypeNameFileSuffix: // icon,cover,avatar,img,img_list,imgList,img_arr,imgArr,image,image_list,imageList,image_arr,imageArr等后缀；	类型：单图片varchar，多图片json或text	// video,video_list,videoList,video_arr,videoArr等后缀；	类型：单视频varchar，多视频json或text	// file,file_list,fileList,file_arr,fileArr等后缀；	类型：单文件varchar，多文件json或text
		case internal.TypeNameArrSuffix: // list,arr等后缀；	类型：json或text；
		}
		/*--------根据字段命名类型处理 结束--------*/

		viewI18n.Add(viewI18nField, tplEM.FieldVar)
	} else {
		viewI18nTmp := myGenViewI18n{}
		for _, v := range tplEM.FieldListOfIdSuffix {
			viewI18nTmp.Add(getViewI18nField(tplEM.tpl, v), v.FieldRaw)
		}
		for _, v := range tplEM.FieldListOfOther {
			viewI18nTmp.Add(getViewI18nField(tplEM.tpl, v), v.FieldRaw)
		}

		viewI18n.name = append(viewI18n.name, internal.GetStrByFieldStyle(tplEM.tplOfTop.FieldStyle, tplEM.FieldVar, ``, `label`)+`: '列表',`)
		viewI18n.name = append(viewI18n.name, tplEM.FieldVar+`: {`+gstr.Join(append([]string{``}, viewI18nTmp.name...), `
            `)+`
        },`)
		/* viewI18n.name = append(viewI18n.name, tplEM.FieldVar+`: {
		    `+tplEM.FieldVar+`: '列表',`+gstr.Join(append([]string{``}, viewI18nTmp.name...), `
		    `)+`
		},`) */
		viewI18n.status = append(viewI18n.status, tplEM.FieldVar+`: {`+gstr.Join(append([]string{``}, viewI18nTmp.status...), `
            `)+`
        },`)
		viewI18n.tip = append(viewI18n.tip, tplEM.FieldVar+`: {`+gstr.Join(append([]string{``}, viewI18nTmp.tip...), `
            `)+`
        },`)
	}
	return
}
