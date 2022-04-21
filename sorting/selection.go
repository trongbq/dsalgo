package sorting

// SelectionSort sorts a list of integer in ascending order
// by selecting smallest number once at a time
func SelectionSort(arr []int) {
	n := len(arr)

	for i := 0; i < n-1; i++ {
		// `min` marks index of the smallest number in arr[i:]
		min := i
		// Select the smallest number in arr[i+1:] to swap with arr[i]
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		// Swap the smallest number with with number at `i` index
		arr[i], arr[min] = arr[min], arr[i]
	}
}
