package main

import (
	"fmt"
	//"math"
)

func main() {
	//cannot operate invalid operands
	// + - * / %
	var num1 float64 = 8
	var num2 int = 4
	answer := num1 + float64(num2)

	//sub_answer := int(num1) / num2
	//follows Bracket , Exponent, Division , Multi.., Add.., Sub...,
	fmt.Printf("%f", answer)
}
