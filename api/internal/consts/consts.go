package consts

const (
	CTX_SCENE_INFO_NAME = `sceneInfo`
	CTX_LOGIN_INFO_NAME = `loginInfo`

	CACHE_TIME_DEFAULT int64 = 7 * 24 * 60 * 60 //默认缓存时间

	CACHE_DB_DATA         = `dbData:%s:%s:%v`     //数据库数据缓存key。参数：db分组，db表名，ID
	CACHE_SALT            = `salt:%s:%s`          //密码盐缓存key。参数：场景ID，手机/邮箱/账号/邮箱
	CACHE_CODE            = `code:%s:%s_%d`       //验证码缓存key。参数：场景ID，手机/邮箱，场景
	CACHE_TOKEN_ACTIVE    = `tokenActive:%s:%s`   //判断Token失活缓存key。参数：场景ID，登录用户ID
	CACHE_TOKEN_IS_UNIQUE = `tokenIsUnique:%s:%s` //判断Token唯一缓存key。参数：场景ID，登录用户ID

	CACHE_WX_GZH_ACCESS_TOKEN = `wxGzhAccessToken:%s` //微信公众号授权Token缓存key。参数：微信公众号AppId

	LOCAL_DB_DATA           = `DB_DATA_%s_%s_%v`  //数据库数据缓存key。参数：db分组，db表名，ID或其它唯一标识
	LOCAL_SERVER_NETWORK_IP = `SERVER_NETWORK_IP` //外网ip
	LOCAL_SERVER_LOCAL_IP   = `SERVER_LOCAL_IP`   //内网ip
)
