package leetcode

func quickSortArr(arr []int) {
	var qsort func(l, r int)
	var part func(l, r int) int
	part = func(l, r int) int {
		val := arr[l]
		i, j := l+1, l+1
		for j <= r {
			if arr[j] < val {
				arr[i], arr[j] = arr[j], arr[i]
				i++
			}
			j++
		}
		arr[l], arr[i-1] = arr[i-1], arr[j]
		return i
	}

	qsort = func(l, r int) {
		if l < r {
			part := part(l, r)
			qsort(l, part-1)
			qsort(part+1, r)
		}
	}
}
