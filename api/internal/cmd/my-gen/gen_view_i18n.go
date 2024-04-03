package my_gen

import (
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
	name   myGenDataStrHandler
	status myGenDataSliceHandler
	tip    myGenDataStrHandler
}

func (viewI18nThis *myGenViewI18n) Add(viewI18nField myGenViewI18nField, field string) {
	if viewI18nField.name.getData() != `` {
		viewI18nThis.name = append(viewI18nThis.name, field+`: `+viewI18nField.name.getData()+`,`)
	}
	if len(viewI18nField.status.getData()) > 0 {
		viewI18nThis.status = append(viewI18nThis.status, field+`: [`+gstr.Join(append([]string{``}, viewI18nField.status.getData()...), `
		`)+`
	],`)
	}
	if viewI18nField.tip.getData() != `` {
		viewI18nThis.tip = append(viewI18nThis.tip, field+`: `+viewI18nField.tip.getData()+`,`)
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
	viewI18n := getViewI18nFieldList(tpl, tpl.FieldArr...)
	for _, v := range tpl.FieldArrAfter {
		viewI18n.Merge(getViewI18nFieldList(tpl, v))
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

func getViewI18nFieldList(tpl myGenTpl, fieldArr ...string) (viewI18n myGenViewI18n) {
	for _, v := range tpl.FieldList {
		if len(fieldArr) > 0 && !garray.NewStrArrayFrom(fieldArr).Contains(v.FieldRaw) {
			continue
		}

		viewI18nField := myGenViewI18nField{}
		viewI18nField.name.Method = ReturnType
		viewI18nField.name.DataType = `'` + v.FieldName + `'`

		/*--------根据字段数据类型处理（注意：这里的代码改动对字段命名类型处理有影响） 开始--------*/
		switch v.FieldType {
		case TypeInt: // `int等类型`
		case TypeIntU: // `int等类型（unsigned）`
		case TypeFloat: // `float等类型`
		case TypeFloatU: // `float等类型（unsigned）`
		case TypeVarchar: // `varchar类型`
		case TypeChar: // `char类型`
		case TypeText: // `text类型`
		case TypeJson: // `json类型`
			viewI18nField.tip.Method = ReturnType
			if v.FieldTip != `` {
				viewI18nField.tip.DataType = `'` + v.FieldTip + `'`
			}
		case TypeTimestamp: // `timestamp类型`
		case TypeDatetime: // `datetime类型`
		case TypeDate: // `date类型`
		}
		/*--------根据字段数据类型处理（注意：这里的代码改动对字段命名类型处理有影响） 结束--------*/

		/*--------根据字段主键类型处理 开始--------*/
		switch v.FieldTypePrimary {
		case TypePrimary: // 独立主键
		case TypePrimaryAutoInc: // 独立主键（自增）
			continue
		case TypePrimaryMany: // 联合主键
		case TypePrimaryManyAutoInc: // 联合主键（自增）
		}
		/*--------根据字段主键类型处理 结束--------*/

		/*--------根据字段命名类型处理 开始--------*/
		switch v.FieldTypeName {
		case TypeNameDeleted: // 软删除字段
			continue
		case TypeNameUpdated: // 更新时间字段
			continue
		case TypeNameCreated: // 创建时间字段
			continue
		case TypeNamePid: // pid；	类型：int等类型；
			viewI18nField.name.Method = ReturnTypeName
			viewI18nField.name.DataTypeName = `'父级'`
		case TypeNameLevel: // level，且pid,level,idPath|id_path同时存在时（才）有效；	类型：int等类型；
		case TypeNameIdPath: // idPath|id_path，且pid,level,idPath|id_path同时存在时（才）有效；	类型：varchar或text；
		case TypeNamePasswordSuffix: // password,passwd后缀；		类型：char(32)；
		case TypeNameSaltSuffix: // salt后缀，且对应的password,passwd后缀存在时（才）有效；	类型：char；
			continue
		case TypeNameNameSuffix: // name,title后缀；	类型：varchar；
		case TypeNameCodeSuffix: // code后缀；	类型：varchar；
		case TypeNameAccountSuffix: // account后缀；	类型：varchar；
		case TypeNamePhoneSuffix: // phone,mobile后缀；	类型：varchar；
		case TypeNameEmailSuffix: // email后缀；	类型：varchar；
		case TypeNameUrlSuffix: // url,link后缀；	类型：varchar；
		case TypeNameIpSuffix: // IP后缀；	类型：varchar；
		case TypeNameIdSuffix: // id后缀；	类型：int等类型；
			relIdObj := tpl.Handle.RelIdMap[v.FieldRaw]
			if relIdObj.tpl.Table != `` && !relIdObj.IsRedundName {
				viewI18nField.name.Method = ReturnTypeName
				viewI18nField.name.DataTypeName = `'` + relIdObj.FieldName + `'`
			}
		case TypeNameSortSuffix, TypeNameSort: // sort,weight等后缀；	类型：int等类型； // sort，且pid,level,idPath|id_path,sort同时存在时（才）有效；	类型：int等类型；
			viewI18nField.tip.Method = ReturnTypeName
			viewI18nField.tip.DataTypeName = `'` + v.FieldTip + `'`
		case TypeNameStatusSuffix: // status,type,method,pos,position,gender等后缀；	类型：int等类型或varchar或char；	注释：多状态之间用[\s,，;；]等字符分隔。示例（状态：0待处理 1已处理 2驳回 yes是 no否）
			viewI18nField.status.Method = ReturnTypeName
			if garray.NewIntArrayFrom([]int{TypeVarchar, TypeChar}).Contains(v.FieldType) {
				for _, status := range v.StatusList {
					viewI18nField.status.DataTypeName = append(viewI18nField.status.DataTypeName, `{ value: '`+status[0]+`', label: '`+status[1]+`' },`)
				}
			} else {
				for _, status := range v.StatusList {
					viewI18nField.status.DataTypeName = append(viewI18nField.status.DataTypeName, `{ value: `+status[0]+`, label: '`+status[1]+`' },`)
				}
			}
		case TypeNameIsPrefix: // is_前缀；		类型：int等类型；注释：多状态之间用[\s,，;；]等字符分隔。示例（停用：0否 1是）
		case TypeNameStartPrefix: // start_前缀；	类型：timestamp或datetime或date；
		case TypeNameEndPrefix: // end_前缀；	类型：timestamp或datetime或date；
		case TypeNameRemarkSuffix: // remark,desc,msg,message,intro,content后缀；	类型：varchar或text；前端对应组件：varchar文本输入框，text富文本编辑器
		case TypeNameImageSuffix: // icon,cover,avatar,img,img_list,imgList,img_arr,imgArr,image,image_list,imageList,image_arr,imageArr等后缀；	类型：单图片varchar，多图片json或text
		case TypeNameVideoSuffix: // video,video_list,videoList,video_arr,videoArr等后缀；		类型：单视频varchar，多视频json或text
		case TypeNameArrSuffix: // list,arr等后缀；	类型：json或text；
		}
		/*--------根据字段命名类型处理 结束--------*/

		viewI18n.Add(viewI18nField, v.FieldRaw)
	}
	return
}

func getViewI18nExtendMiddleOne(tplEM handleExtendMiddle) (viewI18n myGenViewI18n) {
	switch tplEM.TableType {
	case TableTypeExtendOne:
		viewI18n.Merge(getViewI18nFieldList(tplEM.tpl, tplEM.FieldArr...))
	case TableTypeMiddleOne:
		viewI18n.Merge(getViewI18nFieldList(tplEM.tpl, tplEM.FieldArrOfIdSuffix...))
		if len(tplEM.FieldArrOfOther) > 0 {
			viewI18n.Merge(getViewI18nFieldList(tplEM.tpl, tplEM.FieldArrOfOther...))
		}
	}
	return
}

func getViewI18nExtendMiddleMany(tplEM handleExtendMiddle) (viewI18n myGenViewI18n) {
	if len(tplEM.GenFieldArr) == 1 {
		v := tplEM.GenFieldArr[0]

		viewI18nField := myGenViewI18nField{}
		viewI18nField.name.Method = ReturnType
		viewI18nField.name.DataType = `'` + v.FieldName + `'` //`'` + v.FieldName + `列表'`
		if v.FieldTypeName == TypeNameIdSuffix && gstr.ToUpper(gstr.SubStr(v.FieldName, -2)) == `ID` {
			viewI18nField.name.DataType = `'` + gstr.SubStr(v.FieldName, 0, -2) + `'` //`'` + gstr.SubStr(v.FieldName, 0, -2) + `列表'`
		}

		/*--------根据字段数据类型处理（注意：这里的代码改动对字段命名类型处理有影响） 开始--------*/
		switch v.FieldType {
		case TypeInt: // `int等类型`
		case TypeIntU: // `int等类型（unsigned）`
		case TypeFloat: // `float等类型`
		case TypeFloatU: // `float等类型（unsigned）`
		case TypeVarchar: // `varchar类型`
		case TypeChar: // `char类型`
		case TypeText: // `text类型`
		case TypeJson: // `json类型`
			viewI18nField.tip.Method = ReturnType
			if v.FieldTip != `` {
				viewI18nField.tip.DataType = `'` + v.FieldTip + `'`
			}
		case TypeTimestamp: // `timestamp类型`
		case TypeDatetime: // `datetime类型`
		case TypeDate: // `date类型`
		}
		/*--------根据字段数据类型处理（注意：这里的代码改动对字段命名类型处理有影响） 结束--------*/

		/*--------根据字段命名类型处理 开始--------*/
		switch v.FieldTypeName {
		case TypeNameDeleted, TypeNameUpdated, TypeNameCreated: // 软删除字段 // 更新时间字段 // 创建时间字段
		case TypeNamePid: // pid；	类型：int等类型；
		case TypeNameLevel: // level，且pid,level,idPath|id_path同时存在时（才）有效；	类型：int等类型；
		case TypeNameIdPath: // idPath|id_path，且pid,level,idPath|id_path同时存在时（才）有效；	类型：varchar或text；
		case TypeNamePasswordSuffix: // password,passwd后缀；		类型：char(32)；
		case TypeNameSaltSuffix: // salt后缀，且对应的password,passwd后缀存在时（才）有效；	类型：char；
		case TypeNameNameSuffix: // name,title后缀；	类型：varchar；
		case TypeNameCodeSuffix: // code后缀；	类型：varchar；
		case TypeNameAccountSuffix: // account后缀；	类型：varchar；
		case TypeNamePhoneSuffix: // phone,mobile后缀；	类型：varchar；
		case TypeNameEmailSuffix: // email后缀；	类型：varchar；
		case TypeNameUrlSuffix: // url,link后缀；	类型：varchar；
		case TypeNameIpSuffix: // IP后缀；	类型：varchar；
		case TypeNameIdSuffix: // id后缀；	类型：int等类型；
		case TypeNameSortSuffix, TypeNameSort: // sort,weight等后缀；	类型：int等类型； // sort，且pid,level,idPath|id_path,sort同时存在时（才）有效；	类型：int等类型；
			viewI18nField.tip.Method = ReturnTypeName
			viewI18nField.tip.DataTypeName = `'` + v.FieldTip + `'`
		case TypeNameStatusSuffix: // status,type,method,pos,position,gender等后缀；	类型：int等类型或varchar或char；	注释：多状态之间用[\s,，;；]等字符分隔。示例（状态：0待处理 1已处理 2驳回 yes是 no否）
			viewI18nField.status.Method = ReturnTypeName
			if garray.NewIntArrayFrom([]int{TypeVarchar, TypeChar}).Contains(v.FieldType) {
				for _, status := range v.StatusList {
					viewI18nField.status.DataTypeName = append(viewI18nField.status.DataTypeName, `{ value: '`+status[0]+`', label: '`+status[1]+`' },`)
				}
			} else {
				for _, status := range v.StatusList {
					viewI18nField.status.DataTypeName = append(viewI18nField.status.DataTypeName, `{ value: `+status[0]+`, label: '`+status[1]+`' },`)
				}
			}
		case TypeNameIsPrefix: // is_前缀；		类型：int等类型；注释：多状态之间用[\s,，;；]等字符分隔。示例（停用：0否 1是）
		case TypeNameStartPrefix: // start_前缀；	类型：timestamp或datetime或date；
		case TypeNameEndPrefix: // end_前缀；	类型：timestamp或datetime或date；
		case TypeNameRemarkSuffix: // remark,desc,msg,message,intro,content后缀；	类型：varchar或text；前端对应组件：varchar文本输入框，text富文本编辑器
		case TypeNameImageSuffix: // icon,cover,avatar,img,img_list,imgList,img_arr,imgArr,image,image_list,imageList,image_arr,imageArr等后缀；	类型：单图片varchar，多图片json或text
		case TypeNameVideoSuffix: // video,video_list,videoList,video_arr,videoArr等后缀；		类型：单视频varchar，多视频json或text
		case TypeNameArrSuffix: // list,arr等后缀；	类型：json或text；
		}
		/*--------根据字段命名类型处理 结束--------*/

		viewI18n.Add(viewI18nField, v.FieldRaw)     //viewQuery用
		viewI18n.Add(viewI18nField, tplEM.FieldVar) //viewSave用
	} else {
		viewI18nTmp := getViewI18nFieldList(tplEM.tpl, tplEM.FieldArrOfIdSuffix...)
		if len(tplEM.FieldArrOfOther) > 0 {
			viewI18nTmp.Merge(getViewI18nFieldList(tplEM.tpl, tplEM.FieldArrOfOther...))
		}
		viewI18n.name = append(viewI18n.name, tplEM.FieldVar+`Name: '列表',`)
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
