// [Easy] Implement strStr()
// https://leetcode.com/problems/implement-strstr/
//
// Implement strStr().
//
// Return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.
//
// Example 1:
//
//
// Input: haystack = "hello", needle = "ll"
// Output: 2
//
//
// Example 2:
//
//
// Input: haystack = "aaaaa", needle = "bba"
// Output: -1
//
//
// Clarification:
//
// What should we return when needle is an empty string? This is a great question to ask during an interview.
//
// For the purpose of this problem, we will return 0 when needle is an empty string. This is consistent to C's strstr() and Java's indexOf().
package q028_implement_strstr

import (
	"github.com/sunxuia/leetcode-solution-go/src/util/test"
	"testing"
)

func doTest(t *testing.T, method func(string, string) int) {
	th := test.NewTestHelper(t)
	var res int

	defer th.NewTestCase()()
	res = method("hello", "ll")
	th.Expect(2).Assert(res)

	defer th.NewTestCase()()
	res = method("aaaaa", "bba")
	th.Expect(-1).Assert(res)

	defer th.NewTestCase()()
	res = method("", "")
	th.Expect(0).Assert(res)

	defer th.NewTestCase()()
	res = method("mississippi", "issipi")
	th.Expect(-1).Assert(res)
}
