package logic

import (
	daoAuth "api/internal/model/dao/auth"
	"api/internal/service"
	"context"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
)

type sMenu struct{}

func NewMenu() *sMenu {
	return &sMenu{}
}

func init() {
	service.RegisterMenu(NewMenu())
}

// 总数
func (logicMenu *sMenu) Count(ctx context.Context, filter map[string]interface{}) (count int, err error) {
	daoMenu := daoAuth.Menu
	joinTableArr := []string{}
	model := daoMenu.Ctx(ctx)
	if len(filter) > 0 {
		model = model.Handler(daoMenu.ParseFilter(filter, &joinTableArr))
	}
	if len(joinTableArr) > 0 {
		count, err = model.Handler(daoMenu.ParseGroup([]string{"id"}, &joinTableArr)).Distinct().Count(daoMenu.PrimaryKey())
	} else {
		count, err = model.Count()
	}
	return
}

// 列表
func (logicMenu *sMenu) List(ctx context.Context, filter map[string]interface{}, field []string, order [][2]string, offset int, limit int) (list gdb.Result, err error) {
	daoMenu := daoAuth.Menu
	joinTableArr := []string{}
	model := daoMenu.Ctx(ctx)
	if len(field) > 0 {
		model = model.Handler(daoMenu.ParseField(field, &joinTableArr))
	}
	if len(filter) > 0 {
		model = model.Handler(daoMenu.ParseFilter(filter, &joinTableArr))
	}
	if len(order) > 0 {
		model = model.Handler(daoMenu.ParseOrder(order, &joinTableArr))
	}
	if len(joinTableArr) > 0 {
		model = model.Handler(daoMenu.ParseGroup([]string{"id"}, &joinTableArr))
	}
	if limit > 0 {
		model = model.Offset(offset).Limit(limit)
	}
	list, err = model.All()
	return
}

// 详情
func (logicMenu *sMenu) Info(ctx context.Context, filter map[string]interface{}, field []string, order [][2]string) (info gdb.Record, err error) {
	daoMenu := daoAuth.Menu
	joinTableArr := []string{}
	model := daoMenu.Ctx(ctx)
	if len(field) > 0 {
		model = model.Handler(daoMenu.ParseField(field, &joinTableArr))
	}
	if len(filter) > 0 {
		model = model.Handler(daoMenu.ParseFilter(filter, &joinTableArr))
	}
	if len(order) > 0 {
		model = model.Handler(daoMenu.ParseOrder(order, &joinTableArr))
	}
	if len(joinTableArr) > 0 {
		model = model.Handler(daoMenu.ParseGroup([]string{"id"}, &joinTableArr))
	}
	info, err = model.One()
	return
}

// 创建
func (logicMenu *sMenu) Create(ctx context.Context, data []map[string]interface{}) (id int64, err error) {
	daoMenu := daoAuth.Menu
	model := daoMenu.Ctx(ctx)
	if len(data) > 0 {
		model = model.Handler(daoMenu.ParseInsert(data))
	}
	if len(data) == 1 {
		id, err = model.InsertAndGetId()
		return
	}
	result, err := model.Insert()
	if err != nil {
		return
	}
	id, err = result.RowsAffected()
	return
}

// 更新
func (logicMenu *sMenu) Update(ctx context.Context, data map[string]interface{}, filter map[string]interface{}, order [][2]string, offset int, limit int) (row int64, err error) {
	daoMenu := daoAuth.Menu
	joinTableArr := []string{}
	model := daoMenu.Ctx(ctx)
	if len(data) > 0 {
		model = model.Handler(daoMenu.ParseUpdate(data))
	}
	if len(filter) > 0 {
		model = model.Handler(daoMenu.ParseFilter(filter, &joinTableArr))
	}
	if len(order) > 0 {
		model = model.Handler(daoMenu.ParseOrder(order, &joinTableArr))
	}
	if limit > 0 {
		model = model.Offset(offset).Limit(limit)
	}
	result, err := model.Update()
	if err != nil {
		return
	}
	row, err = result.RowsAffected()
	return
}

// 删除
func (logicMenu *sMenu) Delete(ctx context.Context, filter map[string]interface{}, order [][2]string, offset int, limit int) (row int64, err error) {
	daoMenu := daoAuth.Menu
	joinTableArr := []string{}
	model := daoMenu.Ctx(ctx)
	if len(filter) > 0 {
		model = model.Handler(daoMenu.ParseFilter(filter, &joinTableArr))
	}
	if len(order) > 0 {
		model = model.Handler(daoMenu.ParseOrder(order, &joinTableArr))
	}
	if limit > 0 {
		model = model.Offset(offset).Limit(limit)
	}
	result, err := model.Delete()
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
			/* if len(children) == 0 {
				v["children"] = gvar.New(gdb.Result{})
			} */
			tree = append(tree, v)
		}
	}
	return
}
