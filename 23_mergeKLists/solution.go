package mergeklists

import "math"

func mergeKLists(lists []*ListNode) *ListNode {
	N := len(lists)

	if N == 0 {
		return nil
	}

	var out *ListNode = &ListNode{}
	var node *ListNode = out

	for {
		minVal := math.MaxInt
		minIndex := -1

		for i := 0; i < N; i++ {
			iNode := lists[i]

			if iNode != nil && iNode.Val < minVal {
				minVal = iNode.Val
				minIndex = i
			}
		}

		if minIndex == -1 {
			break
		}

		lists[minIndex] = lists[minIndex].Next
		node.Next = &ListNode{minVal, nil}
		node = node.Next
	}

	return out.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
