package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/x509"
	"encoding/base64"
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

// PKCS补码（PKCS7，PKCS5通用）
func PKCSPad(rawByte []byte, padLen int) []byte {
	rawByteLen := len(rawByte)
	fillLen := padLen - (rawByteLen % padLen)
	fillByte := bytes.Repeat([]byte{byte(fillLen)}, fillLen)
	return append(rawByte, fillByte...)
}

// PKCS补码移除（PKCS7，PKCS5通用）
func PKCSUnPad(rawByte []byte, padLen int) ([]byte, error) {
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

type (
	aesPadType    string
	aesCipherType string
)

const (
	AesPadTypeOfPKCS7 aesPadType = `PKCS7`
	AesPadTypeOfPKCS5 aesPadType = `PKCS5`

	AesCipherTypeOfCBC aesCipherType = `CBC`
	AesCipherTypeOfECB aesCipherType = `ECB`
)

// AES加密
func AesEncrypt(rawByte []byte, keyByte []byte, padType aesPadType, padLen int, cipherType aesCipherType, iv ...byte) (cipherByte []byte, err error) {
	block, err := aes.NewCipher(keyByte)
	if err != nil {
		return
	}
	blockSize := block.BlockSize()

	if padLen == 0 {
		padLen = blockSize
	}
	switch padType {
	// case AesPadTypeOfPKCS7, AesPadTypeOfPKCS5:
	default:
		rawByte = PKCSPad(rawByte, padLen)
	}

	rawByteLen := len(rawByte)
	if rawByteLen%blockSize != 0 {
		err = errors.New(`加密串必须是块大小的整数倍`)
		return
	}

	cipherByte = make([]byte, rawByteLen)
	switch cipherType {
	case AesCipherTypeOfECB:
		for i := 0; i < rawByteLen; i += blockSize {
			block.Encrypt(cipherByte[i:i+blockSize], rawByte[i:i+blockSize])
		}
	// case AesCipherTypeOfCBC:
	default:
		if len(iv) == 0 { //默认偏移量
			iv = make([]byte, blockSize)
		}
		blockMode := cipher.NewCBCEncrypter(block, iv)
		blockMode.CryptBlocks(cipherByte, rawByte)
	}
	return
}

// AES解密
func AesDecrypt(cipherByte []byte, keyByte []byte, padType aesPadType, padLen int, cipherType aesCipherType, iv ...byte) (rawByte []byte, err error) {
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
	case AesCipherTypeOfECB:
		for i := 0; i < cipherByteLen; i += blockSize {
			block.Decrypt(rawByte[i:i+blockSize], cipherByte[i:i+blockSize])
		}
	// case AesCipherTypeOfCBC:
	default:
		if len(iv) == 0 { //默认偏移量
			iv = make([]byte, blockSize)
		}
		blockMode := cipher.NewCBCDecrypter(block, iv)
		blockMode.CryptBlocks(rawByte, cipherByte)
	}

	// 补码处理
	if padLen == 0 {
		padLen = blockSize
	}
	switch padType {
	// case AesPadTypeOfPKCS7, AesPadTypeOfPKCS5:
	default:
		rawByte, err = PKCSUnPad(rawByte, padLen)
	}
	return
}

// AES加密（BASE64编码）
func AesEncryptOfBase64(rawByte []byte, keyByte []byte, padType aesPadType, padLen int, cipherType aesCipherType, iv ...byte) (encrypt string, err error) {
	cipherByte, err := AesEncrypt(rawByte, keyByte, padType, padLen, cipherType, iv...)
	if err != nil {
		return
	}
	encrypt = base64.StdEncoding.EncodeToString(cipherByte)
	return
}

// AES解密（BASE64编码）
func AesDecryptOfBase64(encrypt string, keyByte []byte, padType aesPadType, padLen int, cipherType aesCipherType, iv ...byte) (rawByte []byte, err error) {
	encryptByte, err := base64.StdEncoding.DecodeString(encrypt)
	if err != nil {
		return
	}
	rawByte, err = AesDecrypt(encryptByte, keyByte, padType, padLen, cipherType, iv...)
	return
}
