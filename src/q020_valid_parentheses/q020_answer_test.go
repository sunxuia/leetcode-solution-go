package q020_valid_parentheses

import (
	"testing"
)

func TestIsValid(t *testing.T) {
	doTest(t, isValid)
}

func isValid(s string) bool {
	stack := make([]rune, 0)
	for _, c := range s {
		if p, ok := pairs[c]; ok {
			stack = append(stack, p)
		} else {
			size := len(stack)
			if size == 0 || stack[size-1] != c {
				return false
			}
			stack = stack[:size-1]
		}
	}
	return len(stack) == 0
}

var pairs = map[rune]rune{'(': ')', '[': ']', '{': '}'}
