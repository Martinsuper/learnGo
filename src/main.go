package main

import (
	"fmt"
	"mAlgorithm/aSort"
	"mAlgorithm/getData"
	"time"
)

func main() {
	ans := getData.ReadFile("/Users/duanyu/workspace/learnGo/src/mAlgorithm/data/sortData1")
	bubbleTime(&ans)
	insertionTime(&ans)
}

func bubbleTime(ans *[]int) {
	start := time.Now()
	aSort.Sbubble(*ans)
	end := time.Since(start)
	fmt.Printf("用时：")
	fmt.Println(end)
}

func insertionTime(ans *[]int) {
	start := time.Now()
	aSort.Sinsert(*ans)
	end := time.Since(start)
	fmt.Printf("用时：")
	fmt.Println(end)
}
