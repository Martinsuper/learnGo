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
	datapath += "/src/mAlgorithm/data/sortData2"
	ans := getData.ReadFile(datapath)
	//bubbleTime(&ans)
	//insertionTime(&ans)
	aSort.Selection(ans)
	aSort.ExportArray(&ans)
	aSort.IsSorted(&ans)
	//getData.SetData()

}

func bubbleTime(ans *[]int) {
	start := time.Now()
	aSort.Sbubble1(*ans)
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
