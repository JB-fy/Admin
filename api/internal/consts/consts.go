package consts

var (
	CacheSaltFormat  = "salt_%s_%s"  //加密盐缓存key。参数：场景标识，账号
	CacheTokenFormat = "token_%s_%s" //登录后的token缓存key。参数：场景标识，用户标识

	ConstCtxSceneInfoName = "sceneInfo"
	ConstCtxLoginInfoName = "loginInfo"
)
