package aGenerateRand

import (
	"math/rand"
	"time"
)

func GenerateRand() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(10000)
}
