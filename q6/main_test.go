package main

import "testing"

func TestSumOfSquares(t *testing.T) {
	a := sumOfSquares(2)
	if a != 5 {
		t.Error("Expected sOS(2) to be 5, got", a)
	}
	a = sumOfSquares(3)
	if a != 14 {
		t.Error("Expected sOS(3) to be 14, got", a)
	}
}

func TestSquareOfSums(t *testing.T) {
	a := squareOfSum(2)
	if a != 9 {
		t.Error("Expected SqOS(2) to be 9, got", a)
	}
	a = squareOfSum(3)
	if a != 36 {
		t.Error("Expected SqOS(3) to be 36, got", a)
	}
}
