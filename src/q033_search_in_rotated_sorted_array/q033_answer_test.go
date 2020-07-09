package q033_search_in_rotated_sorted_array

import (
	"testing"
)

func TestSearch(t *testing.T) {
	doTest(t, search)
}

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	min := findMinimum(nums)
	if min > 0 && nums[0] <= target {
		return binarySearch(nums, target, 0, min-1)
	} else {
		return binarySearch(nums, target, min, len(nums)-1)
	}
}

func findMinimum(nums []int) int {
	start, end := 0, len(nums)-1
	for start < end {
		mid := (start + end) / 2
		if nums[mid] < nums[end] {
			end = mid
		} else {
			start = mid + 1
		}
	}
	return end
}

func binarySearch(nums []int, target, start, end int) int {
	for start < end {
		mid := (start + end) / 2
		if nums[mid] >= target {
			end = mid
		} else {
			start = mid + 1
		}
	}
	if nums[end] == target {
		return end
	}
	return -1
}
