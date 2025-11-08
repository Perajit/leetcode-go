package swappairs

type SwapPairsTestCase struct {
	name           string
	head           *ListNode
	expectedOutput *ListNode
}

var swapPairsTestCases = []SwapPairsTestCase{
	{
		"#1",
		&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}},
		&ListNode{2, &ListNode{1, &ListNode{4, &ListNode{3, nil}}}},
	},
	{
		"#2",
		nil,
		nil,
	},
	{
		"#3",
		&ListNode{1, nil},
		&ListNode{1, nil},
	},
	{
		"#4",
		&ListNode{1, &ListNode{2, &ListNode{3, nil}}},
		&ListNode{2, &ListNode{1, &ListNode{3, nil}}},
	},
}
