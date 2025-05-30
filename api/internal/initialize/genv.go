package initialize

import (
	"api/internal/consts"
	// daoPlatform "api/internal/dao/platform"
	"api/internal/utils"
	"context"
	"os"

	"github.com/gogf/gf/v2/os/genv"
)

func initGenv(ctx context.Context) {
	/**--------记录当前服务器IP 开始--------**/
	serverNetworkIp := os.Getenv(consts.ENV_SERVER_NETWORK_IP) //docker容器设置环境变量启动时有值：-e SERVER_NETWORK_IP=$(curl -s --max-time 3 ifconfig.me || curl -s --max-time 3 https://ipinfo.io/ip || curl -s --max-time 3 https://checkip.amazonaws.com || curl -s --max-time 3 https://icanhazip.com || curl -s --max-time 3 https://api.ipify.org)
	if serverNetworkIp == `` {
		serverNetworkIp := utils.GetServerNetworkIp()
		genv.Set(consts.ENV_SERVER_NETWORK_IP, serverNetworkIp) //设置服务器外网ip
	}
	serverLocalIp := os.Getenv(consts.ENV_SERVER_LOCAL_IP) //docker容器设置环境变量启动时有值：-e SERVER_LOCAL_IP=$(hostname -I | awk '{printf "%s", $1}')
	if serverLocalIp == `` {
		serverLocalIp := utils.GetServerLocalIp()
		genv.Set(consts.ENV_SERVER_LOCAL_IP, serverLocalIp) //设置服务器内网ip
	}

	/* daoPlatform.Server.CtxDaoModel(ctx).Data(g.Map{
		daoPlatform.Server.Columns().NetworkIp: serverNetworkIp,
		daoPlatform.Server.Columns().LocalIp:   serverLocalIp,
	}).OnConflict(daoPlatform.Server.Columns().NetworkIp).Save() */
	/**--------记录当前服务器IP 结束--------**/
}
