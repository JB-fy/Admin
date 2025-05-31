package initialize

import (
	"api/internal/cache"
	"api/internal/consts"
	"api/internal/utils"
	"context"
	"fmt"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/genv"
	"github.com/gogf/gf/v2/os/gtime"
)

func initCron(ctx context.Context) {
	// myCronThis := myCron{}
	/*--------数据库中某些配置表极少修改，统一缓存在本机内存中，能极大增加服务器性能，减少数据库压力 开始--------*/
	gcron.AddSingleton(ctx, `0 0 4 * * *`, cache.DbDataLocal.Flush, `DbDataLocalFlush`) //每天晚上4点清空内存缓存
	/*--------数据库中某些配置表极少修改，统一缓存在本机内存中，能极大增加服务器性能，减少数据库压力 结束--------*/

	// 部分定时任务不允许全部服务器都开启，只有指定IP的服务器才能开启。比如任务存在数据库先读后改的逻辑时，多服务器同时开启任务，会存在重复处理的问题
	if !utils.IsDev(ctx) && g.Cfg().MustGet(ctx, `cronServerNetworkIp`).String() != genv.Get(consts.ENV_SERVER_NETWORK_IP).String() {
		return
	}

	// gcron.AddSingleton(ctx, `*/5 * * * * *`, myCronThis.Test, `Test`)
}

type myCron struct{}

func (myCron) Test(ctx context.Context) {
	fmt.Println(gtime.Now().Format(`Y-m-d H:i:s`))
}
