package aSort

import "fmt"

func Sinsert(ans []int) {
	ansLen := len(ans)
	for i := 1; i < ansLen; i++ {
		for j := i - 1; j >= 0; j-- {
			if ans[j+1] < ans[j] {
				ans[j+1], ans[j] = ans[j], ans[j+1]
			} else {
				break
			}
		}
	}
	fmt.Printf("直接插入排序：%d个随机数\n", ansLen)
	//for i := 0; i < ansLen; i++ {
	//	fmt.Printf("%d\n", ans[i])
	//}
}