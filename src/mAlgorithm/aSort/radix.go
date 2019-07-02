package aSort

import (
	"strconv"
)

func RadixSort(ans []int) {
	length := getMaxNumLen(ans)
	mod := 10
	for i := 1; i <= length; i++ {
		tmp := setBucket(ans, mod)
		mod *= 10
		getArray(tmp, ans)
	}
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

func setBucket(ans []int, num int) [][]int {
	tmp := make([][]int, 10000000)
	for j := 0; j < len(ans); j++ {
		bucketValue := ans[j] % num / (num / 10)
		tmp[bucketValue] = append(tmp[bucketValue], ans[j])
	}
	return tmp
}

func getArray(tmp [][]int, ans []int) {
	cnt := 0
	for i := 0; i < 10; i++ {
		for j := 0; j < len(tmp[i]); j++ {
			ans[cnt] = tmp[i][j]
			cnt++
		}
	}
	//ExportArray(ans)
}
