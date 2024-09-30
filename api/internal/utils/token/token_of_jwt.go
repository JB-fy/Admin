package token

import (
	"api/internal/utils/common"
	"context"
	"errors"
	"time"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/golang-jwt/jwt/v5"
)

type TokenOfJwt struct {
	Ctx        context.Context
	SignType   string `json:"sign_type"`
	PrivateKey string `json:"private_key"`
	PublicKey  string `json:"public_key"`
	ExpireTime uint   `json:"expire_time"`
	SignMethod jwt.SigningMethod
}

func NewTokenOfJwt(ctx context.Context, config map[string]any) *TokenOfJwt {
	tokenObj := TokenOfJwt{
		Ctx:        ctx,
		SignMethod: jwt.SigningMethodHS256,
	}
	gconv.Struct(config, &tokenObj)
	if tokenObj.SignType == `` || tokenObj.PrivateKey == `` || tokenObj.ExpireTime == 0 || (tokenObj.PublicKey == `` && garray.NewStrArrayFrom([]string{`RS256`, `RS384`, `RS512`}).Contains(tokenObj.SignType)) {
		panic(`缺少配置：token-Jwt`)
	}

	signMethodMap := map[string]jwt.SigningMethod{
		`HS256`: jwt.SigningMethodHS256,
		`HS384`: jwt.SigningMethodHS384,
		`HS512`: jwt.SigningMethodHS512,
		`RS256`: jwt.SigningMethodRS256,
		`RS384`: jwt.SigningMethodRS384,
		`RS512`: jwt.SigningMethodRS512,
		/* `ES256`: jwt.SigningMethodES256,
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
	privateKeyFunc := func() (privateKey any) {
		switch tokenThis.SignType {
		case `RS256`, `RS384`, `RS512`:
			privateKey, _ = common.ParsePrivateKeyOfRSA(tokenThis.PrivateKey)
		/* case `ES256`, `ES384`, `ES512`:
		privateKey, _ = common.ParsePrivateKeyOfRSA(tokenThis.PrivateKey) */
		// case `HS256`, `HS384`, `HS512`:
		default:
			privateKey = []byte(tokenThis.PrivateKey)
		}
		return
	}
	token, err = jwt.NewWithClaims(tokenThis.SignMethod, jwt.RegisteredClaims{
		ID:        tokenInfo.LoginId,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(tokenThis.ExpireTime) * time.Second)), // 过期时间
		IssuedAt:  jwt.NewNumericDate(time.Now()),                                                        // 签发时间
		NotBefore: jwt.NewNumericDate(time.Now()),                                                        // 生效时间
	}).SignedString(privateKeyFunc())
	return
}

func (tokenThis *TokenOfJwt) Parse(token string) (tokenInfo TokenInfo, err error) {
	jwtToken, err := jwt.ParseWithClaims(token, &jwt.RegisteredClaims{}, func(token *jwt.Token) (any, error) {
		switch tokenThis.SignType {
		case `RS256`, `RS384`, `RS512`:
			return common.ParsePublicKeyOfRSA(tokenThis.PublicKey)
		/* case `ES256`, `ES384`, `ES512`:
		return common.ParsePublicKeyOfRSA(tokenThis.PublicKey) */
		// case `HS256`, `HS384`, `HS512`:
		default:
			return []byte(tokenThis.PrivateKey), nil
		}
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
