package consts

var (
	CacheSmsFormat   = `sms_%s_%s_%d` //短信缓存key。参数：场景标识，手机，使用场景
	CacheSaltFormat  = `salt_%s_%s`   //加密盐缓存key。参数：场景标识，账号/手机
	CacheTokenFormat = `token_%s_%d`  //登录后的token缓存key。参数：场景标识，用户ID

	ConstCtxSceneInfoName = `sceneInfo`
	ConstCtxLoginInfoName = `loginInfo`
)
