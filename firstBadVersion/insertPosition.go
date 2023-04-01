package main

func searchInsert(nums []int, target int) int {
	l := 0
	r := len(nums)
	m := 0
	for r >= l && l < len(nums) && r >= 0 {
		m = (r + l) >> 1
		if nums[m] == target {
			return m
		} else if nums[m] > target {
			r = m - 1
		} else {
			l = m + 1
		}

	}
	if nums[m] < target {
		m++
	}
	return m
}
