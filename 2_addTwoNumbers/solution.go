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
			v1 = n1.val
			n1 = n1.next
		}

		if n2 != nil {
			v2 = n2.val
			n2 = n2.next
		}

		sum := v1 + v2 + carry
		carry = sum / 10

		n.next = &ListNode{sum % 10, nil}
		n = n.next
	}

	if carry > 0 {
		n.next = &ListNode{carry, nil}
	}

	return root.next
}

type ListNode struct {
	val  int
	next *ListNode
}
