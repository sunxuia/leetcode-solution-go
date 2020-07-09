package q035_search_insert_position

import (
	"testing"
)

func TestSearchInsert(t *testing.T) {
	doTest(t, searchInsert)
}

func searchInsert(nums []int, target int) int {
	start, end := 0, len(nums)
	for start < end {
		mid := (start + end) / 2
		if nums[mid] < target {
			start = mid + 1
		} else {
			end = mid
		}
	}
	return end
}
