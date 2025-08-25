package cache

import (
	"api/internal/cache/internal"
	"api/internal/consts"
	"api/internal/dao"
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/patrickmn/go-cache"
)

var DbDataLocal = dbDataLocal{
	goCacheMap: sync.Map{}, //第一个缓存库默认凌晨3点清空
	cacheKeyMap: map[string]uint8{ // 可根据 数据库或表 使用不同缓存库。不同缓存库每隔（值*consts.CACHE_LOCAL_INTERVAL_MINUTE）时间会删除缓存，可配合定时器使全部服务器同步缓存
		// `default:auth_scene`:  0, //默认为0，即第一个缓存库
		`default:auth_menu`:   1,
		`default:auth_action`: 1,
		// `default:upload`:      2,
		// `default:pay`:         1,
		// `default:pay_channel`: 1,
		// `default:pay_scene`:   1,
		// `default:app`:         0,
	},
	methodCode:        ``,
	methodCodeOfArr:   `arr_`,
	methodCodeOfSet:   `set_`,
	methodCodeOfPluck: `pluck_`,
	methodCodeOfInfo:  `info_`,
	methodCodeOfList:  `list_`,
	methodCodeOfTree:  `tree_`,
}

type dbDataLocal struct {
	goCacheMap        sync.Map
	cacheKeyMap       map[string]uint8
	methodCode        string
	methodCodeOfArr   string
	methodCodeOfSet   string
	methodCodeOfPluck string
	methodCodeOfInfo  string
	methodCodeOfList  string
	methodCodeOfTree  string
}

func (cacheThis *dbDataLocal) Flush(ctx context.Context) {
	nowTime := gtime.Now()
	hour := nowTime.Hour()
	minute := nowTime.Minute()
	totalMinute := hour*60 + minute
	cacheThis.goCacheMap.Range(func(key, value any) bool {
		if cacheKey := key.(uint8); cacheKey > 0 {
			if totalMinute%(int(cacheKey)*consts.CACHE_LOCAL_INTERVAL_MINUTE) == 0 {
				value.(*cache.Cache).Flush()
			}
		} else if hour == 3 && minute == 0 { //第一个缓存库默认凌晨3点清空
			value.(*cache.Cache).Flush()
		}
		return true
	})
}

