package q005_longest_palindromic_substring

import (
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	doTest(t, longestPalindrome)
}

// 简单的循环查找, 时间复杂度 O(N^2)
func longestPalindrome(s string) string {
	var res string
	for i := 0; i < len(s); i++ {
		for offset := 0; offset <= 1; offset++ {
			left, right := i, i+offset
			for left >= 0 && right < len(s) {
				if s[left] != s[right] {
					break
				}
				left--
				right++
			}
			if len(res) < right-left-1 {
				res = s[left+1 : right]
			}
		}
	}
	return res
}
