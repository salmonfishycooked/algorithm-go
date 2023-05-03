package sort

// MergeSort [l, r)
func MergeSort(arr []int, l int, r int) {
	if l+1 >= r {
		return
	}
	mid := l + (r-l)>>1
	MergeSort(arr, l, mid)
	MergeSort(arr, mid, r)
	merge(arr, l, r)
}

// merge [l, r)
func merge(arr []int, l int, r int) {
	var help []int
	mid := l + (r-l)>>1
	i, j := l, mid
	for i < mid && j < r {
		if arr[i] < arr[j] {
			help = append(help, arr[i])
			i++
		} else {
			help = append(help, arr[j])
			j++
		}
	}

	for i < mid {
		help = append(help, arr[i])
		i++
	}
	for j < r {
		help = append(help, arr[j])
		j++
	}

	// copy help to arr
	copy(arr[l:r], help)
}
