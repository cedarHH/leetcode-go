package assessment

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestXingyeAss1(t *testing.T) {

	testCase := []struct {
		inputArr1 []int
		inputArr2 []int
		expected  float32
	}{
		{[]int{}, []int{1}, 1.0},
		{[]int{1, 3}, []int{2}, 2.0},
		{[]int{1, 2}, []int{3, 4}, 2.5},
		{[]int{1, 2, 3}, []int{4, 5}, 3.0},
	}

	for _, tc := range testCase {
		result := xingyeAss1(
			tc.inputArr1,
			tc.inputArr2)
		require.Equal(t, tc.expected, result,
			"XingyeAss1(%v, %v) was expected to return %f but got %f",
			tc.inputArr1, tc.inputArr2, tc.expected, result)
	}
}
