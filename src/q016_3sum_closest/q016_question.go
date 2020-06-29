// Package q016_3sum_closest [Medium] 3Sum Closest
// https://leetcode.com/problems/3sum-closest/
//
// Given an array nums of n integers and an integer target, find three integers in nums such that the sum is closest to target. Return the sum of the three integers. You may assume that each input would have exactly one solution.
//
//
// Example 1:
//
//
// Input: nums = [-1,2,1,-4], target = 1
// Output: 2
// Explanation: The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).
//
//
//
// Constraints:
//
//
// 	3 <= nums.length <= 10^3
// 	-10^3 <= nums[i] <= 10^3
// 	-10^4 <= target <= 10^4
package q016_3sum_closest

import (
	"testing"

	"github.com/sunxuia/leetcode-solution-go/src/util/test"
)

func doTest(t *testing.T, method func([]int, int) int) {
	ah := test.NewAssertHelper()
	var res int

	res = method([]int{-1, 2, 1, -4}, 1)
	ah.Assert(2, res)
}
