package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/pem"

	"github.com/wenzhenxi/gorsa"
)

func DecryptKey(encrypt string) string {
	encryptArray := []byte(encrypt)
	word := 8
	length := len(encryptArray) / word
	var keyArray = make([]byte, length)
	for i := 0; i < length; i++ {
		for j := 0; j < word; j++ {
			keyArray[i] |= encryptArray[i*word+j] & (128 >> j)
		}
	}
	return string(keyArray)
}

func DecryptRSA(cipherByte []byte, privateKey string) (plainText string, err error) {
	priBlock, _ := pem.Decode([]byte(privateKey))
	priKey, err := x509.ParsePKCS1PrivateKey(priBlock.Bytes)
	if err != nil {
		panic(err)
	}
	decryptOAEP, err := rsa.DecryptOAEP(sha1.New(), rand.Reader, priKey, cipherByte, nil)
	if err != nil {
		panic(err)
	}
	plainText = string(decryptOAEP)
	return
}

func PublicDecrypt(cipherPlain string, publicKey string) (string, error) {
	return gorsa.PublicDecrypt(cipherPlain, publicKey)
}