// 解析缓存分库
func (cacheThis *dbDataLocal) parseCache(cacheKey uint8) *cache.Cache {
	tmp, _ := cacheThis.goCacheMap.LoadOrStore(cacheKey, cache.New(0, 0))
	return tmp.(*cache.Cache)
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
		case []*gvar.Var:
			notExist = len(val) == 0
		case map[*gvar.Var]struct{}:
			notExist = len(val) == 0
		case gdb.Record:
			notExist = val.IsEmpty()
		case gdb.Result:
			notExist = len(val) == 0
		case g.List:
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

func (cacheThis *dbDataLocal) del(ctx context.Context, daoModel *dao.DaoModel, method string, code any) {
	cacheThis.cache(daoModel).Delete(cacheThis.key(daoModel, method, code))
}

func (cacheThis *dbDataLocal) delById(ctx context.Context, daoModel *dao.DaoModel, method string, idArr []any) {
	for index := range idArr {
		cacheThis.cache(daoModel).Delete(cacheThis.key(daoModel, method, idArr[index]))
	}
}

func (cacheThis *dbDataLocal) GetOrSet(ctx context.Context, daoModel *dao.DaoModel, code any, dbSelFunc func(daoModel *dao.DaoModel) (value *gvar.Var, ttl time.Duration, err error)) (value *gvar.Var, err error) {
	valueTmp, _, err := cacheThis.getOrSet(ctx, daoModel, cacheThis.methodCode, code, func(daoModel *dao.DaoModel) (value any, ttl time.Duration, err error) {
		value, ttl, err = dbSelFunc(daoModel)
		return
	})
	value, _ = valueTmp.(*gvar.Var)
	return
}

func (cacheThis *dbDataLocal) GetOrSetArr(ctx context.Context, daoModel *dao.DaoModel, code any, dbSelFunc func(daoModel *dao.DaoModel) (value []*gvar.Var, ttl time.Duration, err error)) (value []*gvar.Var, err error) {
	valueTmp, _, err := cacheThis.getOrSet(ctx, daoModel, cacheThis.methodCodeOfArr, code, func(daoModel *dao.DaoModel) (value any, ttl time.Duration, err error) {
		value, ttl, err = dbSelFunc(daoModel)
		return
	})
	value, _ = valueTmp.([]*gvar.Var)
	return
}

func (cacheThis *dbDataLocal) GetOrSetSet(ctx context.Context, daoModel *dao.DaoModel, code any, dbSelFunc func(daoModel *dao.DaoModel) (value map[*gvar.Var]struct{}, ttl time.Duration, err error)) (value map[*gvar.Var]struct{}, err error) {
	valueTmp, _, err := cacheThis.getOrSet(ctx, daoModel, cacheThis.methodCodeOfSet, code, func(daoModel *dao.DaoModel) (value any, ttl time.Duration, err error) {
		value, ttl, err = dbSelFunc(daoModel)
		return
	})
	value, _ = valueTmp.(map[*gvar.Var]struct{})
	return
}

func (cacheThis *dbDataLocal) GetOrSetPluck(ctx context.Context, daoModel *dao.DaoModel, code any, dbSelFunc func(daoModel *dao.DaoModel) (value gdb.Record, ttl time.Duration, err error)) (value gdb.Record, err error) {
	valueTmp, _, err := cacheThis.getOrSet(ctx, daoModel, cacheThis.methodCodeOfPluck, code, func(daoModel *dao.DaoModel) (value any, ttl time.Duration, err error) {
		value, ttl, err = dbSelFunc(daoModel)
		return
	})
	value, _ = valueTmp.(gdb.Record)
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

func (cacheThis *dbDataLocal) GetOrSetTree(ctx context.Context, daoModel *dao.DaoModel, code any, dbSelFunc func(daoModel *dao.DaoModel) (value g.List, ttl time.Duration, err error)) (value g.List, err error) {
	valueTmp, _, err := cacheThis.getOrSet(ctx, daoModel, cacheThis.methodCodeOfTree, code, func(daoModel *dao.DaoModel) (value any, ttl time.Duration, err error) {
		value, ttl, err = dbSelFunc(daoModel)
		return
	})
	value, _ = valueTmp.(g.List)
	return
}

func (cacheThis *dbDataLocal) Del(ctx context.Context, daoModel *dao.DaoModel, code any) {
	cacheThis.del(ctx, daoModel, cacheThis.methodCode, code)
}

func (cacheThis *dbDataLocal) DelArr(ctx context.Context, daoModel *dao.DaoModel, code any) {
	cacheThis.del(ctx, daoModel, cacheThis.methodCodeOfArr, code)
}

func (cacheThis *dbDataLocal) DelSet(ctx context.Context, daoModel *dao.DaoModel, code any) {
	cacheThis.del(ctx, daoModel, cacheThis.methodCodeOfSet, code)
}

func (cacheThis *dbDataLocal) DelPluck(ctx context.Context, daoModel *dao.DaoModel, code any) {
	cacheThis.del(ctx, daoModel, cacheThis.methodCodeOfPluck, code)
}

func (cacheThis *dbDataLocal) DelInfo(ctx context.Context, daoModel *dao.DaoModel, code any) {
	cacheThis.del(ctx, daoModel, cacheThis.methodCodeOfInfo, code)
}

func (cacheThis *dbDataLocal) DelList(ctx context.Context, daoModel *dao.DaoModel, code any) {
	cacheThis.del(ctx, daoModel, cacheThis.methodCodeOfList, code)
}

func (cacheThis *dbDataLocal) GetOrSetById(ctx context.Context, daoModel *dao.DaoModel, id any, ttlD time.Duration, field string) (value *gvar.Var, err error) {
	valueTmp, _, err := cacheThis.getOrSet(ctx, daoModel, cacheThis.methodCode, id, func(daoModel *dao.DaoModel) (value any, ttl time.Duration, err error) {
		value, err = daoModel.Master().FilterPri(id).Value(field)
		ttl = ttlD
		return
	})
	value, _ = valueTmp.(*gvar.Var)
	return
}

func (cacheThis *dbDataLocal) GetOrSetPluckById(ctx context.Context, daoModel *dao.DaoModel, idArr []any, ttlD time.Duration, field string) (value gdb.Record, err error) {
	var valueTmp any
	var notExist bool
	value = gdb.Record{}
	for index := range idArr {
		valueTmp, notExist, err = cacheThis.getOrSet(ctx, daoModel, cacheThis.methodCode, idArr[index], func(daoModel *dao.DaoModel) (value any, ttl time.Duration, err error) {
			value, err = daoModel.ResetNew().Master().FilterPri(idArr[index]).Value(field)
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

func (cacheThis *dbDataLocal) GetOrSetInfoById(ctx context.Context, daoModel *dao.DaoModel, id any, ttlD time.Duration, fieldArr ...string) (value gdb.Record, err error) {
	valueTmp, _, err := cacheThis.getOrSet(ctx, daoModel, cacheThis.methodCodeOfInfo, id, func(daoModel *dao.DaoModel) (value any, ttl time.Duration, err error) {
		value, err = daoModel.Master().FilterPri(id).Fields(fieldArr...).One()
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
			value, err = daoModel.ResetNew().Master().FilterPri(idArr[index]).Fields(fieldArr...).One()
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

func (cacheThis *dbDataLocal) DelById(ctx context.Context, daoModel *dao.DaoModel, idArr ...any) {
	cacheThis.delById(ctx, daoModel, cacheThis.methodCode, idArr)
}

func (cacheThis *dbDataLocal) DelInfoById(ctx context.Context, daoModel *dao.DaoModel, idArr ...any) {
	cacheThis.delById(ctx, daoModel, cacheThis.methodCodeOfInfo, idArr)
}
