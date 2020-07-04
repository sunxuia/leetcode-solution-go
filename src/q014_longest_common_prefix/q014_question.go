// [Easy] Longest Common Prefix
// https://leetcode.com/problems/longest-common-prefix/
//
// Write a function to find the longest common prefix string amongst an array of strings.
//
// If there is no common prefix, return an empty string "".
//
// Example 1:
//
//
// Input: ["flower","flow","flight"]
// Output: "fl"
//
//
// Example 2:
//
//
// Input: ["dog","racecar","car"]
// Output: ""
// Explanation: There is no common prefix among the input strings.
//
//
// Note:
//
// All given inputs are in lowercase letters a-z.
package q014_longest_common_prefix

import (
	"testing"

	"github.com/sunxuia/leetcode-solution-go/src/util/test"
)

func doTest(t *testing.T, method func([]string) string) {
	th := test.NewTestHelper(t)
	var res string

	defer th.NewTestCase()()
	res = method([]string{"flower", "flow", "flight"})
	th.AssertEqual("fl", res)

	defer th.NewTestCase()()
	res = method([]string{"dog", "racecar", "car"})
	th.AssertEqual("", res)

	defer th.NewTestCase()()
	res = method([]string{"a"})
	th.AssertEqual("a", res)

	defer th.NewTestCase()()
	res = method([]string{"a", "ac"})
	th.AssertEqual("a", res)
}
