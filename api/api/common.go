package api

type CommonNoDataRes struct {
}

type CommonCreateRes struct {
	Id int64 `json:"id" dc:"ID"`
}

type CommonSaltRes struct {
	SaltStatic  string `json:"salt_static" dc:"静态加密盐"`
	SaltDynamic string `json:"salt_dynamic" dc:"动态加密盐"`
}

type CommonTokenRes struct {
	Token string `json:"token" dc:"登录授权token"`
}
