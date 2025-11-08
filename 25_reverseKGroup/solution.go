package reversekgroup

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	}

	root := &ListNode{}
	root.Next = head

	var node *ListNode = head
	var prevNode *ListNode = root

	for node != nil {
		groupNodes := make([]*ListNode, k)
		groupNode := node
		groupIndex := 0

		for groupNode != nil && groupIndex < k {
			groupNodes[groupIndex] = groupNode
			groupNode = groupNode.Next
			groupIndex++
		}

		// break if the number of left-out nodes is less than k
		if groupIndex < k {
			break
		}

		// swap
		for i := 0; i < k; i++ {
			groupNode = groupNodes[i]

			if i == k-1 {
				// last node of the group: link prev node to the node
				prevNode.Next = groupNode
			}

			if i == 0 {
				// first node of the group: link the node to next node of new group
				groupNode.Next = groupNodes[k-1].Next
			} else {
				// default: reverse link
				groupNode.Next = groupNodes[i-1]
			}
		}

		// update for next iteration
		prevNode = groupNodes[0]
		node = prevNode.Next
	}

	return root.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
