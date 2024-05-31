package my_gen

import (
	"api/internal/cmd/my-gen/internal"
	"api/internal/utils"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
)

type myGenLogic struct {
	importDao []string

	verifyData        []string
	verifyDataFunc    string
	verifyDataFuncRun string
	create            []string
	update            []string
	delete            []string
}

type myGenLogicField struct {
	importDao     []string
	verifyDataStr string
}

func (logicThis *myGenLogic) Add(logicField myGenLogicField) {
	logicThis.importDao = append(logicThis.importDao, logicField.importDao...)
	if logicField.verifyDataStr != `` {
		logicThis.verifyData = append(logicThis.verifyData, logicField.verifyDataStr)
	}
}

func (logicThis *myGenLogic) Unique() {
	logicThis.importDao = garray.NewStrArrayFrom(logicThis.importDao).Unique().Slice()
}

// logic生成
func genLogic(option myGenOption, tpl myGenTpl) {
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
		logic.create = append(logic.create, `if _, ok := data[`+daoPath+`.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.Pid)+`]; ok && gconv.Uint(data[`+daoPath+`.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.Pid)+`]) > 0 {
		pInfo, _ := daoModelThis.CloneNew().Filter(`+daoPath+`.Columns().`+tpl.Handle.Id.List[0].FieldCaseCamel+`, data[`+daoPath+`.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.Pid)+`]).One()
		if pInfo.IsEmpty() {
			err = utils.NewErrorCode(ctx, 29999997, `+"``"+`, g.Map{`+"`errValues`"+`: []any{g.I18n().T(ctx, `+"`name.pid`"+`)}})
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
			err = utils.NewErrorCode(ctx, 29999997, `+"``"+`, g.Map{`+"`errValues`"+`: []any{g.I18n().T(ctx, `+"`name.pid`"+`)}})
			return
		}`+updateAddStr+`
	}`)

		logic.delete = append(logic.delete, `if count, _ := daoModelThis.CloneNew().Filter(`+daoPath+`.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.Pid)+`, daoModelThis.IdArr).Count(); count > 0 {
		err = utils.NewErrorCode(ctx, 29999994, `+"``"+`)
		return
	}`)
	}
	if len(logic.verifyData) > 0 {
		logic.verifyDataFunc = `

// 验证数据（create和update共用）
func (logicThis *s` + tpl.LogicStructName + `) verifyData(ctx context.Context, data map[string]any) (err error) {` + gstr.Join(append([]string{``}, logic.verifyData...), `
	`) + `
	return
}`
		logic.verifyDataFuncRun = `
	if err = logicThis.verifyData(ctx, data); err != nil {
		return
	}`
	}
	logic.Unique()

	tplLogic := `package logic

import (
	dao` + tpl.ModuleDirCaseCamel + ` "api/internal/dao/` + tpl.ModuleDirCaseKebab + `"` + gstr.Join(append([]string{``}, logic.importDao...), `
	`) + `
	"api/internal/service"
	"api/internal/utils"
	"context"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/database/gdb"
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

	daoModelThis.Filters(filter).SetIdArr()
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

	daoModelThis.Filters(filter).SetIdArr()
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
}

func getLogicField(tpl myGenTpl, v myGenField) (logicField myGenLogicField) {
	daoPath := `dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel

	switch v.FieldTypeName {
	case internal.TypeNameIdSuffix: // id后缀；	类型：int等类型；
		relIdObj := tpl.Handle.RelIdMap[v.FieldRaw]
		if relIdObj.tpl.Table != `` {
			logicField.importDao = append(logicField.importDao, `dao`+relIdObj.tpl.ModuleDirCaseCamel+` "api/internal/dao/`+relIdObj.tpl.ModuleDirCaseKebab+`"`)
			daoPathRel := `dao` + relIdObj.tpl.ModuleDirCaseCamel + `.` + relIdObj.tpl.TableCaseCamel
			logicField.verifyDataStr = `if _, ok := data[` + daoPath + `.Columns().` + v.FieldCaseCamel + `]; ok && gconv.Uint(data[` + daoPath + `.Columns().` + v.FieldCaseCamel + `]) > 0 {
		if count, _ := ` + daoPathRel + `.CtxDaoModel(ctx).Filter(` + daoPathRel + `.Columns().` + relIdObj.tpl.Handle.Id.List[0].FieldCaseCamel + `, data[` + daoPath + `.Columns().` + v.FieldCaseCamel + `]).Count(); count == 0 {
			err = utils.NewErrorCode(ctx, 29999998, ` + "``" + `)
			return
		}
	}`
		}
	}
	return
}
