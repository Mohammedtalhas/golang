package main

import (
	"fmt"
	//"math"
)

func main() {
	x := 5
	y := 6
	val := x >= y

	if val {
		fmt.Println("In If Condition")
	} else {
		fmt.Println("In Else Condition")
	}
}
