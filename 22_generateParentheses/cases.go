package generateparentheses

type GenerateParenthesisTestCase struct {
	name           string
	n              int
	expectedOutput []string
}

var generateParenthesisTestCases = []GenerateParenthesisTestCase{
	{"#1", 1, []string{"()"}},
	{"#2", 2, []string{"(())", "()()"}},
	{"#3", 3, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}},
	{"#4", 4, []string{
		"(((())))", "((()()))", "((())())", "(()(()))", "(()()())",
		"(())(())", "((()))()", "()((()))", "(()())()", "()(()())",
		"(())()()", "()(())()", "()()(())",
		"()()()()",
	}},
}
