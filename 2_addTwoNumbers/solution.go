package addtwonumbers

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	root := &ListNode{}
	n := root
	n1 := l1
	n2 := l2
	var carry int

	for n1 != nil || n2 != nil {
		var v1, v2 int

		if n1 != nil {
			v1 = n1.Val
			n1 = n1.Next
		}

		if n2 != nil {
			v2 = n2.Val
			n2 = n2.Next
		}

		sum := v1 + v2 + carry
		carry = sum / 10

		n.Next = &ListNode{sum % 10, nil}
		n = n.Next
	}

	if carry > 0 {
		n.Next = &ListNode{carry, nil}
	}

	return root.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
