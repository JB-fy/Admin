package logic

import (
	daoAuth "api/internal/model/dao/auth"
	daoPlatform "api/internal/model/dao/platform"
	"api/internal/service"
	"api/internal/utils"
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/util/gconv"
)

type sAdmin struct{}

func NewAdmin() *sAdmin {
	return &sAdmin{}
}

func init() {
	service.RegisterAdmin(NewAdmin())
}

// 总数
func (logicThis *sAdmin) Count(ctx context.Context, filter map[string]interface{}) (count int, err error) {
	daoThis := daoPlatform.Admin
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
func (logicThis *sAdmin) List(ctx context.Context, filter map[string]interface{}, field []string, order [][2]string, page int, limit int) (list gdb.Result, err error) {
	daoThis := daoPlatform.Admin
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
func (logicThis *sAdmin) Info(ctx context.Context, filter map[string]interface{}, field ...[]string) (info gdb.Record, err error) {
	daoThis := daoPlatform.Admin
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
func (logicThis *sAdmin) Create(ctx context.Context, data map[string]interface{}) (id int64, err error) {
	daoThis := daoPlatform.Admin
	id, err = daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseInsert([]map[string]interface{}{data})).InsertAndGetId()
	if err != nil {
		return
	}

	_, okRoleIdArr := data["roleIdArr"]
	if okRoleIdArr {
		daoThis.SaveRelRole(ctx, gconv.SliceInt(data["roleIdArr"]), int(id))
	}
	return
}

// 更新
func (logicThis *sAdmin) Update(ctx context.Context, data map[string]interface{}, filter map[string]interface{}) (row int64, err error) {
	daoThis := daoPlatform.Admin

	_, okCheckPassword := data["checkPassword"]
	if okCheckPassword {
		password, _ := daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseFilter(filter, &[]string{})).Value("password")
		if gconv.String(data["checkPassword"]) != password.String() {
			err = utils.NewErrorCode(ctx, 39990003, "")
			return
		}
	}

	_, okRoleIdArr := data["roleIdArr"]
	if okRoleIdArr {
		idArr, _ := daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseFilter(filter, &[]string{})).Array(daoThis.PrimaryKey())
		for _, v := range idArr {
			daoThis.SaveRelRole(ctx, gconv.SliceInt(data["roleIdArr"]), v.Int())
		}
		daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseUpdate(data), daoThis.ParseFilter(filter, &[]string{})).Update() //有可能只改roleIdArr
		return
	}

	result, err := daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseUpdate(data), daoThis.ParseFilter(filter, &[]string{})).Update()
	if err != nil {
		return
	}
	row, err = result.RowsAffected()
	return
}

// 删除
func (logicThis *sAdmin) Delete(ctx context.Context, filter map[string]interface{}) (row int64, err error) {
	daoThis := daoPlatform.Admin
	idArr, _ := daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseFilter(filter, &[]string{})).Array(daoThis.PrimaryKey())
	result, err := daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseFilter(filter, &[]string{})).Delete()
	if err != nil {
		return
	}
	row, err = result.RowsAffected()
	if row > 0 {
		daoAuth.RoleRelOfPlatformAdmin.ParseDbCtx(ctx).Where("adminId", idArr).Delete()
	}
	return
}
