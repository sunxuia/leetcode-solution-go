package q017_letter_combinations_of_a_phone_number

import (
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	doTest(t, letterCombinations)
}

func letterCombinations(digits string) []string {
	res := make([]string, 0)
	if len(digits) == 0 {
		return res
	}
	res = append(res, "")
	for _, d := range digits {
		str := nums[d]
		if len(str) == 0 {
			return []string{}
		}
		size := len(res)
		for k := 0; k < size; k++ {
			for j := 0; j < len(str); j++ {
				res = append(res, res[k]+str[j:j+1])
			}
		}
		res = res[size:]
	}
	return res
}

var nums = map[rune]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}
