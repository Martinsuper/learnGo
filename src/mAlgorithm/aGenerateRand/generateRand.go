package aGenerateRand

import (
	"math/rand"
	"time"
)

func GenerateRand() int {
	rand.Seed(time.Now().UnixNano())
	//for i:=0;i<10;i++{
	//	rnd := rand.Intn(1000000000)
	//	fmt.Printf("%v\n",rnd)
	//}
	return rand.Intn(10000)
}
