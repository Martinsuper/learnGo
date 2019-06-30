package aSort

import "fmt"

func Sinsert(ans []int) {
	ansLen := len(ans)
	for i := 1; i < ansLen; i++ {
		for j := i - 1; j >= 0; j-- {
			if ans[j+1] < ans[j] {
				ans[j+1], ans[j] = ans[j], ans[j+1]
			} else {
				//ExportArray(&ans)
				break
			}
		}
	}
	fmt.Printf("直接插入排序：%d个随机数\n", ansLen)
}

func Shellsort(ans []int) {
	ansLen := len(ans)
	for increment := ansLen / 2; increment >= 1; increment = increment / 2 {
		for i := increment; i < ansLen; i++ {
			for j := i - increment; j >= 0 && ans[j] > ans[j+increment]; j -= increment {
				ans[j], ans[j+increment] = ans[j+increment], ans[j]
				//ExportArray(&ans)
			}
		}
	}
	fmt.Printf("希尔排序：%d个随机数\n", ansLen)
}
