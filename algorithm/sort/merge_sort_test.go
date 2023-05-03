package sort

import (
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {
	arr := []int{10, 6, 2, 5, 8, 20, 12}
	MergeSort(arr, 0, len(arr))
	fmt.Println("merge sort: ", arr)
}
