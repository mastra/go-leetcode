package main

import "fmt"

func main() {
	a := []int{1, 3, 5, 6, 9}
	//fmt.Println(searchInsert(a, 10))
	a = []int{1, 3}
	fmt.Println(searchInsert(a, 0))

	fmt.Println("------------")
	//fmt.Println(isBadVersion(9))
	//fmt.Println(firstBadVersion(9))
}

func isBadVersion(i int) bool {
	return i >= 9
}

func firstBadVersion(n int) int {
	r := n
	l := 1
	m := r
	for r >= l {
		m = (r + l) >> 1
		if isBadVersion(m) == false {
			l = m
		} else {
			r = m - 1
		}
	}
	if isBadVersion(m) == false {
		m = m + 1
	}
	return m
}
