// Package q005_longest_palindromic_substring [Medium] Longest Palindromic Substring
// https://leetcode.com/problems/longest-palindromic-substring/
//
// Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.
//
// Example 1:
//
//
// Input: "babad"
// Output: "bab"
// Note: "aba" is also a valid answer.
//
//
// Example 2:
//
//
// Input: "cbbd"
// Output: "bb"
package q005_longest_palindromic_substring

import (
	"testing"

	"github.com/sunxuia/leetcode-solution-go/src/util/test"
)

func doTest(t *testing.T, method func(string) string) {
	ah := test.NewAssertHelper()
	var res string

	res = method("babad")
	ah.AssertInRange(res, "aba", "bab")

	res = method("cbbd")
	ah.Assert("bb", res)
}
