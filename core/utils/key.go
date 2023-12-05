package utils

import "time"

var next = time.Now().UnixMilli()

func rand(bound int64) int64 {
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
			rand1 := rand(2)
			var temp int64 = 0
			if rand1 == 0 {
				temp = 26
			} else {
				temp = 10
			}
			rand2 := rand(int64(temp))
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
