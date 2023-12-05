package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/pem"
	"time"

	"github.com/wenzhenxi/gorsa"
)

var next = time.Now().UnixMilli()

func random(bound int64) int64 {
	next = next*1103515245 + 12345
	rand := (next / 65536) % 32768
	if rand < 0 {
		rand = -rand
	}
	return (rand % bound)
}

func EncryptKey(key string) string {
	keyArray := []byte(key)
	word := 8
	length := len(keyArray)
	var encryptArray = make([]byte, length*word)
	for i := 0; i < length; i++ {
		step := i * word
		for j, k := 0, word-1; j < word; j, k = j+1, k-1 {
			rand1 := random(2)
			var temp int64 = 0
			if rand1 == 0 {
				temp = 26
			} else {
				temp = 10
			}
			rand2 := random(int64(temp))
			if rand1 == 0 {
				temp = 97
			} else {
				temp = 48
			}
			encryptArray[step+j] = byte(rand2 + temp)
			flag := (keyArray[i] & (1 << k)) == 0
			if flag {
				b := (int(keyArray[i]) | ^(i<<k)&0xFF)
				encryptArray[step+j] &= byte(b)
			} else {
				encryptArray[step+j] |= 1 << k
			}
		}
	}
	return string(encryptArray)
}

func EncryptRSA(plain string, publicKey string) (cipherByte []byte, err error) {
	msg := []byte(plain)
	pubBlock, _ := pem.Decode([]byte(publicKey))
	pubKeyValue, err := x509.ParsePKIXPublicKey(pubBlock.Bytes)
	if err != nil {
		panic(err)
	}
	pub := pubKeyValue.(*rsa.PublicKey)
	encryptOAEP, err := rsa.EncryptOAEP(sha1.New(), rand.Reader, pub, msg, nil)
	if err != nil {
		panic(err)
	}
	cipherByte = encryptOAEP
	return
}

func PriKeyEncrypt(plain string, privateKey string) (string, error) {
	return gorsa.PriKeyEncrypt(plain, privateKey)
}
