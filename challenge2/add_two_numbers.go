package challenge2

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var carry int
	l3 := &ListNode{}
	temp := l3
	var sum int
	for l1 != nil || l2 != nil {
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		sum += carry
		carry = sum / 10
		temp.Next = &ListNode{Val: sum % 10}
		temp = temp.Next
		sum = 0
	}
	if carry > 0 {
		temp.Next = &ListNode{Val: carry}
	}
	return l3.Next
}

// Most efficient solution
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := new(ListNode)
	head := result

	var val int
	var carry int

	curL1 := l1
	curL2 := l2
	for curL1 != nil || curL2 != nil || carry > 0 {
		sumNum := carry
		if curL1 != nil {
			sumNum += curL1.Val
		}

		if curL2 != nil {
			sumNum += curL2.Val
		}

		val = sumNum % 10
		carry = sumNum / 10

		newNode := &ListNode{Val: val}
		head.Next = newNode
		head = newNode

		if curL1 != nil {
			curL1 = curL1.Next
		}

		if curL2 != nil {
			curL2 = curL2.Next
		}
	}

	return result.Next
}
