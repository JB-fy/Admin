package initialize

import (
	"api/internal/consts"
	"api/internal/utils"
	"context"
	"fmt"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

func initCron(ctx context.Context) {
	if !utils.IsDev(ctx) && g.Cfg().MustGet(ctx, `cronServerNetworkIp`).String() != g.Cfg().MustGetWithEnv(ctx, consts.SERVER_NETWORK_IP).String() {
		return
	}

	// myCronThis := myCron{}
	// gcron.AddSingleton(ctx, `*/5 * * * * *`, myCronThis.Test, `Test`)
}

type myCron struct{}

func (myCron) Test(ctx context.Context) {
	fmt.Println(gtime.Now().Format(`Y-m-d H:i:s`))
}
