package search_test

import (
	"testing"

	"github.com/dsalgo/search"
	"github.com/stretchr/testify/assert"
)

func TestKarbinKarp(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want int
	}{
		{
			"s is larger than t",
			"abcd",
			"abc",
			-1,
		},
		{
			"s is t",
			"abcd",
			"abcd",
			0,
		},
		{
			"s in the beginning of t",
			"abc",
			"abcdefgh",
			0,
		},
		{
			"s in the middle of t",
			"abde",
			"abdfseabdeage",
			6,
		},
		{
			"s is in the end of t",
			"tyqqa",
			"abcdeaoratyqqstyqqa",
			14,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := search.RabinKarp(tc.s, tc.t)
			assert.Equal(t, tc.want, got)
		})
	}

}
