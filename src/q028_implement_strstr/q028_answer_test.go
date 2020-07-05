package q028_implement_strstr

import (
	"testing"
)

func TestStrStr(t *testing.T) {
	doTest(t, strStr)
}

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	k := kmp(needle)
	j := 0
	for i := 0; i < len(haystack); i++ {
		j = k[j][haystack[i]-'a']
		if j == len(needle) {
			return i - j + 1
		}
	}
	return -1
}

func kmp(p string) [][26]int {
	res := make([][26]int, len(p))
	res[0][p[0]-'a'] = 1
	reset := &res[0]
	for i := 1; i < len(p); i++ {
		copy(res[i][:], reset[:])
		res[i][p[i]-'a'] = i + 1
		reset = &res[reset[p[i]-'a']]
	}
	return res
}
