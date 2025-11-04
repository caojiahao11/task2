package util

func DoubleSliceWithRangeImpl(ptr *[]int) {
	if ptr == nil || *ptr == nil {
		return
	}

	// 先解引用，然后使用range
	slice := *ptr
	for i := range slice {
		slice[i] *= 2
	}
}
