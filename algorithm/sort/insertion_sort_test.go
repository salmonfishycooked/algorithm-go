package sort

import (
	"fmt"
	"testing"
)

func TestInsertSort(t *testing.T) {
	arr := []int{10, 6, 2, 5, 8, 20, 12}
	InsertionSort(arr)
	fmt.Println("insertion sort: ", arr)
}
