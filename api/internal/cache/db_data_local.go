package cache

import (
	"api/internal/cache/internal"
	"api/internal/consts"
	"api/internal/dao"
	"context"
	"fmt"
	"time"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/patrickmn/go-cache"
)

var DbDataLocal = dbDataLocal{
	goCache:  cache.New(0, 0),
	goCache1: cache.New(consts.CACHE_TIME_DEFAULT, 2*time.Hour),
	cacheKeyMap: map[string]uint8{
		`default:item_task`:                 1,
		`default:item_task_rel_to_plt_auth`: 1,
	},
	methodCode:       ``,
	methodCodeOfInfo: `info_`,
	methodCodeOfList: `list_`,
}

type dbDataLocal struct {
	goCache          *cache.Cache
	goCache1         *cache.Cache
	cacheKeyMap      map[string]uint8
	methodCode       string
	methodCodeOfInfo string
	methodCodeOfList string
}

var cacheMap = map[uint8]*cache.Cache{
	1: DbDataLocal.goCache1,
}

func (cacheThis *dbDataLocal) Flush(cacheKey uint8) {
	cacheThis.parseCache(cacheKey).Flush()
}

// 解析缓存分库
func (cacheThis *dbDataLocal) parseCache(cacheKey uint8) *cache.Cache {
	if _, ok := cacheMap[cacheKey]; ok {
		return cacheMap[cacheKey]
	}
	return cacheThis.goCache
}

func (cacheThis *dbDataLocal) cache(daoModel *dao.DaoModel) *cache.Cache {
	return cacheThis.parseCache(cacheThis.cacheKeyMap[daoModel.DbGroup+`:`+daoModel.DbTable])
}

func (cacheThis *dbDataLocal) key(daoModel *dao.DaoModel, method string, idOrCode any) string {
	return fmt.Sprintf(consts.CACHE_DB_DATA, daoModel.DbGroup, daoModel.DbTable, method, idOrCode)
}

func (cacheThis *dbDataLocal) getOrSet(ctx context.Context, daoModel *dao.DaoModel, method string, code any, dbSelFunc func(daoModel *dao.DaoModel) (value any, ttl time.Duration, err error)) (value any, notExist bool, err error) {
	key := cacheThis.key(daoModel, method, code)
	value, notExist, err = internal.GetOrSetLocal.GetOrSetLocal(ctx, key, func() (value any, notExist bool, err error) {
		value, ttl, err := dbSelFunc(daoModel)
		if err != nil {
			return
		}
		switch val := value.(type) {
		case *gvar.Var:
			notExist = val.IsNil()
		case gdb.Record:
			notExist = val.IsEmpty()
		case gdb.Result:
			notExist = len(val) == 0
		default:
			notExist = val == nil
		}
		if notExist {
			return
		}
		cacheThis.cache(daoModel).Set(key, value, ttl)
		return
	}, func() (value any, notExist bool, err error) {
		value, notExist = cacheThis.cache(daoModel).Get(key)
		notExist = !notExist
		return
	})
	return
}

func (cacheThis *dbDataLocal) GetOrSet(ctx context.Context, daoModel *dao.DaoModel, code any, dbSelFunc func(daoModel *dao.DaoModel) (value *gvar.Var, ttl time.Duration, err error)) (value *gvar.Var, err error) {
	valueTmp, _, err := cacheThis.getOrSet(ctx, daoModel, cacheThis.methodCode, code, func(daoModel *dao.DaoModel) (value any, ttl time.Duration, err error) {
		value, ttl, err = dbSelFunc(daoModel)
		return
	})
	value, _ = valueTmp.(*gvar.Var)
	return
}

