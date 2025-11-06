package removenthfromend

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	nodeLookup := make(map[int]*ListNode)
	node := head
	count := 0

	for node != nil {
		nodeLookup[count] = node
		count++
		node = node.Next
	}

	// to remove first node
	if n == count {
		return head.Next
	}

	var prevNode *ListNode = nodeLookup[count-n-1]
	var nextNode *ListNode = nil

	// to remove node other than last node
	if n > 1 {
		nextNode = nodeLookup[count-n+1]
	}

	prevNode.Next = nextNode

	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}
