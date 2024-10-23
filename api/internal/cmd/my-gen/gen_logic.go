package my_gen

import (
	"api/internal/cmd/my-gen/internal"
	"api/internal/utils"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
)

type myGenLogic struct {
	importDao         []string
	i18n              myGenI18n
	verifyData        []string
	verifyDataFunc    string
	verifyDataFuncRun string
	create            []string
	update            []string
	delete            []string
}

type myGenLogicField struct {
	importDao     []string
	i18nField     myGenI18nField
	verifyDataStr string
}

func (logicThis *myGenLogic) Add(logicField myGenLogicField) {
	logicThis.importDao = append(logicThis.importDao, logicField.importDao...)
	logicThis.i18n.Add(logicField.i18nField)
	if logicField.verifyDataStr != `` {
		logicThis.verifyData = append(logicThis.verifyData, logicField.verifyDataStr)
	}
}

func (logicThis *myGenLogic) Merge(logicOther myGenLogic) {
	logicThis.importDao = append(logicThis.importDao, logicOther.importDao...)
	logicThis.i18n.Merge(logicOther.i18n)
	logicThis.verifyData = append(logicThis.verifyData, logicOther.verifyData...)
	logicThis.create = append(logicThis.create, logicOther.create...)
	logicThis.update = append(logicThis.update, logicOther.update...)
	logicThis.delete = append(logicThis.delete, logicOther.delete...)
}

func (logicThis *myGenLogic) Unique() {
	logicThis.importDao = garray.NewStrArrayFrom(logicThis.importDao).Unique().Slice()
	logicThis.i18n.Unique()
}

// logic生成
func genLogic(option myGenOption, tpl myGenTpl) (i18n myGenI18n) {
	saveFile := gfile.SelfDir() + `/internal/logic/` + gstr.Replace(tpl.ModuleDirCaseKebab, `/`, `-`) + `/` + tpl.TableCaseSnake + `.go`
	if !option.IsResetLogic && gfile.IsFile(saveFile) {
		return
	}

	daoPath := `dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel
	logic := myGenLogic{}
	for _, v := range tpl.FieldList {
		logic.Add(getLogicField(tpl, v))
	}
	if tpl.Handle.Pid.Pid != `` {
		i18nField := myGenI18nField{
			key: `name.pid`,
			val: `父级`,
		}
		logic.i18n.Add(i18nField)
		logic.create = append(logic.create, `if _, ok := data[`+daoPath+`.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.Pid)+`]; ok && gconv.Uint(data[`+daoPath+`.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.Pid)+`]) > 0 {
		pInfo, _ := daoModelThis.CloneNew().Filter(`+daoPath+`.Columns().`+tpl.Handle.Id.List[0].FieldCaseCamel+`, data[`+daoPath+`.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.Pid)+`]).One()
		if pInfo.IsEmpty() {
			err = utils.NewErrorCode(ctx, 29999997, `+"``"+`, g.Map{`+"`i18nValues`"+`: []any{g.I18n().T(ctx, `+"`"+i18nField.key+"`"+`)}})
			return
		}
	}`)

		updateAddStr := ``
		if tpl.Handle.Pid.IsCoexist {
			updateAddStr = `
		for _, id := range daoModelThis.IdArr {
			if garray.NewStrArrayFrom(gstr.Split(pInfo[` + daoPath + `.Columns().` + gstr.CaseCamel(tpl.Handle.Pid.IdPath) + `].String(), ` + "`-`" + `)).Contains(gconv.String(id)) {
				err = utils.NewErrorCode(ctx, 29999995, ` + "``" + `)
				return
			}
		}`
		}
		logic.update = append(logic.update, `if _, ok := data[`+daoPath+`.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.Pid)+`]; ok && gconv.Uint(data[`+daoPath+`.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.Pid)+`]) > 0 {
		if garray.NewArrayFrom(gconv.SliceAny(gconv.SliceUint(daoModelThis.IdArr))).Contains(gconv.Uint(data[`+daoPath+`.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.Pid)+`])) {
			err = utils.NewErrorCode(ctx, 29999996, `+"``"+`)
			return
		}
		pInfo, _ := daoModelThis.CloneNew().Filter(`+daoPath+`.Columns().`+tpl.Handle.Id.List[0].FieldCaseCamel+`, data[`+daoPath+`.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.Pid)+`]).One()
		if pInfo.IsEmpty() {
			err = utils.NewErrorCode(ctx, 29999997, `+"``"+`, g.Map{`+"`i18nValues`"+`: []any{g.I18n().T(ctx, `+"`"+i18nField.key+"`"+`)}})
			return
		}`+updateAddStr+`
	}`)

		logic.delete = append(logic.delete, `if count, _ := daoModelThis.CloneNew().Filter(`+daoPath+`.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.Pid)+`, daoModelThis.IdArr).Count(); count > 0 {
		err = utils.NewErrorCode(ctx, 29999994, `+"``"+`)
		return
	}`)
	}
	for _, v := range tpl.Handle.ExtendTableOneList {
		logic.Merge(getLogicExtendMiddleOne(v))
	}
	for _, v := range tpl.Handle.MiddleTableOneList {
		logic.Merge(getLogicExtendMiddleOne(v))
	}
	for _, v := range tpl.Handle.ExtendTableManyList {
		logic.Merge(getLogicExtendMiddleMany(v))
	}
	for _, v := range tpl.Handle.MiddleTableManyList {
		logic.Merge(getLogicExtendMiddleMany(v))
	}
	for _, v := range tpl.Handle.OtherRelTableList {
		logic.Merge(getLogicOtherRel(v))
	}

	if len(logic.verifyData) > 0 {
		logic.verifyDataFunc = `

