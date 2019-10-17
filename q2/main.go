package main

import "fmt"

func main() {
	//Sum of even Fibonacci terms whose values are less than 4 mil.
	sum := 0

	first := 0
	second := 1
	third := 1

	for first < 4000000 { //4 million
		first, second, third = second, third, second+third
		if first%2 == 0 {
			fmt.Println("Adding", first, "to sum")
			fmt.Println(second, third) //do something with these
			sum += first
		}
	}
	fmt.Println("Sum is:", sum)
}
