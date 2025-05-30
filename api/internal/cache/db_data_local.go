package cache

import (
	"api/internal/consts"
	"api/internal/dao"
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/patrickmn/go-cache"
)

var DbDataLocal = dbDataLocal{goCache: cache.New(0, 0)}

type dbDataLocal struct {
	goCache *cache.Cache
	muMap   sync.Map //存放所有缓存KEY的锁（当前服务器用）
}

func (cacheThis *dbDataLocal) cache() *cache.Cache {
	return cacheThis.goCache
}

func (cacheThis *dbDataLocal) key(daoModel *dao.DaoModel, id any) string {
	return fmt.Sprintf(consts.CACHE_DB_DATA, daoModel.DbGroup, daoModel.DbTable, id)
}

func (cacheThis *dbDataLocal) Set(ctx context.Context, daoModel *dao.DaoModel, id any, value string, ttl time.Duration) {
	cacheThis.cache().Set(cacheThis.key(daoModel, id), gvar.New(value), time.Duration(ttl)*time.Second)
}

func (cacheThis *dbDataLocal) Get(ctx context.Context, daoModel *dao.DaoModel, id any) (value *gvar.Var) {
	if valueTmp, ok := cacheThis.cache().Get(cacheThis.key(daoModel, id)); ok {
		value = valueTmp.(*gvar.Var)
		return
	}
	return
}

func (cacheThis *dbDataLocal) GetInfo(ctx context.Context, daoModel *dao.DaoModel, id any) (info gdb.Record) {
	cacheThis.Get(ctx, daoModel, id).Scan(&info)
	return
}

func (cacheThis *dbDataLocal) GetList(ctx context.Context, daoModel *dao.DaoModel, id any) (list gdb.Result) {
	cacheThis.Get(ctx, daoModel, id).Scan(&list)
	return
}

func (cacheThis *dbDataLocal) Del(ctx context.Context, daoModel *dao.DaoModel, idArr ...any) {
	for index := range idArr {
		cacheThis.cache().Delete(cacheThis.key(daoModel, idArr[index]))
	}
}

// ttlOrField是字符串类型时，确保是能从数据库查询结果中获得，且值必须是数字或时间类型
func (cacheThis *dbDataLocal) getOrSet(ctx context.Context, daoModel *dao.DaoModel, id any, ttlOrField any, field ...string) (value *gvar.Var, notExist bool, err error) {
	key := cacheThis.key(daoModel, id)
	if valueTmp, ok := cacheThis.cache().Get(key); ok { //先读一次（不加锁）
		value = valueTmp.(*gvar.Var)
		return
	}
	// 防止当前服务器并发
	getOrSetLocalMuTmp, _ := cacheThis.muMap.LoadOrStore(key, &sync.Mutex{})
	getOrSetLocalMu := getOrSetLocalMuTmp.(*sync.Mutex)
	getOrSetLocalMu.Lock()
	defer getOrSetLocalMu.Unlock()
	if valueTmp, ok := cacheThis.cache().Get(key); ok { // 再读一次（加锁），防止重复初始化
		value = valueTmp.(*gvar.Var)
		return
	}

	fieldArr := field
	ttlField, ok := ttlOrField.(string)
	isTtlField := ok && ttlField != ``
	if len(fieldArr) > 0 && isTtlField {
		fieldArr = append(fieldArr, ttlField)
	}
	info, err := daoModel.FilterPri(id).Fields(fieldArr...).One()
	if err != nil {
		return
	}
	if info.IsEmpty() {
		notExist = true
		return
	}
	var ttl int64
	if isTtlField {
		ttl = info[ttlField].GTime().Unix()
		if nowTime := gtime.Now().Unix(); ttl > nowTime {
			ttl = ttl - nowTime
		}
	} else {
		ttl = gconv.Int64(ttlOrField)
	}
	if ttl <= 0 || ttl > consts.CACHE_TIME_DEFAULT { //缓存时间不能超过默认缓存时间
		ttl = consts.CACHE_TIME_DEFAULT
	}
	if len(field) == 1 {
		value = gvar.New(info[field[0]].String())
	} else {
		value = gvar.New(info.Json())
	}
	cacheThis.cache().Set(key, value, time.Duration(ttl)*time.Second)
	return
}

func (cacheThis *dbDataLocal) GetOrSet(ctx context.Context, daoModel *dao.DaoModel, id any, ttlOrField any, field ...string) (value *gvar.Var, err error) {
	value, _, err = cacheThis.getOrSet(ctx, daoModel, id, ttlOrField, field...)
	return
}

func (cacheThis *dbDataLocal) GetOrSetMany(ctx context.Context, daoModel *dao.DaoModel, idArr []any, ttlOrField any, field ...string) (list gdb.Result, err error) {
	var value *gvar.Var
	var notExist bool
	for _, id := range idArr {
		value, notExist, err = cacheThis.getOrSet(ctx, daoModel.ResetNew(), id, ttlOrField, field...)
		if err != nil {
			return
		}
		if notExist { //缓存的是数据库数据，就需要和数据库SQL查询一样。故无数据时不返回
			continue
		}
		list = append(list, gdb.Record{})
		value.Scan(&list[len(list)-1])
	}
	return
}

func (cacheThis *dbDataLocal) GetOrSetPluck(ctx context.Context, daoModel *dao.DaoModel, idArr []any, ttlOrField any, field ...string) (record gdb.Record, err error) {
	var value *gvar.Var
	var notExist bool
	record = gdb.Record{}
	for _, id := range idArr {
		value, notExist, err = cacheThis.getOrSet(ctx, daoModel.ResetNew(), id, ttlOrField, field...)
		if err != nil {
			return
		}
		if notExist { //缓存的是数据库数据，就需要和数据库SQL查询一样。故无数据时不返回
			continue
		}
		record[gconv.String(id)] = value
	}
	return
}
