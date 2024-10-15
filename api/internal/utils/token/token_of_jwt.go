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

/*
密钥生成方式：

	HS密钥：
		任意字符串即可
	RS密钥：
		私钥命令：openssl genrsa -out rsa-private-key.pem 2048
		公钥命令：openssl rsa -in rsa-private-key.pem -pubout -out rsa-public-key.pem
	ES密钥：
		私钥命令：openssl ecparam -genkey -name prime256r1|secp384r1|secp521r1 -noout -out ecc-private-key.pem
		公钥命令：openssl ec -in ecc-private-key.pem -pubout -out ecc-public-key.pem
*/
type TokenOfJwt struct {
	Ctx        context.Context
	ExpireTime uint   `json:"expire_time"`
	SignType   string `json:"sign_type"`
	PrivateKey string `json:"private_key"`
	PublicKey  string `json:"public_key"`
	SignMethod jwt.SigningMethod
}

func NewTokenOfJwt(ctx context.Context, config map[string]any) *TokenOfJwt {
	tokenObj := TokenOfJwt{
		Ctx:        ctx,
		SignMethod: jwt.SigningMethodHS256,
	}
	gconv.Struct(config, &tokenObj)
	if tokenObj.ExpireTime == 0 || tokenObj.SignType == `` || tokenObj.PrivateKey == `` || (tokenObj.PublicKey == `` && garray.NewStrArrayFrom([]string{`RS256`, `RS384`, `RS512`}).Contains(tokenObj.SignType)) {
		panic(`缺少配置：token-Jwt`)
	}

	signMethodMap := map[string]jwt.SigningMethod{
		`HS256`: jwt.SigningMethodHS256,
		`HS384`: jwt.SigningMethodHS384,
		`HS512`: jwt.SigningMethodHS512,
		`RS256`: jwt.SigningMethodRS256,
		`RS384`: jwt.SigningMethodRS384,
		`RS512`: jwt.SigningMethodRS512,
		`ES256`: jwt.SigningMethodES256,
		`ES384`: jwt.SigningMethodES384,
		`ES512`: jwt.SigningMethodES512,
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
		switch tokenThis.SignMethod {
		case jwt.SigningMethodHS256, jwt.SigningMethodHS384, jwt.SigningMethodHS512:
			privateKey = []byte(tokenThis.PrivateKey)
		default:
			privateKey, _ = common.ParsePrivateKey(tokenThis.PrivateKey)
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
	jwtToken, err := jwt.ParseWithClaims(token, &jwt.RegisteredClaims{}, func(jwtToken *jwt.Token) (any, error) {
		switch jwtToken.Method {
		case jwt.SigningMethodHS256, jwt.SigningMethodHS384, jwt.SigningMethodHS512:
			return []byte(tokenThis.PrivateKey), nil
		default:
			return common.ParsePublicKey(tokenThis.PublicKey)
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