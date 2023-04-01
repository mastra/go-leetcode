package main

import (
	"fmt"
)

func main() {
	fmt.Println(isPalindrome(1221))
	fmt.Println(isPalindromeSolution(1221))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	var digits []int
	divisor := x
	for divisor > 0 {
		digit1 := divisor % 10
		digits = append(digits, int(digit1))
		divisor = divisor / 10
	}

	l := len(digits)
	for i := 0; i < l/2; i++ {
		if digits[i] != digits[l-i-1] {
			return false
		}
	}
	return true
}

func isPalindromeSolution(x int) bool {
	if x < 0 {
		return false
	}
	xx := x
	rev := 0
	for x > 0 {
		rev = rev*10 + x%10
		x = x / 10
	}
	return rev == xx || xx == rev/10
}
