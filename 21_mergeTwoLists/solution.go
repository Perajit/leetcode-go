package mergetwolists

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	node1 := list1
	node2 := list2
	root := &ListNode{}
	node := root

	for node1 != nil || node2 != nil {
		if node1 != nil && (node2 == nil || node1.Val <= node2.Val) {
			node.Next = &ListNode{node1.Val, nil}
			node1 = node1.Next
		} else {
			node.Next = &ListNode{node2.Val, nil}
			node2 = node2.Next
		}

		node = node.Next
	}

	return root.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
