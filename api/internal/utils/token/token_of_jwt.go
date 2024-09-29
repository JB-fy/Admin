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
	SignType   string `json:"sign_type"`
	SignKey    []byte `json:"sign_key"`
	ExpireTime uint   `json:"expire_time"`
	SignMethod jwt.SigningMethod
}

func NewTokenOfJwt(ctx context.Context, config map[string]any) *TokenOfJwt {
	tokenObj := TokenOfJwt{
		Ctx:        ctx,
		SignMethod: jwt.SigningMethodHS256,
	}
	gconv.Struct(config, &tokenObj)
	if tokenObj.SignKey == nil || tokenObj.SignType == `` || tokenObj.ExpireTime == 0 {
		panic(`缺少配置：token-Jwt`)
	}

	signMethodMap := map[string]jwt.SigningMethod{
		`HS256`: jwt.SigningMethodHS256,
		`HS384`: jwt.SigningMethodHS384,
		`HS512`: jwt.SigningMethodHS512,
		/* // SignKeyToParse []byte `json:"SignKeyOfParse"`
		`RS256`: jwt.SigningMethodRS256,
		`RS384`: jwt.SigningMethodRS384,
		`RS512`: jwt.SigningMethodRS512,
		`ES256`: jwt.SigningMethodES256,
		`ES384`: jwt.SigningMethodES384,
		`ES512`: jwt.SigningMethodES512, */
	}
	if signMethod, ok := signMethodMap[tokenObj.SignType]; ok {
		tokenObj.SignMethod = signMethod
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
