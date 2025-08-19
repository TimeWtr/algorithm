package heap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKSmallestPairs(t *testing.T) {
	array1 := []int{4, 8, 10, 12, 23, 1, 18}
	array2 := []int{8, 3, 13, 30, 20, 24, 18}
	sortedArray1 := sorted(array1, true)
	sortedArray2 := sorted(array2, true)
	testCases := []struct {
		name    string
		k       int
		wantRes int
	}{
		{
			name:    "3K",
			k:       3,
			wantRes: 20, // [[1, 3], [4, 3], [1, 8]]
		},
		{
			name:    "2K",
			k:       2,
			wantRes: 11, // [[1, 3], [4, 3]]
		},
		{
			name:    "4K",
			k:       4,
			wantRes: 31, // [[1, 3], [4, 3], [1, 8], [8, 3]]
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := kSmallestPairs(sortedArray1, sortedArray2, tc.k)
			assert.Equal(t, tc.wantRes, res)
		})
	}
}
