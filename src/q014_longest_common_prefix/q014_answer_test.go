package q014_longest_common_prefix

import (
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	doTest(t, longestCommonPrefix)
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for i := 0; i < len(prefix); i++ {
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || prefix[i] != strs[j][i] {
				return prefix[0:i]
			}
		}
	}
	return prefix
}
