package q009_palindrome_number

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	doTest(t, isPalindrome)
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	var reverse int
	for i := x; i > 0; i /= 10 {
		reverse = reverse*10 + i%10
	}
	return reverse == x
}
