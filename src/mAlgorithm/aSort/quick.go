package aSort

func QuickSort(ans []int, low, high int) {
	if low < high {
		pivot := partition(ans, low, high)
		QuickSort(ans, low, pivot-1)
		QuickSort(ans, pivot+1, high)
	}
}

func partition(ans []int, low, high int) int {
	pivot := ans[low]
	for low < high {
		for ; low < high && ans[high] >= pivot; high-- {
		}
		ans[low] = ans[high]
		for ; low < high && ans[low] <= pivot; low++ {
		}
		ans[high] = ans[low]
	}
	ans[low] = pivot
	return low
}
