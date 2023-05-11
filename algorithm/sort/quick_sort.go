package sort

import (
	"math/rand"
)

// QuickSort [l, r)
func QuickSort(arr []int, l int, r int) {
	if l+1 >= r {
		return
	}

	// random pivot
	idx := rand.Intn(r-l) + l
	pivot := arr[idx]

	arr[l], arr[idx] = arr[idx], arr[l]

	left, right := l, r-1
	for left != right {
		for arr[right] >= pivot {
			right--
			if left == right {
				break
			}
		}
		arr[right], arr[left] = arr[left], arr[right]
		for arr[left] < pivot {
			left++
			if left == right {
				break
			}
		}
		arr[right], arr[left] = arr[left], arr[right]
	}
	arr[left] = pivot

	QuickSort(arr, 0, left)
	QuickSort(arr, left+1, r)
}
