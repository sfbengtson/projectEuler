package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Error, expected n to be an int, got", n)
		os.Exit(1)
	}
	ans := findNthPrime(n)
	fmt.Println("The nth prime is", ans)
}

func findNthPrime(n int) int {
	if n == 1 {
		return 2
	}
	ans, count := 2, 1
	for i := 3; count < n; i += 2 {
		if isPrime(i) {
			ans = i
			count++
		}
	}
	return ans
}

func isPrime(n int) bool {
	if n == 2 {
		return true
	}
	if n < 2 {
		return false
	}
	for i := 3; i < n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	fmt.Println(n, "is prime!")
	return true
}
