func BubbleSort(sli []int) {
        for i := range sli {
                for j := 0; j < len(sli)-i-1; j++ {
                        if sli[j] > sli[j+1] {
                                Swap(sli, j)
                        }
                }
        }
}

func Swap(sli []int, i int) {
        tmp := sli[i]
        sli[i], sli[i+1] = sli[i+1], tmp
}
