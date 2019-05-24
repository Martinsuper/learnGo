package main

import (
	"mAlgorithm/aSort"
	"mAlgorithm/getData"
)

func main() {
	ans := getData.ReadFile("/Users/duanyu/workspace/learnGo/src/mAlgorithm/data/sortData")
	aSort.Sbubble(ans)
}
