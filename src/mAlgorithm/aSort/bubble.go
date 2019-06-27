package aSort

import "fmt"

func Sbubble(ans []int) {
	ansLen := len(ans)
	for i := 0; i < ansLen-1; i++ {
		for j := i + 1; j < ansLen; j++ {
			if ans[i] > ans[j] {
				ans[i], ans[j] = ans[j], ans[i]
			}
		}
	}
	//for i := 0; i < ansLen; i++ {
	//	fmt.Printf("%d ", ans[i])
	//}
	fmt.Printf("冒泡排序：%d个随机数\n", ansLen)
}
