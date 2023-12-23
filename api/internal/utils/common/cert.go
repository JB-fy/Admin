package common

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

// 解析RSA私钥
func ParsePrivateKeyOfRSA(privateKeyStr string) (privateKey *rsa.PrivateKey, err error) {
	block, _ := pem.Decode([]byte(privateKeyStr))
	if block == nil {
		err = errors.New(`私钥解码错误`)
		return
	}

	switch block.Type {
	/* case `EC PRIVATE KEY`:
	privateKey, err = x509.ParseECPrivateKey(block.Bytes)
	return */
	case `RSA PRIVATE KEY`:
		privateKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
		return
	case `PRIVATE KEY`:
		priKey, errTmp := x509.ParsePKCS8PrivateKey(block.Bytes)
		if errTmp != nil {
			err = errors.New(`私钥解析错误：` + errTmp.Error())
			return
		}
		privateKeyTmp, ok := priKey.(*rsa.PrivateKey)
		if !ok {
			err = errors.New(`不是RSA私钥`)
			return
		}
		privateKey = privateKeyTmp
		return
	default:
		err = errors.New(`PEM类型错误`)
		return
	}
}

// 解析RSA公钥
func ParsePublicKeyOfRSA(publicKeyStr string) (publicKey *rsa.PublicKey, err error) {
	block, _ := pem.Decode([]byte(publicKeyStr))
	if block == nil {
		err = errors.New(`公钥解码错误`)
		return
	}

	switch block.Type {
	case `RSA PUBLIC KEY`:
		publicKey, err = x509.ParsePKCS1PublicKey(block.Bytes)
		return
	case `PUBLIC KEY`:
		pubKey, errTmp := x509.ParsePKIXPublicKey(block.Bytes)
		if errTmp != nil {
			err = errors.New(`公钥解析错误：` + errTmp.Error())
			return
		}
		publicKeyTmp, ok := pubKey.(*rsa.PublicKey)
		if !ok {
			err = errors.New(`不是RSA公钥`)
			return
		}
		publicKey = publicKeyTmp
		return
	default:
		err = errors.New(`PEM类型错误`)
		return
	}
}
