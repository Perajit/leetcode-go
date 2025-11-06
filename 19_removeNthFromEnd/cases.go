package removenthfromend

type RemoveNthFromEndTestCase struct {
	name           string
	head           *ListNode
	n              int
	expectedOutput *ListNode
}

var removeNthFromEndTestCases = []RemoveNthFromEndTestCase{
	{
		"#1",
		&ListNode{1, nil},
		1,
		nil,
	},
	{
		"#2",
		// [1,2,3,4,5]
		&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
		1,
		// [1,2,3,4]
		&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}},
	},
	{
		"#3",
		// [1,2,3,4,5]
		&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
		2,
		// [1,2,3,5]
		&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{5, nil}}}},
	},
	{
		"#4",
		// [1,2,3,4,5]
		&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
		5,
		// [2,3,4,5]
		&ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}},
	},
}
