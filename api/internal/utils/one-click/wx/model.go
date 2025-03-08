package wx

type AccessToken struct {
	Unionid        string `json:"unionid"`         //用户统一标识（全局唯一），只有当scope为"snsapi_userinfo"时返回
	Openid         string `json:"openid"`          //用户唯一标识（相对于公众号、开放平台下的应用唯一）
	AccessToken    string `json:"access_token"`    //网页授权接口调用凭证,注意：此access_token与基础支持的access_token不同
	RefreshToken   string `json:"refresh_token"`   //用户刷新access_token
	Scope          string `json:"scope"`           //用户授权的作用域，使用逗号（,）分隔
	ExpiresIn      int64  `json:"expires_in"`      //access_token 接口调用凭证超时时间，单位（秒）
	IsSnapshotuser uint8  `json:"is_snapshotuser"` //快照页模式虚拟账号：0否 1是。只有当用户是快照页模式虚拟账号时返回
}

type UserInfo struct {
	Unionid   string `json:"unionid"`    //用户统一标识（全局唯一）
	Openid    string `json:"openid"`     //用户唯一标识（相对于公众号、开放平台下的应用唯一）
	Nickname  string `json:"nickname"`   //昵称
	Avatar    string `json:"headimgurl"` //头像。最后一个数值代表正方形头像大小，有0、46、64、96、132数值可选，0代表640*640正方形头像
	Country   string `json:"country"`    //国家，如中国为CN
	Province  string `json:"province"`   //用户个人资料填写的省份
	City      string `json:"city"`       //用户个人资料填写的城市
	Privilege string `json:"privilege"`  //用户特权信息，json 数组，如微信沃卡用户为（chinaunicom）
	Gender    uint8  `json:"sex"`        //性别：0未知 1男 2女
}

type RefreshToken struct {
	Openid       string `json:"openid"`        //网页授权接口调用凭证,注意：此access_token与基础支持的access_token不同
	AccessToken  string `json:"access_token"`  //授权Token
	RefreshToken string `json:"refresh_token"` //用户刷新access_token
	Scope        string `json:"scope"`         //用户授权的作用域，使用逗号（,）分隔
	ExpiresIn    int64  `json:"expires_in"`    //access_token接口调用凭证超时时间，单位（秒）
}
