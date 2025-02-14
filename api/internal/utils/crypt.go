package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/x509"
	"encoding/pem"
	"errors"

	"github.com/gogf/gf/v2/util/gconv"
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

// PKCS补码（PKCS5，PKCS7通用）
func PKCS5Pad(rawByte []byte, padLen int) []byte {
	rawByteLen := len(rawByte)
	fillLen := padLen - (rawByteLen % padLen)
	fillByte := bytes.Repeat([]byte{byte(fillLen)}, fillLen)
	return append(rawByte, fillByte...)
}

// PKCS补码移除（PKCS5，PKCS7通用）
func PKCS5UnPad(rawByte []byte, padLen int) ([]byte, error) {
	rawByteLen := len(rawByte)
	fillByte := rawByte[rawByteLen-1]
	fillLen := int(fillByte)
	if fillLen > padLen || fillLen == 0 {
		return nil, errors.New(`无效的填充长度：` + gconv.String(fillLen))
	}
	fillPosition := rawByteLen - fillLen
	for _, v := range rawByte[fillPosition:] {
		if v != fillByte {
			return nil, errors.New(`无效的填充位：预期` + gconv.String(fillLen) + `实际` + gconv.String(v))
		}
	}
	return rawByte[:fillPosition], nil
}

// AES加密
func AesEncrypt(rawByte []byte, keyByte []byte, cipherType string, iv ...byte) (cipherByte []byte, err error) {
	block, err := aes.NewCipher(keyByte)
	if err != nil {
		return
	}
	blockSize := block.BlockSize()
	rawByteLen := len(rawByte)
	if rawByteLen%blockSize != 0 {
		err = errors.New(`加密串必须是块大小的整数倍`)
		return
	}
	cipherByte = make([]byte, rawByteLen)
	switch cipherType {
	case `ECB`:
		for i := 0; i < rawByteLen; i += blockSize {
			block.Encrypt(cipherByte[i:i+blockSize], rawByte[i:i+blockSize])
		}
	// case `CBC`:
	default:
		blockMode := cipher.NewCBCEncrypter(block, iv)
		blockMode.CryptBlocks(cipherByte, rawByte)
	}
	return
}

// AES解密
func AesDecrypt(cipherByte []byte, keyByte []byte, cipherType string, iv ...byte) (rawByte []byte, err error) {
	block, err := aes.NewCipher(keyByte)
	if err != nil {
		return
	}
	blockSize := block.BlockSize()
	cipherByteLen := len(cipherByte)
	if cipherByteLen%blockSize != 0 {
		err = errors.New(`解密串必须是块大小的整数倍`)
		return
	}
	rawByte = make([]byte, cipherByteLen)
	switch cipherType {
	case `ECB`:
		for i := 0; i < cipherByteLen; i += blockSize {
			block.Decrypt(rawByte[i:i+blockSize], cipherByte[i:i+blockSize])
		}
	// case `CBC`:
	default:
		blockMode := cipher.NewCBCDecrypter(block, iv)
		blockMode.CryptBlocks(rawByte, cipherByte)
	}
	return
}
