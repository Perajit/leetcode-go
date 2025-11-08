package reversekgroup

type ReverseKGroupTestCase struct {
	name           string
	head           *ListNode
	k              int
	expectedOutput *ListNode
}

var reverseKGroupTestCases = []ReverseKGroupTestCase{
	{
		"#1",
		// [1,2,3,4,5]
		&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
		2,
		// [2,1,4,3,5]
		&ListNode{2, &ListNode{1, &ListNode{4, &ListNode{3, &ListNode{5, nil}}}}},
	},
	{
		"#2",
		// [1,2,3,4,5]
		&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
		3,
		// [3,2,1,4.5]
		&ListNode{3, &ListNode{2, &ListNode{1, &ListNode{4, &ListNode{5, nil}}}}},
	},
	{
		"#3",
		// [1,2,3,4,5,6,7,8]
		&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, &ListNode{7, &ListNode{8, nil}}}}}}}},
		4,
		// [4,3,2,1,8,7,6,5]
		&ListNode{4, &ListNode{3, &ListNode{2, &ListNode{1, &ListNode{8, &ListNode{7, &ListNode{6, &ListNode{5, nil}}}}}}}},
	},
	//[1,2,3,4]	=>	[4,3,2,1]
	//[1,2] 		=>	[2,1]
}
