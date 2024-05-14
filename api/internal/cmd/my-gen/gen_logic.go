package my_gen

import (
	"api/internal/cmd/my-gen/internal"
	"api/internal/utils"

	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
)

// logic生成
func genLogic(option myGenOption, tpl myGenTpl) {
	saveFile := gfile.SelfDir() + `/internal/logic/` + gstr.Replace(tpl.ModuleDirCaseKebab, `/`, `-`) + `/` + tpl.TableCaseSnake + `.go`
	if !option.IsResetLogic && gfile.IsFile(saveFile) {
		return
	}

	tplLogic := `package logic

import (
	dao` + tpl.ModuleDirCaseCamel + ` "api/internal/dao/` + tpl.ModuleDirCaseKebab + `"
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
}

// 新增
func (logicThis *s` + tpl.LogicStructName + `) Create(ctx context.Context, data map[string]interface{}) (id int64, err error) {
	daoThis := dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `
	daoModelThis := daoThis.CtxDaoModel(ctx)
`
	if tpl.Handle.Pid.Pid != `` {
		tplLogic += `
	if _, ok := data[daoThis.Columns().` + gstr.CaseCamel(tpl.Handle.Pid.Pid) + `]; ok {
		pid := gconv.Uint(data[daoThis.Columns().` + gstr.CaseCamel(tpl.Handle.Pid.Pid) + `])
		if pid > 0 {
			pInfo, _ := daoModelThis.CloneNew().Filter(daoThis.Columns().` + tpl.Handle.Id.List[0].FieldCaseCamel + `, pid).One()
			if pInfo.IsEmpty() {
				err = utils.NewErrorCode(ctx, 29999997, ` + "``" + `)
				return
			}
		}
	}
`
	}
	tplLogic += `
	id, err = daoModelThis.HookInsert(data).InsertAndGetId()
	return
}

// 修改
func (logicThis *s` + tpl.LogicStructName + `) Update(ctx context.Context, filter map[string]interface{}, data map[string]interface{}) (row int64, err error) {
	daoThis := dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `
	daoModelThis := daoThis.CtxDaoModel(ctx)

	daoModelThis.Filters(filter).SetIdArr()
	if len(daoModelThis.IdArr) == 0 {
		err = utils.NewErrorCode(ctx, 29999998, ` + "``" + `)
		return
	}
`
	if tpl.Handle.Pid.Pid != `` {
		tplLogic += `
	if _, ok := data[daoThis.Columns().` + gstr.CaseCamel(tpl.Handle.Pid.Pid) + `]; ok {
		pid := gconv.Uint(data[daoThis.Columns().` + gstr.CaseCamel(tpl.Handle.Pid.Pid) + `])
		if pid > 0 {
			if garray.NewArrayFrom(gconv.SliceAny(gconv.SliceUint(daoModelThis.IdArr))).Contains(pid) {
				err = utils.NewErrorCode(ctx, 29999996, ` + "``" + `)
				return
			}
			pInfo, _ := daoModelThis.CloneNew().Filter(daoThis.Columns().` + tpl.Handle.Id.List[0].FieldCaseCamel + `, pid).One()
			if pInfo.IsEmpty() {
				err = utils.NewErrorCode(ctx, 29999997, ` + "``" + `)
				return
			}`
		if tpl.Handle.Pid.IsCoexist {
			tplLogic += `
			for _, id := range daoModelThis.IdArr {
				if garray.NewStrArrayFrom(gstr.Split(pInfo[daoThis.Columns().` + gstr.CaseCamel(tpl.Handle.Pid.IdPath) + `].String(), ` + "`-`" + `)).Contains(gconv.String(id)) {
					err = utils.NewErrorCode(ctx, 29999995, ` + "``" + `)
					return
				}
			}`
		}
		tplLogic += `
		}
	}
`
	}
	tplLogic += `
	row, err = daoModelThis.HookUpdate(data).UpdateAndGetAffected()
	return
}

// 删除
func (logicThis *s` + tpl.LogicStructName + `) Delete(ctx context.Context, filter map[string]interface{}) (row int64, err error) {
	daoThis := dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `
	daoModelThis := daoThis.CtxDaoModel(ctx)

	daoModelThis.Filters(filter).SetIdArr()
	if len(daoModelThis.IdArr) == 0 {
		err = utils.NewErrorCode(ctx, 29999998, ` + "``" + `)
		return
	}
`
	if tpl.Handle.Pid.Pid != `` {
		tplLogic += `
	count, _ := daoModelThis.CloneNew().Filter(daoThis.Columns().` + gstr.CaseCamel(tpl.Handle.Pid.Pid) + `, daoModelThis.IdArr).Count()
	if count > 0 {
		err = utils.NewErrorCode(ctx, 29999994, ` + "``" + `)
		return
	}
`
	}
	tplLogic += `
	row, err = daoModelThis.HookDelete().DeleteAndGetAffected()
	return
}
`

	gfile.PutContents(saveFile, tplLogic)
	utils.GoFileFmt(saveFile)
	internal.Command(`service生成`, true, ``, `gf`, `gen`, `service`)
}
