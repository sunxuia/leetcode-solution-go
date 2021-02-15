package q043_multiply_strings

import (
	"strings"
	"testing"
)

func TestMultiply(t *testing.T) {
	doTest(t, multiply)
}

func multiply(num1 string, num2 string) string {
	m, n := len(num1), len(num2)
	// 乘法规则相乘
	nums := make([]int, m+n)
	for i := 0; i < m; i++ {
		carry := 0
		for j := 0; j < n; j++ {
			a := int(num1[m-1-i] - '0')
			b := int(num2[n-1-j] - '0')
			nums[i+j] += a*b + carry
			carry = nums[i+j] / 10
			nums[i+j] %= 10
		}
		nums[i+n] += carry
	}

	// 数组 -> 字符串
	sb := strings.Builder{}
	last := m + n - 1
	for last > 0 && nums[last] == 0 {
		last--
	}
	for ; last >= 0; last-- {
		sb.WriteByte(byte(nums[last] + '0'))
	}
	return sb.String()
}
