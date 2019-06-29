package aSort

import "fmt"

/*
最基本版冒泡排序，每次选出一个最小的数放到前面，即每次都是从第i位置开始，找到最小的数放到第i位置上
*/
func Sbubble(ans []int) {
	ansLen := len(ans)
	for i := 0; i < ansLen-1; i++ {
		for j := i + 1; j < ansLen; j++ {
			if ans[i] > ans[j] {
				ans[i], ans[j] = ans[j], ans[i]
				//ExportArray(&ans)
			}
		}
	}
	fmt.Printf("冒泡排序：%d个随机数\n", ansLen)
}

func Sbubble1(ans []int) {
	ansLen := len(ans)
	for i := 0; i < ansLen; i++ {
		for j := 0; j < ansLen-1-i; j++ {
			if ans[j] > ans[j+1] {
				ans[j], ans[j+1] = ans[j+1], ans[j]
				//ExportArray(&ans)
			}
		}
	}
	fmt.Printf("冒泡排序：%d个随机数\n", ansLen)
}
func Sbubble2(ans []int) {
	ansLen := len(ans)
	flag := true
	for i := 0; i < ansLen && flag; i++ {
		flag = false
		for j := 0; j < ansLen-1-i; j++ {
			if ans[j] > ans[j+1] {
				ans[j], ans[j+1] = ans[j+1], ans[j]
				flag = true
				//ExportArray(&ans)
			}
		}
		//fmt.Println(i,ansLen)
	}
	fmt.Printf("冒泡排序：%d个随机数\n", ansLen)
}

func Sbubble3(ans []int) {
	ansLen := len(ans)
	for i := 0; i < ansLen-1; i++ {
		for j := ansLen - 1; j > i; j-- {
			if ans[j] < ans[i] {
				ans[j], ans[i] = ans[i], ans[j]
				//ExportArray(&ans)
			}
		}
		//fmt.Println(i,ansLen)
	}
	fmt.Printf("冒泡排序：%d个随机数\n", ansLen)
}