func (cacheThis *dbDataLocal) GetOrSetInfo(ctx context.Context, daoModel *dao.DaoModel, code any, dbSelFunc func(daoModel *dao.DaoModel) (value gdb.Record, ttl time.Duration, err error)) (value gdb.Record, err error) {
	valueTmp, _, err := cacheThis.getOrSet(ctx, daoModel, cacheThis.methodCodeOfInfo, code, func(daoModel *dao.DaoModel) (value any, ttl time.Duration, err error) {
		value, ttl, err = dbSelFunc(daoModel)
		return
	})
	value, _ = valueTmp.(gdb.Record)
	return
}

func (cacheThis *dbDataLocal) GetOrSetList(ctx context.Context, daoModel *dao.DaoModel, code any, dbSelFunc func(daoModel *dao.DaoModel) (value gdb.Result, ttl time.Duration, err error)) (value gdb.Result, err error) {
	valueTmp, _, err := cacheThis.getOrSet(ctx, daoModel, cacheThis.methodCodeOfList, code, func(daoModel *dao.DaoModel) (value any, ttl time.Duration, err error) {
		value, ttl, err = dbSelFunc(daoModel)
		return
	})
	value, _ = valueTmp.(gdb.Result)
	return
}

func (cacheThis *dbDataLocal) GetOrSetById(ctx context.Context, daoModel *dao.DaoModel, id any, ttlD time.Duration, field string) (value *gvar.Var, err error) {
	valueTmp, _, err := cacheThis.getOrSet(ctx, daoModel, cacheThis.methodCode, id, func(daoModel *dao.DaoModel) (value any, ttl time.Duration, err error) {
		value, err = daoModel.FilterPri(id).Value(field)
		ttl = ttlD
		return
	})
	value, _ = valueTmp.(*gvar.Var)
	return
}

func (cacheThis *dbDataLocal) GetOrSetInfoById(ctx context.Context, daoModel *dao.DaoModel, id any, ttlD time.Duration, fieldArr ...string) (value gdb.Record, err error) {
	valueTmp, _, err := cacheThis.getOrSet(ctx, daoModel, cacheThis.methodCodeOfInfo, id, func(daoModel *dao.DaoModel) (value any, ttl time.Duration, err error) {
		value, err = daoModel.FilterPri(id).Fields(fieldArr...).One()
		ttl = ttlD
		return
	})
	value, _ = valueTmp.(gdb.Record)
	return
}

func (cacheThis *dbDataLocal) GetOrSetListById(ctx context.Context, daoModel *dao.DaoModel, idArr []any, ttlD time.Duration, fieldArr ...string) (value gdb.Result, err error) {
	var valueTmp any
	var notExist bool
	for index := range idArr {
		valueTmp, notExist, err = cacheThis.getOrSet(ctx, daoModel, cacheThis.methodCodeOfInfo, idArr[index], func(daoModel *dao.DaoModel) (value any, ttl time.Duration, err error) {
			value, err = daoModel.ResetNew().FilterPri(idArr[index]).Fields(fieldArr...).One()
			ttl = ttlD
			return
		})
		if err != nil {
			return
		}
		if notExist { //缓存的是数据库数据，就需要和数据库SQL查询一样。故无数据时不返回
			continue
		}
		value = append(value, valueTmp.(gdb.Record))
	}
	return
}

func (cacheThis *dbDataLocal) GetOrSetPluckById(ctx context.Context, daoModel *dao.DaoModel, idArr []any, ttlD time.Duration, field string) (value gdb.Record, err error) {
	var valueTmp any
	var notExist bool
	value = gdb.Record{}
	for index := range idArr {
		valueTmp, notExist, err = cacheThis.getOrSet(ctx, daoModel, cacheThis.methodCode, idArr[index], func(daoModel *dao.DaoModel) (value any, ttl time.Duration, err error) {
			value, err = daoModel.ResetNew().FilterPri(idArr[index]).Value(field)
			ttl = ttlD
			return
		})
		if err != nil {
			return
		}
		if notExist { //缓存的是数据库数据，就需要和数据库SQL查询一样。故无数据时不返回
			continue
		}
		value[gconv.String(idArr[index])], _ = valueTmp.(*gvar.Var)
	}
	return
}
