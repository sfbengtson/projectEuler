package main

import "testing"

func TestFindSmallerPrimes(t *testing.T) {

	var ans []int

	ans = findSmallerPrimes(1)
	if ans[0] != 1 {
		t.Error("Expected answer for 1 to be 1, got", ans)
	}

	if len(ans) != 1 {
		t.Error("Expected length of answer to be 1, got", len(ans))
	}

	ans = findSmallerPrimes(2)
	if ans[0] != 2 {
		t.Error("Expected answer for 1 to be 2, got", ans)
	}

	if len(ans) != 1 {
		t.Error("Expected length of answer to be 1, got", len(ans))
	}

	ans = findSmallerPrimes(11)
	if len(ans) != 5 {
		t.Error("Expected length of array for n=11 to be 5, got", len(ans))
	}
	if ans[len(ans)-1] != 11 {
		t.Error("Expected last element when n = 11 to be 11, got", ans[len(ans)-1])
	}
}

func TestFindSmallerPowersOfPrimes(t *testing.T) {

	a := []int{2, 3, 5, 7}
	var ans []int

	ans = findSmallerPowersOfPrimes(a, 9)
	if ans[0] != 8 {
		t.Error("Expected first integer to be 8, got", ans[0])
	}

	if ans[1] != 9 {
		t.Error("Expected second integer to be 9, got", ans[1])
	}

	if len(ans) != 4 {
		t.Error("Expected length of ans to be 4, got", len(ans))
	}
}

func TestSmallestMultiple(t *testing.T) {
	// Testing Project Euler example.
	a := smallestMultiple(10)
	if a != 2520 {
		t.Error("Expected smallest multiple of 10 to be 2520, got", a)
	}
}
