package sorting

func MergeSort(arr []int) {
	helper := make([]int, len(arr))
	mergeSort(arr, 0, len(arr)-1, helper)
}

func mergeSort(arr []int, l, h int, helper []int) {
	if l < h {
		m := (l + h) / 2
		mergeSort(arr, l, m, helper)
		mergeSort(arr, m+1, h, helper)
		merge(arr, l, m, h, helper)
	}
}

func merge(arr []int, l, m, h int, helper []int) {
	// Copy left and right array into helper array
	for i := l; i <= h; i++ {
		helper[i] = arr[i]
	}

	index := l
	lindex := l
	hindex := m + 1
	// Compare and put back to main array until one of array is completed
	for lindex <= m && hindex <= h {
		if helper[lindex] <= helper[hindex] {
			arr[index] = helper[lindex]
			lindex++
		} else {
			arr[index] = helper[hindex]
			hindex++
		}
		index++
	}

	// Copy remaining values of left array
	// No need to worry with right array since it already in its position
	// if there are any remaining
	for i := 0; i < (m - lindex + 1); i++ {
		arr[index+i] = helper[lindex+i]
	}
}
