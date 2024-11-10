package my_gen

import (
	"api/internal/utils"

	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
)

// 后端路由生成
func genRouter(option myGenOption, tpl myGenTpl) {
	saveFile := gfile.SelfDir() + `/internal/router/` + option.SceneId + `.go`
	tplRouter := gfile.GetContents(saveFile)

	moduleName := tpl.GetModuleName(`controller`)
	importControllerStr := `"api/internal/controller/` + option.SceneId + `/` + tpl.ModuleDirCaseKebab + `"`
	if gstr.Pos(tplRouter, importControllerStr) == -1 {
		tplRouter = gstr.Replace(tplRouter, `"api/internal/middleware"`, importControllerStr+`
	"api/internal/middleware"`, 1)
		// 路由生成（controller未导入）
		tplRouter = gstr.Replace(tplRouter, `/*--------后端路由自动代码生成锚点（不允许修改和删除，否则将不能自动生成路由）--------*/`, `group.Group(`+"`"+`/`+tpl.ModuleDirCaseKebab+"`"+`, func(group *ghttp.RouterGroup) {
				group.Bind(`+moduleName+`.New`+tpl.TableCaseCamel+`())
			})

			/*--------后端路由自动代码生成锚点（不允许修改和删除，否则将不能自动生成路由）--------*/`, 1)
		gfile.PutContents(saveFile, tplRouter)
	} else {
		// 路由生成（controller已导入，但路由不存在）
		if gstr.Pos(tplRouter, `group.Bind(`+moduleName+`.New`+tpl.TableCaseCamel+`())`) == -1 {
			tplRouter = gstr.Replace(tplRouter, `group.Group(`+"`"+`/`+tpl.ModuleDirCaseKebab+"`"+`, func(group *ghttp.RouterGroup) {`, `group.Group(`+"`"+`/`+tpl.ModuleDirCaseKebab+"`"+`, func(group *ghttp.RouterGroup) {
				group.Bind(`+moduleName+`.New`+tpl.TableCaseCamel+`())`, 1)
			gfile.PutContents(saveFile, tplRouter)
		}
	}

	utils.GoFileFmt(saveFile)
}
