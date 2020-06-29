// Package q011_container_with_most_water [Medium] Container With Most Water
// https://leetcode.com/problems/container-with-most-water/
//
// Given n non-negative integers a1, a2, ..., an , where each represents a point at coordinate (i, ai). n vertical lines are drawn such that the two endpoints of line i is at (i, ai) and (i, 0). Find two lines, which together with x-axis forms a container, such that the container contains the most water.
//
// Note: You may not slant the container and n is at least 2.
//
//
//
// (图 q011_pic.jpg)
//
// The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.
//
//
//
// Example:
//
//
// Input: [1,8,6,2,5,4,8,3,7]
// Output: 49
package q011_container_with_most_water

import (
	"testing"

	"github.com/sunxuia/leetcode-solution-go/src/util/test"
)

func doTest(t *testing.T, method func([]int) int) {
	ah := test.NewAssertHelper()
	var res int

	res = method([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	ah.Assert(49, res)
}