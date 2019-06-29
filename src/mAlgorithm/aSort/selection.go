package aSort

func Selection(ans []int) {
	ansLen := len(ans)
	for i := 0; i < ansLen-1; i++ {
		minpos := i
		for j := i + 1; j < ansLen; j++ {
			if ans[j] < ans[minpos] {
				minpos = j
			}
		}
		if minpos != i {
			ans[minpos], ans[i] = ans[i], ans[minpos]
		}
	}
}
