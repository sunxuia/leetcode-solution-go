// [Easy] Search Insert Position
// https://leetcode.com/problems/search-insert-position/
//
// Given a sorted array and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.
//
// You may assume no duplicates in the array.
//
// Example 1:
//
//
// Input: [1,3,5,6], 5
// Output: 2
//
//
// Example 2:
//
//
// Input: [1,3,5,6], 2
// Output: 1
//
//
// Example 3:
//
//
// Input: [1,3,5,6], 7
// Output: 4
//
//
// Example 4:
//
//
// Input: [1,3,5,6], 0
// Output: 0
package q035_search_insert_position

import (
	"github.com/sunxuia/leetcode-solution-go/src/util/test"
	"testing"
)

func doTest(t *testing.T, method func([]int, int) int) {
	th := test.NewTestHelper(t)
	var res int

	defer th.NewTestCase()()
	res = method([]int{1, 3, 5, 6}, 5)
	th.Expect(2).Assert(res)

	defer th.NewTestCase()()
	res = method([]int{1, 3, 5, 6}, 2)
	th.Expect(1).Assert(res)

	defer th.NewTestCase()()
	res = method([]int{1, 3, 5, 6}, 7)
	th.Expect(4).Assert(res)

	defer th.NewTestCase()()
	res = method([]int{1, 3, 5, 6}, 0)
	th.Expect(0).Assert(res)
}
