// [Hard] Median of Two Sorted Arrays
// https://leetcode.com/problems/median-of-two-sorted-arrays/
//
// There are two sorted arrays nums1 and nums2 of size m and n respectively.
//
// Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).
//
// You may assume nums1 and nums2 cannot be both empty.
//
// Example 1:
//
//
// nums1 = [1, 3]
// nums2 = [2]
//
// The median is 2.0
//
//
// Example 2:
//
//
// nums1 = [1, 2]
// nums2 = [3, 4]
//
// The median is (2 + 3)/2 = 2.5
package q004_median_of_two_sorted_arrays

import (
	"testing"

	"github.com/sunxuia/leetcode-solution-go/src/util/test"
)

func doTest(t *testing.T, method func([]int, []int) float64) {
	th := test.NewTestHelper(t)
	var res float64

	defer th.NewTestCase()()
	res = method([]int{1, 3}, []int{2})
	th.Expect(2.0).Assert(res)

	defer th.NewTestCase()()
	res = method([]int{1, 2}, []int{3, 4})
	th.Expect(2.5).Assert(res)
}
