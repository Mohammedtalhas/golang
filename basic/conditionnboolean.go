package main

import (
	"fmt"
	//"math"
)

func main() {
	x := 5
	y := 6
	a := "A" //65
	b := "b" //98
	val := x >= y
	asciival := a <= b
	fmt.Printf("%t", val)
	fmt.Printf("%t", asciival)
	fmt.Printf("%t", asciival || val)
	/*
		conditional operator
			<
			>
			<=
			>=
			==
			!=

		Logical Operator
		  &&
		  ||
		  !
	*/
}
