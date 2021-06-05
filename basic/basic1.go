package main

import "fmt"

func main() {
	var name string //explicit declaration
	var number = 10 //implicit declaration
	name = "My self Mohammed"
	num := 15 //declaration types

	fmt.Printf("%d", num)
	fmt.Printf("%T", number)
	fmt.Println("Hello World!", name)

	var x = fmt.Sprintln("Hello World!", name) //storing values
	fmt.Println(x)

}
