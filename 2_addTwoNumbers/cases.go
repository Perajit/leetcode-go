package addtwonumbers

type AddTwoNumbersTestCase struct {
	name           string
	l1             *ListNode
	l2             *ListNode
	expectedOutput *ListNode
}

var addTwoNumbersTestCases = []AddTwoNumbersTestCase{
	{
		"#1",
		&ListNode{},
		&ListNode{},
		&ListNode{},
	},
	{
		"#2",
		&ListNode{2, &ListNode{4, &ListNode{3, nil}}},
		&ListNode{5, &ListNode{6, &ListNode{4, nil}}},
		&ListNode{7, &ListNode{0, &ListNode{8, nil}}},
	},
	{
		"#3",
		&ListNode{9, &ListNode{5, &ListNode{1, nil}}},
		&ListNode{4, &ListNode{8, nil}},
		&ListNode{3, &ListNode{4, &ListNode{2, nil}}},
	},
	{
		"#4",
		&ListNode{5, nil},
		&ListNode{5, nil},
		&ListNode{0, &ListNode{1, nil}},
	},
	{
		"#5",
		nil,
		&ListNode{1, nil},
		&ListNode{1, nil},
	},
	{
		"#6",
		nil,
		nil,
		nil,
	},
}
