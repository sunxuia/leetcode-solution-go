// [Easy] Reverse Integer
// https://leetcode.com/problems/reverse-integer/
//
// Given a 32-bit signed integer, reverse digits of an integer.
//
// Example 1:
//
//
// Input: 123
// Output: 321
//
//
// Example 2:
//
//
// Input: -123
// Output: -321
//
//
// Example 3:
//
//
// Input: 120
// Output: 21
//
//
// Note:
// Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−231,  231 − 1]. For the purpose of this problem, assume that your function returns 0 when the reversed integer overflows.
package q007_reverse_integer

import (
	"testing"

	"github.com/sunxuia/leetcode-solution-go/src/util/test"
)

func doTest(t *testing.T, method func(int) int) {
	th := test.NewTestHelper(t)
	var res int

	defer th.NewTestCase()()
	res = method(123)
	th.Expect(321).Assert(res)

	defer th.NewTestCase()()
	res = method(-123)
	th.Expect(-321).Assert(res)

	defer th.NewTestCase()()
	res = method(120)
	th.Expect(21).Assert(res)

	defer th.NewTestCase()()
	// 结果超过整数范围则返回0
	res = method(1534236469)
	th.Expect(0).Assert(res)
}
