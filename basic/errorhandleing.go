package main

import (
	"fmt"
	//"math"
)

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, fmt.Errorf("math: square root of negative number %g", f)
	}
	return f * f, nil
}
func main() {
	f, err := Sqrt(20)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(f)

	}
}
