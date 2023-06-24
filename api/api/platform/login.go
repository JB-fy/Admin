package api

type AdminUpdateSelfReq struct {
	Account       *string `c:"account,omitempty" json:"account" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$"`
	Phone         *string `c:"phone,omitempty" json:"phone" v:"phone"`
	Nickname      *string `c:"nickname,omitempty" json:"nickname" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$"`
	Avatar        *string `c:"avatar,omitempty" json:"avatar" v:"url|length:1,120"`
	Password      *string `c:"password,omitempty" json:"password" v:"size:32|regex:^[\\p{L}\\p{N}_-]+$|different:CheckPassword"`
	CheckPassword *string `c:"checkPassword,omitempty" json:"checkPassword" v:"required-with:account,phone,password|size:32|regex:^[\\p{L}\\p{N}_-]+$"`
}
