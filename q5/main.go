package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

// Finds the smallest number that can be evenly divided by the positive integers smaller than or equal to n.
func main() {
	//Design:
	//smallest multiple is multiplication of primes, then multiply by a prime for each "power of"
	//ex: 10 = 2 * 3 * 5 * 2 * 2 * 9; OR 5 * 8 * 9 * 7
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Error, argument", os.Args[1], "Not a number!")
		os.Exit(1)
	}
	if n <= 0 {
		fmt.Println("Error: Invalid Arugment.", os.Args[1], "must be a positive integer greater than 0!")
	}
	//answer is 1 if n is 1, 2 if n is 2.
	ans := n
	if n > 2 {
		ans = smallestMultiple(n)
	}
	fmt.Println("Smallest number evenly divisible by all integers 1 to", n, "is:", ans)
}

//Finds primes smaller than
func smallestMultiple(n int) int {
	p := 1
	var arr []int
	arr = findSmallerPrimes(n)
	arr = findSmallerPowersOfPrimes(arr, n)
	for _, num := range arr {
		p = p * num
	}
	return p
}

// n is the upper bound of numbers.
func findSmallerPrimes(n int) []int {
	arr := []int{}
	if n < 3 {
		arr = append(arr, n)
		return arr
	}
	// since we know n > 2, we can put n in the slice and iterate every-other number.
	arr = append(arr, 2)
	for i := 3; i <= n; i += 2 {
		prime := true
		//Check if prime, exit loop if found not prime
		for j := 3; j < i; j += 2 {
			if i%j == 0 {
				prime = false
			}
		}
		if prime == true {
			arr = append(arr, i)
		}
	}
	return arr
}

//n = upper bound of numbers, arr = int slice storing the numbers we need.
func findSmallerPowersOfPrimes(arr []int, n int) []int {
	for i, num := range arr {
		new := num
		for new <= n {
			if new*num <= n {
				new = new * num
			} else {
				arr[i] = new
				new = math.MaxInt16
			}
		}
	}
	return arr
}
