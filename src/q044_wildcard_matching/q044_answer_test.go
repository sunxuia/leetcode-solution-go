package q044_wildcard_matching

import (
	"testing"
)

func TestIsMatch(t *testing.T) {
	doTest(t, isMatch)
}

func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n+1)
	}

	// 边界条件
	dp[0][0] = true
	for i := 0; i < n && p[i] == '*'; i++ {
		dp[0][i+1] = true
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if p[j] == '*' {
				dp[i+1][j+1] = dp[i][j] || dp[i+1][j] || dp[i][j+1]
			} else if p[j] == '?' {
				dp[i+1][j+1] = dp[i][j]
			} else {
				dp[i+1][j+1] = dp[i][j] && s[i] == p[j]
			}
		}
	}
	return dp[m][n]
}
