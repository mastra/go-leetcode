package main

import "fmt"

func main() {
	a := []int{0, 1, 2, 2, 3, 0, 4, 2}
	//fmt.Println(removeElement(a, 2))
	//a = []int{0, 1, 2, 2, 3, 0, 4, 2}
	//fmt.Println(removeElementSolution(a, 2))
	//a = []int{3, 2, 2, 3}
	//fmt.Println(removeElement(a, 3))
	//a = []int{2}
	//fmt.Println(removeElement(a, 3))

	a = []int{3, 2, 2, 3}
	fmt.Println(removeElementBest(a, 3))

}

func removeElement(nums []int, val int) int {
	k := len(nums)
	for i := 0; i < k; i++ {
		if nums[i] == val {
			for k > 0 && nums[k-1] == val {
				k--
			}
			if k > i {
				nums[i] = nums[k-1]
				nums[k-1] = val
			}
		}
	}
	fmt.Println(nums)
	return k
}

func removeElementBest(nums []int, val int) int {
	k := len(nums)
	i := 0
	for i < k {
		if nums[i] == val {
			nums[i] = nums[k-1]
			k--
		} else {
			i++
		}
	}
	fmt.Println(nums)
	return k
}

func removeElementSolution(nums []int, val int) int {
	i := 0
	for j := 0; j < len(nums); j++ {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}
	}
	fmt.Println(nums)
	return i
}
