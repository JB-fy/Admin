package initialize

import (
	"context"

	"github.com/gogf/gf/v2/os/gtime"
)

func initGtime(ctx context.Context) {
	gtime.SetTimeZone(`Asia/Shanghai`)
}
