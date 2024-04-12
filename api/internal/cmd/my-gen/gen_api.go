package my_gen

import (
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

	filterType myGenDataStrHandler
	createType myGenDataStrHandler
	updateType myGenDataStrHandler
	resType    myGenDataStrHandler

	isRequired bool
	filterRule myGenDataSliceHandler
	saveRule   myGenDataSliceHandler
}

func (apiThis *myGenApi) Add(apiField myGenApiField, field myGenField, tableType myGenTableType) {
	apiThis.filterOfFixed = append(apiThis.filterOfFixed, apiField.filterOfFixed...)
	apiThis.resOfAdd = append(apiThis.resOfAdd, apiField.resOfAdd...)
	if apiField.filterType.getData() != `` {
		apiThis.filter = append(apiThis.filter, field.FieldCaseCamel+` `+apiField.filterType.getData()+` `+"`"+`json:"`+field.FieldRaw+`,omitempty" v:"`+gstr.Join(apiField.filterRule.getData(), `|`)+`" dc:"`+field.FieldDesc+`"`+"`")
	}
	if apiField.createType.getData() != `` {
		saveRuleArr := apiField.saveRule.getData()
		if apiField.isRequired && garray.NewFrom([]interface{}{TableTypeDefault, TableTypeExtendOne, TableTypeMiddleOne}).Contains(tableType) {
			saveRuleArr = append([]string{`required`}, saveRuleArr...)
		}
		apiThis.create = append(apiThis.create, field.FieldCaseCamel+` `+apiField.createType.getData()+` `+"`"+`json:"`+field.FieldRaw+`,omitempty" v:"`+gstr.Join(saveRuleArr, `|`)+`" dc:"`+field.FieldDesc+`"`+"`")
	}
	if apiField.updateType.getData() != `` {
		apiThis.update = append(apiThis.update, field.FieldCaseCamel+` `+apiField.updateType.getData()+` `+"`"+`json:"`+field.FieldRaw+`,omitempty" v:"`+gstr.Join(apiField.saveRule.getData(), `|`)+`" dc:"`+field.FieldDesc+`"`+"`")
	}
	if apiField.resType.getData() != `` {
		apiThis.res = append(apiThis.res, field.FieldCaseCamel+` `+apiField.resType.getData()+` `+"`"+`json:"`+field.FieldRaw+`,omitempty" dc:"`+field.FieldDesc+`"`+"`")
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
		api.Add(getApiField(tpl, v), v, TableTypeDefault)
	}
	for _, v := range tpl.FieldListOfAfter {
		api.Add(getApiField(tpl, v), v, TableTypeDefault)
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
	api.Unique()

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

type ` + tpl.TableCaseCamel + `ListFilter struct {` + gstr.Join(append([]string{``}, api.filterOfFixed...), `
	`) + gstr.Join(append([]string{``}, api.filter...), `
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

type ` + tpl.TableCaseCamel + `ListItem struct {` + gstr.Join(append([]string{``}, api.res...), `
	`) + gstr.Join(append([]string{``}, api.resOfAdd...), `
	`) + `
}

/*--------列表 结束--------*/

`
	}
	if option.IsInfo {
		tplApi += `/*--------详情 开始--------*/
type ` + tpl.TableCaseCamel + `InfoReq struct {
	g.Meta ` + "`" + `path:"/` + tpl.TableCaseKebab + `/info" method:"post" tags:"` + option.SceneInfo[daoAuth.Scene.Columns().SceneName].String() + `/` + option.CommonName + `" sm:"详情"` + "`" + `
	Field  []string ` + "`" + `json:"field" v:"distinct|foreach|min-length:1" dc:"查询字段，传值参考返回的字段名，默认返回全部字段。注意：如前端页面所需字段较少，建议传指定字段，可大幅减轻服务器及数据库压力"` + "`" + gstr.Join(append([]string{``}, api.info...), `
	`) + `
}

type ` + tpl.TableCaseCamel + `InfoRes struct {
	Info ` + tpl.TableCaseCamel + `Info ` + "`" + `json:"info" dc:"详情"` + "`" + `
}

type ` + tpl.TableCaseCamel + `Info struct {` + gstr.Join(append([]string{``}, api.res...), `
	`) + `
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
	Filter ` + tpl.TableCaseCamel + `ListFilter ` + "`" + `json:"filter" dc:"过滤条件"` + "`" + `
}

type ` + tpl.TableCaseCamel + `TreeRes struct {
	Tree []` + tpl.TableCaseCamel + `TreeItem ` + "`" + `json:"tree" dc:"列表（树状）"` + "`" + `
}

type ` + tpl.TableCaseCamel + `TreeItem struct {` + gstr.Join(append([]string{``}, api.res...), `
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

func getApiIdAndLabel(tpl myGenTpl) (api myGenApi) {
	if len(tpl.Handle.Id.List) == 1 {
		switch tpl.Handle.Id.List[0].FieldType {
		case TypeInt:
			api.filterOfFixed = append(api.filterOfFixed,
				`Id *int `+"`"+`json:"id,omitempty" v:"" dc:"ID"`+"`",
				`IdArr []int `+"`"+`json:"idArr,omitempty" v:"distinct" dc:"ID数组"`+"`",
				`ExcId *int `+"`"+`json:"excId,omitempty" v:"" dc:"排除ID"`+"`",
				`ExcIdArr []int `+"`"+`json:"excIdArr,omitempty" v:"distinct" dc:"排除ID数组"`+"`",
			)
			api.info = append(api.info, `Id int `+"`"+`json:"id" v:"required" dc:"ID"`+"`")
			api.update = append(api.update, `IdArr []int `+"`"+`json:"idArr,omitempty" v:"required|distinct" dc:"ID数组"`+"`")
			api.delete = append(api.delete, `IdArr []int `+"`"+`json:"idArr,omitempty" v:"required|distinct" dc:"ID数组"`+"`")
			api.res = append(api.res, `Id *int `+"`"+`json:"id,omitempty" dc:"ID"`+"`")
		case TypeIntU:
			api.filterOfFixed = append(api.filterOfFixed,
				`Id *uint `+"`"+`json:"id,omitempty" v:"min:1" dc:"ID"`+"`",
				`IdArr []uint `+"`"+`json:"idArr,omitempty" v:"distinct|foreach|min:1" dc:"ID数组"`+"`",
				`ExcId *uint `+"`"+`json:"excId,omitempty" v:"min:1" dc:"排除ID"`+"`",
				`ExcIdArr []uint `+"`"+`json:"excIdArr,omitempty" v:"distinct|foreach|min:1" dc:"排除ID数组"`+"`",
			)
			api.info = append(api.info, `Id uint `+"`"+`json:"id" v:"required|min:1" dc:"ID"`+"`")
			api.update = append(api.update, `IdArr []uint `+"`"+`json:"idArr,omitempty" v:"required|distinct|foreach|min:1" dc:"ID数组"`+"`")
			api.delete = append(api.delete, `IdArr []uint `+"`"+`json:"idArr,omitempty" v:"required|distinct|foreach|min:1" dc:"ID数组"`+"`")
			api.res = append(api.res, `Id *uint `+"`"+`json:"id,omitempty" dc:"ID"`+"`")
		default:
			api.filterOfFixed = append(api.filterOfFixed,
				`Id string `+"`"+`json:"id,omitempty" v:"max-length:`+tpl.Handle.Id.List[0].FieldLimitStr+`" dc:"ID"`+"`",
				`IdArr []string `+"`"+`json:"idArr,omitempty" v:"distinct|foreach|length:1,`+tpl.Handle.Id.List[0].FieldLimitStr+`" dc:"ID数组"`+"`",
				`ExcId string `+"`"+`json:"excId,omitempty" v:"max-length:`+tpl.Handle.Id.List[0].FieldLimitStr+`" dc:"排除ID"`+"`",
				`ExcIdArr []string `+"`"+`json:"excIdArr,omitempty" v:"distinct|foreach|length:1,`+tpl.Handle.Id.List[0].FieldLimitStr+`" dc:"排除ID数组"`+"`",
			)
			api.info = append(api.info, `Id string `+"`"+`json:"id" v:"required|max-length:`+tpl.Handle.Id.List[0].FieldLimitStr+`" dc:"ID"`+"`")
			api.update = append(api.update, `IdArr []string `+"`"+`json:"idArr,omitempty" v:"required|distinct|foreach|length:1,`+tpl.Handle.Id.List[0].FieldLimitStr+`" dc:"ID数组"`+"`")
			api.delete = append(api.delete, `IdArr []string `+"`"+`json:"idArr,omitempty" v:"required|distinct|foreach|length:1,`+tpl.Handle.Id.List[0].FieldLimitStr+`" dc:"ID数组"`+"`")
			api.res = append(api.res, `Id *string `+"`"+`json:"id,omitempty" dc:"ID"`+"`")
		}
	} else {
		api.filterOfFixed = append(api.filterOfFixed,
			`Id string `+"`"+`json:"id,omitempty" v:"" dc:"ID"`+"`",
			`IdArr []string `+"`"+`json:"idArr,omitempty" v:"distinct|foreach|min-length:1" dc:"ID数组"`+"`",
			`ExcId string `+"`"+`json:"excId,omitempty" v:"" dc:"排除ID"`+"`",
			`ExcIdArr []string `+"`"+`json:"excIdArr,omitempty" v:"distinct|foreach|min-length:1" dc:"排除ID数组"`+"`",
		)
		api.info = append(api.info, `Id string `+"`"+`json:"id" v:"required" dc:"ID"`+"`")
		api.update = append(api.update, `IdArr []string `+"`"+`json:"idArr,omitempty" v:"required|distinct|foreach|min-length:1" dc:"ID数组"`+"`")
		api.delete = append(api.delete, `IdArr []string `+"`"+`json:"idArr,omitempty" v:"required|distinct|foreach|min-length:1" dc:"ID数组"`+"`")
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
	case TypeInt: // `int等类型` // `int等类型（unsigned）`
		// apiField.filterType.Method = ReturnType
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
		if gconv.Uint(v.FieldLimitStr) <= configMaxLenOfStrFilter {
			apiField.filterType.Method = ReturnType
			apiField.filterType.DataType = `string`
		}
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
	case TypeChar: // `char类型`
		if gconv.Uint(v.FieldLimitStr) <= configMaxLenOfStrFilter {
			apiField.filterType.Method = ReturnType
			apiField.filterType.DataType = `string`
		}
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
	case TypeText, TypeJson: // `text类型` // `json类型`
		// apiField.filterType.Method = ReturnType
		apiField.filterType.DataType = `string`
		apiField.createType.Method = ReturnType
		apiField.createType.DataType = `*string`
		apiField.updateType.Method = ReturnType
		apiField.updateType.DataType = `*string`
		apiField.resType.Method = ReturnType
		apiField.resType.DataType = `*string`

		if !v.IsNull {
			apiField.isRequired = true
		}
		if v.FieldType == TypeJson {
			apiField.filterRule.Method = ReturnType
			apiField.filterRule.DataType = append(apiField.filterRule.DataType, `json`)
			apiField.saveRule.Method = ReturnType
			apiField.saveRule.DataType = append(apiField.saveRule.DataType, `json`)
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

	/*--------根据字段主键类型处理 开始--------*/
	switch v.FieldTypePrimary {
	case TypePrimary: // 独立主键
		if v.FieldRaw == `id` {
			return myGenApiField{}
		}
		apiField.filterType.Method = ReturnType
		// 创建和更新按数据类型和命名类型处理。但这种主键一般会在程序内封装ID生成逻辑（可在代码生成后修改）
	case TypePrimaryAutoInc: // 独立主键（自增）
		if v.FieldRaw == `id` {
			return myGenApiField{}
		}
		apiField.filterType.Method = ReturnType
		apiField.createType.Method = ReturnEmpty
		apiField.updateType.Method = ReturnEmpty

		apiField.filterRule.Method = ReturnUnion
		apiField.filterRule.DataTypeName = append(apiField.filterRule.DataTypeName, `min:1`)
		// 跳过命名类型处理
		return
	case TypePrimaryMany: // 联合主键
		apiField.filterType.Method = ReturnType
		// 创建和更新按数据类型和命名类型处理
	case TypePrimaryManyAutoInc: // 联合主键（自增）
		apiField.filterType.Method = ReturnType
		apiField.createType.Method = ReturnEmpty
		apiField.updateType.Method = ReturnEmpty

		apiField.filterRule.Method = ReturnUnion
		apiField.filterRule.DataTypeName = append(apiField.filterRule.DataTypeName, `min:1`)
		// 跳过命名类型处理
		return
	}
	/*--------根据字段主键类型处理 结束--------*/

	/*--------根据字段命名类型处理 开始--------*/
	switch v.FieldTypeName {
	case TypeNameDeleted: // 软删除字段
		return myGenApiField{}
	case TypeNameUpdated: // 更新时间字段
		apiField.filterType.Method = ReturnEmpty
		apiField.createType.Method = ReturnEmpty
		apiField.updateType.Method = ReturnEmpty
	case TypeNameCreated: // 创建时间字段
		apiField.filterType.Method = ReturnEmpty
		apiField.createType.Method = ReturnEmpty
		apiField.updateType.Method = ReturnEmpty

		apiField.filterOfFixed = append(apiField.filterOfFixed,
			`TimeRangeStart *gtime.Time `+"`"+`json:"timeRangeStart,omitempty" v:"date-format:Y-m-d H:i:s" dc:"开始时间：YYYY-mm-dd HH:ii:ss"`+"`",
			`TimeRangeEnd   *gtime.Time `+"`"+`json:"timeRangeEnd,omitempty" v:"date-format:Y-m-d H:i:s|after-equal:TimeRangeStart" dc:"结束时间：YYYY-mm-dd HH:ii:ss"`+"`",
		)
	case TypeNamePid: // pid；	类型：int等类型；
		apiField.filterType.Method = ReturnType

		apiField.resOfAdd = append(apiField.resOfAdd, `P`+gstr.CaseCamel(tpl.Handle.LabelList[0])+` *string `+"`"+`json:"p`+gstr.CaseCamel(tpl.Handle.LabelList[0])+`,omitempty" dc:"父级"`+"`")
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
		return myGenApiField{}
	case TypeNameNameSuffix: // name,title后缀；	类型：varchar；
		if gstr.CaseCamel(tpl.Handle.LabelList[0]) == v.FieldCaseCamel {
			apiField.isRequired = true
		}
	case TypeNameCodeSuffix: // code后缀；	类型：varchar；
		apiField.filterRule.Method = ReturnUnion
		apiField.filterRule.DataTypeName = append(apiField.filterRule.DataTypeName, `regex:^[\\p{L}\\p{N}_-]+$`)
		apiField.saveRule.Method = ReturnUnion
		apiField.saveRule.DataTypeName = append(apiField.saveRule.DataTypeName, `regex:^[\\p{L}\\p{N}_-]+$`)
	case TypeNameAccountSuffix: // account后缀；	类型：varchar；
		/* apiField.filterRule.Method = ReturnUnion
		apiField.filterRule.DataTypeName = append(apiField.filterRule.DataTypeName, `passport`)
		apiField.saveRule.Method = ReturnUnion
		apiField.saveRule.DataTypeName = append(apiField.saveRule.DataTypeName, `passport`) */
		apiField.filterRule.Method = ReturnUnion
		apiField.filterRule.DataTypeName = append(apiField.filterRule.DataTypeName, `regex:^[\\p{L}][\\p{L}\\p{N}_]{3,}$`)
		apiField.saveRule.Method = ReturnUnion
		apiField.saveRule.DataTypeName = append(apiField.saveRule.DataTypeName, `regex:^[\\p{L}][\\p{L}\\p{N}_]{3,}$`)
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
		if apiField.isRequired {
			apiField.saveRule.DataTypeName = append(apiField.saveRule.DataTypeName, `min:1`)
		} else {
			apiField.saveRule.DataTypeName = append(apiField.saveRule.DataTypeName, `min:0`)
		}

		relIdObj := tpl.Handle.RelIdMap[v.FieldRaw]
		if relIdObj.tpl.Table != `` && !relIdObj.IsRedundName {
			apiField.resOfAdd = append(apiField.resOfAdd, gstr.CaseCamel(relIdObj.tpl.Handle.LabelList[0])+gstr.CaseCamel(relIdObj.Suffix)+` *string `+"`"+`json:"`+relIdObj.tpl.Handle.LabelList[0]+relIdObj.Suffix+`,omitempty" dc:"`+relIdObj.FieldName+`"`+"`")
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
	}
	/*--------根据字段命名类型处理 结束--------*/
	return
}

func getApiExtendMiddleOne(tplEM handleExtendMiddle) (api myGenApi) {
	switch tplEM.TableType {
	case TableTypeExtendOne:
		for _, v := range tplEM.FieldList {
			api.Add(getApiField(tplEM.tpl, v), v, tplEM.TableType)
		}
	case TableTypeMiddleOne:
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
		case TypeInt: // `int等类型` // `int等类型（unsigned）`
			apiField.createType.Method = ReturnType
			apiField.createType.DataType = `*[]int`
			apiField.updateType.Method = ReturnType
			apiField.updateType.DataType = `*[]int`
			apiField.resType.Method = ReturnType
			apiField.resType.DataType = `[]int`
		case TypeIntU: // `int等类型（unsigned）`
			apiField.createType.Method = ReturnType
			apiField.createType.DataType = `*[]uint`
			apiField.updateType.Method = ReturnType
			apiField.updateType.DataType = `*[]uint`
			apiField.resType.Method = ReturnType
			apiField.resType.DataType = `[]uint`
		case TypeFloat, TypeFloatU: // `float等类型` // `float等类型（unsigned）`
			apiField.createType.Method = ReturnType
			apiField.createType.DataType = `*[]float64`
			apiField.updateType.Method = ReturnType
			apiField.updateType.DataType = `*[]float64`
			apiField.resType.Method = ReturnType
			apiField.resType.DataType = `[]float64`
			if v.FieldType == TypeFloatU {
				apiField.saveRule.Method = ReturnType
				apiField.saveRule.DataType = append(apiField.saveRule.DataType, `foreach`, `min:0`)
			}
		/* // 注释掉的类型当作字符串处理
		case TypeTimestamp, TypeDatetime: // `timestamp类型` // `datetime类型`
			apiField.createType.Method = ReturnType
			apiField.createType.DataType = `*[]gtime.Time`
			apiField.updateType.Method = ReturnType
			apiField.updateType.DataType = `*[]gtime.Time`
			apiField.resType.Method = ReturnType
			apiField.resType.DataType = `*[]gtime.Time`

			apiField.saveRule.Method = ReturnType
			apiField.saveRule.DataType = append(apiField.saveRule.DataType, `date-format:Y-m-d H:i:s`)
		case TypeDate: // `date类型`
			apiField.createType.Method = ReturnType
			apiField.createType.DataType = `*[]gtime.Time`
			apiField.updateType.Method = ReturnType
			apiField.updateType.DataType = `*[]gtime.Time`
			apiField.resType.Method = ReturnType
			apiField.resType.DataType = `[]string`

			apiField.saveRule.Method = ReturnType
			apiField.saveRule.DataType = append(apiField.saveRule.DataType, `date-format:Y-m-d`) */
		default:
			apiField.createType.Method = ReturnType
			apiField.createType.DataType = `*[]string`
			apiField.updateType.Method = ReturnType
			apiField.updateType.DataType = `*[]string`
			apiField.resType.Method = ReturnType
			apiField.resType.DataType = `[]string`

			switch v.FieldType {
			case TypeVarchar:
				apiField.saveRule.Method = ReturnType
				apiField.saveRule.DataType = append(apiField.saveRule.DataType, `foreach`, `max-length:`+v.FieldLimitStr)
			case TypeChar:
				apiField.saveRule.Method = ReturnType
				apiField.saveRule.DataType = append(apiField.saveRule.DataType, `foreach`, `size:`+v.FieldLimitStr)
			case TypeJson:
				apiField.saveRule.Method = ReturnType
				apiField.saveRule.DataType = append(apiField.saveRule.DataType, `foreach`, `json`)
			case TypeTimestamp, TypeDatetime:
				apiField.saveRule.Method = ReturnType
				apiField.saveRule.DataType = append(apiField.saveRule.DataType, `date-format:Y-m-d H:i:s`)
			case TypeDate:
				apiField.saveRule.Method = ReturnType
				apiField.saveRule.DataType = append(apiField.saveRule.DataType, `date-format:Y-m-d`)
			}
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
			apiField.saveRule.Method = ReturnUnion
			apiField.saveRule.DataTypeName = append(apiField.saveRule.DataTypeName, `foreach`, `regex:^[\\p{L}\\p{N}_-]+$`)
		case TypeNameAccountSuffix: // account后缀；	类型：varchar；
			/* apiField.saveRule.Method = ReturnUnion
			apiField.saveRule.DataTypeName = append(apiField.saveRule.DataTypeName, `foreach`, `passport`) */
			apiField.saveRule.Method = ReturnUnion
			apiField.saveRule.DataTypeName = append(apiField.saveRule.DataTypeName, `foreach`, `regex:^[\\p{L}][\\p{L}\\p{N}_]{3,}$`)
		case TypeNamePhoneSuffix: // phone,mobile后缀；	类型：varchar；
			apiField.saveRule.Method = ReturnUnion
			apiField.saveRule.DataTypeName = append(apiField.saveRule.DataTypeName, `foreach`, `phone`)
		case TypeNameEmailSuffix: // email后缀；	类型：varchar；
			apiField.saveRule.Method = ReturnUnion
			apiField.saveRule.DataTypeName = append(apiField.saveRule.DataTypeName, `foreach`, `email`)
		case TypeNameUrlSuffix: // url,link后缀；	类型：varchar；
			apiField.saveRule.Method = ReturnUnion
			apiField.saveRule.DataTypeName = append(apiField.saveRule.DataTypeName, `foreach`, `url`)
		case TypeNameIpSuffix: // IP后缀；	类型：varchar；
			apiField.saveRule.Method = ReturnUnion
			apiField.saveRule.DataTypeName = append(apiField.saveRule.DataTypeName, `foreach`, `ip`)
		case TypeNameIdSuffix: // id后缀；	类型：int等类型；
			apiField.saveRule.Method = ReturnUnion
			apiField.saveRule.DataTypeName = append(apiField.saveRule.DataTypeName, `foreach`, `min:1`)
		case TypeNameSortSuffix, TypeNameSort: // sort,weight等后缀；	类型：int等类型； // sort，且pid,level,idPath|id_path,sort同时存在时（才）有效；	类型：int等类型；
			apiField.saveRule.Method = ReturnUnion
			apiField.saveRule.DataTypeName = append(apiField.saveRule.DataTypeName, `foreach`, `between:0,100`)
		case TypeNameStatusSuffix: // status,type,method,pos,position,gender等后缀；	类型：int等类型或varchar或char；	注释：多状态之间用[\s,，;；]等字符分隔。示例（状态：0待处理 1已处理 2驳回 yes是 no否）
			statusArr := make([]string, len(v.StatusList))
			for index, item := range v.StatusList {
				statusArr[index] = item[0]
			}
			statusStr := gstr.Join(statusArr, `,`)
			apiField.saveRule.Method = ReturnUnion
			apiField.saveRule.DataTypeName = append(apiField.saveRule.DataTypeName, `foreach`, `in:`+statusStr)
		case TypeNameIsPrefix: // is_前缀；		类型：int等类型；注释：多状态之间用[\s,，;；]等字符分隔。示例（停用：0否 1是）
			apiField.saveRule.Method = ReturnUnion
			apiField.saveRule.DataTypeName = append(apiField.saveRule.DataTypeName, `foreach`, `in:0,1`)
		case TypeNameStartPrefix: // start_前缀；	类型：timestamp或datetime或date；
		case TypeNameEndPrefix: // end_前缀；	类型：timestamp或datetime或date；
		case TypeNameRemarkSuffix: // remark,desc,msg,message,intro,content后缀；	类型：varchar或text；前端对应组件：varchar文本输入框，text富文本编辑器
		case TypeNameImageSuffix, TypeNameVideoSuffix: // icon,cover,avatar,img,img_list,imgList,img_arr,imgArr,image,image_list,imageList,image_arr,imageArr等后缀；	类型：单图片varchar，多图片json或text	// video,video_list,videoList,video_arr,videoArr等后缀；		类型：单视频varchar，多视频json或text
			if v.FieldType == TypeVarchar {
				apiField.saveRule.Method = ReturnUnion
				apiField.saveRule.DataTypeName = append(apiField.saveRule.DataTypeName, `foreach`, `url`)
			}
		case TypeNameArrSuffix: // list,arr等后缀；	类型：json或text；
		}
		/*--------根据字段命名类型处理 结束--------*/

		apiField.saveRule.DataTypeName = append([]string{`distinct`}, apiField.saveRule.DataTypeName...)
		if apiField.createType.getData() != `` {
			api.create = append(api.create, gstr.CaseCamel(tplEM.FieldVar)+` `+apiField.createType.getData()+` `+"`"+`json:"`+tplEM.FieldVar+`,omitempty" v:"`+gstr.Join(apiField.saveRule.getData(), `|`)+`" dc:"`+v.FieldDesc+`列表"`+"`")
		}
		if apiField.updateType.getData() != `` {
			api.update = append(api.update, gstr.CaseCamel(tplEM.FieldVar)+` `+apiField.updateType.getData()+` `+"`"+`json:"`+tplEM.FieldVar+`,omitempty" v:"`+gstr.Join(apiField.saveRule.getData(), `|`)+`" dc:"`+v.FieldDesc+`列表"`+"`")
		}
		if apiField.resType.getData() != `` {
			api.res = append(api.res, gstr.CaseCamel(tplEM.FieldVar)+` `+apiField.resType.getData()+` `+"`"+`json:"`+tplEM.FieldVar+`,omitempty" dc:"`+v.FieldDesc+`列表"`+"`")
		}
	} else {
		api.create = append(api.create, gstr.CaseCamel(tplEM.FieldVar)+` *[]struct {`+gstr.Join(append([]string{``}, apiTmp.create...), `
		`)+`
	} `+"`"+`json:"`+tplEM.FieldVar+`,omitempty" v:"" dc:"列表"`+"`")
		api.update = append(api.update, gstr.CaseCamel(tplEM.FieldVar)+` *[]struct {`+gstr.Join(append([]string{``}, apiTmp.update...), `
		`)+`
	} `+"`"+`json:"`+tplEM.FieldVar+`,omitempty" v:"" dc:"列表"`+"`")
		api.res = append(api.res, gstr.CaseCamel(tplEM.FieldVar)+` []struct {`+gstr.Join(append([]string{``}, apiTmp.res...), `
		`)+`
	} `+"`"+`json:"`+tplEM.FieldVar+`,omitempty" dc:"列表"`+"`")
	}
	return
}
