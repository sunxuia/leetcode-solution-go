package q022_generate_parentheses

import (
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {
	doTest(t, generateParenthesis)
}

func generateParenthesis(n int) []string {
	res := make([]string, 0)
	if n == 0 {
		return res
	}

	sb := make([]byte, n*2)
	recursion(&res, sb, 0, 0)
	return res
}

func recursion(res *[]string, sb []byte, open int, close int) {
	if close+open == len(sb) {
		*res = append(*res, string(sb))
		return
	}
	if open < len(sb)/2 {
		sb[open+close] = '('
		recursion(res, sb, open+1, close)
	}
	if open > close {
		sb[open+close] = ')'
		recursion(res, sb, open, close+1)
	}
}
