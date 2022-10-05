func BubbleSort(num []int) []int {
	for i := len(num); i > 0; i-- {
		for j := 1; j < i; j++ {
			if num[j-1] > num[j] {
				num[j-1], num[j] = num[j], num[j-1]
			}
		}
	}

	return num
}
