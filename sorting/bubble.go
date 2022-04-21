package sorting

// BubbleSort sorts array in ascending order
// by move let largest value bubble up to the tail of the array
func BubbleSort(arr []int) {
	n := len(arr)

	// Bubble up n-1 values, last value already in its position
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
