// [Medium] Multiply Strings
// https://leetcode.com/problems/multiply-strings/
//
// Given two non-negative integers num1 and num2 represented as strings, return the product of num1 and num2, also represented as a string.
//
// Example 1:
//
//
// Input: num1 = "2", num2 = "3"
// Output: "6"
//
// Example 2:
//
//
// Input: num1 = "123", num2 = "456"
// Output: "56088"
//
//
// Note:
//
//
// 	The length of both num1 and num2 is < 110.
// 	Both num1 and num2 contain only digits 0-9.
// 	Both num1 and num2 do not contain any leading zero, except the number 0 itself.
// 	You must not use any built-in BigInteger library or convert the inputs to integer directly.
package q043_multiply_strings

import (
	"github.com/sunxuia/leetcode-solution-go/src/util/test"
	"testing"
)

func doTest(t *testing.T, method func(string, string) string) {
	th := test.NewTestHelper(t)
	var res string

	defer th.NewTestCase()()
	res = method("2", "3")
	th.Expect("6").Assert(res)

	defer th.NewTestCase()()
	res = method("123", "456")
	th.Expect("56088").Assert(res)

	defer th.NewTestCase()()
	res = method("9133", "0")
	th.Expect("0").Assert(res)
}
