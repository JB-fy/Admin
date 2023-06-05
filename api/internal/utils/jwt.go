package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	SigningKey []byte
}

func NewJWT() *JWT {
	return &JWT{
		//[]byte(global.GVA_CONFIG.JWT.SigningKey),
	}
}

type CustomClaims struct {
	LoginId  uint   `json:"loginId"`
	Account  string `json:"account"`
	NickName string `json:"nickName"`
	jwt.RegisteredClaims
}

// 创建一个token
func (j *JWT) CreateToken(claims CustomClaims) (tokenString string, err error) {
	claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(2 * time.Hour)) // 过期时间2小时
	claims.IssuedAt = jwt.NewNumericDate(time.Now())                     // 签发时间
	claims.NotBefore = jwt.NewNumericDate(time.Now())                    // 生效时间
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(j.SigningKey)
	return
}

// 解析 token
func (j *JWT) ParseToken(tokenString string) (claims *CustomClaims, err error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return
	}
	if !token.Valid {
		err = errors.New("token无效")
		return
	}
	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		err = errors.New("claims无效")
		return
	}
	return
}
