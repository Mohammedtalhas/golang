package main

import (
	"fmt"
	//"math"
)

func main() {
	var arr [5]int // cannot be of multiple data type like python and javascript
	arr[2] = 100
	arrnew := [5]int{1, 2, 3, 4, 5}
	arr2D := [2][2]int{{1, 2}, {1, 3}}
	arr2D[1][1] = 2
	fmt.Println(arrnew)
	fmt.Println(arr2D)
}
