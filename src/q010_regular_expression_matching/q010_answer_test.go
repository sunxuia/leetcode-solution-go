package q010_regular_expression_matching

import (
	"regexp"
	"testing"
)

func TestIsMatch(t *testing.T) {
	doTest(t, isMatch)
}

// 使用自带的正则表达式
func isMatch(s string, p string) bool {
	return regexp.MustCompile("^" + p + "$").MatchString(s)
}

func TestIsMatch2(t *testing.T) {
	doTest(t, isMatch2)
}

// dp 的方法
func isMatch2(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true

	// 开头的* 匹配
	for i := 0; i < n; i++ {
		if p[i] == '*' && dp[0][i-1] {
			dp[0][i+1] = true
		}
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' {
				// 0 匹配. 此时的值表示前一个匹配结果.
				match := dp[i][j-2]
				// 1或多次匹配. [si, pi -1] 表示* 前面的匹配与当前字符匹配, 这是第2 次匹配,
				// [si -1, pi] 表示前一个字符与* 匹配, 这是第N 次匹配.
				if !match && (dp[i][j-1] || dp[i-1][j]) {
					// .* 的多次匹配.
					match = p[j-2] == '.'
					// * 的多次匹配, 与上一个字符相等.
					match = match || s[i-1] == p[j-2]
				}
				dp[i][j] = match
			} else {
				dp[i][j] = dp[i-1][j-1] && (p[j-1] == '.' || s[i-1] == p[j-1])
			}
		}
	}
	return dp[m][n]
}
