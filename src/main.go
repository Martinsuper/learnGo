package main

import (
	"fmt"
	"mAlgorithm/aSort"
	"mAlgorithm/getData"
	"path/filepath"
	"time"
)

func main() {
	datapath, _ := filepath.Abs(filepath.Dir("./"))
	datapath += "/src/mAlgorithm/data/sortData1"
	ans := getData.ReadFile(datapath)
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
