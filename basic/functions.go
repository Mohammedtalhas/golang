package main

import (
	"fmt"
	//"math"
)

func Sum(a, b float64) (float64, error) {
	if a < 0 && b < 0 {
		return 0, fmt.Errorf("math: square root of negative number %g %g", a, b)
	}
	return a + b, nil
}
func main() {
	f, err := Sum(20, 10)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(f)

	}
}
