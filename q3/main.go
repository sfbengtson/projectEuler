package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	//Large number prime factorization:
	//store largest prime factor in list.
	//divide by 2 if even, continue until odd.
	//attempt to divide by 3, 5, etc;
	//continue until 'divider' is greater than sqrt of remaining number.

	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Error converting", os.Args[1], "to int!")
		os.Exit(1)
	}
	answer := primeFactor(num)
	fmt.Println(answer)

}

func primeFactor(num int) int {
	ans := num
	for num%2 == 0 {
		num = num / 2
		ans = 2
	}

	for i := 3; i <= num; i += 2 {
		if num%i == 0 {
			ans = i
			num = num / i
		}
	}
	fmt.Println("answer is", ans)
	return ans
}

// Project Euler: answer: 6857
