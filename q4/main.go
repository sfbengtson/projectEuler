package main

import (
	"fmt"
	"strconv"
)

func main() {
	//Palindromic product of 2 3digit ints.
	// Func to
	ans := 0
	for i := 100; i < 1000; i++ {
		for j := 100; j < 1000; j++ {
			ans = largerPalindrome(i*j, ans)
		}
	}
	fmt.Println("Answer is:", ans)
}

func largerPalindrome(new int, old int) int {
	if new > old && palindrome(new) {
		fmt.Println("new largest is:", new)
		return new
	}
	return old
}

func palindrome(i int) bool {
	a := strconv.Itoa(i)
	for i = 0; i < len(a); i++ {
		if a[i] != a[len(a)-i-1] {
			return false
		}
	}
	return true
}
