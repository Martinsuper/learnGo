package main

import (
	"fmt"
	"mAlgorithm/aSort"
	"mAlgorithm/getData"
	"time"
)

func main() {
	bubbleTime()
}

func bubbleTime() {
	ans := getData.ReadFile("/Users/duanyu/workspace/learnGo/src/mAlgorithm/data/sortData1")
	start := time.Now()
	aSort.Sbubble(ans)
	end := time.Since(start)
	fmt.Println(end)
}
