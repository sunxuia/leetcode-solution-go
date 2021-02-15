// [Hard] Trapping Rain Water
// https://leetcode.com/problems/trapping-rain-water/
//
// Given n non-negative integers representing an elevation map where the width of each bar is 1, compute how much water it is able to trap after raining.
//
// <img src="Q042_PIC.png">
// The above elevation map is represented by array [0,1,0,2,1,0,1,3,2,1,2,1]. In this case, 6 units of rain water (blue section) are being trapped. Thanks Marcos for contributing this image!
//
// Example:
//
//
// Input: [0,1,0,2,1,0,1,3,2,1,2,1]
// Output: 6
package q042_trapping_rain_water

import (
	"github.com/sunxuia/leetcode-solution-go/src/util/test"
	"testing"
)

func doTest(t *testing.T, method func([]int) int) {
	th := test.NewTestHelper(t)
	var res int

	defer th.NewTestCase()()
	res = method([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})
	th.Expect(6).Assert(res)

	defer th.NewTestCase()()
	res = method([]int{4, 2, 3})
	th.Expect(1).Assert(res)
}
