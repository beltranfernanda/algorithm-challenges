package challenge4

import "testing"

func TestFindMedianSortedArrays(t *testing.T) {
	// Test cases
	testCases := []struct {
		nums1    []int
		nums2    []int
		expected float64
	}{
		{[]int{1, 3}, []int{2}, 2.0},
		{[]int{1, 2}, []int{3, 4}, 2.5},
		{[]int{0, 0}, []int{0, 0}, 0.0},
		{[]int{}, []int{1}, 1.0},
		{[]int{2}, []int{}, 2.0},
		{[]int{1, 2, 3}, []int{4, 5, 6}, 3.5},
	}

	// Iterate over test cases
	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			result := FindMedianSortedArrays(tc.nums1, tc.nums2)
			if result != tc.expected {
				t.Errorf("Expected %f but got %f", tc.expected, result)
			}
		})
	}
}
