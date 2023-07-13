package initialize

// 环境变量设置。如：记录当前服务器IP
import (
	"api/internal/utils"

	daoPlatform "api/internal/dao/platform"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/genv"
)

func init() {
	ctx := gctx.New()

	/**--------记录当前服务器IP 开始--------**/
	serverNetworkIp := utils.GetServerNetworkIp()
	serverLocalIp := utils.GetServerLocalIp()
	genv.Set(`SERVER_NETWORK_IP`, serverNetworkIp) //设置服务器外网ip（key必须由大写和_组成，才能用g.Cfg().MustGetWithEnv()方法读取）
	genv.Set(`SERVER_LOCAL_IP`, serverLocalIp)     //设置服务器内网ip（key必须由大写和_组成，才能用g.Cfg().MustGetWithEnv()方法读取）

	daoPlatform.Server.ParseDbCtx(ctx).Data(g.Map{
		daoPlatform.Server.Columns().NetworkIp: serverNetworkIp,
		daoPlatform.Server.Columns().LocalIp:   serverLocalIp,
	}).Save()
	/**--------记录当前服务器IP 结束--------**/
}
