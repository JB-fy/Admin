package consts

var (
	CacheSaltFormat  = `salt_%s_%s`   //密码盐缓存key。参数：场景标识，手机/邮箱/账号/邮箱
	CacheCodeFormat  = `sms_%s_%s_%d` //验证码缓存key。参数：场景标识，手机/邮箱，场景
	CacheTokenFormat = `token_%s_%d`  //登录后的token缓存key。参数：场景标识，用户ID

	ConstCtxSceneInfoName = `sceneInfo`
	ConstCtxLoginInfoName = `loginInfo`

	CacheWxGzhAccessToken = `wxGzhAccessToken_%s` //微信公众号授权Token缓存key。参数：微信公众号AppId
)
