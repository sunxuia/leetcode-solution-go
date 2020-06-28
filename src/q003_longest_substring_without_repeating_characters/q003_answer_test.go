package q003_longest_substring_without_repeating_characters

import (
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	doTest(t, lengthOfLongestSubstring)
}

func lengthOfLongestSubstring(s string) int {
	occupied := [128]int{}
	res, last := 0, 0
	for i, c := range s {
		if occupied[c] > last {
			if res < i-last {
				res = i - last
			}
			last = occupied[c]
		}
		occupied[c] = i + 1
	}
	if res < len(s)-last {
		res = len(s) - last
	}
	return res
}
