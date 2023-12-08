package utils

import (
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
				b := (keyArray[i] | ^(1<<k)&0xFF)
				encryptArray[step+j] &= byte(b)
			} else {
				encryptArray[step+j] |= 1 << k
			}
		}
	}
	return string(encryptArray)
}

func EncryptRSA(plainText string, publicKey string) (string, error) {
	return gorsa.PublicEncrypt(plainText, publicKey)
}

func PriKeyEncrypt(plainText string, privateKey string) (string, error) {
	return gorsa.PriKeyEncrypt(plainText, privateKey)
}
