// Package q007_reverse_integer [Easy] Reverse Integer
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
	ah := test.NewAssertHelper()
	var res int

	res = method(123)
	ah.Assert(321, res)

	res = method(-123)
	ah.Assert(-321, res)

	res = method(120)
	ah.Assert(21, res)

	// 结果超过整数范围则返回0
	res = method(1534236469)
	ah.Assert(0, res)
}
