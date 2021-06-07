package main

import (
	"fmt"
	"palindrome"
	"strings"
)

func main() {
	// Get a greeting message and print it.
	var str string
	fmt.Print("Enter a string: ")
	fmt.Scan(&str)
	if palindrome.VerifyPalindrome(strings.ToUpper(str)) == true {
		fmt.Print(str, " is a palindrome.")
	} else {
		fmt.Print(str, " is not a palindrome.")
	}
}
