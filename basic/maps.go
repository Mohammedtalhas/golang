package main

import (
	"fmt"
	//"math"
)

func main() {

	//key value object
	var mp map[string]int = map[string]int{
		"apple":  1,
		"oramge": 5,
	}
	decmp := make(map[string]int)
	decmp = map[string]int{"a": 1, "c": 2, "b": 3}
	delete(decmp, "a")
	val, ok := mp["apple"]
	fmt.Println(val, ok)
	fmt.Println(mp)
	fmt.Println(decmp)

}
