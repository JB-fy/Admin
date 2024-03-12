package my_gen

import (
	daoAuth "api/internal/dao/auth"
	"api/internal/utils"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

type myGenApi struct {
	filter   []string
	create   []string
	update   []string
	res      []string
	resOfAdd []string
}

type myGenApiField struct {
	filterType myGenDataStrHandler
	createType myGenDataStrHandler
	updateType myGenDataStrHandler
	resType    myGenDataStrHandler

	filterRule myGenDataSliceHandler
	saveRule   myGenDataSliceHandler
	isRequired bool
}

func genApi(option myGenOption, tpl myGenTpl) {
	apiObj := getApiFieldList(tpl)

	tplApi := `package api

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

`
	if option.IsList {
		tplApi += `
/*--------列表 开始--------*/
type ` + tpl.TableCaseCamel + `ListReq struct {
	g.Meta ` + "`" + `path:"/` + tpl.TableCaseKebab + `/list" method:"post" tags:"` + option.SceneInfo[daoAuth.Scene.Columns().SceneName].String() + `/` + option.CommonName + `" sm:"列表"` + "`" + `
	Filter ` + tpl.TableCaseCamel + `ListFilter ` + "`" + `json:"filter" dc:"过滤条件"` + "`" + `
	Field  []string        ` + "`" + `json:"field" v:"distinct|foreach|min-length:1" dc:"查询字段，传值参考返回的字段名，默认返回全部字段。注意：如前端页面所需字段较少，建议传指定字段，可大幅减轻服务器及数据库压力"` + "`" + `
	Sort   string          ` + "`" + `json:"sort" default:"id DESC" dc:"排序"` + "`" + `
	Page   int             ` + "`" + `json:"page" v:"min:1" default:"1" dc:"页码"` + "`" + `
	Limit  int             ` + "`" + `json:"limit" v:"min:0" default:"10" dc:"每页数量。可传0取全部"` + "`" + `
}

type ` + tpl.TableCaseCamel + `ListFilter struct {
	Id             *uint       ` + "`" + `json:"id,omitempty" v:"min:1" dc:"ID"` + "`" + `
	IdArr          []uint      ` + "`" + `json:"idArr,omitempty" v:"distinct|foreach|min:1" dc:"ID数组"` + "`" + `
	ExcId          *uint       ` + "`" + `json:"excId,omitempty" v:"min:1" dc:"排除ID"` + "`" + `
	ExcIdArr       []uint      ` + "`" + `json:"excIdArr,omitempty" v:"distinct|foreach|min:1" dc:"排除ID数组"` + "`" + gstr.Join(append([]string{``}, apiObj.filter...), `
	`) + `
}

type ` + tpl.TableCaseCamel + `ListRes struct {`
		if option.IsCount {
			tplApi += `
	Count int         ` + "`" + `json:"count" dc:"总数"` + "`"
		}
		tplApi += `
	List  []` + tpl.TableCaseCamel + `ListItem ` + "`" + `json:"list" dc:"列表"` + "`" + `
}

type ` + tpl.TableCaseCamel + `ListItem struct {
	Id          *uint       ` + "`" + `json:"id,omitempty" dc:"ID"` + "`" + gstr.Join(append([]string{``}, apiObj.res...), `
	`) + gstr.Join(append([]string{``}, apiObj.resOfAdd...), `
	`) + `
}

/*--------列表 结束--------*/

`
	}
	if option.IsInfo {
		tplApi += `/*--------详情 开始--------*/
type ` + tpl.TableCaseCamel + `InfoReq struct {
	g.Meta ` + "`" + `path:"/` + tpl.TableCaseKebab + `/info" method:"post" tags:"` + option.SceneInfo[daoAuth.Scene.Columns().SceneName].String() + `/` + option.CommonName + `" sm:"详情"` + "`" + `
	Id     uint     ` + "`" + `json:"id" v:"required|min:1" dc:"ID"` + "`" + `
	Field  []string ` + "`" + `json:"field" v:"distinct|foreach|min-length:1" dc:"查询字段，传值参考返回的字段名，默认返回全部字段。注意：如前端页面所需字段较少，建议传指定字段，可大幅减轻服务器及数据库压力"` + "`" + `
}

type ` + tpl.TableCaseCamel + `InfoRes struct {
	Info ` + tpl.TableCaseCamel + `Info ` + "`" + `json:"info" dc:"详情"` + "`" + `
}

type ` + tpl.TableCaseCamel + `Info struct {
	Id          *uint       ` + "`" + `json:"id,omitempty" dc:"ID"` + "`" + gstr.Join(append([]string{``}, apiObj.res...), `
	`) + `
}

/*--------详情 结束--------*/

`
	}
	if option.IsCreate {
		tplApi += `/*--------新增 开始--------*/
type ` + tpl.TableCaseCamel + `CreateReq struct {
	g.Meta      ` + "`" + `path:"/` + tpl.TableCaseKebab + `/create" method:"post" tags:"` + option.SceneInfo[daoAuth.Scene.Columns().SceneName].String() + `/` + option.CommonName + `" sm:"新增"` + "`" + gstr.Join(append([]string{``}, apiObj.create...), `
	`) + `
}

/*--------新增 结束--------*/

`
	}

	if option.IsUpdate {
		tplApi += `/*--------修改 开始--------*/
type ` + tpl.TableCaseCamel + `UpdateReq struct {
	g.Meta      ` + "`" + `path:"/` + tpl.TableCaseKebab + `/update" method:"post" tags:"` + option.SceneInfo[daoAuth.Scene.Columns().SceneName].String() + `/` + option.CommonName + `" sm:"修改"` + "`" + `
	IdArr       []uint  ` + "`" + `json:"idArr,omitempty" v:"required|distinct|foreach|min:1" dc:"ID数组"` + "`" + gstr.Join(append([]string{``}, apiObj.update...), `
	`) + `
}

/*--------修改 结束--------*/

`
	}

	if option.IsDelete {
		tplApi += `/*--------删除 开始--------*/
type ` + tpl.TableCaseCamel + `DeleteReq struct {
	g.Meta ` + "`" + `path:"/` + tpl.TableCaseKebab + `/del" method:"post" tags:"` + option.SceneInfo[daoAuth.Scene.Columns().SceneName].String() + `/` + option.CommonName + `" sm:"删除"` + "`" + `
	IdArr  []uint ` + "`" + `json:"idArr,omitempty" v:"required|distinct|foreach|min:1" dc:"ID数组"` + "`" + `
}

/*--------删除 结束--------*/
`
	}

	if option.IsList && tpl.Handle.Pid.Pid != `` {
		tplApi += `
/*--------列表（树状） 开始--------*/
type ` + tpl.TableCaseCamel + `TreeReq struct {
	g.Meta ` + "`" + `path:"/` + tpl.TableCaseKebab + `/tree" method:"post" tags:"` + option.SceneInfo[daoAuth.Scene.Columns().SceneName].String() + `/` + option.CommonName + `" sm:"列表（树状）"` + "`" + `
	Field  []string       ` + "`" + `json:"field" v:"foreach|min-length:1"` + "`" + `
	Filter ` + tpl.TableCaseCamel + `ListFilter ` + "`" + `json:"filter" dc:"过滤条件"` + "`" + `
}

type ` + tpl.TableCaseCamel + `TreeRes struct {
	Tree []` + tpl.TableCaseCamel + `TreeItem ` + "`" + `json:"tree" dc:"列表（树状）"` + "`" + `
}

type ` + tpl.TableCaseCamel + `TreeItem struct {
	Id       *uint       ` + "`" + `json:"id,omitempty" dc:"ID"` + "`" + gstr.Join(append([]string{``}, apiObj.res...), `
	`) + `
	Children []` + tpl.TableCaseCamel + `TreeItem ` + "`" + `json:"children" dc:"子级列表"` + "`" + `
}

/*--------列表（树状） 结束--------*/
`
	}

	saveFile := gfile.SelfDir() + `/api/` + option.SceneCode + `/` + tpl.ModuleDirCaseKebab + `/` + tpl.TableCaseSnake + `.go`
	gfile.PutContents(saveFile, tplApi)
	utils.GoFileFmt(saveFile)
}

func getApiFieldList(tpl myGenTpl) (api myGenApi) {
	if len(tpl.Handle.LabelList) > 0 {
		api.filter = append(api.filter, `Label string `+"`"+`json:"label,omitempty" v:"max-length:30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"标签。常用于前端组件"`+"`")
		api.res = append(api.res, `Label *string `+"`"+`json:"label,omitempty" dc:"标签。常用于前端组件"`+"`")
	}

	for _, v := range tpl.FieldList {
		apiField := myGenApiField{}

		/*--------根据字段数据类型处理（注意：这里的代码改动对字段命名类型处理有影响） 开始--------*/
		switch v.FieldType {
		case TypeInt: // `int等类型` // `int等类型（unsigned）`
			// apiField.filterType.Method = ReturnEmpty
			apiField.filterType.DataType = `*int`
			apiField.createType.Method = ReturnType
			apiField.createType.DataType = `*int`
			apiField.updateType.Method = ReturnType
			apiField.updateType.DataType = `*int`
			apiField.resType.Method = ReturnType
			apiField.resType.DataType = `*int`
		case TypeIntU: // `int等类型（unsigned）`
			// apiField.filterType.Method = ReturnType
			apiField.filterType.DataType = `*uint`
			apiField.createType.Method = ReturnType
			apiField.createType.DataType = `*uint`
			apiField.updateType.Method = ReturnType
			apiField.updateType.DataType = `*uint`
			apiField.resType.Method = ReturnType
			apiField.resType.DataType = `*uint`
		case TypeFloat: // `float等类型`
			// apiField.filterType.Method = ReturnType
			apiField.filterType.DataType = `*float64`
			apiField.createType.Method = ReturnType
			apiField.createType.DataType = `*float64`
			apiField.updateType.Method = ReturnType
			apiField.updateType.DataType = `*float64`
			apiField.resType.Method = ReturnType
			apiField.resType.DataType = `*float64`
		case TypeFloatU: // `float等类型（unsigned）`
			// apiField.filterType.Method = ReturnType
			apiField.filterType.DataType = `*float64`
			apiField.createType.Method = ReturnType
			apiField.createType.DataType = `*float64`
			apiField.updateType.Method = ReturnType
			apiField.updateType.DataType = `*float64`
			apiField.resType.Method = ReturnType
			apiField.resType.DataType = `*float64`

			apiField.filterRule.Method = ReturnType
			apiField.filterRule.DataType = append(apiField.filterRule.DataType, `min:0`)
			apiField.saveRule.Method = ReturnType
			apiField.saveRule.DataType = append(apiField.saveRule.DataType, `min:0`)
		case TypeVarchar: // `varchar类型`
			apiField.filterType.Method = ReturnType
			apiField.filterType.DataType = `string`
			apiField.createType.Method = ReturnType
			apiField.createType.DataType = `*string`
			apiField.updateType.Method = ReturnType
			apiField.updateType.DataType = `*string`
			apiField.resType.Method = ReturnType
			apiField.resType.DataType = `*string`

			apiField.filterRule.Method = ReturnType
			apiField.filterRule.DataType = append(apiField.filterRule.DataType, `max-length:`+v.FieldLimitStr)
			apiField.saveRule.Method = ReturnType
			apiField.saveRule.DataType = append(apiField.saveRule.DataType, `max-length:`+v.FieldLimitStr)
			if v.IndexRaw == `UNI` && !v.IsNull {
				apiField.isRequired = true
			}
		case TypeChar: // `char类型`
			apiField.filterType.Method = ReturnType
			apiField.filterType.DataType = `string`
			apiField.createType.Method = ReturnType
			apiField.createType.DataType = `*string`
			apiField.updateType.Method = ReturnType
			apiField.updateType.DataType = `*string`
			apiField.resType.Method = ReturnType
			apiField.resType.DataType = `*string`

			apiField.filterRule.Method = ReturnType
			apiField.filterRule.DataType = append(apiField.filterRule.DataType, `max-length:`+v.FieldLimitStr)
			apiField.saveRule.Method = ReturnType
			apiField.saveRule.DataType = append(apiField.saveRule.DataType, `size:`+v.FieldLimitStr)
			if v.IndexRaw == `UNI` && !v.IsNull {
				apiField.isRequired = true
			}
		case TypeText: // `text类型`
			// apiField.filterType.Method = ReturnType
			apiField.filterType.DataType = `string`
			apiField.createType.Method = ReturnType
			apiField.createType.DataType = `*string`
			apiField.updateType.Method = ReturnType
			apiField.updateType.DataType = `*string`
			apiField.resType.Method = ReturnType
			apiField.resType.DataType = `*string`
			/* if !v.IsNull {
				apiField.isRequired = true
			} */
		case TypeJson: // `json类型`
			// apiField.filterType.Method = ReturnType
			apiField.filterType.DataType = `string`
			apiField.createType.Method = ReturnType
			apiField.createType.DataType = `*string`
			apiField.updateType.Method = ReturnType
			apiField.updateType.DataType = `*string`
			apiField.resType.Method = ReturnType
			apiField.resType.DataType = `*string`

			apiField.filterRule.Method = ReturnType
			apiField.filterRule.DataType = append(apiField.filterRule.DataType, `json`)
			apiField.saveRule.Method = ReturnType
			apiField.saveRule.DataType = append(apiField.saveRule.DataType, `json`)
			if !v.IsNull {
				apiField.isRequired = true
			}
		case TypeTimestamp, TypeDatetime: // `timestamp类型` // `datetime类型`
			// apiField.filterType.Method = ReturnType
			apiField.filterType.DataType = `*gtime.Time`
			apiField.createType.Method = ReturnType
			apiField.createType.DataType = `*gtime.Time`
			apiField.updateType.Method = ReturnType
			apiField.updateType.DataType = `*gtime.Time`
			apiField.resType.Method = ReturnType
			apiField.resType.DataType = `*gtime.Time`

			apiField.filterRule.Method = ReturnType
			apiField.filterRule.DataType = append(apiField.filterRule.DataType, `date-format:Y-m-d H:i:s`)
			apiField.saveRule.Method = ReturnType
			apiField.saveRule.DataType = append(apiField.saveRule.DataType, `date-format:Y-m-d H:i:s`)
			if !v.IsNull && gconv.String(v.Default) == `` {
				apiField.isRequired = true
			}
		case TypeDate: // `date类型`
			apiField.filterType.Method = ReturnType
			apiField.filterType.DataType = `*gtime.Time`
			apiField.createType.Method = ReturnType
			apiField.createType.DataType = `*gtime.Time`
			apiField.updateType.Method = ReturnType
			apiField.updateType.DataType = `*gtime.Time`
			apiField.resType.Method = ReturnType
			apiField.resType.DataType = `*string`

			apiField.filterRule.Method = ReturnType
			apiField.filterRule.DataType = append(apiField.filterRule.DataType, `date-format:Y-m-d`)
			apiField.saveRule.Method = ReturnType
			apiField.saveRule.DataType = append(apiField.saveRule.DataType, `date-format:Y-m-d`)
			if !v.IsNull && gconv.String(v.Default) == `` {
				apiField.isRequired = true
			}
		default:
			apiField.filterType.Method = ReturnType
			apiField.filterType.DataType = `string`
			apiField.createType.Method = ReturnType
			apiField.createType.DataType = `*string`
			apiField.updateType.Method = ReturnType
			apiField.updateType.DataType = `*string`
			apiField.resType.Method = ReturnType
			apiField.resType.DataType = `*string`
		}
		/*--------根据字段数据类型处理（注意：这里的代码改动对字段命名类型处理有影响） 结束--------*/

		/*--------根据字段命名类型处理 开始--------*/
		switch v.FieldTypeName {
		case TypeNameDeleted: // 软删除字段
			continue
		case TypeNameUpdated: // 更新时间字段
			apiField.filterType.Method = ReturnEmpty
			apiField.createType.Method = ReturnEmpty
			apiField.updateType.Method = ReturnEmpty
		case TypeNameCreated: // 创建时间字段
			apiField.filterType.Method = ReturnEmpty
			apiField.createType.Method = ReturnEmpty
			apiField.updateType.Method = ReturnEmpty

			api.filter = append(api.filter,
				`TimeRangeStart *gtime.Time `+"`"+`json:"timeRangeStart,omitempty" v:"date-format:Y-m-d H:i:s" dc:"开始时间：YYYY-mm-dd HH:ii:ss"`+"`",
				`TimeRangeEnd   *gtime.Time `+"`"+`json:"timeRangeEnd,omitempty" v:"date-format:Y-m-d H:i:s|after-equal:TimeRangeStart" dc:"结束时间：YYYY-mm-dd HH:ii:ss"`+"`",
			)
		case TypeNamePri: // 主键
		case TypeNamePriAutoInc: // 主键（自增）
			if v.FieldRaw == `id` {
				continue
			}
			apiField.createType.Method = ReturnEmpty
			apiField.updateType.Method = ReturnEmpty

			apiField.filterRule.Method = ReturnUnion
			apiField.filterRule.DataTypeName = append(apiField.filterRule.DataTypeName, `min:1`)
		case TypeNamePid: // pid；	类型：int等类型；
			apiField.filterType.Method = ReturnType

			if len(tpl.Handle.LabelList) > 0 {
				api.resOfAdd = append(api.resOfAdd, `P`+gstr.CaseCamel(tpl.Handle.LabelList[0])+` *string `+"`"+`json:"p`+gstr.CaseCamel(tpl.Handle.LabelList[0])+`,omitempty" dc:"父级"`+"`")
			}
		case TypeNameLevel: // level，且pid,level,idPath|id_path同时存在时（才）有效；	类型：int等类型；
			apiField.filterType.Method = ReturnType
			apiField.createType.Method = ReturnEmpty
			apiField.updateType.Method = ReturnEmpty

			apiField.filterRule.Method = ReturnUnion
			apiField.filterRule.DataTypeName = append(apiField.filterRule.DataTypeName, `min:1`)
		case TypeNameIdPath: // idPath|id_path，且pid,level,idPath|id_path同时存在时（才）有效；	类型：varchar或text；
			apiField.filterType.Method = ReturnEmpty
			apiField.createType.Method = ReturnEmpty
			apiField.updateType.Method = ReturnEmpty
		case TypeNamePasswordSuffix: // password,passwd后缀；		类型：char(32)；
			apiField.filterType.Method = ReturnEmpty
			apiField.resType.Method = ReturnEmpty

			apiField.isRequired = true
		case TypeNameSaltSuffix: // salt后缀，且对应的password,passwd后缀存在时（才）有效；	类型：char；
			continue
		case TypeNameNameSuffix: // name,title后缀；	类型：varchar；
			/* // 不验证该规则。有时会用到特殊符号
			apiField.filterRule.Method = ReturnUnion
			apiField.filterRule.DataTypeName = append(apiField.filterRule.DataTypeName, `regex:^[\\p{L}\\p{M}\\p{N}_-]+$`)
			apiField.createRule.Method = ReturnUnion
			apiField.createRule.DataTypeName = append(apiField.createRule.DataTypeName, `regex:^[\\p{L}\\p{M}\\p{N}_-]+$`)
			apiField.updateRule.Method = ReturnUnion
			apiField.updateRule.DataTypeName = append(apiField.updateRule.DataTypeName, `regex:^[\\p{L}\\p{M}\\p{N}_-]+$`) */
			if len(tpl.Handle.LabelList) > 0 && gstr.CaseCamel(tpl.Handle.LabelList[0]) == v.FieldCaseCamel {
				apiField.isRequired = true
			}
		case TypeNameCodeSuffix: // code后缀；	类型：varchar；
			apiField.filterRule.Method = ReturnUnion
			apiField.filterRule.DataTypeName = append(apiField.filterRule.DataTypeName, `regex:^[\\p{L}\\p{M}\\p{N}_-]+$`)
			apiField.saveRule.Method = ReturnUnion
			apiField.saveRule.DataTypeName = append(apiField.saveRule.DataTypeName, `regex:^[\\p{L}\\p{M}\\p{N}_-]+$`)
		case TypeNameAccountSuffix: // account后缀；	类型：varchar；
			/* apiField.filterRule.Method = ReturnUnion
			apiField.filterRule.DataTypeName = append(apiField.filterRule.DataTypeName, `passport`)
			apiField.saveRule.Method = ReturnUnion
			apiField.saveRule.DataTypeName = append(apiField.saveRule.DataTypeName, `passport`) */
			apiField.filterRule.Method = ReturnUnion
			apiField.filterRule.DataTypeName = append(apiField.filterRule.DataTypeName, `regex:^(?!\\d*$)[\\p{L}\\p{M}\\p{N}_]+$`)
			apiField.saveRule.Method = ReturnUnion
			apiField.saveRule.DataTypeName = append(apiField.saveRule.DataTypeName, `regex:^(?!\\d*$)[\\p{L}\\p{M}\\p{N}_]+$`)
		case TypeNamePhoneSuffix: // phone,mobile后缀；	类型：varchar；
			apiField.filterRule.Method = ReturnUnion
			apiField.filterRule.DataTypeName = append(apiField.filterRule.DataTypeName, `phone`)
			apiField.saveRule.Method = ReturnUnion
			apiField.saveRule.DataTypeName = append(apiField.saveRule.DataTypeName, `phone`)
		case TypeNameEmailSuffix: // email后缀；	类型：varchar；
			apiField.filterRule.Method = ReturnUnion
			apiField.filterRule.DataTypeName = append(apiField.filterRule.DataTypeName, `email`)
			apiField.saveRule.Method = ReturnUnion
			apiField.saveRule.DataTypeName = append(apiField.saveRule.DataTypeName, `email`)
		case TypeNameUrlSuffix: // url,link后缀；	类型：varchar；
			apiField.filterRule.Method = ReturnUnion
			apiField.filterRule.DataTypeName = append(apiField.filterRule.DataTypeName, `url`)
			apiField.saveRule.Method = ReturnUnion
			apiField.saveRule.DataTypeName = append(apiField.saveRule.DataTypeName, `url`)
		case TypeNameIpSuffix: // IP后缀；	类型：varchar；
			apiField.filterRule.Method = ReturnUnion
			apiField.filterRule.DataTypeName = append(apiField.filterRule.DataTypeName, `ip`)
			apiField.saveRule.Method = ReturnUnion
			apiField.saveRule.DataTypeName = append(apiField.saveRule.DataTypeName, `ip`)
		case TypeNameIdSuffix: // id后缀；	类型：int等类型；
			apiField.filterType.Method = ReturnType

			apiField.filterRule.Method = ReturnUnion
			apiField.filterRule.DataTypeName = append(apiField.filterRule.DataTypeName, `min:1`)
			apiField.saveRule.Method = ReturnUnion
			apiField.saveRule.DataTypeName = append(apiField.saveRule.DataTypeName, `min:1`)

			if tpl.Handle.RelIdMap[v.FieldRaw].tpl.Table != `` && !tpl.Handle.RelIdMap[v.FieldRaw].IsRedundName {
				relIdObj := tpl.Handle.RelIdMap[v.FieldRaw]
				api.resOfAdd = append(api.resOfAdd, gstr.CaseCamel(relIdObj.tpl.Handle.LabelList[0])+gstr.CaseCamel(relIdObj.Suffix)+` *string `+"`"+`json:"`+relIdObj.tpl.Handle.LabelList[0]+relIdObj.Suffix+`,omitempty" dc:"`+relIdObj.FieldName+`"`+"`")
			}
		case TypeNameSortSuffix, TypeNameSort: // sort,weight等后缀；	类型：int等类型； // sort，且pid,level,idPath|id_path,sort同时存在时（才）有效；	类型：int等类型；
			apiField.saveRule.Method = ReturnUnion
			apiField.saveRule.DataTypeName = append(apiField.saveRule.DataTypeName, `between:0,100`)
		case TypeNameStatusSuffix: // status,type,method,pos,position,gender等后缀；	类型：int等类型或varchar或char；	注释：多状态之间用[\s,，;；]等字符分隔。示例（状态：0待处理 1已处理 2驳回 yes是 no否）
			apiField.filterType.Method = ReturnType

			statusArr := make([]string, len(v.StatusList))
			for index, item := range v.StatusList {
				statusArr[index] = item[0]
			}
			statusStr := gstr.Join(statusArr, `,`)
			apiField.filterRule.Method = ReturnUnion
			apiField.filterRule.DataTypeName = append(apiField.filterRule.DataTypeName, `in:`+statusStr)
			apiField.saveRule.Method = ReturnUnion
			apiField.saveRule.DataTypeName = append(apiField.saveRule.DataTypeName, `in:`+statusStr)
		case TypeNameIsPrefix: // is_前缀；		类型：int等类型；注释：多状态之间用[\s,，;；]等字符分隔。示例（停用：0否 1是）
			apiField.filterType.Method = ReturnType

			/* TODO 可改成状态一样处理，同时需要修改前端开关组件属性设置（暂时不改）*/
			apiField.filterRule.Method = ReturnUnion
			apiField.filterRule.DataTypeName = append(apiField.filterRule.DataTypeName, `in:0,1`)
			apiField.saveRule.Method = ReturnUnion
			apiField.saveRule.DataTypeName = append(apiField.saveRule.DataTypeName, `in:0,1`)
		case TypeNameStartPrefix: // start_前缀；	类型：timestamp或datetime或date；
			apiField.filterType.Method = ReturnType
		case TypeNameEndPrefix: // end_前缀；	类型：timestamp或datetime或date；
			apiField.filterType.Method = ReturnType
		case TypeNameRemarkSuffix: // remark,desc,msg,message,intro,content后缀；	类型：varchar或text；前端对应组件：varchar文本输入框，text富文本编辑器
			apiField.filterType.Method = ReturnEmpty
		case TypeNameImageSuffix, TypeNameVideoSuffix: // icon,cover,avatar,img,img_list,imgList,img_arr,imgArr,image,image_list,imageList,image_arr,imageArr等后缀；	类型：单图片varchar，多图片json或text	// video,video_list,videoList,video_arr,videoArr等后缀；		类型：单视频varchar，多视频json或text
			if v.FieldType == TypeVarchar {
				apiField.filterType.Method = ReturnEmpty

				apiField.saveRule.Method = ReturnUnion
				apiField.saveRule.DataTypeName = append(apiField.saveRule.DataTypeName, `url`)
			} else {
				apiField.createType.Method = ReturnTypeName
				apiField.createType.DataTypeName = `*[]string`
				apiField.updateType.Method = ReturnTypeName
				apiField.updateType.DataTypeName = `*[]string`
				apiField.resType.Method = ReturnTypeName
				apiField.resType.DataTypeName = `*[]string`

				apiField.saveRule.Method = ReturnUnion
				apiField.saveRule.DataTypeName = append(apiField.saveRule.DataTypeName, `distinct`, `foreach`, `url`, `foreach`, `min-length:1`)
				if !v.IsNull {
					apiField.isRequired = true
				}
			}
		case TypeNameArrSuffix: // list,arr等后缀；	类型：json或text；
			apiField.createType.Method = ReturnTypeName
			apiField.createType.DataTypeName = `*[]interface{}`
			apiField.updateType.Method = ReturnTypeName
			apiField.updateType.DataTypeName = `*[]interface{}`
			apiField.resType.Method = ReturnTypeName
			apiField.resType.DataTypeName = `*[]interface{}`

			apiField.saveRule.Method = ReturnUnion
			apiField.saveRule.DataTypeName = append(apiField.saveRule.DataTypeName, `distinct`)
			if !v.IsNull {
				apiField.isRequired = true
			}
		}
		/*--------根据字段命名类型处理 结束--------*/

		if apiField.filterType.getData() != `` {
			api.filter = append(api.filter, v.FieldCaseCamel+` `+apiField.filterType.getData()+` `+"`"+`json:"`+v.FieldRaw+`,omitempty" v:"`+gstr.Join(apiField.filterRule.getData(), `|`)+`" dc:"`+v.FieldDesc+`"`+"`")
		}
		if apiField.createType.getData() != `` {
			if apiField.isRequired {
				api.create = append(api.create, v.FieldCaseCamel+` `+apiField.createType.getData()+` `+"`"+`json:"`+v.FieldRaw+`,omitempty" v:"`+gstr.Join(append([]string{`required`}, apiField.saveRule.getData()...), `|`)+`" dc:"`+v.FieldDesc+`"`+"`")
			} else {
				api.create = append(api.create, v.FieldCaseCamel+` `+apiField.createType.getData()+` `+"`"+`json:"`+v.FieldRaw+`,omitempty" v:"`+gstr.Join(apiField.saveRule.getData(), `|`)+`" dc:"`+v.FieldDesc+`"`+"`")
			}
		}
		if apiField.updateType.getData() != `` {
			api.update = append(api.update, v.FieldCaseCamel+` `+apiField.updateType.getData()+` `+"`"+`json:"`+v.FieldRaw+`,omitempty" v:"`+gstr.Join(apiField.saveRule.getData(), `|`)+`" dc:"`+v.FieldDesc+`"`+"`")
		}
		if apiField.resType.getData() != `` {
			api.res = append(api.res, v.FieldCaseCamel+` `+apiField.resType.getData()+` `+"`"+`json:"`+v.FieldRaw+`,omitempty" dc:"`+v.FieldDesc+`"`+"`")
		}
	}

	// 做一次去重
	api.filter = garray.NewStrArrayFrom(api.filter).Unique().Slice()
	api.create = garray.NewStrArrayFrom(api.create).Unique().Slice()
	api.update = garray.NewStrArrayFrom(api.update).Unique().Slice()
	api.res = garray.NewStrArrayFrom(api.res).Unique().Slice()
	api.resOfAdd = garray.NewStrArrayFrom(api.resOfAdd).Unique().Slice()
	return
}
