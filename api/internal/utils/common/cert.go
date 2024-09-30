package common

import (
	"crypto/x509"
	"encoding/pem"
	"errors"
)

// 解析私钥
func ParsePrivateKey(privateKeyStr string) (privateKey any /* *rsa.PrivateKey	*ecdsa.PrivateKey */, err error) {
	block, _ := pem.Decode([]byte(privateKeyStr))
	if block == nil {
		err = errors.New(`私钥解码错误`)
		return
	}

	switch block.Type {
	case `EC PRIVATE KEY`:
		privateKey, err = x509.ParseECPrivateKey(block.Bytes)
	case `RSA PRIVATE KEY`:
		privateKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
	case `PRIVATE KEY`:
		privateKey, err = x509.ParsePKCS8PrivateKey(block.Bytes)
	default:
		err = errors.New(`PEM类型错误`)
	}
	return
}

// 解析公钥
func ParsePublicKey(publicKeyStr string) (publicKey any /* *rsa.PublicKey	*ecdsa.PublicKey */, err error) {
	block, _ := pem.Decode([]byte(publicKeyStr))
	if block == nil {
		err = errors.New(`公钥解码错误`)
		return
	}

	switch block.Type {
	case `RSA PUBLIC KEY`:
		publicKey, err = x509.ParsePKCS1PublicKey(block.Bytes)
	case `PUBLIC KEY`:
		publicKey, err = x509.ParsePKIXPublicKey(block.Bytes)
	default:
		err = errors.New(`PEM类型错误`)
	}
	return
}
