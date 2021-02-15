// [Hard] Wildcard Matching
// https://leetcode.com/problems/wildcard-matching/
//
// Given an input string (s) and a pattern (p), implement wildcard pattern matching with support for '?' and '*'.
//
//
// '?' Matches any single character.
// '*' Matches any sequence of characters (including the empty sequence).
//
//
// The matching should cover the entire input string (not partial).
//
// Note:
//
//
// 	s could be empty and contains only lowercase letters a-z.
// 	p could be empty and contains only lowercase letters a-z, and characters like ? or *.
//
//
// Example 1:
//
//
// Input:
// s = "aa"
// p = "a"
// Output: false
// Explanation: "a" does not match the entire string "aa".
//
//
// Example 2:
//
//
// Input:
// s = "aa"
// p = "*"
// Output: true
// Explanation: '*' matches any sequence.
//
//
// Example 3:
//
//
// Input:
// s = "cb"
// p = "?a"
// Output: false
// Explanation: '?' matches 'c', but the second letter is 'a', which does not match 'b'.
//
//
// Example 4:
//
//
// Input:
// s = "adceb"
// p = "*a*b"
// Output: true
// Explanation: The first '*' matches the empty sequence, while the second '*' matches the substring "dce".
//
//
// Example 5:
//
//
// Input:
// s = "acdcb"
// p = "a*c?b"
// Output: false
package q044_wildcard_matching

import (
	"github.com/sunxuia/leetcode-solution-go/src/util/test"
	"testing"
)

func doTest(t *testing.T, method func(string, string) bool) {
	th := test.NewTestHelper(t)
	var res bool

	defer th.NewTestCase()()
	res = method("aa", "a")
	th.Expect(false).Assert(res)

	defer th.NewTestCase()()
	res = method("aa", "*")
	th.Expect(true).Assert(res)

	defer th.NewTestCase()()
	res = method("cb", "?a")
	th.Expect(false).Assert(res)

	defer th.NewTestCase()()
	res = method("adceb", "*a*b")
	th.Expect(true).Assert(res)

	defer th.NewTestCase()()
	res = method("acdcb", "a*c?b")
	th.Expect(false).Assert(res)
}
