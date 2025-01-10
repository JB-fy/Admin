package cache

import (
	"api/internal/cache/internal"
	"api/internal/consts"
	"api/internal/dao"
	"context"
	"fmt"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

var DbData = dbData{redis: g.Redis()}

type dbData struct{ redis *gredis.Redis }

func (cacheThis *dbData) cache() *gredis.Redis {
	return cacheThis.redis
}

func (cacheThis *dbData) key(daoModel *dao.DaoModel, id any) string {
	return fmt.Sprintf(consts.CacheDbDataFormat, daoModel.DbGroup, daoModel.DbTable, id)
}

func (cacheThis *dbData) GetOrSet(ctx context.Context, dao dao.DaoInterface, id any, ttl int64, field ...string) (value *gvar.Var, noExistOfDb bool, err error) {
	daoModel := dao.CtxDaoModel(ctx)
	redis := cacheThis.cache()
	key := cacheThis.key(daoModel, id)
	valueFunc := func() (value *gvar.Var, noSetCache bool, err error) {
		info, err := daoModel.FilterPri(id).Fields(field...).One()
		if err != nil {
			return
		}
		if info.IsEmpty() {
			noSetCache = true
			return
		}
		if len(field) == 1 {
			value = info[field[0]]
		} else {
			value = gvar.New(info.Json())
		}
		return
	}
	return internal.GetOrSet.GetOrSet(ctx, redis, key, valueFunc, ttl, 0, 0, 0)
}

func (cacheThis *dbData) GetOrSetMany(ctx context.Context, dao dao.DaoInterface, idArr []any, ttl int64, field ...string) (list gdb.Result, err error) {
	for _, id := range idArr {
		value, noExistOfDb, errTmp := cacheThis.GetOrSet(ctx, dao, id, ttl, field...)
		if errTmp != nil {
			err = errTmp
			return
		}
		if noExistOfDb { //缓存的是数据库数据，就需要和数据库SQL查询一样。故无数据时不返回
			continue
		}
		var info gdb.Record
		value.Scan(&info)
		list = append(list, info)
	}
	return
}

func (cacheThis *dbData) GetOrSetPluck(ctx context.Context, dao dao.DaoInterface, idArr []any, ttl int64, field ...string) (record gdb.Record, err error) {
	record = gdb.Record{}
	for _, id := range idArr {
		value, noExistOfDb, errTmp := cacheThis.GetOrSet(ctx, dao, id, ttl, field...)
		if errTmp != nil {
			err = errTmp
			return
		}
		if noExistOfDb { //缓存的是数据库数据，就需要和数据库SQL查询一样。故无数据时不返回
			continue
		}
		record[gconv.String(id)] = value
	}
	return
}

func (cacheThis *dbData) Del(ctx context.Context, dao dao.DaoInterface, idArr ...any) (row int64, err error) {
	daoModel := dao.CtxDaoModel(ctx)
	keyArr := make([]string, len(idArr))
	for index, id := range idArr {
		keyArr[index] = cacheThis.key(daoModel, id)
	}
	row, err = internal.GetOrSet.Del(ctx, cacheThis.cache(), keyArr...)
	return
}
