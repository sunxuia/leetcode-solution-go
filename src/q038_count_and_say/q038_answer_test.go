package q038_count_and_say

import (
	"strconv"
	"strings"
	"testing"
)

func TestCountAndSay(t *testing.T) {
	doTest(t, countAndSay)
}

func countAndSay(n int) string {
	str := "1"
	for i := 1; i < n; i++ {
		start, sb := 0, strings.Builder{}
		for j := 1; j < len(str); j++ {
			if str[start] != str[j] {
				sb.WriteString(strconv.Itoa(j - start))
				sb.WriteByte(str[start])
				start = j
			}
		}
		sb.WriteString(strconv.Itoa(len(str) - start))
		sb.WriteByte(str[start])
		str = sb.String()
	}
	return str
}
