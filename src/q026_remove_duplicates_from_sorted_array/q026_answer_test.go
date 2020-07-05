package q026_remove_duplicates_from_sorted_array

import (
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	doTest(t, removeDuplicates)
}

func removeDuplicates(nums []int) int {
	size := 0
	for i := 0; i < len(nums); i++ {
		if i == 0 || nums[size-1] != nums[i] {
			nums[size] = nums[i]
			size++
		}
	}
	return size
}
