package cache

import (
	"api/internal/cache/common"
	"api/internal/consts"
	"api/internal/dao"
	"api/internal/utils/jbredis"
	"context"
	"fmt"
	"time"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/redis/go-redis/v9"
)

var DbData = dbData{
	methodCode:        ``,
	methodCodeOfArr:   `arr_`,
	methodCodeOfSet:   `set_`,
	methodCodeOfPluck: `pluck_`,
	methodCodeOfInfo:  `info_`,
	methodCodeOfList:  `list_`,
	methodCodeOfTree:  `tree_`,
}

type dbData struct {
	methodCode        string
	methodCodeOfArr   string
	methodCodeOfSet   string
	methodCodeOfPluck string
	methodCodeOfInfo  string
	methodCodeOfList  string
	methodCodeOfTree  string
}

func (cacheThis *dbData) cache() redis.UniversalClient {
	return jbredis.DB()
}

func (cacheThis *dbData) key(daoModel *dao.DaoModel, method string, idOrCode any) string {
	return fmt.Sprintf(consts.CACHE_DB_DATA, daoModel.DbGroup, daoModel.DbTable, method, idOrCode)
}

func (cacheThis *dbData) getOrSet(ctx context.Context, daoModel *dao.DaoModel, method string, code any, dbSelFunc func(daoModel *dao.DaoModel) (value any, ttl time.Duration, err error), numLock, numRead uint8, oneTime time.Duration) (value any, notExist bool, err error) {
	key := cacheThis.key(daoModel, method, code)
	value, notExist, err = common.GetOrSet.GetOrSet(ctx, key, func() (value any, notExist bool, err error) {
		// 查询时如果刚好有更新或删除时，可能出现先删除缓存再保存旧数据的情况
		// 解决方法：开启事务，且dbSelFunc方法返回的与value有关的数据，都必须使用LockUpdate()上锁做查询
		isCache := false
		err = daoModel.Master().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
			var ttl time.Duration
			value, ttl, err = dbSelFunc(daoModel.Ctx(ctx).LockUpdate()) //先上锁，防止dbSelFunc方法内忘记上锁。注意：在dbSelFunc方法内使用ResetNew()等方法时，此时的锁将失效
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
			if ttl < time.Second {
				ttl = consts.CACHE_TIME_DEFAULT
			}
			err = cacheThis.cache().SetEx(ctx, key, gconv.String(value), ttl).Err()
			isCache = true
			return
		})
		if isCache && err != nil {
			g.Log().Error(ctx, `数据库事物报错：`+err.Error())
			_, errTmp := cacheThis.cache().Del(ctx, key).Result()
			if errTmp != nil {
				return
			}
			common.GetOrSet.Del(ctx, key)
		}
		return
	}, func() (value any, notExist bool, err error) {
		value, err = cacheThis.cache().Get(ctx, key).Result()
		if err == nil {
			value = gvar.New(value)
		} else {
			value = gvar.New(nil)
			notExist = err == redis.Nil
			if notExist {
				err = nil
			}
		}
		return
	}, numLock, numRead, oneTime)
	return
}

func (cacheThis *dbData) del(ctx context.Context, daoModel *dao.DaoModel, method string, code any) (row int64, err error) {
	key := cacheThis.key(daoModel, method, code)
	row, err = cacheThis.cache().Del(ctx, key).Result()
	if err != nil {
		return
	}
	common.GetOrSet.Del(ctx, key)
	return
}

func (cacheThis *dbData) delById(ctx context.Context, daoModel *dao.DaoModel, method string, idArr []any) (row int64, err error) {
	keyArr := make([]string, len(idArr))
	for index := range idArr {
		keyArr[index] = cacheThis.key(daoModel, method, idArr[index])
	}
	row, err = cacheThis.cache().Del(ctx, keyArr...).Result()
	if err != nil {
		return
	}
	common.GetOrSet.Del(ctx, keyArr...)
	return
}

