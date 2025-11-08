package mergetwolists

type MergeTwoListsTestCase struct {
	name           string
	list1          *ListNode
	list2          *ListNode
	expectedOutput *ListNode
}

var mergeTwoListsTestCases = []MergeTwoListsTestCase{
	{
		"#1",
		// [1,2,4]
		&ListNode{1, &ListNode{2, &ListNode{4, nil}}},
		// [1,3,4]
		&ListNode{1, &ListNode{3, &ListNode{4, nil}}},
		// [1,1,2,3,4,4]
		&ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{4, nil}}}}}},
	},
	{
		"#2",
		// [1,2,3]
		&ListNode{1, &ListNode{2, &ListNode{3, nil}}},
		// [2,3,4]
		&ListNode{2, &ListNode{3, &ListNode{4, nil}}},
		// [1,2,2,3,3,4]
		&ListNode{1, &ListNode{2, &ListNode{2, &ListNode{3, &ListNode{3, &ListNode{4, nil}}}}}},
	},
	{
		"#3",
		// [1,4,5]
		&ListNode{1, &ListNode{4, &ListNode{5, nil}}},
		// [2,3,4]
		&ListNode{2, &ListNode{3, &ListNode{4, nil}}},
		// [1,2,3,4,4,5]
		&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{4, &ListNode{5, nil}}}}}},
	},
	{
		"#4",
		// [1,4,5]
		&ListNode{1, &ListNode{4, &ListNode{5, nil}}},
		// [2,3]
		&ListNode{2, &ListNode{3, nil}},
		// [1,2,3,4,5]
		&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
	},
	{
		"#5",
		// [1,4]
		&ListNode{1, &ListNode{4, nil}},
		// [2,3,5]
		&ListNode{2, &ListNode{3, &ListNode{5, nil}}},
		// [1,2,3,4,5]
		&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
	},
	{
		"#6",
		&ListNode{1, nil},
		nil,
		&ListNode{1, nil},
	},
	{
		"#7",
		nil,
		&ListNode{2, nil},
		&ListNode{2, nil},
	},
	{
		"#8",
		nil,
		nil,
		nil,
	},
}
