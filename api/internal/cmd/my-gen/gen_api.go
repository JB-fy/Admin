package my_gen

import (
	"api/internal/cmd/my-gen/internal"
	daoAuth "api/internal/dao/auth"
	"api/internal/utils"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

type myGenApi struct {
	filterOfFixed []string
	resOfAdd      []string
	filter        []string
	info          []string
	create        []string
	update        []string
	delete        []string
	res           []string
}

type myGenApiField struct {
	filterOfFixed []string
	resOfAdd      []string

	filterType internal.MyGenDataStrHandler
	createType internal.MyGenDataStrHandler
	updateType internal.MyGenDataStrHandler
	resType    internal.MyGenDataStrHandler

	isRequired bool
	filterRule internal.MyGenDataSliceHandler
	saveRule   internal.MyGenDataSliceHandler
}

func (apiThis *myGenApi) Add(apiField myGenApiField, field myGenField, tableType internal.MyGenTableType) {
	apiThis.filterOfFixed = append(apiThis.filterOfFixed, apiField.filterOfFixed...)
	apiThis.resOfAdd = append(apiThis.resOfAdd, apiField.resOfAdd...)
	if apiField.filterType.GetData() != `` {
		apiThis.filter = append(apiThis.filter, field.FieldCaseCamel+` `+apiField.filterType.GetData()+` `+"`"+`json:"`+field.FieldRaw+`,omitempty" v:"`+gstr.Join(apiField.filterRule.GetData(), `|`)+`" dc:"`+field.FieldDesc+`"`+"`")
	}
	if apiField.createType.GetData() != `` {
		saveRuleArr := apiField.saveRule.GetData()
		if apiField.isRequired && garray.NewFrom([]any{internal.TableTypeDefault, internal.TableTypeExtendOne, internal.TableTypeMiddleOne}).Contains(tableType) {
			saveRuleArr = append([]string{`required`}, saveRuleArr...)
		}
		apiThis.create = append(apiThis.create, field.FieldCaseCamel+` `+apiField.createType.GetData()+` `+"`"+`json:"`+field.FieldRaw+`,omitempty" v:"`+gstr.Join(saveRuleArr, `|`)+`" dc:"`+field.FieldDesc+`"`+"`")
	}
	if apiField.updateType.GetData() != `` {
		apiThis.update = append(apiThis.update, field.FieldCaseCamel+` `+apiField.updateType.GetData()+` `+"`"+`json:"`+field.FieldRaw+`,omitempty" filter:"-" v:"`+gstr.Join(apiField.saveRule.GetData(), `|`)+`" dc:"`+field.FieldDesc+`"`+"`")
	}
	if apiField.resType.GetData() != `` {
		apiThis.res = append(apiThis.res, field.FieldCaseCamel+` `+apiField.resType.GetData()+` `+"`"+`json:"`+field.FieldRaw+`,omitempty" dc:"`+field.FieldDesc+`"`+"`")
	}
}

func (apiThis *myGenApi) Merge(apiOther myGenApi) {
	apiThis.filterOfFixed = append(apiThis.filterOfFixed, apiOther.filterOfFixed...)
	apiThis.filter = append(apiThis.filter, apiOther.filter...)
	apiThis.info = append(apiThis.info, apiOther.info...)
	apiThis.create = append(apiThis.create, apiOther.create...)
	apiThis.update = append(apiThis.update, apiOther.update...)
	apiThis.delete = append(apiThis.delete, apiOther.delete...)
	apiThis.res = append(apiThis.res, apiOther.res...)
	apiThis.resOfAdd = append(apiThis.resOfAdd, apiOther.resOfAdd...)
}

func (apiThis *myGenApi) Unique() {
	// apiThis.filterOfFixed = garray.NewStrArrayFrom(apiThis.filterOfFixed).Unique().Slice()
	// apiThis.filter = garray.NewStrArrayFrom(apiThis.filter).Unique().Slice()
	// apiThis.info = garray.NewStrArrayFrom(apiThis.info).Unique().Slice()
	// apiThis.create = garray.NewStrArrayFrom(apiThis.create).Unique().Slice()
	// apiThis.update = garray.NewStrArrayFrom(apiThis.update).Unique().Slice()
	// apiThis.delete = garray.NewStrArrayFrom(apiThis.delete).Unique().Slice()
	// apiThis.res = garray.NewStrArrayFrom(apiThis.res).Unique().Slice()
	// apiThis.resOfAdd = garray.NewStrArrayFrom(apiThis.resOfAdd).Unique().Slice()
}

// api生成
func genApi(option myGenOption, tpl myGenTpl) {
	api := getApiIdAndLabel(tpl)
	for _, v := range tpl.FieldListOfDefault {
		api.Add(getApiField(tpl, v), v, internal.TableTypeDefault)
	}
	for _, v := range tpl.FieldListOfAfter1 {
		api.Add(getApiField(tpl, v), v, internal.TableTypeDefault)
	}
	for _, v := range tpl.Handle.ExtendTableOneList {
		api.Merge(getApiExtendMiddleOne(v))
	}
	for _, v := range tpl.Handle.MiddleTableOneList {
		api.Merge(getApiExtendMiddleOne(v))
	}
	for _, v := range tpl.Handle.ExtendTableManyList {
		api.Merge(getApiExtendMiddleMany(v))
	}
	for _, v := range tpl.Handle.MiddleTableManyList {
		api.Merge(getApiExtendMiddleMany(v))
	}
	for _, v := range tpl.FieldListOfAfter2 {
		api.Add(getApiField(tpl, v), v, internal.TableTypeDefault)
	}
	api.Unique()

	tplApi := `package ` + tpl.GetModuleName(`api`) + `

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// 共用详情。list,info,tree等接口返回时用，但返回默认字段有差异。可根据需要在controller对应的defaultField中补充所需字段
type ` + tpl.TableCaseCamel + `Info struct {` + gstr.Join(append([]string{``}, api.res...), `
	`) + gstr.Join(append([]string{``}, api.resOfAdd...), `
	`)
	if option.IsList && tpl.Handle.Pid.Pid != `` {
		tplApi += `
	Children []` + tpl.TableCaseCamel + `Info ` + "`" + `json:"children" dc:"子级列表"` + "`"
	}
	tplApi += `
}

`
	if option.IsList {
		tplApi += `
type ` + tpl.TableCaseCamel + `Filter struct {` + gstr.Join(append([]string{``}, api.filterOfFixed...), `
	`) + gstr.Join(append([]string{``}, api.filter...), `
	`) + `
}

/*--------列表 开始--------*/
type ` + tpl.TableCaseCamel + `ListReq struct {
	g.Meta ` + "`" + `path:"/` + tpl.TableCaseKebab + `/list" method:"post" tags:"` + option.SceneInfo[daoAuth.Scene.Columns().SceneName].String() + `/` + option.CommonName + `" sm:"列表"` + "`" + `
	Filter ` + tpl.TableCaseCamel + `Filter ` + "`" + `json:"filter" dc:"过滤条件"` + "`" + `
	Field  []string        ` + "`" + `json:"field" v:"distinct|foreach|min-length:1" dc:"查询字段，传值参考返回的字段名，默认返回常用字段，如果所需字段较少或需特别字段时，可使用。特别注意：所需字段较少时使用，可大幅减轻数据库压力"` + "`" + `
	Sort   string          ` + "`" + `json:"sort" default:"id DESC" dc:"排序"` + "`" + `
	Page   int             ` + "`" + `json:"page" v:"min:1" default:"1" dc:"页码"` + "`" + `
	Limit  int             ` + "`" + `json:"limit" v:"min:0" default:"10" dc:"每页数量。可传0取全部"` + "`" + `
}

type ` + tpl.TableCaseCamel + `ListRes struct {`
		if option.IsCount {
			tplApi += `
	Count int         ` + "`" + `json:"count" dc:"总数"` + "`"
		}
		tplApi += `
	List  []` + tpl.TableCaseCamel + `Info ` + "`" + `json:"list" dc:"列表"` + "`" + `
}

/*--------列表 结束--------*/

`
	}
	if option.IsInfo {
		tplApi += `/*--------详情 开始--------*/
type ` + tpl.TableCaseCamel + `InfoReq struct {
	g.Meta ` + "`" + `path:"/` + tpl.TableCaseKebab + `/info" method:"post" tags:"` + option.SceneInfo[daoAuth.Scene.Columns().SceneName].String() + `/` + option.CommonName + `" sm:"详情"` + "`" + `
	Field  []string ` + "`" + `json:"field" v:"distinct|foreach|min-length:1" dc:"查询字段，传值参考返回的字段名，默认返回常用字段，如果所需字段较少或需特别字段时，可使用。特别注意：所需字段较少时使用，可大幅减轻数据库压力"` + "`" + gstr.Join(append([]string{``}, api.info...), `
	`) + `
}

type ` + tpl.TableCaseCamel + `InfoRes struct {
	Info ` + tpl.TableCaseCamel + `Info ` + "`" + `json:"info" dc:"详情"` + "`" + `
}

/*--------详情 结束--------*/

`
	}
	if option.IsCreate {
		tplApi += `/*--------新增 开始--------*/
type ` + tpl.TableCaseCamel + `CreateReq struct {
	g.Meta      ` + "`" + `path:"/` + tpl.TableCaseKebab + `/create" method:"post" tags:"` + option.SceneInfo[daoAuth.Scene.Columns().SceneName].String() + `/` + option.CommonName + `" sm:"新增"` + "`" + gstr.Join(append([]string{``}, api.create...), `
	`) + `
}

/*--------新增 结束--------*/

`
	}

	if option.IsUpdate {
		tplApi += `/*--------修改 开始--------*/
type ` + tpl.TableCaseCamel + `UpdateReq struct {
	g.Meta      ` + "`" + `path:"/` + tpl.TableCaseKebab + `/update" method:"post" tags:"` + option.SceneInfo[daoAuth.Scene.Columns().SceneName].String() + `/` + option.CommonName + `" sm:"修改"` + "`" + gstr.Join(append([]string{``}, api.update...), `
	`) + `
}

/*--------修改 结束--------*/

`
	}

	if option.IsDelete {
		tplApi += `/*--------删除 开始--------*/
type ` + tpl.TableCaseCamel + `DeleteReq struct {
	g.Meta ` + "`" + `path:"/` + tpl.TableCaseKebab + `/del" method:"post" tags:"` + option.SceneInfo[daoAuth.Scene.Columns().SceneName].String() + `/` + option.CommonName + `" sm:"删除"` + "`" + gstr.Join(append([]string{``}, api.delete...), `
	`) + `
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
	Filter ` + tpl.TableCaseCamel + `Filter ` + "`" + `json:"filter" dc:"过滤条件"` + "`" + `
}

type ` + tpl.TableCaseCamel + `TreeRes struct {
	Tree []` + tpl.TableCaseCamel + `Info ` + "`" + `json:"tree" dc:"列表（树状）"` + "`" + `
}

/*--------列表（树状） 结束--------*/
`
	}

	saveFile := gfile.SelfDir() + `/api/` + option.SceneId + `/` + tpl.ModuleDirCaseKebab + `/` + tpl.TableCaseSnake + `.go`
	gfile.PutContents(saveFile, tplApi)
	utils.GoFileFmt(saveFile)
}

func getApiIdAndLabel(tpl myGenTpl) (api myGenApi) {
	fieldStyleOfIdArr := internal.GetStrByFieldStyle(tpl.FieldStyle, `id_arr`)
	fieldStyleOfExcId := internal.GetStrByFieldStyle(tpl.FieldStyle, `exc_id`)
	fieldStyleOfExcIdArr := internal.GetStrByFieldStyle(tpl.FieldStyle, `exc_id_arr`)
	if len(tpl.Handle.Id.List) == 1 {
		switch tpl.Handle.Id.List[0].FieldType {
		case internal.TypeInt, internal.TypeIntU:
			dataType := `int`
			if tpl.Handle.Id.List[0].FieldType == internal.TypeIntU {
				dataType = `uint`
			}
			ruleOfId := []string{}
			ruleOfIdArr := []string{`distinct`}
			if tpl.Handle.Id.List[0].IsAutoInc || tpl.Handle.Id.List[0].FieldTypeName == internal.TypeNameIdSuffix {
				ruleOfId = append(ruleOfId, `between:1,`+tpl.Handle.Id.List[0].FieldLimitInt.Max)
				ruleOfIdArr = append(ruleOfIdArr, `foreach`, `between:1,`+tpl.Handle.Id.List[0].FieldLimitInt.Max)
			} else {
				ruleOfId = append(ruleOfId, `between:`+tpl.Handle.Id.List[0].FieldLimitInt.Min+`,`+tpl.Handle.Id.List[0].FieldLimitInt.Max)
				ruleOfIdArr = append(ruleOfIdArr, `foreach`, `between:`+tpl.Handle.Id.List[0].FieldLimitInt.Min+`,`+tpl.Handle.Id.List[0].FieldLimitInt.Max)
			}
			api.filterOfFixed = append(api.filterOfFixed,
				`Id *`+dataType+" `"+`json:"id,omitempty" v:"`+gstr.Join(ruleOfId, `|`)+`" dc:"ID"`+"`",
				`IdArr []`+dataType+" `"+`json:"`+fieldStyleOfIdArr+`,omitempty" v:"`+gstr.Join(ruleOfIdArr, `|`)+`" dc:"ID数组"`+"`",
				`ExcId *`+dataType+" `"+`json:"`+fieldStyleOfExcId+`,omitempty" v:"`+gstr.Join(ruleOfId, `|`)+`" dc:"排除ID"`+"`",
				`ExcIdArr []`+dataType+" `"+`json:"`+fieldStyleOfExcIdArr+`,omitempty" v:"`+gstr.Join(ruleOfIdArr, `|`)+`" dc:"排除ID数组"`+"`",
			)
			api.info = append(api.info, `Id `+dataType+" `"+`json:"id" v:"`+gstr.Join(append([]string{`required`}, ruleOfId...), `|`)+`" dc:"ID"`+"`")
			api.update = append(api.update, `Id `+dataType+" `"+`json:"-" filter:"id,omitempty" v:"`+gstr.Join(append([]string{`required-without:IdArr`}, ruleOfId...), `|`)+`" dc:"ID"`+"`")
			api.update = append(api.update, `IdArr []`+dataType+" `"+`json:"-" filter:"`+fieldStyleOfIdArr+`,omitempty" v:"`+gstr.Join(append([]string{`required-without:Id`}, ruleOfIdArr...), `|`)+`" dc:"ID数组"`+"`")
			api.delete = append(api.delete, `Id `+dataType+" `"+`json:"id,omitempty" v:"`+gstr.Join(append([]string{`required-without:IdArr`}, ruleOfId...), `|`)+`" dc:"ID"`+"`")
			api.delete = append(api.delete, `IdArr []`+dataType+" `"+`json:"`+fieldStyleOfIdArr+`,omitempty" v:"`+gstr.Join(append([]string{`required-without:Id`}, ruleOfIdArr...), `|`)+`" dc:"ID数组"`+"`")
			api.res = append(api.res, `Id *`+dataType+" `"+`json:"id,omitempty" dc:"ID"`+"`")
		default:
			api.filterOfFixed = append(api.filterOfFixed,
				`Id string `+"`"+`json:"id,omitempty" v:"max-length:`+tpl.Handle.Id.List[0].FieldLimitStr+`" dc:"ID"`+"`",
				`IdArr []string `+"`"+`json:"`+fieldStyleOfIdArr+`,omitempty" v:"distinct|foreach|length:1,`+tpl.Handle.Id.List[0].FieldLimitStr+`" dc:"ID数组"`+"`",
				`ExcId string `+"`"+`json:"`+fieldStyleOfExcId+`,omitempty" v:"max-length:`+tpl.Handle.Id.List[0].FieldLimitStr+`" dc:"排除ID"`+"`",
				`ExcIdArr []string `+"`"+`json:"`+fieldStyleOfExcIdArr+`,omitempty" v:"distinct|foreach|length:1,`+tpl.Handle.Id.List[0].FieldLimitStr+`" dc:"排除ID数组"`+"`",
			)
			api.info = append(api.info, `Id string `+"`"+`json:"id" v:"required|max-length:`+tpl.Handle.Id.List[0].FieldLimitStr+`" dc:"ID"`+"`")
			api.update = append(api.update, `Id string `+"`"+`json:"-" filter:"id,omitempty" v:"required-without:IdArr|length:1,`+tpl.Handle.Id.List[0].FieldLimitStr+`" dc:"ID"`+"`")
			api.update = append(api.update, `IdArr []string `+"`"+`json:"-" filter:"`+fieldStyleOfIdArr+`,omitempty" v:"required-without:Id|distinct|foreach|length:1,`+tpl.Handle.Id.List[0].FieldLimitStr+`" dc:"ID数组"`+"`")
			api.delete = append(api.delete, `Id string `+"`"+`json:"id,omitempty" v:"required-without:IdArr|length:1,`+tpl.Handle.Id.List[0].FieldLimitStr+`" dc:"ID"`+"`")
			api.delete = append(api.delete, `IdArr []string `+"`"+`json:"`+fieldStyleOfIdArr+`,omitempty" v:"required-without:Id|distinct|foreach|length:1,`+tpl.Handle.Id.List[0].FieldLimitStr+`" dc:"ID数组"`+"`")
			api.res = append(api.res, `Id *string `+"`"+`json:"id,omitempty" dc:"ID"`+"`")
		}
	} else {
		api.filterOfFixed = append(api.filterOfFixed,
			`Id string `+"`"+`json:"id,omitempty" v:"" dc:"ID"`+"`",
			`IdArr []string `+"`"+`json:"`+fieldStyleOfIdArr+`,omitempty" v:"distinct|foreach|min-length:1" dc:"ID数组"`+"`",
			`ExcId string `+"`"+`json:"`+fieldStyleOfExcId+`,omitempty" v:"" dc:"排除ID"`+"`",
			`ExcIdArr []string `+"`"+`json:"`+fieldStyleOfExcIdArr+`,omitempty" v:"distinct|foreach|min-length:1" dc:"排除ID数组"`+"`",
		)
		api.info = append(api.info, `Id string `+"`"+`json:"id" v:"required" dc:"ID"`+"`")
		api.update = append(api.update, `Id string `+"`"+`json:"-" filter:"id,omitempty" v:"required-without:IdArr|min-length:1" dc:"ID"`+"`")
		api.update = append(api.update, `IdArr []string `+"`"+`json:"-" filter:"`+fieldStyleOfIdArr+`,omitempty" v:"required-without:Id|distinct|foreach|min-length:1" dc:"ID数组"`+"`")
		api.delete = append(api.delete, `Id string `+"`"+`json:"id,omitempty" v:"required-without:IdArr|min-length:1" dc:"ID"`+"`")
		api.delete = append(api.delete, `IdArr []string `+"`"+`json:"`+fieldStyleOfIdArr+`,omitempty" v:"required-without:Id|distinct|foreach|min-length:1" dc:"ID数组"`+"`")
		api.res = append(api.res, `Id *string `+"`"+`json:"id,omitempty" dc:"ID"`+"`")
	}

	api.filterOfFixed = append(api.filterOfFixed, `Label string `+"`"+`json:"label,omitempty" v:"max-length:30|regex:^[\\p{L}\\p{N}_-]+$" dc:"标签。常用于前端组件"`+"`")
	api.res = append(api.res, `Label *string `+"`"+`json:"label,omitempty" dc:"标签。常用于前端组件"`+"`")
	return
}

func getApiField(tpl myGenTpl, v myGenField) (apiField myGenApiField) {
	if !v.IsNull && (gvar.New(v.Default).IsNil() || v.IsUnique) {
		apiField.isRequired = true
	}
	/*--------根据字段数据类型处理（注意：这里的代码改动对字段命名类型处理有影响） 开始--------*/
	switch v.FieldType {
	case internal.TypeInt, internal.TypeIntU: // `int等类型`	// `int等类型（unsigned）`
		dataType := `int`
		if v.FieldType == internal.TypeIntU {
			dataType = `uint`
		}
		// apiField.filterType.Method = internal.ReturnType
		apiField.filterType.DataType = `*` + dataType
		apiField.createType.Method = internal.ReturnType
		apiField.createType.DataType = `*` + dataType
		apiField.updateType.Method = internal.ReturnType
		apiField.updateType.DataType = `*` + dataType
		apiField.resType.Method = internal.ReturnType
		apiField.resType.DataType = `*` + dataType

		apiField.filterRule.Method = internal.ReturnType
		apiField.filterRule.DataType = append(apiField.filterRule.DataType, `between:`+v.FieldLimitInt.Min+`,`+v.FieldLimitInt.Max)
		apiField.saveRule.Method = internal.ReturnType
		apiField.saveRule.DataType = append(apiField.saveRule.DataType, `between:`+v.FieldLimitInt.Min+`,`+v.FieldLimitInt.Max)
	case internal.TypeFloat, internal.TypeFloatU: // `float等类型`	// `float等类型（unsigned）`
		// apiField.filterType.Method = internal.ReturnType
		apiField.filterType.DataType = `*float64`
		apiField.createType.Method = internal.ReturnType
		apiField.createType.DataType = `*float64`
		apiField.updateType.Method = internal.ReturnType
		apiField.updateType.DataType = `*float64`
		apiField.resType.Method = internal.ReturnType
		apiField.resType.DataType = `*float64`

		if v.FieldLimitFloat.Min != `` && v.FieldLimitFloat.Max != `` {
			apiField.filterRule.Method = internal.ReturnType
			apiField.filterRule.DataType = append(apiField.filterRule.DataType, `between:`+v.FieldLimitFloat.Min+`,`+v.FieldLimitFloat.Max)
			apiField.saveRule.Method = internal.ReturnType
			apiField.saveRule.DataType = append(apiField.saveRule.DataType, `between:`+v.FieldLimitFloat.Min+`,`+v.FieldLimitFloat.Max)
		} else if v.FieldLimitFloat.Min != `` {
			apiField.filterRule.Method = internal.ReturnType
			apiField.filterRule.DataType = append(apiField.filterRule.DataType, `min:`+v.FieldLimitFloat.Min)
			apiField.saveRule.Method = internal.ReturnType
			apiField.saveRule.DataType = append(apiField.saveRule.DataType, `min:`+v.FieldLimitFloat.Min)
		} else if v.FieldLimitFloat.Max != `` {
			apiField.filterRule.Method = internal.ReturnType
			apiField.filterRule.DataType = append(apiField.filterRule.DataType, `max:`+v.FieldLimitFloat.Max)
			apiField.saveRule.Method = internal.ReturnType
			apiField.saveRule.DataType = append(apiField.saveRule.DataType, `max:`+v.FieldLimitFloat.Max)
		}
	case internal.TypeVarchar, internal.TypeChar: // `varchar类型`	// `char类型`
		if v.IsUnique || gconv.Uint(v.FieldLimitStr) <= internal.ConfigMaxLenOfStrFilter {
			apiField.filterType.Method = internal.ReturnType
			apiField.filterType.DataType = `string`
		}
		apiField.createType.Method = internal.ReturnType
		apiField.createType.DataType = `*string`
		apiField.updateType.Method = internal.ReturnType
		apiField.updateType.DataType = `*string`
		apiField.resType.Method = internal.ReturnType
		apiField.resType.DataType = `*string`

		apiField.filterRule.Method = internal.ReturnType
		apiField.filterRule.DataType = append(apiField.filterRule.DataType, `max-length:`+v.FieldLimitStr)
		apiField.saveRule.Method = internal.ReturnType
		if v.FieldType == internal.TypeChar {
			apiField.saveRule.DataType = append(apiField.saveRule.DataType, `size:`+v.FieldLimitStr)
		} else {
			apiField.saveRule.DataType = append(apiField.saveRule.DataType, `max-length:`+v.FieldLimitStr)
		}
	case internal.TypeText, internal.TypeJson: // `text类型` // `json类型`
		// apiField.filterType.Method = internal.ReturnType
		apiField.filterType.DataType = `string`
		apiField.createType.Method = internal.ReturnType
		apiField.createType.DataType = `*string`
		apiField.updateType.Method = internal.ReturnType
		apiField.updateType.DataType = `*string`
		apiField.resType.Method = internal.ReturnType
		apiField.resType.DataType = `*string`

		if !v.IsNull {
			apiField.isRequired = true
		}
		if v.FieldType == internal.TypeJson {
			apiField.filterRule.Method = internal.ReturnType
			apiField.filterRule.DataType = append(apiField.filterRule.DataType, `json`)
			apiField.saveRule.Method = internal.ReturnType
			apiField.saveRule.DataType = append(apiField.saveRule.DataType, `json`)
		}
	case internal.TypeDatetime, internal.TypeTimestamp: // `datetime类型`	// `timestamp类型`
		// apiField.filterType.Method = internal.ReturnType
		apiField.filterType.DataType = `*gtime.Time`
		apiField.createType.Method = internal.ReturnType
		apiField.createType.DataType = `*gtime.Time`
		apiField.updateType.Method = internal.ReturnType
		apiField.updateType.DataType = `*gtime.Time`
		apiField.resType.Method = internal.ReturnType
		apiField.resType.DataType = `*gtime.Time`

		apiField.filterRule.Method = internal.ReturnType
		apiField.filterRule.DataType = append(apiField.filterRule.DataType, `date-format:Y-m-d H:i:s`)
		apiField.saveRule.Method = internal.ReturnType
		apiField.saveRule.DataType = append(apiField.saveRule.DataType, `date-format:Y-m-d H:i:s`)
	case internal.TypeDate: // `date类型`
		apiField.filterType.Method = internal.ReturnType
		apiField.filterType.DataType = `*gtime.Time`
		apiField.createType.Method = internal.ReturnType
		apiField.createType.DataType = `*gtime.Time`
		apiField.updateType.Method = internal.ReturnType
		apiField.updateType.DataType = `*gtime.Time`
		apiField.resType.Method = internal.ReturnType
		apiField.resType.DataType = `*string`

		apiField.filterRule.Method = internal.ReturnType
		apiField.filterRule.DataType = append(apiField.filterRule.DataType, `date-format:Y-m-d`)
		apiField.saveRule.Method = internal.ReturnType
		apiField.saveRule.DataType = append(apiField.saveRule.DataType, `date-format:Y-m-d`)
	case internal.TypeTime: // `time类型`
		// apiField.filterType.Method = internal.ReturnType
		apiField.filterType.DataType = `string`
		apiField.createType.Method = internal.ReturnType
		apiField.createType.DataType = `*string`
		apiField.updateType.Method = internal.ReturnType
		apiField.updateType.DataType = `*string`
		apiField.resType.Method = internal.ReturnType
		apiField.resType.DataType = `*string`

		apiField.filterRule.Method = internal.ReturnType
		apiField.filterRule.DataType = append(apiField.filterRule.DataType, `date-format:H:i:s`)
		apiField.saveRule.Method = internal.ReturnType
		apiField.saveRule.DataType = append(apiField.saveRule.DataType, `date-format:H:i:s`)
	default:
		apiField.filterType.Method = internal.ReturnType
		apiField.filterType.DataType = `string`
		apiField.createType.Method = internal.ReturnType
		apiField.createType.DataType = `*string`
		apiField.updateType.Method = internal.ReturnType
		apiField.updateType.DataType = `*string`
		apiField.resType.Method = internal.ReturnType
		apiField.resType.DataType = `*string`
	}
	/*--------根据字段数据类型处理（注意：这里的代码改动对字段命名类型处理有影响） 结束--------*/

	/*--------根据字段主键类型处理 开始--------*/
	switch v.FieldTypePrimary {
	case internal.TypePrimary: // 独立主键
		if v.FieldRaw == `id` {
			return myGenApiField{}
		}
		apiField.filterType.Method = internal.ReturnType
		apiField.updateType.Method = internal.ReturnEmpty
		// 创建和更新按数据类型和命名类型处理。但这种主键一般会在程序内封装ID生成逻辑（可在代码生成后修改）
	case internal.TypePrimaryAutoInc: // 独立主键（自增）
		if v.FieldRaw == `id` {
			return myGenApiField{}
		}
		apiField.filterType.Method = internal.ReturnType
		apiField.createType.Method = internal.ReturnEmpty
		apiField.updateType.Method = internal.ReturnEmpty
		// 跳过命名类型处理
		return
	case internal.TypePrimaryMany: // 联合主键
		apiField.filterType.Method = internal.ReturnType
		// 创建和更新按数据类型和命名类型处理
	case internal.TypePrimaryManyAutoInc: // 联合主键（自增）
		apiField.filterType.Method = internal.ReturnType
		apiField.createType.Method = internal.ReturnEmpty
		apiField.updateType.Method = internal.ReturnEmpty
		// 跳过命名类型处理
		return
	}
	/*--------根据字段主键类型处理 结束--------*/

	/*--------根据字段命名类型处理 开始--------*/
	switch v.FieldTypeName {
	case internal.TypeNameDeleted: // 软删除字段
		return myGenApiField{}
	case internal.TypeNameUpdated: // 更新时间字段
		apiField.filterType.Method = internal.ReturnEmpty
		apiField.createType.Method = internal.ReturnEmpty
		apiField.updateType.Method = internal.ReturnEmpty
	case internal.TypeNameCreated: // 创建时间字段
		apiField.filterType.Method = internal.ReturnEmpty
		apiField.createType.Method = internal.ReturnEmpty
		apiField.updateType.Method = internal.ReturnEmpty

		apiField.filterOfFixed = append(apiField.filterOfFixed,
			`TimeRangeStart *gtime.Time `+"`"+`json:"`+internal.GetStrByFieldStyle(tpl.FieldStyle, `time_range_start`)+`,omitempty" v:"date-format:Y-m-d H:i:s" dc:"开始时间：YYYY-mm-dd HH:ii:ss"`+"`",
			`TimeRangeEnd   *gtime.Time `+"`"+`json:"`+internal.GetStrByFieldStyle(tpl.FieldStyle, `time_range_end`)+`,omitempty" v:"date-format:Y-m-d H:i:s|after-equal:TimeRangeStart" dc:"结束时间：YYYY-mm-dd HH:ii:ss"`+"`",
		)
	case internal.TypeNamePid: // pid；	类型：int等类型；
		apiField.filterType.Method = internal.ReturnType

		apiField.resOfAdd = append(apiField.resOfAdd,
			internal.GetStrByFieldStyle(internal.FieldStyleCaseCamel, tpl.Handle.LabelList[0], `p`)+` *string `+"`"+`json:"`+internal.GetStrByFieldStyle(tpl.FieldStyle, tpl.Handle.LabelList[0], `p`)+`,omitempty" dc:"父级"`+"`",
			internal.GetStrByFieldStyle(internal.FieldStyleCaseCamel, `is_has_child`)+` *uint `+"`"+`json:"`+internal.GetStrByFieldStyle(tpl.FieldStyle, `is_has_child`)+`,omitempty" dc:"有子级：0否 1是"`+"`",
		)
	case internal.TypeNameLevel: // level，且pid,level,id_path|idPath同时存在时（才）有效；	类型：int等类型；
		apiField.filterType.Method = internal.ReturnType
		apiField.createType.Method = internal.ReturnEmpty
		apiField.updateType.Method = internal.ReturnEmpty
	case internal.TypeNameIdPath: // id_path|idPath，且pid,level,id_path|idPath同时存在时（才）有效；	类型：varchar或text；
		apiField.filterType.Method = internal.ReturnEmpty
		apiField.createType.Method = internal.ReturnEmpty
		apiField.updateType.Method = internal.ReturnEmpty
	case internal.TypeNamePasswordSuffix: // password,passwd后缀；	类型：char(32)；
		apiField.filterType.Method = internal.ReturnEmpty
		apiField.resType.Method = internal.ReturnEmpty

		apiField.isRequired = true
	case internal.TypeNameSaltSuffix: // salt后缀，且对应的password,passwd后缀存在时（才）有效；	类型：char；
		return myGenApiField{}
	case internal.TypeNameNameSuffix: // name,title后缀；	类型：varchar；
		if gstr.CaseCamel(tpl.Handle.LabelList[0]) == v.FieldCaseCamel {
			apiField.isRequired = true
		}
	case internal.TypeNameCodeSuffix: // code后缀；	类型：varchar；
		apiField.filterRule.Method = internal.ReturnUnion
		apiField.filterRule.DataTypeName = append(apiField.filterRule.DataTypeName, `regex:^[\\p{L}\\p{N}_-]+$`)
		apiField.saveRule.Method = internal.ReturnUnion
		apiField.saveRule.DataTypeName = append(apiField.saveRule.DataTypeName, `regex:^[\\p{L}\\p{N}_-]+$`)
	case internal.TypeNameAccountSuffix: // account后缀；	类型：varchar；
		/* apiField.filterRule.Method = internal.ReturnUnion
		apiField.filterRule.DataTypeName = append(apiField.filterRule.DataTypeName, `passport`)
		apiField.saveRule.Method = internal.ReturnUnion
		apiField.saveRule.DataTypeName = append(apiField.saveRule.DataTypeName, `passport`) */
		apiField.filterRule.Method = internal.ReturnUnion
		apiField.filterRule.DataTypeName = append(apiField.filterRule.DataTypeName, `regex:^[\\p{L}][\\p{L}\\p{N}_]{3,}$`)
		apiField.saveRule.Method = internal.ReturnUnion
		apiField.saveRule.DataTypeName = append(apiField.saveRule.DataTypeName, `regex:^[\\p{L}][\\p{L}\\p{N}_]{3,}$`)
	case internal.TypeNamePhoneSuffix: // phone,mobile后缀；	类型：varchar；
		apiField.filterRule.Method = internal.ReturnUnion
		apiField.filterRule.DataTypeName = append(apiField.filterRule.DataTypeName, `phone`)
		apiField.saveRule.Method = internal.ReturnUnion
		apiField.saveRule.DataTypeName = append(apiField.saveRule.DataTypeName, `phone`)
	case internal.TypeNameEmailSuffix: // email后缀；	类型：varchar；
		apiField.filterRule.Method = internal.ReturnUnion
		apiField.filterRule.DataTypeName = append(apiField.filterRule.DataTypeName, `email`)
		apiField.saveRule.Method = internal.ReturnUnion
		apiField.saveRule.DataTypeName = append(apiField.saveRule.DataTypeName, `email`)
	case internal.TypeNameUrlSuffix: // url,link后缀；	类型：varchar；
		apiField.filterRule.Method = internal.ReturnUnion
		apiField.filterRule.DataTypeName = append(apiField.filterRule.DataTypeName, `url`)
		apiField.saveRule.Method = internal.ReturnUnion
		apiField.saveRule.DataTypeName = append(apiField.saveRule.DataTypeName, `url`)
	case internal.TypeNameIpSuffix: // IP后缀；	类型：varchar；
		apiField.filterRule.Method = internal.ReturnUnion
		apiField.filterRule.DataTypeName = append(apiField.filterRule.DataTypeName, `ip`)
		apiField.saveRule.Method = internal.ReturnUnion
		apiField.saveRule.DataTypeName = append(apiField.saveRule.DataTypeName, `ip`)
	case internal.TypeNameColorSuffix: // color后缀；	类型：varchar；
	case internal.TypeNameIdSuffix: // id后缀；	类型：int等类型或varchar或char；
		apiField.filterType.Method = internal.ReturnType

		relIdObj := tpl.Handle.RelIdMap[v.FieldRaw]
		if relIdObj.tpl.Table != `` {
			if !apiField.isRequired && garray.NewIntArrayFrom([]int{internal.TypeInt, internal.TypeIntU}).Contains(v.FieldType) && relIdObj.tpl.KeyList[0].FieldList[0].IsAutoInc {
				for index, rule := range apiField.saveRule.DataType {
					if rule == `between:`+v.FieldLimitInt.Min+`,`+v.FieldLimitInt.Max {
						apiField.saveRule.DataType[index] = `between:0,` + v.FieldLimitInt.Max
					}
				}
			}

			if !relIdObj.IsRedundName {
				apiField.resOfAdd = append(apiField.resOfAdd, gstr.CaseCamel(relIdObj.tpl.Handle.LabelList[0])+relIdObj.SuffixCaseCamel+` *string `+"`"+`json:"`+relIdObj.tpl.Handle.LabelList[0]+relIdObj.Suffix+`,omitempty" dc:"`+relIdObj.FieldName+`"`+"`")
			}
		}
	case internal.TypeNameStatusSuffix: // status,type,scene,method,pos,position,gender,currency等后缀；	类型：int等类型或varchar或char；	注释：多状态之间用[\s,，.。;；]等字符分隔。示例（状态：0待处理 1已处理 2驳回 yes是 no否）
		apiField.filterType.Method = internal.ReturnType

		statusArr := make([]string, len(v.StatusList))
		for index, item := range v.StatusList {
			statusArr[index] = item[0]
		}
		statusStr := gstr.Join(statusArr, `,`)
		apiField.filterRule.Method = internal.ReturnTypeName
		apiField.filterRule.DataTypeName = append(apiField.filterRule.DataTypeName, `in:`+statusStr)
		apiField.saveRule.Method = internal.ReturnTypeName
		apiField.saveRule.DataTypeName = append(apiField.saveRule.DataTypeName, `in:`+statusStr)
	case internal.TypeNameIsPrefix: // is_前缀；	类型：int等类型；注释：多状态之间用[\s,，.。;；]等字符分隔。示例（停用：0否 1是）
		apiField.filterType.Method = internal.ReturnType

		apiField.filterRule.Method = internal.ReturnTypeName
		apiField.filterRule.DataTypeName = append(apiField.filterRule.DataTypeName, `in:0,1`)
		apiField.saveRule.Method = internal.ReturnTypeName
		apiField.saveRule.DataTypeName = append(apiField.saveRule.DataTypeName, `in:0,1`)
	case internal.TypeNameSortSuffix, internal.TypeNameNoSuffix: // sort,num,number,weight等后缀；	类型：int等类型；	// no,level,rank等后缀；	类型：int等类型；
	case internal.TypeNameStartPrefix: // start_前缀；	类型：datetime或date或timestamp或time；
		apiField.filterType.Method = internal.ReturnType
	case internal.TypeNameEndPrefix: // end_前缀；	类型：datetime或date或timestamp或time；
		apiField.filterType.Method = internal.ReturnType
	case internal.TypeNameRemarkSuffix: // remark,desc,msg,message,intro,content后缀；	类型：varchar或text；前端对应组件：varchar文本输入框，text富文本编辑器
		apiField.filterType.Method = internal.ReturnEmpty
	case internal.TypeNameImageSuffix, internal.TypeNameVideoSuffix, internal.TypeNameAudioSuffix, internal.TypeNameFileSuffix: // icon,cover,avatar,img,img_list,imgList,img_arr,imgArr,image,image_list,imageList,image_arr,imageArr等后缀；	类型：单图片varchar，多图片json或text	// video,video_list,videoList,video_arr,videoArr等后缀；	类型：单视频varchar，多视频json或text	// audio,audio_list,audioList,audio_arr,audioArr等后缀；	类型：单音频varchar，多音频json或text	// file,file_list,fileList,file_arr,fileArr等后缀；	类型：单文件varchar，多文件json或text
		if v.FieldType == internal.TypeVarchar {
			apiField.filterType.Method = internal.ReturnEmpty

			apiField.saveRule.Method = internal.ReturnUnion
			apiField.saveRule.DataTypeName = append(apiField.saveRule.DataTypeName, `url`)
		} else {
			apiField.createType.Method = internal.ReturnTypeName
			apiField.createType.DataTypeName = `*[]string`
			apiField.updateType.Method = internal.ReturnTypeName
			apiField.updateType.DataTypeName = `*[]string`
			apiField.resType.Method = internal.ReturnTypeName
			apiField.resType.DataTypeName = `*[]string`

			apiField.saveRule.Method = internal.ReturnTypeName
			apiField.saveRule.DataTypeName = append(apiField.saveRule.DataTypeName, `distinct`, `foreach`, `url`, `foreach`, `min-length:1`)
		}
	case internal.TypeNameArrSuffix: // list,arr等后缀；	类型：json或text；
		apiField.createType.Method = internal.ReturnTypeName
		apiField.createType.DataTypeName = `*[]any`
		apiField.updateType.Method = internal.ReturnTypeName
		apiField.updateType.DataTypeName = `*[]any`
		apiField.resType.Method = internal.ReturnTypeName
		apiField.resType.DataTypeName = `*[]any`

		apiField.saveRule.Method = internal.ReturnTypeName
		apiField.saveRule.DataTypeName = append(apiField.saveRule.DataTypeName, `distinct`)
	}
	/*--------根据字段命名类型处理 结束--------*/
	return
}

func getApiExtendMiddleOne(tplEM handleExtendMiddle) (api myGenApi) {
	switch tplEM.TableType {
	case internal.TableTypeExtendOne:
		for _, v := range tplEM.FieldList {
			api.Add(getApiField(tplEM.tpl, v), v, tplEM.TableType)
		}
	case internal.TableTypeMiddleOne:
		for _, v := range tplEM.FieldListOfIdSuffix {
			api.Add(getApiField(tplEM.tpl, v), v, tplEM.TableType)
		}
		for _, v := range tplEM.FieldListOfOther {
			api.Add(getApiField(tplEM.tpl, v), v, tplEM.TableType)
		}
	}
	return
}

func getApiExtendMiddleMany(tplEM handleExtendMiddle) (api myGenApi) {
	apiTmp := myGenApi{}
	for _, v := range tplEM.FieldList {
		apiTmp.Add(getApiField(tplEM.tpl, v), v, tplEM.TableType)
	}
	api.filter = append(api.filter, apiTmp.filter...)
	if len(tplEM.FieldList) == 1 {
		v := tplEM.FieldList[0]

		apiField := myGenApiField{}
		/*--------根据字段数据类型处理（注意：这里的代码改动对字段命名类型处理有影响） 开始--------*/
		switch v.FieldType {
		case internal.TypeInt, internal.TypeIntU: // `int等类型`	// `int等类型（unsigned）`
			dataType := `int`
			if v.FieldType == internal.TypeIntU {
				dataType = `uint`
			}
			apiField.createType.Method = internal.ReturnType
			apiField.createType.DataType = `*[]` + dataType
			apiField.updateType.Method = internal.ReturnType
			apiField.updateType.DataType = `*[]` + dataType
			apiField.resType.Method = internal.ReturnType
			apiField.resType.DataType = `[]` + dataType

			apiField.filterRule.Method = internal.ReturnType
			apiField.filterRule.DataType = append(apiField.filterRule.DataType, `foreach`, `between:`+v.FieldLimitInt.Min+`,`+v.FieldLimitInt.Max)
			apiField.saveRule.Method = internal.ReturnType
			apiField.saveRule.DataType = append(apiField.saveRule.DataType, `foreach`, `between:`+v.FieldLimitInt.Min+`,`+v.FieldLimitInt.Max)
		case internal.TypeFloat, internal.TypeFloatU: // `float等类型` // `float等类型（unsigned）`
			apiField.createType.Method = internal.ReturnType
			apiField.createType.DataType = `*[]float64`
			apiField.updateType.Method = internal.ReturnType
			apiField.updateType.DataType = `*[]float64`
			apiField.resType.Method = internal.ReturnType
			apiField.resType.DataType = `[]float64`

			if v.FieldLimitFloat.Min != `` && v.FieldLimitFloat.Max != `` {
				apiField.filterRule.Method = internal.ReturnType
				apiField.filterRule.DataType = append(apiField.filterRule.DataType, `foreach`, `between:`+v.FieldLimitFloat.Min+`,`+v.FieldLimitFloat.Max)
				apiField.saveRule.Method = internal.ReturnType
				apiField.saveRule.DataType = append(apiField.saveRule.DataType, `foreach`, `between:`+v.FieldLimitFloat.Min+`,`+v.FieldLimitFloat.Max)
			} else if v.FieldLimitFloat.Min != `` {
				apiField.filterRule.Method = internal.ReturnType
				apiField.filterRule.DataType = append(apiField.filterRule.DataType, `foreach`, `min:`+v.FieldLimitFloat.Min)
				apiField.saveRule.Method = internal.ReturnType
				apiField.saveRule.DataType = append(apiField.saveRule.DataType, `foreach`, `min:`+v.FieldLimitFloat.Min)
			} else if v.FieldLimitFloat.Max != `` {
				apiField.filterRule.Method = internal.ReturnType
				apiField.filterRule.DataType = append(apiField.filterRule.DataType, `foreach`, `max:`+v.FieldLimitFloat.Max)
				apiField.saveRule.Method = internal.ReturnType
				apiField.saveRule.DataType = append(apiField.saveRule.DataType, `foreach`, `max:`+v.FieldLimitFloat.Max)
			}
		/* // 注释掉的类型当作字符串处理
		case internal.TypeDatetime, internal.TypeTimestamp: // `datetime类型`	// `timestamp类型`
			apiField.createType.Method = internal.ReturnType
			apiField.createType.DataType = `*[]gtime.Time`
			apiField.updateType.Method = internal.ReturnType
			apiField.updateType.DataType = `*[]gtime.Time`
			apiField.resType.Method = internal.ReturnType
			apiField.resType.DataType = `*[]gtime.Time`

			apiField.saveRule.Method = internal.ReturnType
			apiField.saveRule.DataType = append(apiField.saveRule.DataType, `date-format:Y-m-d H:i:s`)
		case internal.TypeDate: // `date类型`
			apiField.createType.Method = internal.ReturnType
			apiField.createType.DataType = `*[]gtime.Time`
			apiField.updateType.Method = internal.ReturnType
			apiField.updateType.DataType = `*[]gtime.Time`
			apiField.resType.Method = internal.ReturnType
			apiField.resType.DataType = `[]string`

			apiField.saveRule.Method = internal.ReturnType
			apiField.saveRule.DataType = append(apiField.saveRule.DataType, `date-format:Y-m-d`)
		case internal.TypeTime: // `time类型`
			apiField.createType.Method = internal.ReturnType
			apiField.createType.DataType = `*[]string`
			apiField.updateType.Method = internal.ReturnType
			apiField.updateType.DataType = `*[]string`
			apiField.resType.Method = internal.ReturnType
			apiField.resType.DataType = `[]string`

			apiField.saveRule.Method = internal.ReturnType
			apiField.saveRule.DataType = append(apiField.saveRule.DataType, `date-format:H:i:s`) */
		default:
			apiField.createType.Method = internal.ReturnType
			apiField.createType.DataType = `*[]string`
			apiField.updateType.Method = internal.ReturnType
			apiField.updateType.DataType = `*[]string`
			apiField.resType.Method = internal.ReturnType
			apiField.resType.DataType = `[]string`

			switch v.FieldType {
			case internal.TypeVarchar:
				apiField.saveRule.Method = internal.ReturnType
				apiField.saveRule.DataType = append(apiField.saveRule.DataType, `foreach`, `max-length:`+v.FieldLimitStr)
			case internal.TypeChar:
				apiField.saveRule.Method = internal.ReturnType
				apiField.saveRule.DataType = append(apiField.saveRule.DataType, `foreach`, `size:`+v.FieldLimitStr)
			case internal.TypeJson:
				apiField.saveRule.Method = internal.ReturnType
				apiField.saveRule.DataType = append(apiField.saveRule.DataType, `foreach`, `json`)
			case internal.TypeDatetime, internal.TypeTimestamp:
				apiField.saveRule.Method = internal.ReturnType
				apiField.saveRule.DataType = append(apiField.saveRule.DataType, `date-format:Y-m-d H:i:s`)
			case internal.TypeDate:
				apiField.saveRule.Method = internal.ReturnType
				apiField.saveRule.DataType = append(apiField.saveRule.DataType, `date-format:Y-m-d`)
			case internal.TypeTime:
				apiField.saveRule.Method = internal.ReturnType
				apiField.saveRule.DataType = append(apiField.saveRule.DataType, `date-format:H:i:s`)
			}
		}
		/*--------根据字段数据类型处理（注意：这里的代码改动对字段命名类型处理有影响） 结束--------*/

		/*--------根据字段命名类型处理 开始--------*/
		switch v.FieldTypeName {
		case internal.TypeNameDeleted, internal.TypeNameUpdated, internal.TypeNameCreated: // 软删除字段 // 更新时间字段 // 创建时间字段
		case internal.TypeNamePid: // pid；	类型：int等类型；
		case internal.TypeNameLevel: // level，且pid,level,id_path|idPath同时存在时（才）有效；	类型：int等类型；
		case internal.TypeNameIdPath: // id_path|idPath，且pid,level,id_path|idPath同时存在时（才）有效；	类型：varchar或text；
		case internal.TypeNamePasswordSuffix: // password,passwd后缀；	类型：char(32)；
		case internal.TypeNameSaltSuffix: // salt后缀，且对应的password,passwd后缀存在时（才）有效；	类型：char；
		case internal.TypeNameNameSuffix: // name,title后缀；	类型：varchar；
		case internal.TypeNameCodeSuffix: // code后缀；	类型：varchar；
			apiField.saveRule.Method = internal.ReturnUnion
			apiField.saveRule.DataTypeName = append(apiField.saveRule.DataTypeName, `foreach`, `regex:^[\\p{L}\\p{N}_-]+$`)
		case internal.TypeNameAccountSuffix: // account后缀；	类型：varchar；
			/* apiField.saveRule.Method = internal.ReturnUnion
			apiField.saveRule.DataTypeName = append(apiField.saveRule.DataTypeName, `foreach`, `passport`) */
			apiField.saveRule.Method = internal.ReturnUnion
			apiField.saveRule.DataTypeName = append(apiField.saveRule.DataTypeName, `foreach`, `regex:^[\\p{L}][\\p{L}\\p{N}_]{3,}$`)
		case internal.TypeNamePhoneSuffix: // phone,mobile后缀；	类型：varchar；
			apiField.saveRule.Method = internal.ReturnUnion
			apiField.saveRule.DataTypeName = append(apiField.saveRule.DataTypeName, `foreach`, `phone`)
		case internal.TypeNameEmailSuffix: // email后缀；	类型：varchar；
			apiField.saveRule.Method = internal.ReturnUnion
			apiField.saveRule.DataTypeName = append(apiField.saveRule.DataTypeName, `foreach`, `email`)
		case internal.TypeNameUrlSuffix: // url,link后缀；	类型：varchar；
			apiField.saveRule.Method = internal.ReturnUnion
			apiField.saveRule.DataTypeName = append(apiField.saveRule.DataTypeName, `foreach`, `url`)
		case internal.TypeNameIpSuffix: // IP后缀；	类型：varchar；
			apiField.saveRule.Method = internal.ReturnUnion
			apiField.saveRule.DataTypeName = append(apiField.saveRule.DataTypeName, `foreach`, `ip`)
		case internal.TypeNameColorSuffix: // color后缀；	类型：varchar；
		case internal.TypeNameIdSuffix: // id后缀；	类型：int等类型或varchar或char；
		case internal.TypeNameStatusSuffix: // status,type,scene,method,pos,position,gender,currency等后缀；	类型：int等类型或varchar或char；	注释：多状态之间用[\s,，.。;；]等字符分隔。示例（状态：0待处理 1已处理 2驳回 yes是 no否）
			statusArr := make([]string, len(v.StatusList))
			for index, item := range v.StatusList {
				statusArr[index] = item[0]
			}
			statusStr := gstr.Join(statusArr, `,`)
			apiField.saveRule.Method = internal.ReturnTypeName
			apiField.saveRule.DataTypeName = append(apiField.saveRule.DataTypeName, `foreach`, `in:`+statusStr)
		case internal.TypeNameIsPrefix: // is_前缀；	类型：int等类型；注释：多状态之间用[\s,，.。;；]等字符分隔。示例（停用：0否 1是）
			apiField.saveRule.Method = internal.ReturnTypeName
			apiField.saveRule.DataTypeName = append(apiField.saveRule.DataTypeName, `foreach`, `in:0,1`)
		case internal.TypeNameSortSuffix, internal.TypeNameNoSuffix: // sort,num,number,weight等后缀；	类型：int等类型；	// no,level,rank等后缀；	类型：int等类型；
		case internal.TypeNameStartPrefix: // start_前缀；	类型：datetime或date或timestamp或time；
		case internal.TypeNameEndPrefix: // end_前缀；	类型：datetime或date或timestamp或time；
		case internal.TypeNameRemarkSuffix: // remark,desc,msg,message,intro,content后缀；	类型：varchar或text；前端对应组件：varchar文本输入框，text富文本编辑器
		case internal.TypeNameImageSuffix, internal.TypeNameVideoSuffix, internal.TypeNameAudioSuffix, internal.TypeNameFileSuffix: // icon,cover,avatar,img,img_list,imgList,img_arr,imgArr,image,image_list,imageList,image_arr,imageArr等后缀；	类型：单图片varchar，多图片json或text	// video,video_list,videoList,video_arr,videoArr等后缀；	类型：单视频varchar，多视频json或text	// audio,audio_list,audioList,audio_arr,audioArr等后缀；	类型：单音频varchar，多音频json或text	// file,file_list,fileList,file_arr,fileArr等后缀；	类型：单文件varchar，多文件json或text
			if v.FieldType == internal.TypeVarchar {
				apiField.saveRule.Method = internal.ReturnUnion
				apiField.saveRule.DataTypeName = append(apiField.saveRule.DataTypeName, `foreach`, `url`)
			}
		case internal.TypeNameArrSuffix: // list,arr等后缀；	类型：json或text；
		}
		/*--------根据字段命名类型处理 结束--------*/

		// apiField.saveRule.DataTypeName = append([]string{`distinct`}, apiField.saveRule.GetData()...)
		if apiField.createType.GetData() != `` {
			api.create = append(api.create, gstr.CaseCamel(tplEM.FieldVar)+` `+apiField.createType.GetData()+` `+"`"+`json:"`+tplEM.FieldVar+`,omitempty" v:"`+gstr.Join(append([]string{`distinct`}, apiField.saveRule.GetData()...), `|`)+`" dc:"`+v.FieldDesc+`"`+"`")
		}
		if apiField.updateType.GetData() != `` {
			api.update = append(api.update, gstr.CaseCamel(tplEM.FieldVar)+` `+apiField.updateType.GetData()+` `+"`"+`json:"`+tplEM.FieldVar+`,omitempty" filter:"-" v:"`+gstr.Join(append([]string{`distinct`}, apiField.saveRule.GetData()...), `|`)+`" dc:"`+v.FieldDesc+`"`+"`")
		}
		if apiField.resType.GetData() != `` {
			api.res = append(api.res, gstr.CaseCamel(tplEM.FieldVar)+` `+apiField.resType.GetData()+` `+"`"+`json:"`+tplEM.FieldVar+`,omitempty" dc:"`+v.FieldDesc+`"`+"`")
		}
	} else {
		api.create = append(api.create, gstr.CaseCamel(tplEM.FieldVar)+` *[]struct {`+gstr.Join(append([]string{``}, apiTmp.create...), `
		`)+`
	} `+"`"+`json:"`+tplEM.FieldVar+`,omitempty" filter:"-" v:"" dc:"列表"`+"`")
		api.update = append(api.update, gstr.CaseCamel(tplEM.FieldVar)+` *[]struct {`+gstr.Join(append([]string{``}, apiTmp.update...), `
		`)+`
	} `+"`"+`json:"`+tplEM.FieldVar+`,omitempty" v:"" dc:"列表"`+"`")
		api.res = append(api.res, gstr.CaseCamel(tplEM.FieldVar)+` []struct {`+gstr.Join(append([]string{``}, apiTmp.res...), `
		`)+`
	} `+"`"+`json:"`+tplEM.FieldVar+`,omitempty" dc:"列表"`+"`")
	}
	return
}
