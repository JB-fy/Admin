package logic

import (
	daoAuth "api/internal/model/dao/auth"
	"api/internal/service"
	"api/internal/utils"
	"context"
	"fmt"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

type sMenu struct{}

func NewMenu() *sMenu {
	return &sMenu{}
}

func init() {
	service.RegisterMenu(NewMenu())
}

// 总数
func (logicThis *sMenu) Count(ctx context.Context, filter map[string]interface{}) (count int, err error) {
	daoThis := daoAuth.Menu
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
func (logicThis *sMenu) List(ctx context.Context, filter map[string]interface{}, field []string, order [][2]string, page int, limit int) (list gdb.Result, err error) {
	daoThis := daoAuth.Menu
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
func (logicThis *sMenu) Info(ctx context.Context, filter map[string]interface{}, field ...[]string) (info gdb.Record, err error) {
	daoThis := daoAuth.Menu
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
func (logicThis *sMenu) Create(ctx context.Context, data map[string]interface{}) (id int64, err error) {
	daoThis := daoAuth.Menu
	var pInfo gdb.Record
	fmt.Println(data["pid"])
	fmt.Println(gconv.Int(data["pid"]))
	pid := gconv.Int(data["pid"])
	if pid > 0 {
		joinTableArr := []string{}
		field := []string{"pidPath", "level"}
		filterTmp := g.Map{"menuId": data["pid"], "sceneId": data["sceneId"]}
		pInfo, _ = daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseFilter(filterTmp, &joinTableArr), daoThis.ParseField(field, &joinTableArr)).One()
		if len(pInfo) == 0 {
			err = utils.NewErrorCode(ctx, 29999998, "")
			return
		}
	}

	id, err = daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseInsert([]map[string]interface{}{data})).InsertAndGetId()
	if err != nil {
		return
	}

	update := map[string]interface{}{
		"pidPath": "0-" + gconv.String(id),
		"level":   1,
	}
	if pid > 0 {
		update = map[string]interface{}{
			"pidPath": pInfo["pidPath"].String() + "-" + gconv.String(id),
			"level":   pInfo["level"].Int() + 1,
		}
	}
	daoThis.ParseDbCtx(ctx).Data(update).Where(daoThis.PrimaryKey(), id).Update()
	return
}

// 更新
func (logicThis *sMenu) Update(ctx context.Context, data map[string]interface{}, filter map[string]interface{}) (row int64, err error) {
	daoThis := daoAuth.Menu
	_, okPid := data["pid"]
	var oldInfo gdb.Record
	if okPid {
		pid := gconv.Int(data["pid"])
		oldInfo, _ = daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseFilter(filter, &[]string{})).One()
		if pid == oldInfo["menuId"].Int() { //父级不能是自身
			err = utils.NewErrorCode(ctx, 29999997, "")
			return
		}
		if pid != oldInfo["pid"].Int() {
			if pid > 0 {
				joinTableArr := []string{}
				field := []string{"pidPath", "level"}
				filterTmp := g.Map{"menuId": data["pid"], "sceneId": oldInfo["sceneId"]}
				_, okSceneId := data["sceneId"]
				if okSceneId {
					filterTmp["sceneId"] = data["sceneId"]
				}
				pInfo, _ := daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseFilter(filterTmp, &joinTableArr), daoThis.ParseField(field, &joinTableArr)).One()
				if len(pInfo) == 0 {
					err = utils.NewErrorCode(ctx, 29999998, "")
					return
				}
				if garray.NewStrArrayFrom(gstr.Split(pInfo["pidPath"].String(), "-")).Contains(oldInfo["menuId"].String()) { //父级不能是自身的子孙级
					err = utils.NewErrorCode(ctx, 29999996, "")
					return
				}
				data["pidPath"] = pInfo["pidPath"].String() + "-" + oldInfo["menuId"].String()
				data["level"] = pInfo["level"].Int() + 1
			} else {
				data["pidPath"] = "0-" + oldInfo["menuId"].String()
				data["level"] = 1
			}
		} else {
			delete(data, "pid") //未修改则删除，更新后就不用处理data['pid']
		}
	}

	result, err := daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseUpdate(data), daoThis.ParseFilter(filter, &[]string{})).Update()
	if err != nil {
		return
	}
	row, err = result.RowsAffected()
	if row == 0 || err != nil {
		return
	}

	//修改pid时，更新所有子孙级的pidPath和level
	_, okPid1 := data["pid"]
	if okPid1 {
		update := map[string]interface{}{
			"pidPathOfChild": map[string]interface{}{
				"newVal": data["pidPath"],
				"oldVal": oldInfo["pidPath"],
			},
			"levelOfChild": map[string]interface{}{
				"newVal": data["level"],
				"oldVal": oldInfo["level"],
			},
		}
		filter := map[string]interface{}{"pidPath Like ?": oldInfo["pidPath"].String() + "%"}
		daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseUpdate(update), daoThis.ParseFilter(filter, &[]string{})).Update()
	}
	return
}

// 删除
func (logicThis *sMenu) Delete(ctx context.Context, filter map[string]interface{}) (row int64, err error) {
	daoThis := daoAuth.Menu
	idArr, _ := daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseFilter(filter, &[]string{})).Array(daoThis.PrimaryKey())
	count, _ := daoThis.ParseDbCtx(ctx).Where("pid", idArr).Count()
	if count > 0 {
		err = utils.NewErrorCode(ctx, 29999995, "")
		return
	}
	result, err := daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseFilter(filter, &[]string{})).Delete()
	if err != nil {
		return
	}
	row, err = result.RowsAffected()
	return
}

// 菜单树
func (logicMenu *sMenu) Tree(ctx context.Context, list gdb.Result, menuId int) (tree gdb.Result, err error) {
	for _, v := range list {
		//list = append(list[:k], list[(k+1):]...) //删除元素，减少后面递归循环次数（有bug，待处理）
		if v["pid"].Int() == menuId {
			children, _ := logicMenu.Tree(ctx, list, v["menuId"].Int())
			v["children"] = gvar.New(children)
			tree = append(tree, v)
		}
	}
	return
}