func (cacheThis *dbData) GetOrSet(ctx context.Context, daoModel *dao.DaoModel, code any, dbSelFunc func(daoModel *dao.DaoModel) (value *gvar.Var, ttl time.Duration, err error), numLock, numRead uint8, oneTime time.Duration) (value *gvar.Var, err error) {
	valueTmp, _, err := cacheThis.getOrSet(ctx, daoModel, cacheThis.methodCode, code, func(daoModel *dao.DaoModel) (value any, ttl time.Duration, err error) {
		value, ttl, err = dbSelFunc(daoModel)
		return
	}, numLock, numRead, oneTime)
	value, _ = valueTmp.(*gvar.Var)
	return
}

func (cacheThis *dbData) GetOrSetArr(ctx context.Context, daoModel *dao.DaoModel, code any, dbSelFunc func(daoModel *dao.DaoModel) (value []*gvar.Var, ttl time.Duration, err error), numLock, numRead uint8, oneTime time.Duration) (value []*gvar.Var, err error) {
	valueTmp, _, err := cacheThis.getOrSet(ctx, daoModel, cacheThis.methodCodeOfArr, code, func(daoModel *dao.DaoModel) (value any, ttl time.Duration, err error) {
		value, ttl, err = dbSelFunc(daoModel)
		return
	}, numLock, numRead, oneTime)
	value, ok := valueTmp.([]*gvar.Var)
	if !ok {
		valueTmp.(*gvar.Var).Scan(&value)
	}
	return
}

func (cacheThis *dbData) GetOrSetSet(ctx context.Context, daoModel *dao.DaoModel, code any, dbSelFunc func(daoModel *dao.DaoModel) (value map[*gvar.Var]struct{}, ttl time.Duration, err error), numLock, numRead uint8, oneTime time.Duration) (value map[*gvar.Var]struct{}, err error) {
	valueTmp, _, err := cacheThis.getOrSet(ctx, daoModel, cacheThis.methodCodeOfSet, code, func(daoModel *dao.DaoModel) (value any, ttl time.Duration, err error) {
		value, ttl, err = dbSelFunc(daoModel)
		return
	}, numLock, numRead, oneTime)
	value, ok := valueTmp.(map[*gvar.Var]struct{})
	if !ok {
		valueTmp.(*gvar.Var).Scan(&value)
	}
	return
}

func (cacheThis *dbData) GetOrSetPluck(ctx context.Context, daoModel *dao.DaoModel, code any, dbSelFunc func(daoModel *dao.DaoModel) (value gdb.Record, ttl time.Duration, err error), numLock, numRead uint8, oneTime time.Duration) (value gdb.Record, err error) {
	valueTmp, _, err := cacheThis.getOrSet(ctx, daoModel, cacheThis.methodCodeOfPluck, code, func(daoModel *dao.DaoModel) (value any, ttl time.Duration, err error) {
		value, ttl, err = dbSelFunc(daoModel)
		return
	}, numLock, numRead, oneTime)
	value, ok := valueTmp.(gdb.Record)
	if !ok {
		valueTmp.(*gvar.Var).Scan(&value)
	}
	return
}

func (cacheThis *dbData) GetOrSetInfo(ctx context.Context, daoModel *dao.DaoModel, code any, dbSelFunc func(daoModel *dao.DaoModel) (value gdb.Record, ttl time.Duration, err error), numLock, numRead uint8, oneTime time.Duration) (value gdb.Record, err error) {
	valueTmp, _, err := cacheThis.getOrSet(ctx, daoModel, cacheThis.methodCodeOfInfo, code, func(daoModel *dao.DaoModel) (value any, ttl time.Duration, err error) {
		value, ttl, err = dbSelFunc(daoModel)
		return
	}, numLock, numRead, oneTime)
	value, ok := valueTmp.(gdb.Record)
	if !ok {
		valueTmp.(*gvar.Var).Scan(&value)
	}
	return
}

