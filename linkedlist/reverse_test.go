package linkedlist_test

import (
	"testing"

	"github.com/dsalgo/linkedlist"
	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		name string
		head linkedlist.Node
		want []int
	}{
		{
			"Reverse a single node linked list",
			*linkedlist.FromArray([]int{3}),
			[]int{3},
		},
		{
			"Reverse an odd number of nodes linked list",
			*linkedlist.FromArray([]int{3, 4, 5, 6, 7}),
			[]int{7, 6, 5, 4, 3},
		},
		{
			"Reverse an even number of nodes linked list",
			*linkedlist.FromArray([]int{3, 4, 5, 6, 7, 19}),
			[]int{19, 7, 6, 5, 4, 3},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := linkedlist.Reverse(tc.head)
			assert.Equal(t, tc.want, linkedlist.ToArray(&got))
		})
	}

}
