package main

import (
	"fmt"
	//"math"
)

func Oper(a, b float64) (x, y float64) {
	defer fmt.Println("Executed Successfully")
	//labeled return
	x = a + b
	y = a - b
	fmt.Println("Before Execution")
	return
}
func main() {
	a := Oper
	//reference Function
	c, d := a(10, 5)
	fmt.Println(c, d)

}