// 验证数据（create和update共用）
func (logicThis *s` + tpl.LogicStructName + `) verifyData(ctx context.Context, data map[string]any) (err error) {
	` + gstr.Join(logic.verifyData, `

	`) + `
	return
}`
		logic.verifyDataFuncRun = `
	if err = logicThis.verifyData(ctx, data); err != nil {
		return
	}`
	}
	logic.Unique()

	tplLogic := `package ` + tpl.GetModuleName(`logic`) + `

import (
	dao` + tpl.ModuleDirCaseCamel + ` "api/internal/dao/` + tpl.ModuleDirCaseKebab + `"` + gstr.Join(append([]string{``}, logic.importDao...), `
	`) + `
	"api/internal/service"
	"api/internal/utils"
	"context"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

type s` + tpl.LogicStructName + ` struct{}

func New` + tpl.LogicStructName + `() *s` + tpl.LogicStructName + ` {
	return &s` + tpl.LogicStructName + `{}
}

func init() {
	service.Register` + tpl.LogicStructName + `(New` + tpl.LogicStructName + `())
}` + logic.verifyDataFunc + `

// 新增
func (logicThis *s` + tpl.LogicStructName + `) Create(ctx context.Context, data map[string]any) (id int64, err error) {` + logic.verifyDataFuncRun + `
	daoModelThis := ` + daoPath + `.CtxDaoModel(ctx)` + gstr.Join(append([]string{``}, logic.create...), `

	`) + `

	id, err = daoModelThis.HookInsert(data).InsertAndGetId()
	return
}

// 修改
func (logicThis *s` + tpl.LogicStructName + `) Update(ctx context.Context, filter map[string]any, data map[string]any) (row int64, err error) {` + logic.verifyDataFuncRun + `
	daoModelThis := ` + daoPath + `.CtxDaoModel(ctx)

	daoModelThis.SetIdArr(filter)
	if len(daoModelThis.IdArr) == 0 {
		err = utils.NewErrorCode(ctx, 29999998, ` + "``" + `)
		return
	}` + gstr.Join(append([]string{``}, logic.update...), `

	`) + `

	row, err = daoModelThis.HookUpdate(data).UpdateAndGetAffected()
	return
}

// 删除
func (logicThis *s` + tpl.LogicStructName + `) Delete(ctx context.Context, filter map[string]any) (row int64, err error) {
	daoModelThis := ` + daoPath + `.CtxDaoModel(ctx)

	daoModelThis.SetIdArr(filter)
	if len(daoModelThis.IdArr) == 0 {
		err = utils.NewErrorCode(ctx, 29999998, ` + "``" + `)
		return
	}` + gstr.Join(append([]string{``}, logic.delete...), `

	`) + `

	row, err = daoModelThis.HookDelete().DeleteAndGetAffected()
	return
}
`

	gfile.PutContents(saveFile, tplLogic)
	utils.GoFileFmt(saveFile)
	internal.Command(`service生成`, true, ``, `gf`, `gen`, `service`)
	i18n = logic.i18n
	return
}

func getLogicField(tpl myGenTpl, v myGenField) (logicField myGenLogicField) {
	daoPath := `dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel

	switch v.FieldTypeName {
	case internal.TypeNameIdSuffix: // id后缀；	类型：int等类型；
		relIdObj := tpl.Handle.RelIdMap[v.FieldRaw]
		if relIdObj.tpl.Table != `` {
			logicField.importDao = append(logicField.importDao, `dao`+relIdObj.tpl.ModuleDirCaseCamel+` "api/internal/dao/`+relIdObj.tpl.ModuleDirCaseKebab+`"`)
			logicField.i18nField = myGenI18nField{
				key: `name.` + gstr.CaseCamelLower(relIdObj.tpl.ModuleDirCaseCamel) + `.` + gstr.CaseCamelLower(relIdObj.tpl.TableCaseCamel),
				val: relIdObj.FieldName,
			}
			daoPathRel := `dao` + relIdObj.tpl.ModuleDirCaseCamel + `.` + relIdObj.tpl.TableCaseCamel
			logicField.verifyDataStr = `if _, ok := data[` + daoPath + `.Columns().` + v.FieldCaseCamel + `]; ok && gconv.Uint(data[` + daoPath + `.Columns().` + v.FieldCaseCamel + `]) > 0 {
		if count, _ := ` + daoPathRel + `.CtxDaoModel(ctx).Filter(` + daoPathRel + `.Columns().` + relIdObj.tpl.Handle.Id.List[0].FieldCaseCamel + `, data[` + daoPath + `.Columns().` + v.FieldCaseCamel + `]).Count(); count == 0 {
			err = utils.NewErrorCode(ctx, 29999997, ` + "``" + `, g.Map{` + "`i18nValues`" + `: []any{g.I18n().T(ctx, ` + "`" + logicField.i18nField.key + "`" + `)}})
			return
		}
	}`
		}
	}
	return
}

