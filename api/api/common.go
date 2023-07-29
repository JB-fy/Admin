package api

type CommonNoDataRes struct {
}

type CommonCreateRes struct {
	Id int64 `json:"id" dc:"ID"`
}

type CommonSaltRes struct {
	SaltStatic  string `json:"saltStatic" dc:"静态加密盐"`
	SaltDynamic string `json:"saltDynamic" dc:"动态加密盐"`
}

type CommonTokenRes struct {
	Token string `json:"token" dc:"登录授权token"`
}
