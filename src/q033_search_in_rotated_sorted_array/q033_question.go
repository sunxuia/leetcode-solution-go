// [Medium] Search in Rotated Sorted Array
// https://leetcode.com/problems/search-in-rotated-sorted-array/
//
// Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.
//
// (i.e., [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2]).
//
// You are given a target value to search. If found in the array return its index, otherwise return -1.
//
// You may assume no duplicate exists in the array.
//
// Your algorithm's runtime complexity must be in the order of O(log n).
//
// Example 1:
//
//
// Input: nums = [4,5,6,7,0,1,2], target = 0
// Output: 4
//
//
// Example 2:
//
//
// Input: nums = [4,5,6,7,0,1,2], target = 3
// Output: -1
package q033_search_in_rotated_sorted_array

import (
	"github.com/sunxuia/leetcode-solution-go/src/util/test"
	"testing"
)

func doTest(t *testing.T, method func([]int, int) int) {
	th := test.NewTestHelper(t)
	var res int

	defer th.NewTestCase()()
	res = method([]int{4, 5, 6, 7, 0, 1, 2}, 0)
	th.Expect(4).Assert(res)

	defer th.NewTestCase()()
	res = method([]int{4, 5, 6, 7, 0, 1, 2}, 3)
	th.Expect(-1).Assert(res)

	defer th.NewTestCase()()
	res = method([]int{}, 5)
	th.Expect(-1).Assert(res)

	defer th.NewTestCase()()
	res = method([]int{1}, 1)
	th.Expect(0).Assert(res)
}
