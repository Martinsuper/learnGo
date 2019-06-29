package aSort

import (
	"fmt"
	"time"
)

func IsSorted(ans *[]int) {
	for i := 0; i < len(*ans)-1; i++ {
		if (*ans)[i] > (*ans)[i+1] {
			fmt.Println("无序")
			return
		}
	}
	fmt.Println("有序")
}

func ExportArray(array *[]int) {
	for i := 0; i < len(*array); i++ {
		fmt.Printf("%d ", (*array)[i])
	}
	time.Sleep(1000000000)
	fmt.Println()
}
