package swappairs

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	root := &ListNode{}
	root.Next = head

	var node *ListNode = head
	var prevNode *ListNode = root

	for node != nil {
		nextNode := node.Next

		if nextNode != nil {
			node.Next = nextNode.Next
			nextNode.Next = node
			prevNode.Next = nextNode
			prevNode = node
		}

		node = node.Next
	}

	return root.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
