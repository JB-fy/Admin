package my_gen

import (
	"api/internal/utils"
	"strings"

	"github.com/gogf/gf/v2/os/gfile"
)

// 后端路由生成
func genRouter(option myGenOption, tpl *myGenTpl) {
	saveFile := gfile.SelfDir() + `/internal/router/` + option.SceneId + `.go`
	tplRouter := gfile.GetContents(saveFile)

	moduleName := tpl.GetModuleName(`controller`)
	importControllerStr := `"api/internal/controller/` + option.SceneId + `/` + tpl.ModuleDirCaseKebab + `"`
	if strings.Contains(tplRouter, importControllerStr) {
		if strings.Contains(tplRouter, `group.Bind(`+moduleName+`.New`+tpl.TableCaseCamel+`())`) {
			return
		}
		// 路由生成（controller已导入，但路由不存在）
		replacePoint := `group.Group(` + "`" + `/` + tpl.ModuleDirCaseKebab + "`" + `, func(group *ghttp.RouterGroup) {`
		replaceIndex := strings.LastIndex(tplRouter, replacePoint)
		tplRouter = tplRouter[:replaceIndex] + replacePoint + `
				group.Bind(` + moduleName + `.New` + tpl.TableCaseCamel + `())` + tplRouter[replaceIndex+len(replacePoint):]
		utils.FilePutFormat(saveFile, []byte(tplRouter)...)
	} else { // 路由生成（controller未导入）
		tplRouter = strings.Replace(tplRouter, `"api/internal/middleware"`, importControllerStr+`
	"api/internal/middleware"`, 1)

		tplRouter = strings.Replace(tplRouter, `/*--------后端路由自动代码生成锚点（不允许修改和删除，否则将不能自动生成路由）--------*/`, `group.Group(`+"`"+`/`+tpl.ModuleDirCaseKebab+"`"+`, func(group *ghttp.RouterGroup) {
				group.Bind(`+moduleName+`.New`+tpl.TableCaseCamel+`())
			})

			/*--------后端路由自动代码生成锚点（不允许修改和删除，否则将不能自动生成路由）--------*/`, 1)
		utils.FilePutFormat(saveFile, []byte(tplRouter)...)
	}
}
