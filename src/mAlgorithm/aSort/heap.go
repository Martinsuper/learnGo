package aSort

import "fmt"

func HeapSort(ans []int) {
	ansLen := len(ans)
	for i := ansLen/2 - 1; i >= 0; i-- {
		adjustHeap(ans, i, ansLen)
	}
	for j := ansLen - 1; j > 0; j-- {
		ans[0], ans[j] = ans[j], ans[0]
		adjustHeap(ans, 0, j)
	}
	fmt.Printf("希尔排序：%d个随机数\n", ansLen)
}

func adjustHeap(ans []int, i, length int) {
	pos := i
	for k := 2*i + 1; k < length; k = 2*k + 1 {
		if k+1 < length && ans[k] < ans[k+1] {
			k++
		}
		if ans[k] > ans[pos] {
			ans[k], ans[pos] = ans[pos], ans[k]
			pos = k
		} else {
			break
		}
	}
}