func (cacheThis *dbData) GetOrSetList(ctx context.Context, daoModel *dao.DaoModel, code any, dbSelFunc func(daoModel *dao.DaoModel) (value gdb.Result, ttl time.Duration, err error), numLock, numRead uint8, oneTime time.Duration) (value gdb.Result, err error) {
	valueTmp, _, err := cacheThis.getOrSet(ctx, daoModel, cacheThis.methodCodeOfList, code, func(daoModel *dao.DaoModel) (value any, ttl time.Duration, err error) {
		value, ttl, err = dbSelFunc(daoModel)
		return
	}, numLock, numRead, oneTime)
	value, ok := valueTmp.(gdb.Result)
	if !ok {
		valueTmp.(*gvar.Var).Scan(&value)
	}
	return
}

func (cacheThis *dbData) GetOrSetTree(ctx context.Context, daoModel *dao.DaoModel, code any, dbSelFunc func(daoModel *dao.DaoModel) (value g.List, ttl time.Duration, err error), numLock, numRead uint8, oneTime time.Duration) (value g.List, err error) {
	valueTmp, _, err := cacheThis.getOrSet(ctx, daoModel, cacheThis.methodCodeOfTree, code, func(daoModel *dao.DaoModel) (value any, ttl time.Duration, err error) {
		value, ttl, err = dbSelFunc(daoModel)
		return
	}, numLock, numRead, oneTime)
	value, ok := valueTmp.(g.List)
	if !ok {
		valueTmp.(*gvar.Var).Scan(&value)
	}
	return
}

func (cacheThis *dbData) Del(ctx context.Context, daoModel *dao.DaoModel, code any) (int64, error) {
	return cacheThis.del(ctx, daoModel, cacheThis.methodCode, code)
}

func (cacheThis *dbData) DelArr(ctx context.Context, daoModel *dao.DaoModel, code any) (int64, error) {
	return cacheThis.del(ctx, daoModel, cacheThis.methodCodeOfArr, code)
}

func (cacheThis *dbData) DelSet(ctx context.Context, daoModel *dao.DaoModel, code any) (int64, error) {
	return cacheThis.del(ctx, daoModel, cacheThis.methodCodeOfSet, code)
}

func (cacheThis *dbData) DelPluck(ctx context.Context, daoModel *dao.DaoModel, code any) (int64, error) {
	return cacheThis.del(ctx, daoModel, cacheThis.methodCodeOfPluck, code)
}

func (cacheThis *dbData) DelInfo(ctx context.Context, daoModel *dao.DaoModel, code any) (int64, error) {
	return cacheThis.del(ctx, daoModel, cacheThis.methodCodeOfInfo, code)
}

func (cacheThis *dbData) DelList(ctx context.Context, daoModel *dao.DaoModel, code any) (int64, error) {
	return cacheThis.del(ctx, daoModel, cacheThis.methodCodeOfList, code)
}

func (cacheThis *dbData) DelTree(ctx context.Context, daoModel *dao.DaoModel, code any) (int64, error) {
	return cacheThis.del(ctx, daoModel, cacheThis.methodCodeOfTree, code)
}

func (cacheThis *dbData) GetOrSetById(ctx context.Context, daoModel *dao.DaoModel, id any, ttlD time.Duration, field string) (value *gvar.Var, err error) {
	valueTmp, _, err := cacheThis.getOrSet(ctx, daoModel, cacheThis.methodCode, id, func(daoModel *dao.DaoModel) (value any, ttl time.Duration, err error) {
		value, err = daoModel.ResetNew().LockUpdate().FilterPri(id).Value(field)
		ttl = ttlD
		return
	}, 0, 0, 0)
	value, _ = valueTmp.(*gvar.Var)
	return
}

func (cacheThis *dbData) GetOrSetArrById(ctx context.Context, daoModel *dao.DaoModel, idArr []any, ttlD time.Duration, field string) (value []*gvar.Var, err error) {
	var valueTmp any
	var notExist bool
	for index := range idArr {
		valueTmp, notExist, err = cacheThis.getOrSet(ctx, daoModel, cacheThis.methodCode, idArr[index], func(daoModel *dao.DaoModel) (value any, ttl time.Duration, err error) {
			value, err = daoModel.ResetNew().LockUpdate().FilterPri(idArr[index]).Value(field)
			ttl = ttlD
			return
		}, 0, 0, 0)
		if err != nil {
			return
		}
		if notExist { //缓存的是数据库数据，就需要和数据库SQL查询一样。故无数据时不返回
			continue
		}
		value = append(value, valueTmp.(*gvar.Var))
	}
	return
}

