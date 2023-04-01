package main

import "fmt"

func main() {
	a := []int{0, 2, 0, 4, 5, 6, 7}
	//a = []int{-7, -3, 2, 3, 11}
	//a = []int{1, 0, 2, 3}
	moveZeroes(a)
	fmt.Println(a)
}

func moveZeroes(nums []int) {
	z := 0 //z
	for n := 0; n < len(nums); n++ {
		for nums[z] != 0 {
			z++
			if z == len(nums) {
				return
			}
		}
		//for nums[n] != 0 {
		//	n++
		//}
		if z < n {
			nums[z], nums[n] = nums[n], nums[z]
		}
	}
}
