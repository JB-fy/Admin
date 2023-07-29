package api

type CommonNoDataRes struct {
}

type CommonCreateRes struct {
	Id int64 `json:"id" dc:"ID"`
}

type CommonSaltRes struct {
	SaltStatic  string `json:"saltStatic" dc:"加密盐（静态）"`
	SaltDynamic string `json:"saltDynamic" dc:"加密盐（动态）"`
}

type CommonTokenRes struct {
	Token string `json:"token" dc:"登录授权token"`
}
