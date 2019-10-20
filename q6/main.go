package main

import (
	"fmt"
	"os"
	"strconv"
)

//difference between sum of squares and square of sum
func main() {
	var a, b int

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Error:", os.Args[1], "is not an int!")
		os.Exit(1)
	}

	a = sumOfSquares(n)
	b = squareOfSum(n)
	fmt.Println("Difference between sum of squares and square of sums is", b-a)
}

func sumOfSquares(n int) int {
	ans := 0

	for i := 1; i <= n; i++ {
		ans += (i * i)
	}
	fmt.Println("Sum of squares between 1 and", n, "is:", ans)
	return ans
}

func squareOfSum(n int) int {
	ans := 0

	for i := 1; i <= n; i++ {
		ans += i
	}
	ans = ans * ans
	fmt.Println("Square of sum of numbers between 1 and", n, "is:", ans)
	return ans
}
