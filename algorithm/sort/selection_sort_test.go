package sort

import (
	"fmt"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	arr := []int{10, 6, 2, 5, 8, 20, 12}
	SelectionSort(arr)
	fmt.Println("selection sort: ", arr)
}
