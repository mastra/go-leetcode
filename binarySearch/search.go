package main

import "fmt"

func main() {
	nums := []int{2, 5, 8}
	t := 8
	fmt.Println(search(nums, t))
	fmt.Println(search2(nums, t))
}

func search(nums []int, target int) int {
	l := 0
	r := len(nums)
	if r < 1 {
		return -1
	}

	i := 0
	for r-l > 1 {
		i = (r + l) / 2
		t := nums[i]
		if t == target {
			return i
		}
		if target > t {
			l = i
		} else {
			r = i
		}
	}
	if r-l == 1 {
		if nums[0] == target {
			return 0
		}

	}
	return -1
}

func search2(nums []int, target int) int {
	l := 0
	r := len(nums) - 1

	i := 0
	for r >= l {
		i = (r + l) / 2
		t := nums[i]
		if t == target {
			return i
		}
		if target > t {
			l = i + 1
		} else {
			r = i - 1
		}
	}
	return -1
}
