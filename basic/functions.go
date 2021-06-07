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
func Oper(a, b float64) (x, y float64) {
	defer fmt.Println("Executed Successfully")
	//labeled return
	x = a + b
	y = a - b
	fmt.Println("Before Execution")
	return
}
func main() {
	a, b := Oper(20, 10)

	fmt.Println(a, b)

}
