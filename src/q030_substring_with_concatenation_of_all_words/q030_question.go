// [Hard] Substring with Concatenation of All Words
// https://leetcode.com/problems/substring-with-concatenation-of-all-words/
//
// You are given a string, s, and a list of words, words, that are all of the same length. Find all starting indices of substring(s) in s that is a concatenation of each word in words exactly once and without any intervening characters.
//
//
//
// Example 1:
//
//
// Input:
//   s = "barfoothefoobarman",
//   words = ["foo","bar"]
// Output: [0,9]
// Explanation: Substrings starting at index 0 and 9 are "barfoo" and "foobar" respectively.
// The output order does not matter, returning [9,0] is fine too.
//
//
// Example 2:
//
//
// Input:
//   s = "wordgoodgoodgoodbestword",
//   words = ["word","good","best","word"]
// Output: []
package q030_substring_with_concatenation_of_all_words

import (
	"github.com/sunxuia/leetcode-solution-go/src/util/test"
	"testing"
)

func doTest(t *testing.T, method func(string, []string) []int) {
	th := test.NewTestHelper(t)
	var res []int

	defer th.NewTestCase()()
	res = method("barfoothefoobarman", []string{"foo", "bar"})
	th.Expect([]int{0, 9}).Assert(res)

	defer th.NewTestCase()()
	res = method("wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"})
	th.Expect([]int{}).Assert(res)

	defer th.NewTestCase()()
	res = method("wordgoodgoodgoodbestword", []string{"word", "good", "best", "good"})
	th.Expect([]int{8}).Assert(res)
}
