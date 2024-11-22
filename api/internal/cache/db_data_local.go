package cache

import (
	"api/internal/consts"
	"api/internal/dao"
	"context"
	"fmt"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/genv"
	"github.com/gogf/gf/v2/text/gstr"
)

var DbDataLocal = dbDataLocal{}

type dbDataLocal struct{}

func (cacheThis *dbDataLocal) key(daoModel *dao.DaoModel, key string) string {
	return fmt.Sprintf(consts.DB_DATA_FORMAT, gstr.ToUpper(daoModel.DbGroup), gstr.ToUpper(daoModel.DbTable), gstr.ToUpper(key))
}

func (cacheThis *dbDataLocal) Set(ctx context.Context, daoModel *dao.DaoModel, key string, value string) (err error) {
	err = genv.Set(cacheThis.key(daoModel, key), value)
	return
}

func (cacheThis *dbDataLocal) Get(ctx context.Context, daoModel *dao.DaoModel, key string) (value *gvar.Var, err error) {
	value, err = g.Cfg().GetWithEnv(ctx, cacheThis.key(daoModel, key))
	return
}

func (cacheThis *dbDataLocal) GetInfo(ctx context.Context, daoModel *dao.DaoModel, key string) (info gdb.Record, err error) {
	value, err := cacheThis.Get(ctx, daoModel, key)
	if err != nil {
		return
	}
	value.Scan(&info)
	return
}

func (cacheThis *dbDataLocal) GetList(ctx context.Context, daoModel *dao.DaoModel, key string) (list gdb.Result, err error) {
	value, err := cacheThis.Get(ctx, daoModel, key)
	if err != nil {
		return
	}
	value.Scan(&list)
	return
}
