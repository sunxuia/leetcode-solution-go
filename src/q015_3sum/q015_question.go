// [Medium] 3Sum
// https://leetcode.com/problems/3sum/
//
// Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.
//
// Note:
//
// The solution set must not contain duplicate triplets.
//
// Example:
//
//
// Given array nums = [-1, 0, 1, 2, -1, -4],
//
// A solution set is:
// [
//   [-1, 0, 1],
//   [-1, -1, 2]
// ]
package q015_3sum

import (
	"testing"

	"github.com/sunxuia/leetcode-solution-go/src/util/test"
)

func doTest(t *testing.T, method func([]int) [][]int) {
	th := test.NewTestHelper(t)
	th.AssertConfig(test.NewAssertConfig().UnOrder("**"))
	var res [][]int

	defer th.NewTestCase()()
	res = method([]int{-1, 0, 1, 2, -1, -4})
	th.AssertEqual([][]int{{-1, 0, 1}, {-1, -1, 2}}, res)
}