func getLogicExtendMiddleOne(tplEM handleExtendMiddle) (logic myGenLogic) {
	tpl := tplEM.tpl
	logic.importDao = append(logic.importDao, `dao`+tpl.ModuleDirCaseCamel+` "api/internal/dao/`+tpl.ModuleDirCaseKebab+`"`)

	for _, v := range tplEM.FieldList {
		logic.Add(getLogicField(tpl, v))
	}
	return
}

func getLogicExtendMiddleMany(tplEM handleExtendMiddle) (logic myGenLogic) {
	tpl := tplEM.tpl
	if len(tplEM.FieldList) == 1 {
		v := tplEM.FieldList[0]
		switch v.FieldTypeName {
		case internal.TypeNameIdSuffix: // id后缀；	类型：int等类型；
			relIdObj := tpl.Handle.RelIdMap[v.FieldRaw]
			if relIdObj.tpl.Table != `` {
				logic.importDao = append(logic.importDao, `dao`+relIdObj.tpl.ModuleDirCaseCamel+` "api/internal/dao/`+relIdObj.tpl.ModuleDirCaseKebab+`"`)
				i18nField := myGenI18nField{
					key: `name.` + gstr.CaseCamelLower(relIdObj.tpl.ModuleDirCaseCamel) + `.` + gstr.CaseCamelLower(relIdObj.tpl.TableCaseCamel),
					val: relIdObj.FieldName,
				}
				logic.i18n.Add(i18nField)
				daoPathRel := `dao` + relIdObj.tpl.ModuleDirCaseCamel + `.` + relIdObj.tpl.TableCaseCamel
				logic.verifyData = append(logic.verifyData, `if _, ok := data[`+"`"+tplEM.FieldVar+"`"+`]; ok && len(gconv.SliceUint(data[`+"`"+tplEM.FieldVar+"`"+`])) > 0 {
		`+gstr.CaseCamelLower(tplEM.FieldVar)+` := gconv.SliceUint(data[`+"`"+tplEM.FieldVar+"`"+`])
		if count, _ := `+daoPathRel+`.CtxDaoModel(ctx).Filter(`+daoPathRel+`.Columns().`+relIdObj.tpl.Handle.Id.List[0].FieldCaseCamel+`, `+gstr.CaseCamelLower(tplEM.FieldVar)+`).Count(); count != len(`+gstr.CaseCamelLower(tplEM.FieldVar)+`) {
			err = utils.NewErrorCode(ctx, 29999997, `+"``"+`, g.Map{`+"`i18nValues`"+`: []any{g.I18n().T(ctx, `+"`"+i18nField.key+"`"+`)}})
			return
		}
	}`)
			}
		}
	} else {
		verifyDataArr := struct {
			part1 []string
			part2 []string
			part3 []string
		}{}
		for _, v := range tplEM.FieldList {
			switch v.FieldTypeName {
			case internal.TypeNameIdSuffix: // id后缀；	类型：int等类型；
				relIdObj := tpl.Handle.RelIdMap[v.FieldRaw]
				if relIdObj.tpl.Table != `` {
					logic.importDao = append(logic.importDao, `dao`+relIdObj.tpl.ModuleDirCaseCamel+` "api/internal/dao/`+relIdObj.tpl.ModuleDirCaseKebab+`"`)
					i18nField := myGenI18nField{
						key: `name.` + gstr.CaseCamelLower(relIdObj.tpl.ModuleDirCaseCamel) + `.` + gstr.CaseCamelLower(relIdObj.tpl.TableCaseCamel),
						val: relIdObj.FieldName,
					}
					logic.i18n.Add(i18nField)
					daoPathRel := `dao` + relIdObj.tpl.ModuleDirCaseCamel + `.` + relIdObj.tpl.TableCaseCamel
					fieldCaseCamelLower := gstr.CaseCamelLower(v.FieldCaseCamel)
					verifyDataArr.part1 = append(verifyDataArr.part1, fieldCaseCamelLower+`Arr := []uint{}`)
					verifyDataArr.part2 = append(verifyDataArr.part2, `if `+fieldCaseCamelLower+` := gconv.Uint(item[`+daoPathRel+`.Columns().`+v.FieldCaseCamel+`]); `+fieldCaseCamelLower+` > 0 {
						`+fieldCaseCamelLower+`Arr = append(`+fieldCaseCamelLower+`Arr, `+fieldCaseCamelLower+`)
			}`)
					verifyDataArr.part3 = append(verifyDataArr.part3, `if len(`+fieldCaseCamelLower+`Arr) > 0 {
			if count, _ := `+daoPathRel+`.CtxDaoModel(ctx).Filter(`+daoPathRel+`.Columns().`+v.FieldCaseCamel+`, `+fieldCaseCamelLower+`Arr).Count(); count != len(`+fieldCaseCamelLower+`Arr) {
				err = utils.NewErrorCode(ctx, 29999997, `+"``"+`, g.Map{`+"`i18nValues`"+`: []any{g.I18n().T(ctx, `+"`"+i18nField.key+"`"+`)}})
				return
			}
		}`)
				}
			}
		}

		if len(verifyDataArr.part1) > 0 {
			logic.verifyData = append(logic.verifyData, `if _, ok := data[`+"`"+tplEM.FieldVar+"`"+`]; ok && len(gconv.SliceMap(data[`+"`"+tplEM.FieldVar+"`"+`])) > 0 {`+gstr.Join(append([]string{``}, verifyDataArr.part1...), `
		`)+`
		for _, item := range gconv.SliceMap(data[`+"`"+tplEM.FieldVar+"`"+`]) {`+gstr.Join(append([]string{``}, verifyDataArr.part2...), `
			`)+`
		}`+gstr.Join(append([]string{``}, verifyDataArr.part3...), `
		`)+`
	}`)
		}
	}
	return
}

