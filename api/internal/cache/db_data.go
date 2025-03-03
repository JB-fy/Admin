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
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

var DbData = dbData{redis: g.Redis()}

type dbData struct{ redis *gredis.Redis }

func (cacheThis *dbData) cache() *gredis.Redis {
	return cacheThis.redis
}

func (cacheThis *dbData) key(daoModel *dao.DaoModel, id any) string {
	return fmt.Sprintf(consts.CACHE_DB_DATA, daoModel.DbGroup, daoModel.DbTable, id)
}

// ttlOrField是字符串类型时，确保是能从数据库查询结果中获得，且值必须是数字或时间类型
func (cacheThis *dbData) getOrSet(ctx context.Context, daoModel *dao.DaoModel, id any, ttlOrField any, field ...string) (value *gvar.Var, noSetCache bool, err error) {
	redis := cacheThis.cache()
	key := cacheThis.key(daoModel, id)
	valueFunc := func() (value any, ttl int64, noSetCache bool, err error) {
		fieldArr := field
		if len(fieldArr) > 0 {
			if ttlField, ok := ttlOrField.(string); ok {
				fieldArr = append(fieldArr, ttlField)
			}
		}
		info, err := daoModel.FilterPri(id).Fields(fieldArr...).One()
		if err != nil {
			return
		}
		if info.IsEmpty() {
			noSetCache = true
			return
		}
		if ttlField, ok := ttlOrField.(string); ok {
			ttl = info[ttlField].GTime().Unix()
			if nowTime := gtime.Now().Unix(); ttl > nowTime {
				ttl = ttl - nowTime
			} else if ttl <= 0 || ttl > consts.CACHE_TIME_DEFAULT { //比当前时间小时，缓存时间不能超过默认缓存时间
				ttl = consts.CACHE_TIME_DEFAULT
			}
		} else {
			ttl = gconv.Int64(ttlOrField)
		}
		if len(field) == 1 {
			value = info[field[0]].String()
		} else {
			value = info.Json()
		}
		return
	}
	return internal.GetOrSet.GetOrSet(ctx, redis, key, valueFunc, 0, 0, 0)
}

func (cacheThis *dbData) GetOrSet(ctx context.Context, daoModel *dao.DaoModel, id any, ttlOrField any, field ...string) (value *gvar.Var, err error) {
	value, _, err = cacheThis.getOrSet(ctx, daoModel, id, ttlOrField, field...)
	return
}

func (cacheThis *dbData) GetOrSetMany(ctx context.Context, daoModel *dao.DaoModel, idArr []any, ttlOrField any, field ...string) (list gdb.Result, err error) {
	for _, id := range idArr {
		value, noSetCache, errTmp := cacheThis.getOrSet(ctx, daoModel.ResetNew(), id, ttlOrField, field...)
		if errTmp != nil {
			err = errTmp
			return
		}
		if noSetCache { //缓存的是数据库数据，就需要和数据库SQL查询一样。故无数据时不返回
			continue
		}
		var info gdb.Record
		value.Scan(&info)
		list = append(list, info)
	}
	return
}

func (cacheThis *dbData) GetOrSetPluck(ctx context.Context, daoModel *dao.DaoModel, idArr []any, ttlOrField any, field ...string) (record gdb.Record, err error) {
	record = gdb.Record{}
	for _, id := range idArr {
		value, noSetCache, errTmp := cacheThis.getOrSet(ctx, daoModel.ResetNew(), id, ttlOrField, field...)
		if errTmp != nil {
			err = errTmp
			return
		}
		if noSetCache { //缓存的是数据库数据，就需要和数据库SQL查询一样。故无数据时不返回
			continue
		}
		record[gconv.String(id)] = value
	}
	return
}

func (cacheThis *dbData) Del(ctx context.Context, daoModel *dao.DaoModel, idArr ...any) (row int64, err error) {
	keyArr := make([]string, len(idArr))
	for index, id := range idArr {
		keyArr[index] = cacheThis.key(daoModel, id)
	}
	row, err = internal.GetOrSet.Del(ctx, cacheThis.cache(), keyArr...)
	return
}
