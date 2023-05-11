package sort

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	arr := []int{10, 6, 2, 5, 6, 20, 3, 2, 5}
	QuickSort(arr, 0, len(arr))
	fmt.Println("quick sort: ", arr)
}
