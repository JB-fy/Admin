package initialize

import (
	"api/internal/cache"
	"api/internal/consts"
	"api/internal/utils"
	"context"
	"fmt"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/genv"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

func initCron(ctx context.Context) {
	// myCronThis := myCron{}
	/*--------数据库中某些配置表极少修改，统一缓存在本机内存中，能极大增加服务器性能，减少数据库压力 开始--------*/
	gcron.AddSingleton(ctx, `0 */`+gconv.String(consts.CACHE_LOCAL_INTERVAL_MINUTE/time.Minute)+` * * * *`, cache.DbDataLocal.Flush, `DbDataLocalFlush`) //全部服务器使用定时器删除缓存，不同服务器定时器执行存在时间差，也可能出现不一致的情况（但时间差很小，一般可以接受）
	/*--------数据库中某些配置表极少修改，统一缓存在本机内存中，能极大增加服务器性能，减少数据库压力 结束--------*/

	// 只有指定IP的服务器才继续执行，否则会造成重复处理
	if !(utils.IsDev(ctx) || g.Cfg().MustGet(ctx, `masterServerNetworkIpArr.0`).String() == genv.Get(consts.ENV_SERVER_NETWORK_IP).String()) {
		return
	}

	// gcron.AddSingleton(ctx, `*/5 * * * * *`, myCronThis.Test, `Test`)
}

type myCron struct{}

func (myCron) Test(ctx context.Context) {
	fmt.Println(gtime.Now().Format(`Y-m-d H:i:s`))
}
