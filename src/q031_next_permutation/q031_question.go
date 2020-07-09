// [Medium] Next Permutation
// https://leetcode.com/problems/next-permutation/
//
// Implement next permutation, which rearranges numbers into the lexicographically next greater permutation of numbers.
//
// If such arrangement is not possible, it must rearrange it as the lowest possible order (ie, sorted in ascending order).
//
// The replacement must be in-place and use only constant extra memory.
//
// Here are some examples. Inputs are in the left-hand column and its corresponding outputs are in the right-hand column.
//
// 1,2,3 → 1,3,2
// 3,2,1 → 1,2,3
// 1,1,5 → 1,5,1
package q031_next_permutation

import (
	"github.com/sunxuia/leetcode-solution-go/src/util/test"
	"testing"
)

func doTest(t *testing.T, method func([]int)) {
	th := test.NewTestHelper(t)
	var res []int

	defer th.NewTestCase()()
	res = []int{1, 2, 3}
	method(res)
	th.Expect([]int{1, 3, 2}).Assert(res)

	defer th.NewTestCase()()
	res = []int{3, 2, 1}
	method(res)
	th.Expect([]int{1, 2, 3}).Assert(res)

	defer th.NewTestCase()()
	res = []int{1, 1, 5}
	method(res)
	th.Expect([]int{1, 5, 1}).Assert(res)

	defer th.NewTestCase()()
	res = []int{1, 3, 2}
	method(res)
	th.Expect([]int{2, 1, 3}).Assert(res)

	defer th.NewTestCase()()
	res = []int{1, 5, 1}
	method(res)
	th.Expect([]int{5, 1, 1}).Assert(res)
}
