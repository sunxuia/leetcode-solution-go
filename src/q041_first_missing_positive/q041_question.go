// [Hard] First Missing Positive
// https://leetcode.com/problems/first-missing-positive/
//
// Given an unsorted integer array, find the smallest missing positive integer.
//
// Example 1:
//
//
// Input: [1,2,0]
// Output: 3
//
//
// Example 2:
//
//
// Input: [3,4,-1,1]
// Output: 2
//
//
// Example 3:
//
//
// Input: [7,8,9,11,12]
// Output: 1
//
//
// Follow up:
//
// Your algorithm should run in O(n) time and uses constant extra space.
package q041_first_missing_positive

import (
	"github.com/sunxuia/leetcode-solution-go/src/util/test"
	"testing"
)

func doTest(t *testing.T, method func([]int) int) {
	th := test.NewTestHelper(t)
	var res int

	defer th.NewTestCase()()
	res = method([]int{1, 2, 0})
	th.Expect(3).Assert(res)

	defer th.NewTestCase()()
	res = method([]int{3, 4, -1, 1})
	th.Expect(2).Assert(res)

	defer th.NewTestCase()()
	res = method([]int{7, 8, 9, 11, 12})
	th.Expect(1).Assert(res)
}
