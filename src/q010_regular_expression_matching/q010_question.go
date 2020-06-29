// Package q010_regular_expression_matching [Hard] Regular Expression Matching
// https://leetcode.com/problems/regular-expression-matching/
//
// Given an input string (s) and a pattern (p), implement regular expression matching with support for '.' and '*'.
//
//
// '.' Matches any single character.
// '*' Matches zero or more of the preceding element.
//
//
// The matching should cover the entire input string (not partial).
//
// Note:
//
//
// 	s could be empty and contains only lowercase letters a-z.
// 	p could be empty and contains only lowercase letters a-z, and characters like . or *.
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
// p = "a*"
// Output: true
// Explanation: '*' means zero or more of the preceding element, 'a'. Therefore, by repeating 'a' once, it becomes "aa".
//
//
// Example 3:
//
//
// Input:
// s = "ab"
// p = ".*"
// Output: true
// Explanation: ".*" means "zero or more (*) of any character (.)".
//
//
// Example 4:
//
//
// Input:
// s = "aab"
// p = "c*a*b"
// Output: true
// Explanation: c can be repeated 0 times, a can be repeated 1 time. Therefore, it matches "aab".
//
//
// Example 5:
//
//
// Input:
// s = "mississippi"
// p = "mis*is*p*."
// Output: false
package q010_regular_expression_matching

import (
	"testing"

	"github.com/sunxuia/leetcode-solution-go/src/util/test"
)

func doTest(t *testing.T, method func(string, string) bool) {
	ah := test.NewAssertHelper()
	var res bool

	res = method("aa", "a")
	ah.Assert(false, res)

	res = method("aa", "a*")
	ah.Assert(true, res)

	res = method("ab", ".*")
	ah.Assert(true, res)

	res = method("aab", "c*a*b")
	ah.Assert(true, res)

	res = method("mississippi", "mis*is*p*.")
	ah.Assert(false, res)

	res = method("aaa", "ab*a*c*a")
	ah.Assert(true, res)

	res = method("a", "ab*")
	ah.Assert(true, res)

	res = method("abcdede", "ab.*de")
	ah.Assert(true, res)

	res = method("", "c*c*")
	ah.Assert(true, res)

	res = method("aaca", "ab*a*c*a")
	ah.Assert(true, res)

	res = method("cbaacacaaccbaabcb", "c*b*b*.*ac*.*bc*a*")
	ah.Assert(true, res)
}
