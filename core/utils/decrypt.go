package utils

import (
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

func DecryptRSA(cipherText string, privateKey string) (string, error) {
	return gorsa.PriKeyDecrypt(cipherText, privateKey)
}

func PublicDecrypt(cipherText string, publicKey string) (string, error) {
	return gorsa.PublicDecrypt(cipherText, publicKey)
}
