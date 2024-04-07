package challenge2

import "testing"

func TestAddTwoNumbers(t *testing.T) {
	// Test cases
	testCases := []struct {
		l1     *ListNode
		l2     *ListNode
		result *ListNode
	}{
		{
			&ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}},
			&ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}},
			&ListNode{Val: 7, Next: &ListNode{Val: 0, Next: &ListNode{Val: 8}}},
		},
		{
			&ListNode{Val: 0},
			&ListNode{Val: 0},
			&ListNode{Val: 0},
		},
		{
			&ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}},
			&ListNode{Val: 1},
			&ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 1}}}},
		},
	}

	// Iterate over test cases
	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			result := AddTwoNumbers(tc.l1, tc.l2)
			if !isEqual(result, tc.result) {
				t.Errorf("Expected %v but got %v", tc.result, result)
			}
		})
	}
}

// Helper function to compare two ListNode
func isEqual(l1 *ListNode, l2 *ListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return l1 == nil && l2 == nil
}
