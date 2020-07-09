// [Hard] Longest Valid Parentheses
// https://leetcode.com/problems/longest-valid-parentheses/
//
// Given a string containing just the characters '(' and ')', find the length of the longest valid (well-formed) parentheses substring.
//
// Example 1:
//
//
// Input: "(()"
// Output: 2
// Explanation: The longest valid parentheses substring is "()"
//
//
// Example 2:
//
//
// Input: ")()())"
// Output: 4
// Explanation: The longest valid parentheses substring is "()()"
package q032_longest_valid_parentheses

import (
	"github.com/sunxuia/leetcode-solution-go/src/util/test"
	"testing"
)

func doTest(t *testing.T, method func(string) int) {
	th := test.NewTestHelper(t)
	var res int

	defer th.NewTestCase()()
	res = method("(()")
	th.Expect(2).Assert(res)

	defer th.NewTestCase()()
	res = method(")()())")
	th.Expect(4).Assert(res)
}
