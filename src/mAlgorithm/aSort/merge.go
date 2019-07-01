package aSort

func MergeSort(ans []int) []int {
	ansLen := len(ans)
	if ansLen < 2 {
		return ans
	}
	left := ans[0:(ansLen / 2)]
	right := ans[(ansLen/2 + 1):]
	return merge(MergeSort(left), MergeSort(right))
}

func merge(left []int, right []int) []int {
	var res []int
	for len(left) != 0 && len(right) != 0 {
		if left[0] <= right[0] {
			res = append(res, left[0])
			left = left[1:]
		} else {
			res = append(res, right[0])
			right = right[1:]
		}
	}
	for len(left) != 0 {
		res = append(res, left[0])
		left = left[1:]
	}
	for len(right) != 0 {
		res = append(res, right[0])
		right = right[1:]
	}
	return res
}
