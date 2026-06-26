package my

import (
	apiMy "api/api/org/my"
	"api/internal/consts"
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

	/* // 表数据很小，无需这样做，且会导致数据修改无法立即生效。确实需要减轻数据库压力时可以使用
	list, err := daoAuth.Menu.CacheGetListOfSelf(ctx, jbctx.GetSceneId(ctx).String(), loginInfo[consts.CTX_LOGIN_ID_NAME]) */
	field := []string{`id`, `label`, `tree`, `show_menu`}
	filter := map[string]any{
		`self_menu`: map[string]any{
			`scene_id`: jbctx.GetSceneId(ctx),
			`login_id`: loginInfo[consts.CTX_LOGIN_ID_NAME],
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
