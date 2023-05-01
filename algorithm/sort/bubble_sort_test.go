package sort

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	arr := []int{10, 6, 2, 5, 8, 20, 12}
	BubbleSort(arr)
	fmt.Println("bubble sort: ", arr)
}
