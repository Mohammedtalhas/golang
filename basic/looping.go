package main

import (
	"fmt"
	//"math"
)

func main() {
	x := 0
	//y := 6
	for x < 6 {
		fmt.Println(x)
		x++

	}
	for x := 0; x < 6; x++ {
		fmt.Println(x)

	}
	for x := 0; x < 1000; x++ {
		if x != 0 && x%3 == 0 && x%7 == 0 && x%9 == 0 {
			fmt.Println(x)
			continue

		}
		fmt.Println(x, "Not Divisible")

	}

	/* for { //infinite loop
		continue
		break

	} */
}
