package api

type CommonNoDataRes struct {
}

type CommonCreateRes struct {
	Id int64 `json:"id" dc:"ID"`
}

type CommonEncryptStrRes struct {
	EncryptStr string `json:"encryptStr" dc:"加密盐"`
}

type CommonTokenRes struct {
	Token string `json:"token" dc:"登录授权token"`
}
