package main

import (
	"fmt"
	//"math"
)

func main() {

	arr := [5]int{1, 2, 3, 4, 5}
	slice := make([]int, 5)
	var sli []int = []int{1, 2, 3, 4, 5}
	var s []int = arr[1:3]
	sli = append(sli, 20)
	fmt.Println(cap(s))
	fmt.Println(cap(arr))
	fmt.Println(cap(sli))
	fmt.Println(sli)
	fmt.Println(slice)
}
