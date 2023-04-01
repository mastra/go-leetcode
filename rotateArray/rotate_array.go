package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	//a = []int{-7, -3, 2, 3, 11}
	//a = []int{-2, -1, 0}
	rotateMem(a, 3)
	fmt.Println(a)
	//	rotate1(a, 3)
	fmt.Println(a)

	// r := sortedSquaresForLoop(a)
	//fmt.Println(r)

	//fmt.Println(r)
}

func rotate(nums []int, k int) {
	for k > 0 {
		m := nums[len(nums)-1]
		p := 0
		for i := 0; i < len(nums); i += 1 {
			p = nums[i]
			nums[i] = m
			m = p
		}
		k--
	}
}

func rotate1(nums []int, k int) {
	m := nums[len(nums)-1]
	p := 0
	for i := 0; i < len(nums); i += 2 {
		if i+1 < len(nums) {
			p = nums[i+1]
			nums[i+1] = nums[i]
		}
		nums[i] = m
		m = p
	}
}

func rotateMem(nums []int, k int) {
	var result []int
	var l = len(nums)
	//m := nums[len(nums)-1]
	j := 0
	for i := 0; i < len(nums); i += 1 {
		j = (i - k) % l
		if j < 0 {
			j = l + j
		}
		result = append(result, nums[j])
	}
	for i := 0; i < len(nums); i += 1 {
		nums[i] = result[i]
	}
}
