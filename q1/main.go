package main

import "fmt"

func main() {
	a := 0
	for i := 0; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			fmt.Println("Adding", i, "; sum is now", a+i)
			a += i
		}
	}
	fmt.Println("The Answer is:", a)
}
