package consts

var (
	CacheSaltFormat          = `salt:%s:%s`          //密码盐缓存key。参数：场景标识，手机/邮箱/账号/邮箱
	CacheCodeFormat          = `code:%s:%s_%d`       //验证码缓存key。参数：场景标识，手机/邮箱，场景
	CacheTokenActiveFormat   = `tokenActive:%s:%s`   //判断Token失活缓存key。参数：场景标识，登录用户ID
	CacheTokenIsUniqueFormat = `tokenIsUnique:%s:%s` //判断Token唯一缓存key。参数：场景标识，登录用户ID

	ConstCtxSceneInfoName = `sceneInfo`
	ConstCtxLoginInfoName = `loginInfo`

	CacheWxGzhAccessToken = `wxGzhAccessToken:%s` //微信公众号授权Token缓存key。参数：微信公众号AppId
)

var (
	SERVER_NETWORK_IP = `SERVER_NETWORK_IP` //外网ip
	SERVER_LOCAL_IP   = `SERVER_LOCAL_IP`   //内网ip
)
