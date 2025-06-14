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
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

var DbData = dbData{
	redis:             g.Redis(),
	methodCode:        ``,
	methodCodeOfArr:   `arr_`,
	methodCodeOfSet:   `set_`,
	methodCodeOfPluck: `pluck_`,
	methodCodeOfInfo:  `info_`,
	methodCodeOfList:  `list_`,
}

type dbData struct {
	redis             *gredis.Redis
	methodCode        string
	methodCodeOfArr   string
	methodCodeOfSet   string
	methodCodeOfPluck string
	methodCodeOfInfo  string
	methodCodeOfList  string
}

func (cacheThis *dbData) cache() *gredis.Redis {
	return cacheThis.redis
}

func (cacheThis *dbData) key(daoModel *dao.DaoModel, method string, idOrCode any) string {
	return fmt.Sprintf(consts.CACHE_DB_DATA, daoModel.DbGroup, daoModel.DbTable, method, idOrCode)
}

func (cacheThis *dbData) getOrSet(ctx context.Context, daoModel *dao.DaoModel, method string, code any, dbSelFunc func(daoModel *dao.DaoModel) (value any, ttl time.Duration, err error)) (value any, notExist bool, err error) {
	key := cacheThis.key(daoModel, method, code)
	value, notExist, err = internal.GetOrSet.GetOrSet(ctx, key, func() (value any, notExist bool, err error) {
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
		default:
			notExist = val == nil
		}
		if notExist {
			return
		}
		if ttl < time.Second {
			ttl = consts.CACHE_TIME_DEFAULT
		}
		err = cacheThis.cache().SetEX(ctx, key, gconv.String(value), gconv.Int64(ttl/time.Second))
		return
	}, func() (value any, notExist bool, err error) {
		value, err = cacheThis.cache().Get(ctx, key)
		notExist = value.(*gvar.Var).IsNil()
		return
	}, 0, 0, 0)
	return
}

func (cacheThis *dbData) del(ctx context.Context, daoModel *dao.DaoModel, method string, code any) (row int64, err error) {
	key := cacheThis.key(daoModel, method, code)
	row, err = cacheThis.cache().Del(ctx, key)
	if err != nil {
		return
	}
	internal.GetOrSet.Del(ctx, key)
	return
}

func (cacheThis *dbData) delById(ctx context.Context, daoModel *dao.DaoModel, method string, idArr []any) (row int64, err error) {
	keyArr := make([]string, len(idArr))
	for index := range idArr {
		keyArr[index] = cacheThis.key(daoModel, method, idArr[index])
	}
	row, err = cacheThis.cache().Del(ctx, keyArr...)
	if err != nil {
		return
	}
	internal.GetOrSet.Del(ctx, keyArr...)
	return
}

func (cacheThis *dbData) GetOrSet(ctx context.Context, daoModel *dao.DaoModel, code any, dbSelFunc func(daoModel *dao.DaoModel) (value *gvar.Var, ttl time.Duration, err error)) (value *gvar.Var, err error) {
	valueTmp, _, err := cacheThis.getOrSet(ctx, daoModel, cacheThis.methodCode, code, func(daoModel *dao.DaoModel) (value any, ttl time.Duration, err error) {
		value, ttl, err = dbSelFunc(daoModel)
		return
	})
	value, _ = valueTmp.(*gvar.Var)
	return
}

func (cacheThis *dbData) GetOrSetArr(ctx context.Context, daoModel *dao.DaoModel, code any, dbSelFunc func(daoModel *dao.DaoModel) (value []*gvar.Var, ttl time.Duration, err error)) (value []*gvar.Var, err error) {
	valueTmp, _, err := cacheThis.getOrSet(ctx, daoModel, cacheThis.methodCodeOfArr, code, func(daoModel *dao.DaoModel) (value any, ttl time.Duration, err error) {
		value, ttl, err = dbSelFunc(daoModel)
		return
	})
	value, _ = valueTmp.([]*gvar.Var)
	return
}

