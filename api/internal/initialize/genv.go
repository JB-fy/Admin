package initialize

import (
	"api/internal/consts"
	// daoPlatform "api/internal/dao/platform"
	"api/internal/utils"
	"context"

	"github.com/gogf/gf/v2/os/genv"
)

func initGenv(ctx context.Context) {
	/**--------记录当前服务器IP 开始--------**/
	serverNetworkIp := utils.GetServerNetworkIp()
	serverLocalIp := utils.GetServerLocalIp()
	genv.Set(consts.LOCAL_SERVER_NETWORK_IP, serverNetworkIp) //设置服务器外网ip
	genv.Set(consts.LOCAL_SERVER_LOCAL_IP, serverLocalIp)     //设置服务器内网ip

	/* daoPlatform.Server.CtxDaoModel(ctx).Data(g.Map{
		daoPlatform.Server.Columns().NetworkIp: serverNetworkIp,
		daoPlatform.Server.Columns().LocalIp:   serverLocalIp,
	}).OnConflict(daoPlatform.Server.Columns().NetworkIp).Save() */
	/**--------记录当前服务器IP 结束--------**/
}
