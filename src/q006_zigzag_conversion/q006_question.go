// [Medium] ZigZag Conversion
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
	th := test.NewTestHelper(t)
	var res string

	defer th.NewTestCase()()
	res = method("PAYPALISHIRING", 3)
	th.Expect("PAHNAPLSIIGYIR").Assert(res)

	defer th.NewTestCase()()
	res = method("PAYPALISHIRING", 4)
	th.Expect("PINALSIGYAHRPI").Assert(res)

	defer th.NewTestCase()()
	res = method("123456789", 1)
	th.Expect("123456789").Assert(res)

	defer th.NewTestCase()()
	res = method("", 1)
	th.Expect("").Assert(res)
}
