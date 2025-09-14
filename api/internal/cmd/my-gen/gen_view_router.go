package my_gen

import (
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/text/gstr"
)

// 前端路由生成
func genViewRouter(option myGenOption, tpl *myGenTpl) {
	saveFile := gfile.SelfDir() + `/../view/` + option.SceneId + `/src/router/index.ts`
	tplViewRouter := gfile.GetContents(saveFile)

	path := `/` + tpl.ModuleDirCaseKebab + `/` + tpl.TableCaseKebab
	replaceStr := `{
                path: '` + path + `',
                component: async () => {
                    const component = await import('@/views` + path + `/Index.vue')
                    component.default.name = '` + path + `'
                    return component
                },
                meta: { isAuth: true, keepAlive: true, componentName: '` + path + `' }
            },`

	if gstr.Pos(tplViewRouter, `'`+path+`'`) == -1 { //路由不存在时新增
		tplViewRouter = gstr.Replace(tplViewRouter, `/*--------前端路由自动代码生成锚点（不允许修改和删除，否则将不能自动生成路由）--------*/`, replaceStr+`
            /*--------前端路由自动代码生成锚点（不允许修改和删除，否则将不能自动生成路由）--------*/`, 1)
	} else { //路由已存在则替换
		tplViewRouter, _ = gregex.ReplaceString(`\{
                path: '`+path+`',[\s\S]*'`+path+`' \}
            \},`, replaceStr, tplViewRouter)
	}
	gfile.PutContents(saveFile, tplViewRouter)
}
