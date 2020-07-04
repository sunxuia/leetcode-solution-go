// [Medium] 4Sum
// https://leetcode.com/problems/4sum/
//
// Given an array nums of n integers and an integer target, are there elements a, b, c, and d in nums such that a + b + c + d = target? Find all unique quadruplets in the array which gives the sum of target.
//
// Note:
//
// The solution set must not contain duplicate quadruplets.
//
// Example:
//
//
// Given array nums = [1, 0, -1, 0, -2, 2], and target = 0.
//
// A solution set is:
// [
//   [-1,  0, 0, 1],
//   [-2, -1, 1, 2],
//   [-2,  0, 0, 2]
// ]
package q018_4sum

import (
	"testing"

	"github.com/sunxuia/leetcode-solution-go/src/util/test"
)

func doTest(t *testing.T, method func([]int, int) [][]int) {
	th := test.NewTestHelper(t).AssertConfig(test.NewAssertConfig().UnOrder("**"))
	var res [][]int

	defer th.NewTestCase()()
	res = method([]int{1, 0, -1, 0, -2, 2}, 0)
	th.AssertEqual([][]int{
		{-1, 0, 0, 1},
		{-2, -1, 1, 2},
		{-2, 0, 0, 2}}, res)

	defer th.NewTestCase()()
	res = method([]int{0, 0, 0, 0}, 0)
	th.AssertEqual([][]int{{0, 0, 0, 0}}, res)
}
