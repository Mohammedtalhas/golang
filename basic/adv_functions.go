package main

import (
	"fmt"
	//"math"
)

func main() {
	test := func() {
		fmt.Println("Helloo")
	}
	test()
	calc := func(x int) int {
		return x * -1
	}(10)
	fmt.Println(calc)
}