func (cacheThis *dbData) GetOrSetSetById(ctx context.Context, daoModel *dao.DaoModel, idArr []any, ttlD time.Duration, field string) (value map[*gvar.Var]struct{}, err error) {
	var valueTmp any
	var notExist bool
	value = map[*gvar.Var]struct{}{}
	for index := range idArr {
		valueTmp, notExist, err = cacheThis.getOrSet(ctx, daoModel, cacheThis.methodCode, idArr[index], func(daoModel *dao.DaoModel) (value any, ttl time.Duration, err error) {
			value, err = daoModel.ResetNew().LockUpdate().FilterPri(idArr[index]).Value(field)
			ttl = ttlD
			return
		}, 0, 0, 0)
		if err != nil {
			return
		}
		if notExist { //缓存的是数据库数据，就需要和数据库SQL查询一样。故无数据时不返回
			continue
		}
		value[valueTmp.(*gvar.Var)] = struct{}{}
	}
	return
}

func (cacheThis *dbData) GetOrSetPluckById(ctx context.Context, daoModel *dao.DaoModel, idArr []any, ttlD time.Duration, field string) (value gdb.Record, err error) {
	var valueTmp any
	var notExist bool
	value = gdb.Record{}
	for index := range idArr {
		valueTmp, notExist, err = cacheThis.getOrSet(ctx, daoModel, cacheThis.methodCode, idArr[index], func(daoModel *dao.DaoModel) (value any, ttl time.Duration, err error) {
			value, err = daoModel.ResetNew().LockUpdate().FilterPri(idArr[index]).Value(field)
			ttl = ttlD
			return
		}, 0, 0, 0)
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

func (cacheThis *dbData) GetOrSetInfoById(ctx context.Context, daoModel *dao.DaoModel, id any, ttlD time.Duration, fieldArr ...string) (value gdb.Record, err error) {
	valueTmp, _, err := cacheThis.getOrSet(ctx, daoModel, cacheThis.methodCodeOfInfo, id, func(daoModel *dao.DaoModel) (value any, ttl time.Duration, err error) {
		value, err = daoModel.ResetNew().LockUpdate().FilterPri(id).Fields(fieldArr...).One()
		ttl = ttlD
		return
	}, 0, 0, 0)
	value, ok := valueTmp.(gdb.Record)
	if !ok {
		valueTmp.(*gvar.Var).Scan(&value)
	}
	return
}

func (cacheThis *dbData) GetOrSetListById(ctx context.Context, daoModel *dao.DaoModel, idArr []any, ttlD time.Duration, fieldArr ...string) (value gdb.Result, err error) {
	var valueTmp any
	var notExist bool
	var ok bool
	for index := range idArr {
		valueTmp, notExist, err = cacheThis.getOrSet(ctx, daoModel, cacheThis.methodCodeOfInfo, idArr[index], func(daoModel *dao.DaoModel) (value any, ttl time.Duration, err error) {
			value, err = daoModel.ResetNew().LockUpdate().FilterPri(idArr[index]).Fields(fieldArr...).One()
			ttl = ttlD
			return
		}, 0, 0, 0)
		if err != nil {
			return
		}
		if notExist { //缓存的是数据库数据，就需要和数据库SQL查询一样。故无数据时不返回
			continue
		}
		value = append(value, gdb.Record{})
		value[len(value)-1], ok = valueTmp.(gdb.Record)
		if !ok {
			valueTmp.(*gvar.Var).Scan(&value[len(value)-1])
		}
	}
	return
}

func (cacheThis *dbData) DelById(ctx context.Context, daoModel *dao.DaoModel, idArr ...any) (int64, error) {
	return cacheThis.delById(ctx, daoModel, cacheThis.methodCode, idArr)
}

func (cacheThis *dbData) DelInfoById(ctx context.Context, daoModel *dao.DaoModel, idArr ...any) (int64, error) {
	return cacheThis.delById(ctx, daoModel, cacheThis.methodCodeOfInfo, idArr)
}
