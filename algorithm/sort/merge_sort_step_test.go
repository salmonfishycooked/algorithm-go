package sort

import (
	"fmt"
	"testing"
)

func TestMergeSortStep(t *testing.T) {
	arr := []int{10, 6, 2, 5, 6, 20, 3, 2, 5}
	MergeSortStep(arr, 0, len(arr))
	fmt.Println("merge sort step: ", arr)
}
