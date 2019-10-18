package main

import "testing"

func TestPalindrome(t *testing.T) {
	r := palindrome(1)
	if r != true {
		t.Error("Expected 1 to be a palindrome!")
	}
	r = palindrome(22)
	if r != true {
		t.Error("Expected 22 to be a palindrome!")
	}
	r = palindrome(23)
	if r != false {
		t.Error("Expected 23 to not be a palindrome!")
	}
	r = palindrome(232)
	if r != true {
		t.Error("Expected 232 to be a palindrome!")
	}

}
