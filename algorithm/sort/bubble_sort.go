package sort

func BubbleSort(arr []int) {
	l := len(arr)
	for i := l - 1; i > 0; i-- {
		flag := true
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				flag = false
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
		if flag {
			break
		}
	}
}
