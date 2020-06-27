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

import "../util/test"

var runner = test.NewMethodTestRunner().AddTestCase(
	test.NewMethodTestCase("abcabcbb").ExpectReturn(3)).AddTestCase(
	test.NewMethodTestCase("bbbbb").ExpectReturn(1)).AddTestCase(
	test.NewMethodTestCase("pwwkew").ExpectReturn(3)).AddTestCase(
	test.NewMethodTestCase("cdd").ExpectReturn(2))
