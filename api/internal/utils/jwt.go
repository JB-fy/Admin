package utils

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/util/gconv"
	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	Ctx        context.Context
	SignKey    []byte `json:"signKey"`
	ExpireTime uint   `json:"expireTime"`
	SignType   string `json:"signType"`
	SignMethod *jwt.SigningMethodHMAC
}

func NewJWT(ctx context.Context, config map[string]interface{}) *JWT {
	jwtObj := JWT{
		Ctx: ctx,
	}
	gconv.Struct(config, &jwtObj)
	switch jwtObj.SignType {
	case `HS256`:
		jwtObj.SignMethod = jwt.SigningMethodHS256
	case `HS384`:
		jwtObj.SignMethod = jwt.SigningMethodHS384
	case `HS512`:
		jwtObj.SignMethod = jwt.SigningMethodHS512
	default:
		jwtObj.SignMethod = jwt.SigningMethodHS256
	}
	return &jwtObj
}

type CustomClaims struct {
	jwt.RegisteredClaims
	LoginId uint `json:"loginId"`
	// Nickname string `json:"nickname"`
}

// 创建一个token
func (jwtThis *JWT) CreateToken(claims CustomClaims) (tokenString string, err error) {
	claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Duration(jwtThis.ExpireTime) * time.Second)) // 过期时间
	claims.IssuedAt = jwt.NewNumericDate(time.Now())                                                       // 签发时间
	claims.NotBefore = jwt.NewNumericDate(time.Now())                                                      // 生效时间
	token := jwt.NewWithClaims(jwtThis.SignMethod, claims)
	tokenString, err = token.SignedString(jwtThis.SignKey)
	return
}

// 解析 token
func (jwtThis *JWT) ParseToken(tokenString string) (claims *CustomClaims, err error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtThis.SignKey, nil
	})
	if err != nil {
		err = NewErrorCode(jwtThis.Ctx, 39994001, err.Error())
		return
	}
	if !token.Valid {
		err = NewErrorCode(jwtThis.Ctx, 39994001, ``)
		return
	}
	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		err = NewErrorCode(jwtThis.Ctx, 39994001, ``)
		return
	}
	return
}
