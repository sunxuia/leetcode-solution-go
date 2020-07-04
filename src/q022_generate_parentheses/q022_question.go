// [Medium] Generate Parentheses
// https://leetcode.com/problems/generate-parentheses/
//
// Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.
//
//
//
// For example, given n = 3, a solution set is:
//
//
// [
//   "((()))",
//   "(()())",
//   "(())()",
//   "()(())",
//   "()()()"
// ]
package q022_generate_parentheses

import (
	"testing"

	"github.com/sunxuia/leetcode-solution-go/src/util/test"
)

func doTest(t *testing.T, method func(int) []string) {
	th := test.NewTestHelper(t)
	th.AssertConfig(test.NewAssertConfig().UnOrder("*"))
	var res []string

	defer th.NewTestCase()()
	res = method(3)
	th.AssertEqual([]string{"((()))",
		"(()())",
		"(())()",
		"()(())",
		"()()()"}, res)
}
