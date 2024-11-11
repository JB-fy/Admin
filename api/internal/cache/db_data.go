package cache

import (
	"api/internal/consts"
	"api/internal/dao"
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
)

type dbData struct {
	Ctx   context.Context
	Redis *gredis.Redis
	Dao   dao.DaoInterface
	// 以下三个参数只在调用GetOrSet方法时初始化
	DaoModel *dao.DaoModel
	Key      string
	Id       string
}

func NewDbData(ctx context.Context, dao dao.DaoInterface) *dbData {
	//可在这里写分库逻辑
	redis := g.Redis()
	chcheObj := &dbData{
		Ctx:   ctx,
		Redis: redis,
		Dao:   dao,
	}
	return chcheObj
}

var (
	dbDataMuMap sync.Map //存放所有缓存KEY的锁（当前服务器用）
	// （不同服务器竞争redis锁用）依据以下两个时间设置：运行速度 + 缓存写入到能读取之间的时间
	dbDataIsSetKeySuffix       = `_isSet`               //redis锁缓存Key后缀。与原来的缓存KEY拼接
	dbDataIsSetKeyTTL    int64 = 3                      //redis锁缓存Key时间
	dbDataNumLock              = 3                      //获取锁的重试次数。作用：当获取锁的服务器因报错，无法做缓存写入时，允许其它服务器重新获取锁，以保证缓存写入成功
	dbDataNumRead              = 5                      //读取缓存的重试次数。作用：当未获取锁的服务器获取缓存时数据为空时，可以重试多次
	dbDataOneTime              = 200 * time.Millisecond //读取缓存重试间隔时间
)

func (cacheThis *dbData) GetOrSet(id string, field ...string) (value *gvar.Var, noExistOfDb bool, err error) {
	cacheThis.DaoModel = cacheThis.Dao.CtxDaoModel(cacheThis.Ctx)
	cacheThis.Id = id
	cacheThis.Key = fmt.Sprintf(consts.CacheDbDataFormat, cacheThis.DaoModel.DbGroup, cacheThis.DaoModel.DbTable, cacheThis.Id)

	value, noExistOfCache, err := cacheThis.get()
	if err != nil {
		return
	}
	if !noExistOfCache {
		return
	}

	// 防止当前服务器并发
	dbDataMuTmp, _ := dbDataMuMap.LoadOrStore(cacheThis.Key, &sync.Mutex{})
	dbDataMu := dbDataMuTmp.(*sync.Mutex)
	dbDataMu.Lock()
	defer dbDataMu.Unlock()

	// 防止不同服务器并发
	isSetKey := cacheThis.Key + dbDataIsSetKeySuffix
	var isSet int64 = 0
	for i := 0; i < dbDataNumLock; i++ {
		isSet, err = cacheThis.Redis.Incr(cacheThis.Ctx, isSetKey)
		if err != nil {
			return
		}
		if isSet == 1 {
			value, noExistOfDb, err = cacheThis.set(field...)
			if noExistOfDb {
				cacheThis.Redis.Del(cacheThis.Ctx, isSetKey) //数据库不存在时，删除redis锁缓存Key，允许其它服务器重新尝试设置缓存
				return
			}
			if err != nil {
				cacheThis.Redis.Del(cacheThis.Ctx, isSetKey) //报错时，删除redis锁缓存Key，允许其它服务器重新尝试设置缓存
				return
			}
			cacheThis.Redis.Expire(cacheThis.Ctx, isSetKey, dbDataIsSetKeyTTL)
			return
		}

		// 等待读取数据
		for i := 0; i < dbDataNumRead; i++ {
			value, noExistOfCache, err = cacheThis.get()
			if err != nil {
				return
			}
			if !noExistOfCache {
				return
			}
			time.Sleep(dbDataOneTime)
		}
	}
	/*
		出现该错误的情况：
			1：所有服务器执行方法时都报错了。一般不大可能出现这种情况，概率极低
			2：redis服务负载过高，需要及时做优化了。解决方案：扩容或分库
	*/
	err = errors.New(`尝试多次查询缓存失败：` + cacheThis.Key)
	return
}

func (cacheThis *dbData) GetOrSetMany(idArr []string, field ...string) (list gdb.Result, err error) {
	for _, id := range idArr {
		value, noExistOfDb, errTmp := cacheThis.GetOrSet(id, field...)
		if errTmp != nil {
			err = errTmp
			return
		}
		if noExistOfDb { //既然是缓存数据库数据，就要根数据库一样，查不到数据，就没有数据
			continue
		}
		var info gdb.Record
		value.Scan(&info)
		list = append(list, info)
	}
	return
}

func (cacheThis *dbData) GetOrSetPluck(idArr []string, field ...string) (record gdb.Record, err error) {
	record = gdb.Record{}
	for _, id := range idArr {
		value, noExistOfDb, errTmp := cacheThis.GetOrSet(id, field...)
		if errTmp != nil {
			err = errTmp
			return
		}
		if noExistOfDb { //既然是缓存数据库数据，就要根数据库一样，查不到数据，就没有数据
			continue
		}
		record[id] = value
	}
	return
}

func (cacheThis *dbData) Del(idArr ...string) (row int64, err error) {
	daoModel := cacheThis.Dao.CtxDaoModel(cacheThis.Ctx)
	var keyArr []string
	for _, id := range idArr {
		keyArr = append(keyArr, fmt.Sprintf(consts.CacheDbDataFormat, daoModel.DbGroup, daoModel.DbTable, id))
	}
	row, err = cacheThis.Redis.Del(cacheThis.Ctx, keyArr...)
	return
}

func (cacheThis *dbData) set(field ...string) (value *gvar.Var, noExistOfDb bool, err error) {
	info, err := cacheThis.DaoModel.Filter(`id`, cacheThis.Id).Fields(field...).One()
	if err != nil {
		return
	}
	if info.IsEmpty() { //数据不存在，不缓存
		noExistOfDb = true
		return
	}
	if len(field) == 1 {
		value = info[field[0]]
	} else {
		value = gvar.New(info.Json())
	}
	_, err = cacheThis.Redis.Set(cacheThis.Ctx, cacheThis.Key, value.String())
	return
}

func (cacheThis *dbData) get() (value *gvar.Var, noExistOfCache bool, err error) {
	value, err = cacheThis.Redis.Get(cacheThis.Ctx, cacheThis.Key)
	if err != nil {
		return
	}
	if value.String() != `` {
		return
	}
	//为空时增加判断，数据库数据本身就是空字符串，但已缓存在数据库
	exists, err := cacheThis.Redis.Exists(cacheThis.Ctx, cacheThis.Key)
	if err != nil {
		return
	}
	if exists > 0 {
		return
	}
	noExistOfCache = true
	return
}
