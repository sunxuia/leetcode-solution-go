// [Medium] Combination Sum
// https://leetcode.com/problems/combination-sum/
//
// Given a set of candidate numbers (candidates) (without duplicates) and a target number (target), find all unique combinations in candidates where the candidate numbers sums to target.
//
// The same repeated number may be chosen from candidates unlimited number of times.
//
// Note:
//
//
// 	All numbers (including target) will be positive integers.
// 	The solution set must not contain duplicate combinations.
//
//
// Example 1:
//
//
// Input: candidates = [2,3,6,7], target = 7,
// A solution set is:
// [
//   [7],
//   [2,2,3]
// ]
//
//
// Example 2:
//
//
// Input: candidates = [2,3,5], target = 8,
// A solution set is:
// [
//   [2,2,2,2],
//   [2,3,3],
//   [3,5]
// ]
package q039_combination_sum

import (
	"github.com/sunxuia/leetcode-solution-go/src/util/test"
	"testing"
)

func doTest(t *testing.T, method func([]int, int) [][]int) {
	th := test.NewTestHelper(t)
	cfg := test.NewAssertConfig().UnOrder("*")
	var res [][]int

	defer th.NewTestCase()()
	res = method([]int{2, 3, 6, 7}, 7)
	th.Expect([][]int{{7}, {2, 2, 3}}).Config(cfg).Assert(res)

	defer th.NewTestCase()()
	res = method([]int{2, 3, 5}, 8)
	th.Expect([][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}).Config(cfg).Assert(res)

	defer th.NewTestCase()()
	res = method([]int{7, 3, 2}, 18)
	th.Expect([][]int{
		{2, 2, 2, 2, 2, 2, 2, 2, 2},
		{2, 2, 2, 2, 2, 2, 3, 3},
		{2, 2, 2, 2, 3, 7},
		{2, 2, 2, 3, 3, 3, 3},
		{2, 2, 7, 7},
		{2, 3, 3, 3, 7},
		{3, 3, 3, 3, 3, 3},
	}).Config(cfg).Assert(res)

}
