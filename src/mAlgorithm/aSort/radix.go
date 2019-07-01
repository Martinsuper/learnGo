package aSort

import (
	"fmt"
	"strconv"
)

func RadixSort(ans []int) {
	//var darray [10][100]int
	darray := make([][]int, 10000000)
	length := getMaxNumLen(ans)
	mod := 10
	dev := 1
	for i := 0; i < length; {
		for j := 0; j < len(ans); j++ {
			bucketValue := ans[j]%mod/dev + mod
			darray[bucketValue] = append(darray[bucketValue], ans[j])
			fmt.Println(bucketValue, ans[j])
		}
		cnt := 0
		for k := 0; k < 10; k++ {
			for l := 0; l < len(darray[k]); l++ {
				ans[cnt] = darray[k][l]
				cnt++
			}
		}
		//break
		mod *= 10
		dev *= 10
	}
	ExportArray(darray[0])
}

func getMaxNumLen(ans []int) int {
	maxNum := ans[0]
	for i := 0; i < len(ans); i++ {
		if maxNum < ans[i] {
			maxNum = ans[i]
		}
	}
	maxStr := strconv.Itoa(maxNum)
	return len(maxStr)
}

func test() {
	length := len(strconv.Itoa(100))
	fmt.Println(length)
}
