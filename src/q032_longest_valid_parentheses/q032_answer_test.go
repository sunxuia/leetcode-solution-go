package q032_longest_valid_parentheses

import (
	"testing"
)

func TestLongestValidParentheses(t *testing.T) {
	doTest(t, longestValidParentheses)
}

func longestValidParentheses(s string) int {
	res := 0
	stack := make([]int, 0)
	// 栈底表示上一个无法匹配的 ")" 元素位置.
	// 其余元素表示未匹配的 "(" 位置.
	stack = append(stack, -1)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else if len(stack) == 1 {
			stack[0] = i
		} else {
			prev := stack[len(stack)-2]
			if res < i-prev {
				res = i - prev
			}
			stack = stack[:len(stack)-1]
		}
	}
	return res
}
