package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type customWriter struct{}

func main() {
	file, err := os.Open("num.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	bt := ""
	for scanner.Scan() {
		bt += scanner.Text()
	}

	t := string(bt)

	arr := make([]int, 0)

	for i := 0; i < len(t); i++ {
		j, _ := strconv.Atoi(t[i : i+1])
		arr = append(arr, j)
	}

	maxProd := 0

	for i := 0; i < len(arr)-13; i++ {
		prod := 1
		for j := 0; j < 13; j++ {
			prod = prod * arr[i+j]
		}
		if prod > maxProd {
			maxProd = prod
		}
	}

	fmt.Println("Max product is", maxProd)
}
