package main

import (
	"fmt"
	"time"
)

func main() {

	var a []int
	max := 10000
	target := max + (max - 1)
	for i := 0; i <= max; i++ {
		a = append(a, i)
	}
	//a := []int{3, 3}
	//target := 6
	start := time.Now()

	fmt.Println(twoSum(a, target))
	elapsed := time.Since(start)
	fmt.Printf("Twosum took %s", elapsed)

	start = time.Now()
	fmt.Println(twoSumSolution(a, target))
	elapsed = time.Since(start)
	fmt.Printf("TwosumSolution took %s", elapsed)

}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	max := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}
	ignore := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if nums[i]+max < target {
			ignore[i] = 1
		}
	}
	//fmt.Println(max)
	//fmt.Println(ignore)
	for i := 0; i < len(nums); i++ {
		if _, f := ignore[i]; f {
			continue
		}
		for j := 0; j < len(nums); j++ {
			if j == i {
				continue
			}
			//if v, found := m[j]; i == v && found {
			//	continue
			//}
			if nums[i]+nums[j] == target {
				return []int{i, j}
			} else {
				m[i] = j
			}
		}
	}
	return []int{}
}

func twoSumBrute(nums []int, target int) []int {
	m := make(map[int]int)
	fmt.Println(len(nums))
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if j == i {
				continue
			}
			if v, found := m[j]; i == v && found {
				continue
			}
			if nums[i]+nums[j] == target {
				return []int{i, j}
			} else {
				m[i] = j
			}
		}
	}
	return []int{}
}

func twoSumSolution(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = i
	}

	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]
		if pos, f := m[complement]; pos != i && f {
			return []int{pos, i}
		}
	}

	return []int{}
}

func twoSumSolution1(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]
		if pos, f := m[complement]; pos != i && f {
			return []int{pos, i}
		}
		m[nums[i]] = i
	}

	return []int{}
}
