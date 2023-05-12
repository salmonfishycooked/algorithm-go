package segment_tree

import (
	"fmt"
	"testing"
)

func TestSegmentTree(t *testing.T) {
	nums := []int{2, 3, 5, 5, 4, 10, 20, 3, 2, 50}
	Init(nums)
	fmt.Println(Query(2, 5, 0))
}
