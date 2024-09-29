package token

import (
	"context"
	"errors"
	"time"

	"github.com/gogf/gf/v2/util/gconv"
	"github.com/golang-jwt/jwt/v5"
)

type TokenOfJwt struct {
	Ctx        context.Context
	SignKey    []byte `json:"signKey"`
	SignType   string `json:"signType"`
	SignMethod *jwt.SigningMethodHMAC
	ExpireTime uint `json:"expireTime"`
}

func NewTokenOfJwt(ctx context.Context, config map[string]any) *TokenOfJwt {
	tokenObj := TokenOfJwt{Ctx: ctx}
	gconv.Struct(config, &tokenObj)
	switch tokenObj.SignType {
	case `HS256`:
		tokenObj.SignMethod = jwt.SigningMethodHS256
	case `HS384`:
		tokenObj.SignMethod = jwt.SigningMethodHS384
	case `HS512`:
		tokenObj.SignMethod = jwt.SigningMethodHS512
	default:
		tokenObj.SignMethod = jwt.SigningMethodHS256
	}
	return &tokenObj
}

/* type CustomClaims struct {
	jwt.RegisteredClaims
	LoginId uint `json:"login_id"`
} */

func (tokenThis *TokenOfJwt) Create(tokenInfo TokenInfo) (token string, err error) {
	claims := jwt.RegisteredClaims{
		ID:        tokenInfo.LoginId,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(tokenThis.ExpireTime) * time.Second)), // 过期时间
		IssuedAt:  jwt.NewNumericDate(time.Now()),                                                        // 签发时间
		NotBefore: jwt.NewNumericDate(time.Now()),                                                        // 生效时间
	}
	token, err = jwt.NewWithClaims(tokenThis.SignMethod, claims).SignedString(tokenThis.SignKey)
	return
}

func (tokenThis *TokenOfJwt) Parse(token string) (tokenInfo TokenInfo, err error) {
	jwtToken, err := jwt.ParseWithClaims(token, &jwt.RegisteredClaims{}, func(token *jwt.Token) (any, error) {
		return tokenThis.SignKey, nil
	})
	if err != nil {
		return
	}
	if !jwtToken.Valid {
		err = errors.New(`token无效`)
		return
	}
	claims, ok := jwtToken.Claims.(*jwt.RegisteredClaims)
	if !ok {
		err = errors.New(`token载体类型错误`)
		return
	}
	tokenInfo.LoginId = claims.ID
	return
}

func (tokenThis *TokenOfJwt) GetExpireTime() (expireTime int64) {
	return int64(tokenThis.ExpireTime)
}
