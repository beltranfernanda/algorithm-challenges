package challenge1

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	// Test cases
	testCases := []struct {
		nums   []int
		target int
		output []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
		{[]int{1, 2, 3, 4, 5}, 9, []int{3, 4}},
		{[]int{-1, -2, -3, -4, -5}, -8, []int{2, 4}},
		{[]int{0, 4, 3, 0}, 0, []int{0, 3}},
	}

	// Iterate over test cases
	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			result := TwoSum(tc.nums, tc.target)
			if !reflect.DeepEqual(result, tc.output) {
				t.Errorf("Expected %v but got %v", tc.output, result)
			}
		})
	}
}
