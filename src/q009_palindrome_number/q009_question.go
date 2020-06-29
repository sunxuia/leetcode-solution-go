// Package q009_palindrome_number [Easy] Palindrome Number
// https://leetcode.com/problems/palindrome-number/
//
// Determine whether an integer is a palindrome. An integer is a palindrome when it reads the same backward as forward.
//
// Example 1:
//
//
// Input: 121
// Output: true
//
//
// Example 2:
//
//
// Input: -121
// Output: false
// Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
//
//
// Example 3:
//
//
// Input: 10
// Output: false
// Explanation: Reads 01 from right to left. Therefore it is not a palindrome.
//
//
// Follow up:
//
// Coud you solve it without converting the integer to a string?
package q009_palindrome_number

import (
	"testing"

	"github.com/sunxuia/leetcode-solution-go/src/util/test"
)

func doTest(t *testing.T, method func(int) bool) {
	ah := test.NewAssertHelper()
	var res bool

	res = method(121)
	ah.Assert(true, res)

	res = method(-121)
	ah.Assert(false, res)

	res = method(10)
	ah.Assert(false, res)
}
