package common

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

// 解析RSA私钥
func ParsePrivateKeyOfRSA(privateKeyStr string) (privateKey *rsa.PrivateKey, err error) {
	/* if gstr.Pos(privateKeyStr, `-----`) != 0 {
		// privateKeyStr = "-----BEGIN RSA PRIVATE KEY-----\n" + privateKeyStr + "\n-----END RSA PRIVATE KEY-----"
		privateKeyStrTmp := privateKeyStr
		privateKeyStr = "-----BEGIN RSA PRIVATE KEY-----\n"
		for i := 0; i < len(privateKeyStrTmp); i += 64 {
			end := i + 64
			if end > len(privateKeyStrTmp) {
				end = len(privateKeyStrTmp)
			}
			privateKeyStr += privateKeyStrTmp[i:end] + "\n"
		}
		privateKeyStr += "-----END RSA PRIVATE KEY-----"
	} */
	block, _ := pem.Decode([]byte(privateKeyStr))
	if block == nil {
		err = errors.New(`解析私钥错误`)
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
			err = errTmp
			return
		}
		privateKeyTmp, ok := priKey.(*rsa.PrivateKey)
		if !ok {
			err = errors.New(`非RSA私钥`)
			return
		}
		privateKey = privateKeyTmp
		return
	default:
		err = errors.New(`不是RSA私钥`)
		return
	}
}

// 解析RSA公钥
func ParsePublicKeyOfRSA(publicKeyStr string) (publicKey *rsa.PublicKey, err error) {
	/* if gstr.Pos(publicKeyF, `-----`) != 0 {
		publicKeyStr = "-----BEGIN PUBLIC KEY-----\n" + publicKeyStr + "\n-----END PUBLIC KEY-----"
	} */
	block, _ := pem.Decode([]byte(publicKeyStr))
	if block == nil {
		err = errors.New(`解析公钥错误`)
		return
	}

	switch block.Type {
	case `RSA PUBLIC KEY`:
		publicKey, err = x509.ParsePKCS1PublicKey(block.Bytes)
		return
	case `PUBLIC KEY`:
		pubKey, errTmp := x509.ParsePKIXPublicKey(block.Bytes)
		if errTmp != nil {
			err = errTmp
			return
		}
		publicKeyTmp, ok := pubKey.(*rsa.PublicKey)
		if !ok {
			err = errors.New(`非RSA公钥`)
			return
		}
		publicKey = publicKeyTmp
		return
	default:
		err = errors.New(`不是RSA公钥`)
		return
	}
}
