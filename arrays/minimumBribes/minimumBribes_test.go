package mb

import (
	"testing"
)

func TestMinimumBribes(t *testing.T) {
	var tests = []struct {
		in  []int32
		out int
	}{
		{
			in:  []int32{1, 2, 5, 3, 7, 8, 6, 4},
			out: 3,
		},
		// {
		// 	in:  []int32{2, 5, 1, 3, 4},
		// 	out: 3,
		// },
	}

	for _, test := range tests {
		err := minimumBribes(test.in)
		if err != nil {
			t.Error("Error")
		}
	}
}
