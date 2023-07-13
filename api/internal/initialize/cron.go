package initialize

// 定时器设置
import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
)

func init() {
	ctx := gctx.New()

	myCronThis := myCron{}
	gcron.AddSingleton(ctx, `*/5 * * * * *`, myCronThis.Test, `Test`)
}

type myCron struct{}

func (myCron) Test(ctx context.Context) {
	fmt.Println(gtime.Now().Format(`Y-m-d H:i:s`))
}
