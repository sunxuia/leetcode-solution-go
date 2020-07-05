package q027_remove_element

import (
	"testing"
)

func TestRemoveElement(t *testing.T) {
	doTest(t, removeElement)
}

func removeElement(nums []int, val int) int {
	size := 0
	for i := 0; i < len(nums); i++ {
		if val != nums[i] {
			nums[size] = nums[i]
			size++
		}
	}
	return size
}
