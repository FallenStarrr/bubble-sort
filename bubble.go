func BubbleSort(sliceToSort []int) {
	for i := 0; i < len(sliceToSort)-1; i++ {
		for j := 0; j < len(sliceToSort)-1; j++ {
			if sliceToSort[j] > sliceToSort[j+1] {
				Swap(sliceToSort, j)
				//sliceToSort[j], sliceToSort[j+1] = sliceToSort[j+1], sliceToSort[j]
			}
		}
	}
}

func Swap(o []int, i int) {
	if len(o) <= i {
		fmt.Println("error - cannot swap as at end of slice")
	} else {
		o[i], o[i+1] = o[i+1], o[i]
	}
}
