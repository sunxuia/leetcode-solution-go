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

import "testing"

func lengthOfLongestSubstring(s string) int {
	occupied := [128]int{}
	res, last := 0, 0
	for i, c := range s {
		if occupied[c] > last {
			if res < i-last {
				res = i - last
			}
			last = occupied[c]
		}
		occupied[c] = i + 1
	}
	if res < len(s)-last {
		res = len(s) - last
	}
	return res
}

func TestLengthOfLongestSubstring(t *testing.T) {
	runner.Test(t, lengthOfLongestSubstring)
}
