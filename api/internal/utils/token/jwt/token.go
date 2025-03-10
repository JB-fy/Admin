package jwt

import (
	"api/internal/utils"
	"api/internal/utils/token/model"
	"context"
	"errors"
	"slices"
	"time"

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
type Token struct {
	ExpireTime int64  `json:"expire_time"`
	SignType   string `json:"sign_type"`
	PrivateKey string `json:"private_key"`
	PublicKey  string `json:"public_key"`
	SignMethod jwt.SigningMethod
}

func NewToken(ctx context.Context, config map[string]any) model.Token {
	obj := &Token{}
	gconv.Struct(config, obj)
	if obj.ExpireTime == 0 || obj.SignType == `` || obj.PrivateKey == `` || (obj.PublicKey == `` && !slices.Contains([]string{`HS256`, `HS384`, `HS512`}, obj.SignType)) {
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
	ok := false
	obj.SignMethod, ok = signMethodMap[obj.SignType]
	if !ok {
		obj.SignMethod = jwt.SigningMethodHS256
	}
	return obj
}

func (tokenThis *Token) Create(ctx context.Context, tokenInfo model.TokenInfo) (token string, err error) {
	privateKeyFunc := func() (privateKey any) {
		switch tokenThis.SignMethod {
		case jwt.SigningMethodHS256, jwt.SigningMethodHS384, jwt.SigningMethodHS512:
			privateKey = []byte(tokenThis.PrivateKey)
		default:
			privateKey, _ = utils.ParsePrivateKey(tokenThis.PrivateKey)
		}
		return
	}

	claims := tokenClaims{}
	claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Duration(tokenThis.ExpireTime) * time.Second)) // 过期时间
	claims.IssuedAt = jwt.NewNumericDate(time.Now())                                                         // 签发时间
	claims.NotBefore = jwt.NewNumericDate(time.Now())                                                        // 生效时间
	claims.ID = tokenInfo.LoginId
	claims.IP = tokenInfo.IP
	token, err = jwt.NewWithClaims(tokenThis.SignMethod, claims).SignedString(privateKeyFunc())
	return
}

func (tokenThis *Token) Parse(ctx context.Context, token string) (tokenInfo model.TokenInfo, err error) {
	jwtToken, err := jwt.ParseWithClaims(token, &tokenClaims{}, func(jwtToken *jwt.Token) (any, error) {
		switch jwtToken.Method {
		case jwt.SigningMethodHS256, jwt.SigningMethodHS384, jwt.SigningMethodHS512:
			return []byte(tokenThis.PrivateKey), nil
		default:
			return utils.ParsePublicKey(tokenThis.PublicKey)
		}
	})
	if err != nil {
		return
	}
	if !jwtToken.Valid {
		err = errors.New(`token无效`)
		return
	}
	claims, ok := jwtToken.Claims.(*tokenClaims)
	if !ok {
		err = errors.New(`token载体类型错误`)
		return
	}

	tokenInfo.LoginId = claims.ID
	tokenInfo.IP = claims.IP
	return
}

func (tokenThis *Token) GetExpireTime() int64 {
	return tokenThis.ExpireTime
}