func (cacheThis *dbData) GetOrSetSet(ctx context.Context, daoModel *dao.DaoModel, code any, dbSelFunc func(daoModel *dao.DaoModel) (value map[*gvar.Var]struct{}, ttl time.Duration, err error)) (value map[*gvar.Var]struct{}, err error) {
	valueTmp, _, err := cacheThis.getOrSet(ctx, daoModel, cacheThis.methodCodeOfSet, code, func(daoModel *dao.DaoModel) (value any, ttl time.Duration, err error) {
		value, ttl, err = dbSelFunc(daoModel)
		return
	})
	value, _ = valueTmp.(map[*gvar.Var]struct{})
	return
}

func (cacheThis *dbData) GetOrSetPluck(ctx context.Context, daoModel *dao.DaoModel, code any, dbSelFunc func(daoModel *dao.DaoModel) (value gdb.Record, ttl time.Duration, err error)) (value gdb.Record, err error) {
	valueTmp, _, err := cacheThis.getOrSet(ctx, daoModel, cacheThis.methodCodeOfPluck, code, func(daoModel *dao.DaoModel) (value any, ttl time.Duration, err error) {
		value, ttl, err = dbSelFunc(daoModel)
		return
	})
	value, _ = valueTmp.(gdb.Record)
	return
}

func (cacheThis *dbData) GetOrSetInfo(ctx context.Context, daoModel *dao.DaoModel, code any, dbSelFunc func(daoModel *dao.DaoModel) (value gdb.Record, ttl time.Duration, err error)) (value gdb.Record, err error) {
	valueTmp, _, err := cacheThis.getOrSet(ctx, daoModel, cacheThis.methodCodeOfInfo, code, func(daoModel *dao.DaoModel) (value any, ttl time.Duration, err error) {
		value, ttl, err = dbSelFunc(daoModel)
		return
	})
	value, ok := valueTmp.(gdb.Record)
	if !ok {
		valueTmp.(*gvar.Var).Scan(&value)
	}
	return
}

func (cacheThis *dbData) GetOrSetList(ctx context.Context, daoModel *dao.DaoModel, code any, dbSelFunc func(daoModel *dao.DaoModel) (value gdb.Result, ttl time.Duration, err error)) (value gdb.Result, err error) {
	valueTmp, _, err := cacheThis.getOrSet(ctx, daoModel, cacheThis.methodCodeOfList, code, func(daoModel *dao.DaoModel) (value any, ttl time.Duration, err error) {
		value, ttl, err = dbSelFunc(daoModel)
		return
	})
	value, ok := valueTmp.(gdb.Result)
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

func (cacheThis *dbData) GetOrSetById(ctx context.Context, daoModel *dao.DaoModel, id any, ttlD time.Duration, field string) (value *gvar.Var, err error) {
	valueTmp, _, err := cacheThis.getOrSet(ctx, daoModel, cacheThis.methodCode, id, func(daoModel *dao.DaoModel) (value any, ttl time.Duration, err error) {
		value, err = daoModel.FilterPri(id).Value(field)
		ttl = ttlD
		return
	})
	value, _ = valueTmp.(*gvar.Var)
	return
}

/* func (cacheThis *dbData) GetOrSetArrById(ctx context.Context, daoModel *dao.DaoModel, idArr []any, ttlD time.Duration, field string) (value []*gvar.Var, err error) {
	var valueTmp any
	var notExist bool
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
		value[valueTmp.(*gvar.Var)] = struct{}{}
	}
	return
} */

func (cacheThis *dbData) GetOrSetPluckById(ctx context.Context, daoModel *dao.DaoModel, idArr []any, ttlD time.Duration, field string) (value gdb.Record, err error) {
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

func (cacheThis *dbData) GetOrSetInfoById(ctx context.Context, daoModel *dao.DaoModel, id any, ttlD time.Duration, fieldArr ...string) (value gdb.Record, err error) {
	valueTmp, _, err := cacheThis.getOrSet(ctx, daoModel, cacheThis.methodCodeOfInfo, id, func(daoModel *dao.DaoModel) (value any, ttl time.Duration, err error) {
		value, err = daoModel.FilterPri(id).Fields(fieldArr...).One()
		ttl = ttlD
		return
	})
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
