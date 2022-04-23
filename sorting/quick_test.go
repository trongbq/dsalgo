package sorting_test

import (
	"testing"

	"github.com/dsalgo/sorting"
	"github.com/stretchr/testify/assert"
)

func TestQuickSort(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			"Empty array",
			[]int{},
			[]int{},
		},
		// {
		// 	"Array contains unique values",
		// 	[]int{5, 2, 7, 4, 19, 25, 84, 1, 6, 77, 8},
		// 	[]int{1, 2, 4, 5, 6, 7, 8, 19, 25, 77, 84},
		// },
		// {
		// 	"Array contains duplicate values",
		// 	[]int{5, 2, 7, 4, 19, 25, 84, 7, 6, 25, 8},
		// 	[]int{2, 4, 5, 6, 7, 7, 8, 19, 25, 25, 84},
		// },
		{
			"Array contains unique values",
			[]int{5, 2, 7, 4, 19, 25, 84, 1, 6, 77, 8},
			[]int{1, 2, 4, 5, 6, 7, 8, 19, 25, 77, 84},
		},
	}

	for _, tc := range tests {
		sorting.QuickSort(tc.input)

		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, tc.input)
		})
	}

}
