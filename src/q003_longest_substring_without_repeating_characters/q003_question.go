// Package q003_longest_substring_without_repeating_characters [Medium] Longest Substring Without Repeating Characters
// https://leetcode.com/problems/longest-substring-without-repeating-characters/
//
// Given a string, find the length of the longest substring without repeating characters.
//
// Example 1:
//
// Input: "abcabcbb"
// Output: 3
// Explanation: The answer is "abc", with the length of 3.
//
// Example 2:
//
// Input: "bbbbb"
// Output: 1
// Explanation: The answer is "b", with the length of 1.
//
// Example 3:
//
// Input: "pwwkew"
// Output: 3
// Explanation: The answer is "wke", with the length of 3.
//             Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
package q003_longest_substring_without_repeating_characters

import (
	"testing"

	"github.com/sunxuia/leetcode-solution-go/src/util/test"
)

func doTest(t *testing.T, method func(string) int) {
	ah := test.NewAssertHelper()
	var res int

	res = method("abcabcbb")
	ah.Assert(3, res)

	res = method("bbbbb")
	ah.Assert(1, res)

	res = method("pwwkew")
	ah.Assert(3, res)

	res = method("cdd")
	ah.Assert(2, res)
}
