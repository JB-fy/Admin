package cache

import (
	"api/internal/consts"
	"api/internal/dao"
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

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
	numLock = 3                      //获取锁的重试次数。作用：当获取锁的服务器因报错，无法做缓存写入时，允许其它服务器重新获取锁，以保证缓存写入成功
	numRead = 5                      //读取缓存的重试次数。作用：当未获取锁的服务器获取缓存时数据为空时，可以重试多次
	oneTime = 200 * time.Millisecond //读取缓存重试间隔时间
)

func (cacheThis *dbData) GetOrSet(id string, field ...string) (value string, err error) {
	cacheThis.DaoModel = cacheThis.Dao.CtxDaoModel(cacheThis.Ctx)
	cacheThis.Id = id
	cacheThis.Key = fmt.Sprintf(consts.CacheDbDataFormat, cacheThis.DaoModel.DbGroup, cacheThis.DaoModel.DbTable, cacheThis.Id)

	value, err = cacheThis.get()
	if err != nil {
		return
	}
	if value != `` {
		return
	}

	// 防止当前服务器并发
	dbDataMuTmp, _ := dbDataMuMap.LoadOrStore(cacheThis.Key, &sync.Mutex{})
	dbDataMu := dbDataMuTmp.(*sync.Mutex)
	dbDataMu.Lock()
	defer dbDataMu.Unlock()

	// 防止不同服务器并发
	isSetKey := cacheThis.Key + `_isSet`
	var isSet int64 = 0
	for i := 0; i < numLock; i++ {
		isSet, err = cacheThis.Redis.Incr(cacheThis.Ctx, isSetKey)
		if err != nil {
			return
		}

		if isSet == 1 {
			value, err = cacheThis.set(field...)
			if err != nil {
				cacheThis.Redis.Del(cacheThis.Ctx, isSetKey) //报错时删除可设置缓存KEY，允许其它服务器重新尝试设置缓存
				return
			}
			cacheThis.Redis.Expire(cacheThis.Ctx, isSetKey, 3)
			return
		}
		// 等待读取数据
		for i := 0; i < numRead; i++ {
			value, _ = cacheThis.get()
			if value != `` {
				return
			}
			time.Sleep(oneTime)
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

func (cacheThis *dbData) GetOrSetMany(idArr []string, field ...string) (valueArr []string, err error) {
	for _, id := range idArr {
		value, errTmp := cacheThis.GetOrSet(id, field...)
		if errTmp != nil {
			err = errTmp
			return
		}
		valueArr = append(valueArr, value)
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

func (cacheThis *dbData) set(field ...string) (value string, err error) {
	if len(field) == 1 {
		value, err = cacheThis.DaoModel.Filter(`id`, cacheThis.Id).ValueStr(field[0])
		if err != nil {
			return
		}
	} else {
		info, errTmp := cacheThis.DaoModel.Filter(`id`, cacheThis.Id).Fields(field...).One()
		if errTmp != nil {
			err = errTmp
			return
		}
		value = info.Json()
	}
	_, err = cacheThis.Redis.Set(cacheThis.Ctx, cacheThis.Key, value)
	return
}

func (cacheThis *dbData) get() (value string, err error) {
	valueTmp, err := cacheThis.Redis.Get(cacheThis.Ctx, cacheThis.Key)
	if err != nil {
		return
	}
	value = valueTmp.String()
	return
}
