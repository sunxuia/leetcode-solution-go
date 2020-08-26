// [Medium] Combination Sum II
// https://leetcode.com/problems/combination-sum-ii/
//
// Given a collection of candidate numbers (candidates) and a target number (target), find all unique combinations in candidates where the candidate numbers sums to target.
//
// Each number in candidates may only be used once in the combination.
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
// Input: candidates = [10,1,2,7,6,1,5], target = 8,
// A solution set is:
// [
//   [1, 7],
//   [1, 2, 5],
//   [2, 6],
//   [1, 1, 6]
// ]
//
//
// Example 2:
//
//
// Input: candidates = [2,5,2,1,2], target = 5,
// A solution set is:
// [
//   [1,2,2],
//   [5]
// ]
// 题解: 和上一题相比, 每个数字的使用次数变成1 次, 且 candidates 中有重复值
package q040_combination_sum_ii

import (
	"github.com/sunxuia/leetcode-solution-go/src/util/test"
	"testing"
)

func doTest(t *testing.T, method func([]int, int) [][]int) {
	th := test.NewTestHelper(t)
	cfg := test.NewAssertConfig().UnOrder("*")
	var res [][]int

	defer th.NewTestCase()()
	res = method([]int{10, 1, 2, 7, 6, 1, 5}, 8)
	th.Expect([][]int{
		{1, 7},
		{1, 2, 5},
		{2, 6},
		{1, 1, 6},
	}).Config(cfg).Assert(res)

	defer th.NewTestCase()()
	res = method([]int{2, 5, 2, 1, 2}, 5)
	th.Expect([][]int{
		{1, 2, 2},
		{5},
	}).Config(cfg).Assert(res)

	defer th.NewTestCase()()
	res = method([]int{1, 1}, 2)
	th.Expect([][]int{{1, 1}}).Config(cfg).Assert(res)

	defer th.NewTestCase()()
	res = method([]int{2, 2, 2}, 2)
	th.Expect([][]int{{2}}).Config(cfg).Assert(res)

}
