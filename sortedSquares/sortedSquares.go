package main

import (
	"fmt"
)

func main() {
	a := []int{-4, -1, 0, 3, 10}
	a = []int{-7, -3, 2, 3, 11}
	a = []int{-2, -1, 0}
	fmt.Println(sortedSquares(a))
	// r := sortedSquaresForLoop(a)
	//fmt.Println(r)
	r := sortedSquares2pointers(a)
	fmt.Println(r)
	r = sortedSquares2pointersAnother(a)
	fmt.Println(r)
}

func sortedSquares(nums []int) []int {
	var r []int
	var m []int

	n := len(nums)

	for i := 0; i < n; i++ {
		if nums[i] < 0 {
			m = append(m, nums[i]*nums[i])
		} else {
			s := nums[i] * nums[i]
			for len(m) > 0 && s > m[len(m)-1] {
				r = append(r, m[len(m)-1])
				m = m[:len(m)-1]
			}
			r = append(r, s)
		}
	}
	for i := len(m) - 1; i >= 0; i-- {
		r = append(r, m[i])
	}
	return r
}

func sortedSquaresForLoop(nums []int) []int {
	var r []int

	n := len(nums)
	if n == 0 {
		return r
	}
	if n == 1 {
		return append(r, nums[0]*nums[0])
	}
	m := 0

	//if nums[0] >= 0 {
	//	for m = 0; m < n && nums[m] < 0; m++ {
	//		r = append(r, nums[m]*nums[m])
	//	}
	//	return r
	//}

	for m = 0; m < n && nums[m] < 0; m++ {
	}
	m--
	s := nums[m] * nums[m]

	for i := m + 1; i < n; {
		ii := nums[i] * nums[i]
		if ii <= s || m == -1 {
			r = append(r, ii)
			i++
		}

		for s < ii && m >= 0 {
			r = append(r, s)
			m--
			if m >= 0 {
				s = nums[m] * nums[m]
			}
		}
	}

	return r
}
func sortedSquares2pointers(nums []int) []int {
	var r []int

	n := len(nums)
	if n == 0 {
		return r
	}
	if n == 1 {
		return append(r, nums[0]*nums[0])
	}

	i := 0
	j := n - 1

	for i < n && j >= 0 && i <= j {
		ii := nums[i] * nums[i]
		jj := nums[j] * nums[j]

		if ii > jj {
			r = append([]int{ii}, r...)
			i++
		} else {
			r = append([]int{jj}, r...)
			j--
		}

	}

	return r
}

func sortedSquares2pointersAnother(nums []int) []int {
	var r []int
	n := len(nums)

	if n == 0 {
		return r
	}
	if n == 1 {
		return append(r, nums[0]*nums[0])
	}

	i := 0
	j := n - 1

	for p := n - 1; p >= 0; p-- {

		if Abs(nums[i]) > Abs(nums[j]) {
			ii := nums[i] * nums[i]
			r = append([]int{ii}, r...)
			i++
		} else {
			jj := nums[j] * nums[j]
			r = append([]int{jj}, r...)
			j--
		}

	}

	return r
}

func Abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}
