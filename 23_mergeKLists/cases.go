package mergeklists

type MergeKListsTestCase struct {
	name           string
	lists          []*ListNode
	expectedOutput *ListNode
}

var mergeKListsTestCases = []MergeKListsTestCase{
	{
		"#1",
		// [1,4,5],[1,3,4],[2,6]
		[]*ListNode{
			{1, &ListNode{4, &ListNode{5, nil}}},
			{1, &ListNode{3, &ListNode{4, nil}}},
			{2, &ListNode{6, nil}},
		},
		// [1,1,2,3,4,4,5,6]
		&ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{4, &ListNode{5, &ListNode{6, nil}}}}}}}},
	},
	{
		"#2",
		[]*ListNode{{1, &ListNode{4, &ListNode{5, nil}}}},
		&ListNode{1, &ListNode{4, &ListNode{5, nil}}},
	},
	{
		"#3",
		[]*ListNode{},
		nil,
	},
	{
		"#4",
		[]*ListNode{nil},
		nil,
	},
}
