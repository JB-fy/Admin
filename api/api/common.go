package api

type CommonHeaderReq struct {
	Language string `json:"-" v:"in:zh-cn,en" in:"header" d:"zh-cn" dc:"多语言标识"`
}

type CommonAllTokenHeaderReq struct {
	CommonHeaderReq
	PlatformToken string `json:"-" v:"" in:"header" d:"" dc:"平台后台登录token。有些接口（以对应场景开头：/场景/Xxxx）同时用于多个场景时，可能需要传对应场景的登录token"`
	OrgToken      string `json:"-" v:"" in:"header" d:"" dc:"机构后台登录token。有些接口（以对应场景开头：/场景/Xxxx）同时用于多个场景时，可能需要传对应场景的登录token"`
	AppToken      string `json:"-" v:"" in:"header" d:"" dc:"APP登录token。有些接口（以对应场景开头：/场景/Xxxx）同时用于多个场景时，可能需要传对应场景的登录token"`
}

type CommonPlatformHeaderReq struct {
	CommonHeaderReq
	PlatformToken string `json:"-" v:"" in:"header" d:"" dc:"登录token"`
}

type CommonOrgHeaderReq struct {
	CommonHeaderReq
	OrgToken string `json:"-" v:"" in:"header" d:"" dc:"登录token"`
}

type CommonAppHeaderReq struct {
	CommonHeaderReq
	AppToken string `json:"-" v:"" in:"header" d:"" dc:"登录token"`
}

type CommonInfoReq struct {
	Field []string `json:"field" v:"distinct|foreach|min-length:1" dc:"查询字段，传值参考返回的字段名，默认返回常用字段，如果所需字段较少或需特别字段时，可使用。特别注意：所需字段较少时使用，可大幅减轻数据库压力"`
}

type CommonListReq struct {
	CommonInfoReq
	Sort  string `json:"sort" default:"id DESC" dc:"排序"`
	Page  int    `json:"page" v:"min:1" default:"1" dc:"页码"`
	Limit int    `json:"limit" v:"min:0" default:"10" dc:"每页数量。可传0取全部"`
}

type CommonSignReq struct {
	Ts   uint   `json:"ts" v:"" dc:"时间戳。单位：秒"`
	Sign string `json:"sign" v:"" dc:"签名。所有参数按ASCII码排序后拼接（值是对象数组等复杂数据时，要转json格式拼接），再拼接密钥key，最后md5得到签名。公式：md5(参数1=值&参数2=值&....&key=密钥)"`
}

type CommonNoDataRes struct{}

type CommonCreateRes struct {
	Id any `json:"id" dc:"ID"`
}

type CommonSaltRes struct {
	SaltStatic  string `json:"salt_static" dc:"静态密码盐"`
	SaltDynamic string `json:"salt_dynamic" dc:"动态密码盐"`
}

type CommonTokenRes struct {
	Token string `json:"token" dc:"登录token"`
}
