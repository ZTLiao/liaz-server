package utils

import (
	"math/rand"
	"time"
)

func RandomForSix() int {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	return rand.Intn(900000) + 100000
}