func getLogicOtherRel(tplOR handleOtherRel) (logic myGenLogic) {
	logic.importDao = append(logic.importDao, `dao`+tplOR.tpl.ModuleDirCaseCamel+` "api/internal/dao/`+tplOR.tpl.ModuleDirCaseKebab+`"`)

	getFieldName := func(fieldName string) string {
		if gstr.ToUpper(gstr.SubStr(fieldName, -2)) == `ID` {
			fieldName = gstr.SubStr(fieldName, 0, -2)
		}
		return fieldName
	}
	i18nFieldTop := myGenI18nField{
		key: `name.` + gstr.CaseCamelLower(tplOR.tplOfTop.ModuleDirCaseCamel) + `.` + gstr.CaseCamelLower(tplOR.tplOfTop.TableCaseCamel),
		val: getFieldName(tplOR.tplOfTop.Handle.Id.List[0].FieldName),
	}
	logic.i18n.Add(i18nFieldTop)
	i18nField := myGenI18nField{
		key: `name.` + gstr.CaseCamelLower(tplOR.tpl.ModuleDirCaseCamel) + `.` + gstr.CaseCamelLower(tplOR.tpl.TableCaseCamel),
		val: getFieldName(tplOR.tpl.Handle.Id.List[0].FieldName),
	}
	// 有可能是中间表（一对多），是联合主键。所以需要排除tplOR.RelId
	if len(tplOR.tpl.Handle.Id.List) > 1 && tplOR.tpl.Handle.Id.List[0].FieldRaw == tplOR.RelId {
		i18nField.val = getFieldName(tplOR.tpl.Handle.Id.List[1].FieldName)
	}
	logic.i18n.Add(i18nField)

	daoPath := `dao` + tplOR.tpl.ModuleDirCaseCamel + `.` + tplOR.tpl.TableCaseCamel
	logic.delete = append(logic.delete, `if count, _ := `+daoPath+`.CtxDaoModel(ctx).Filter(`+daoPath+`.Columns().`+gstr.CaseCamel(tplOR.RelId)+`, daoModelThis.IdArr).Count(); count > 0 {
		err = utils.NewErrorCode(ctx, 30009999, `+"``"+`, g.Map{`+"`i18nValues`"+`: []any{g.I18n().T(ctx, `+"`"+i18nFieldTop.key+"`"+`), count, g.I18n().T(ctx, `+"`"+i18nField.key+"`"+`)}})
		return
	}`)
	return
}
