package my

import (
	apiMy "api/api/platform/my"
	daoAdmin "api/internal/dao/admin"
	daoAuth "api/internal/dao/auth"
	"api/internal/utils"
	"api/internal/utils/jbctx"
	"context"

	"github.com/gogf/gf/v2/util/gconv"
)

type Menu struct{}

func NewMenu() *Menu {
	return &Menu{}
}

// 列表（树状）
func (controllerThis *Menu) Tree(ctx context.Context, req *apiMy.MenuTreeReq) (res *apiMy.MenuTreeRes, err error) {
	loginInfo := jbctx.GetLoginInfo(ctx)
	sceneInfo := jbctx.GetSceneInfo(ctx)

	/* // 表数据很小，无需这样做，且会导致数据修改无法立即生效。确实需要减轻数据库压力时可以使用
	var list gdb.Result
	if loginInfo[daoAdmin.Admin.Columns().IsSuper].Bool() {
		list, err = daoAuth.Menu.CacheGetListOfNoStop(ctx, sceneInfo[daoAuth.Scene.Columns().SceneId].String())
	} else {
		list, err = daoAuth.Menu.CacheGetListOfSelf(ctx, sceneInfo[daoAuth.Scene.Columns().SceneId].String(), loginInfo[`login_id`])
	} */
	field := []string{`id`, `label`, `tree`, `show_menu`}
	filter := map[string]any{
		`self_menu`: map[string]any{
			`scene_id`: sceneInfo[daoAuth.Scene.Columns().SceneId],
			`login_id`: loginInfo[`login_id`],
			`is_super`: loginInfo[daoAdmin.Admin.Columns().IsSuper].Uint8(),
		},
	}
	list, err := daoAuth.Menu.CtxDaoModel(ctx).Filters(filter).Fields(field...).ListPri()
	if err != nil {
		return
	}
	tree := utils.Tree(list.List(), 0, daoAuth.Menu.Columns().MenuId, daoAuth.Menu.Columns().Pid)

	res = &apiMy.MenuTreeRes{Tree: []apiMy.MenuInfo{}}
	gconv.Structs(tree, &res.Tree)
	return
}
