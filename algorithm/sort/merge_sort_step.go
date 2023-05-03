package sort

// MergeSortStep [l, r)
// Non-recursive method implementation
func MergeSortStep(arr []int, l int, r int) {
	step := 1
	for step < r-l {
		left, right := 0, step

		// merge [left, left+step) and [right, right+step)
		for left < r {
			var help []int
			i, j := left, right
			for i < left+step && j < right+step && j < r {
				if arr[i] < arr[j] {
					help = append(help, arr[i])
					i++
				} else {
					help = append(help, arr[j])
					j++
				}
			}
			for i < left+step && i < r {
				help = append(help, arr[i])
				i++
			}
			for j < right+step && j < r {
				help = append(help, arr[j])
				j++
			}

			if i == r {
				copy(arr[left:i], help)
			} else {
				copy(arr[left:j], help)
			}
			left, right = left+step*2, right+step*2
		}

		// Prevent the value from exceeding the maximum value of int
		if step > (r-l)>>1 {
			break
		}
		step <<= 1
	}
}
