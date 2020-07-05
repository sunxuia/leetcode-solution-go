// [Medium] Divide Two Integers
// https://leetcode.com/problems/divide-two-integers/
//
// Given two integers dividend and divisor, divide two integers without using multiplication, division and mod operator.
//
// Return the quotient after dividing dividend by divisor.
//
// The integer division should truncate toward zero, which means losing its fractional part. For example, truncate(8.345) = 8 and truncate(-2.7335) = -2.
//
// Example 1:
//
//
// Input: dividend = 10, divisor = 3
// Output: 3
// Explanation: 10/3 = truncate(3.33333..) = 3.
//
//
// Example 2:
//
//
// Input: dividend = 7, divisor = -3
// Output: -2
// Explanation: 7/-3 = truncate(-2.33333..) = -2.
//
//
// Note:
//
//
// 	Both dividend and divisor will be 32-bit signed integers.
// 	The divisor will never be 0.
// 	Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−2^31,  2^31 − 1]. For the purpose of this problem, assume that your function returns 2^31 − 1 when the division result overflows.
package q029_divide_two_integers

import (
	"github.com/sunxuia/leetcode-solution-go/src/util/test"
	"testing"
)

func doTest(t *testing.T, method func(int, int) int) {
	th := test.NewTestHelper(t)
	var res int

	defer th.NewTestCase()()
	res = method(10, 3)
	th.Expect(3).Assert(res)

	defer th.NewTestCase()()
	res = method(7, -3)
	th.Expect(-2).Assert(res)

	defer th.NewTestCase()()
	res = method(0, 1)
	th.Expect(0).Assert(res)

	defer th.NewTestCase()()
	res = method(-2147483648, -1)
	th.Expect(2147483647).Assert(res)
}
