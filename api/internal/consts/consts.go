package consts

var (
	CacheSaltFormat  = `salt_%s_%s`  //加密盐缓存key。参数：场景标识，账号/手机
	CacheTokenFormat = `token_%s_%d` //登录后的token缓存key。参数：场景标识，用户ID

	ConstCtxSceneInfoName = `sceneInfo`
	ConstCtxLoginInfoName = `loginInfo`
)
