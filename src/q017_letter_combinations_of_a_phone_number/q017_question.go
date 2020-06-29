// Package q017_letter_combinations_of_a_phone_number [Medium] Letter Combinations of a Phone Number
// https://leetcode.com/problems/letter-combinations-of-a-phone-number/
//
// Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent.
//
// A mapping of digit to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.
//
// (图 Q017_PIC.png)
//
// Example:
//
//
// Input: "23"
// Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
//
//
// Note:
//
// Although the above answer is in lexicographical order, your answer could be in any order you want.
package q017_letter_combinations_of_a_phone_number

import (
	"testing"

	"github.com/sunxuia/leetcode-solution-go/src/util/test"
)

func doTest(t *testing.T, method func(string) []string) {
	ah := test.NewAssertHelper()
	ah.UnOrder("*")
	var res []string

	res = method("23")
	ah.Assert([]string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}, res)

	res = method("234")
	ah.Assert([]string{"adg", "adh", "adi", "aeg", "aeh", "aei", "afg", "afh", "afi", "bdg", "bdh", "bdi", "beg", "beh", "bei", "bfg", "bfh", "bfi", "cdg", "cdh", "cdi", "ceg", "ceh", "cei", "cfg", "cfh", "cfi"}, res)
}
