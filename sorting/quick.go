package sorting

func QuickSort(arr []int) {
	quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, l, h int) {
	if l < h {
		m := partition(arr, l, h)
		quickSort(arr, l, m-1)
		quickSort(arr, m+1, h)
	}
}

func partition(arr []int, l, h int) int {
	// Select pivot at the last value
	pivot := arr[h]

	// We separate array in 3 parts
	// - arr[:i+1]: contains values that are less than or equal pivot
	// - arr[i+1:h]: contains values that is not determined
	// - arr[h]: our pivot
	// We will go through arr[i-1:h] and search for less than or equal pivot values and append them
	// to the first part.
	i := l - 1
	for j := l; j < h; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	// Swap pivot value to its correct position
	arr[i+1], arr[h] = pivot, arr[i+1]

	return i + 1
}
