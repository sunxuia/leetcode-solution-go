package q034_find_first_and_last_position_of_element_in_sorted_array

import (
	"testing"
)

func TestSearchRange(t *testing.T) {
	doTest(t, searchRange)
}

func searchRange(nums []int, target int) []int {
	start, end := -1, len(nums)-1
	for start < end {
		mid := (start + end + 1) / 2
		if target <= nums[mid] {
			end = mid - 1
		} else {
			start = mid
		}
	}

	res := []int{-1, -1}
	if start > len(nums)-2 || nums[start+1] != target {
		return res
	}
	res[0] = start + 1

	start, end = res[0]+1, len(nums)
	for start < end {
		mid := (start + end) / 2
		if nums[mid] <= target {
			start = mid + 1
		} else {
			end = mid
		}
	}
	res[1] = start - 1
	return res
}
