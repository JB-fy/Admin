package logic

import (
	daoAuth "api/internal/model/dao/auth"
	"api/internal/service"
	"api/internal/utils"
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type sRole struct{}

func NewRole() *sRole {
	return &sRole{}
}

func init() {
	service.RegisterRole(NewRole())
}

// 总数
func (logicThis *sRole) Count(ctx context.Context, filter map[string]interface{}) (count int, err error) {
	daoThis := daoAuth.Role
	joinTableArr := []string{}
	model := daoThis.ParseDbCtx(ctx)
	if len(filter) > 0 {
		model = model.Handler(daoThis.ParseFilter(filter, &joinTableArr))
	}
	if len(joinTableArr) > 0 {
		count, err = model.Handler(daoThis.ParseGroup([]string{"id"}, &joinTableArr)).Distinct().Count(daoThis.PrimaryKey())
	} else {
		count, err = model.Count()
	}
	return
}

// 列表
func (logicThis *sRole) List(ctx context.Context, filter map[string]interface{}, field []string, order [][2]string, page int, limit int) (list gdb.Result, err error) {
	daoThis := daoAuth.Role
	joinTableArr := []string{}
	model := daoThis.ParseDbCtx(ctx)
	if len(filter) > 0 {
		model = model.Handler(daoThis.ParseFilter(filter, &joinTableArr))
	}
	if len(field) > 0 {
		model = model.Handler(daoThis.ParseField(field, &joinTableArr))
	}
	if len(order) > 0 {
		model = model.Handler(daoThis.ParseOrder(order, &joinTableArr))
	}
	if len(joinTableArr) > 0 {
		model = model.Handler(daoThis.ParseGroup([]string{"id"}, &joinTableArr))
	}
	if limit > 0 {
		model = model.Offset((page - 1) * limit).Limit(limit)
	}
	list, err = model.All()
	return
}

// 详情
func (logicThis *sRole) Info(ctx context.Context, filter map[string]interface{}, field ...[]string) (info gdb.Record, err error) {
	daoThis := daoAuth.Role
	joinTableArr := []string{}
	model := daoThis.ParseDbCtx(ctx)
	model = model.Handler(daoThis.ParseFilter(filter, &joinTableArr))
	if len(field) > 0 && len(field[0]) > 0 {
		model = model.Handler(daoThis.ParseField(field[0], &joinTableArr))
	}
	if len(joinTableArr) > 0 {
		model = model.Handler(daoThis.ParseGroup([]string{"id"}, &joinTableArr))
	}
	info, err = model.One()
	return
}

// 创建
func (logicThis *sRole) Create(ctx context.Context, data map[string]interface{}) (id int64, err error) {
	daoThis := daoAuth.Role
	_, okMenuIdArr := data["menuIdArr"]
	if okMenuIdArr {
		menuIdArr := gconv.SliceInt(data["menuIdArr"])
		filter := g.Map{"menuId": data["menuIdArr"], "sceneId": data["sceneId"]}
		menuIdArrCount, _ := daoAuth.Menu.ParseDbCtx(ctx).Handler(daoAuth.Menu.ParseFilter(filter, &[]string{})).Count()
		if len(menuIdArr) != menuIdArrCount {
			err = utils.NewErrorCode(ctx, 89999998, "")
			return
		}
	}
	_, okActionIdArr := data["actionIdArr"]
	if okActionIdArr {
		actionIdArr := gconv.SliceInt(data["actionIdArr"])
		filter := g.Map{"actionId": data["actionIdArr"], "sceneId": data["sceneId"]}
		actionIdArrCount, _ := daoAuth.ActionRelToScene.ParseDbCtx(ctx).Handler(daoAuth.ActionRelToScene.ParseFilter(filter, &[]string{})).Count()
		if len(actionIdArr) != actionIdArrCount {
			err = utils.NewErrorCode(ctx, 89999998, "")
			return
		}
	}

	id, err = daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseInsert([]map[string]interface{}{data})).InsertAndGetId()
	if err != nil {
		return
	}

	if okMenuIdArr {
		daoThis.SaveRelMenu(ctx, gconv.SliceInt(data["menuIdArr"]), int(id))
	}
	if okActionIdArr {
		daoThis.SaveRelAction(ctx, gconv.SliceInt(data["actionIdArr"]), int(id))
	}
	return
}

// 更新
func (logicThis *sRole) Update(ctx context.Context, data map[string]interface{}, filter map[string]interface{}) (row int64, err error) {
	daoThis := daoAuth.Role
	_, okMenuIdArr := data["menuIdArr"]
	_, okActionIdArr := data["actionIdArr"]
	if okMenuIdArr || okActionIdArr {
		oldInfo, _ := daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseFilter(filter, &[]string{})).One()
		if okMenuIdArr {
			menuIdArr := gconv.SliceInt(data["menuIdArr"])
			filter := g.Map{"menuId": data["menuIdArr"], "sceneId": oldInfo["sceneId"]}
			_, okSceneId := data["sceneId"]
			if okSceneId {
				filter["sceneId"] = data["sceneId"]
			}
			menuIdArrCount, _ := daoAuth.Menu.ParseDbCtx(ctx).Handler(daoAuth.Menu.ParseFilter(filter, &[]string{})).Count()
			if len(menuIdArr) != menuIdArrCount {
				err = utils.NewErrorCode(ctx, 89999998, "")
				return
			}
		}
	}

	// if (isset($data['menuIdArr']) || isset($data['actionIdArr'])) {
	// 	$oldInfo = $this->getDao()->parseFilter($filter)->info();
	// 	if (isset($data['menuIdArr'])) {
	// 		if (count($data['menuIdArr']) != getDao(Menu::class)->parseFilter(['id' => $data['menuIdArr'], 'sceneId' => $data['sceneId'] ?? $oldInfo->sceneId])->getBuilder()->count()) {
	// 			throwFailJson(89999998);
	// 		}
	// 		$this->container->get(AuthRole::class)->saveRelMenu($data['menuIdArr'], $oldInfo->roleId);
	// 		$this->getDao()->parseFilter($filter)->parseUpdate($data)->update();    //有可能只改menuIdArr
	// 	}
	// 	if (isset($data['actionIdArr'])) {
	// 		if (count($data['actionIdArr']) != getDao(ActionRelToScene::class)->parseFilter(['actionId' => $data['actionIdArr'], 'sceneId' => $data['sceneId'] ?? $oldInfo->sceneId])->getBuilder()->count()) {
	// 			throwFailJson(89999998);
	// 		}
	// 		$this->container->get(AuthRole::class)->saveRelAction($data['actionIdArr'], $oldInfo->roleId);
	// 		$this->getDao()->parseFilter($filter)->parseUpdate($data)->update();    //有可能只改actionIdArr
	// 	}
	// } else {
	// 	$result = $this->getDao()->parseFilter($filter)->parseUpdate($data)->update();
	// 	if (empty($result)) {
	// 		throwFailJson();
	// 	}
	// }

	result, err := daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseUpdate(data), daoThis.ParseFilter(filter, &[]string{})).Update()
	if err != nil {
		return
	}
	row, err = result.RowsAffected()
	return
}

// 删除
func (logicThis *sRole) Delete(ctx context.Context, filter map[string]interface{}) (row int64, err error) {
	daoThis := daoAuth.Role
	result, err := daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseFilter(filter, &[]string{})).Delete()
	if err != nil {
		return
	}
	row, err = result.RowsAffected()
	return
}
