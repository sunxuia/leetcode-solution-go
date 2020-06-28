// Package q001_two_sum [Easy] Two Sum
// https://leetcode.com/problems/two-sum/
//
// Given an array of integers, return indices of the two numbers such that they add up to a specific target.
//
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
//
// Example:
//
//
// Given nums = [2, 7, 11, 15], target = 9,
//
// Because nums[0] + nums[1] = 2 + 7 = 9,
// return [0, 1].
//
//
package q001_two_sum

import (
	"testing"

	"github.com/sunxuia/leetcode-solution-go/src/util/test"
)

func doTest(t *testing.T, method func([]int, int) []int) {
	ah := test.NewAssertHelper()
	var res []int

	res = method([]int{2, 7, 11, 15}, 9)
	ah.Assert([]int{0, 1}, res)
}
