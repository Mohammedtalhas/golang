package main

import (
	"fmt"
	//"math"
)

func main() {
	x := 5
	ans := x%3 == 0 || x%5 == 0 || x%7 == 0
	switch ans {
	case false, true:
		fmt.Println(x, "Is not divisible by 3 or 5 or 7")
	case true:
		fmt.Println(x, "Is divisible by 3 or 5 or 7")
	default:
		fmt.Println("Error Occured,Neither Both")

	}
	switch {
	case ans == false:
		fmt.Println(x, "Is not divisible by 3 or 5 or 7")
	case ans == true:
		fmt.Println(x, "Is divisible by 3 or 5 or 7")
	default:
		fmt.Println("Error Occured,Neither Both")

	}
}
