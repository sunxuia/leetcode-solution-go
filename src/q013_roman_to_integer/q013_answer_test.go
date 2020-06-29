package q013_roman_to_integer

import (
	"math"
	"testing"
)

func TestRomanToInt(t *testing.T) {
	doTest(t, romanToInt)
}

func romanToInt(s string) int {
	res, prev := 0, math.MaxInt32
	for _, c := range s {
		cur := romans[c]
		res += cur
		if prev < cur {
			res -= prev << 1
		}
		prev = cur
	}
	return res
}

var romans = map[rune]int{
	'M': 1000,
	'D': 500,
	'C': 100,
	'L': 50,
	'X': 10,
	'V': 5,
	'I': 1,
}
