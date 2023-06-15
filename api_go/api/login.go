package api

type LoginEncryptReq struct {
	Account string `p:"account"  v:"required|length:4,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$"`
}

type LoginLoginReq struct {
	Account  string `p:"account"  v:"required|length:4,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$"`
	Password string `p:"password"  v:"required|size:32|regex:^[\\p{L}\\p{N}]+$"`
}
