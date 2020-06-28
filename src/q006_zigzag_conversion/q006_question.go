// Package q006_zigzag_conversion [Medium] ZigZag Conversion
// https://leetcode.com/problems/zigzag-conversion/
//
// The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)
//
//
// P   A   H   N
// A P L S I I G
// Y   I   R
//
//
// And then read line by line: "PAHNAPLSIIGYIR"
//
// Write the code that will take a string and make this conversion given a number of rows:
//
//
// string convert(string s, int numRows);
//
// Example 1:
//
//
// Input: s = "PAYPALISHIRING", numRows = 3
// Output: "PAHNAPLSIIGYIR"
//
//
// Example 2:
//
//
// Input: s = "PAYPALISHIRING", numRows = 4
// Output: "PINALSIGYAHRPI"
// Explanation:
//
// P     I    N
// A   L S  I G
// Y A   H R
// P     I
package q006_zigzag_conversion

import (
	"testing"

	"github.com/sunxuia/leetcode-solution-go/src/util/test"
)

func doTest(t *testing.T, method func(string, int) string) {
	ah := test.NewAssertHelper()
	var res string

	res = method("PAYPALISHIRING", 3)
	ah.Assert("PAHNAPLSIIGYIR", res)

	res = method("PAYPALISHIRING", 4)
	ah.Assert("PINALSIGYAHRPI", res)

	res = method("123456789", 1)
	ah.Assert("123456789", res)

	res = method("", 1)
	ah.Assert("", res)
}
