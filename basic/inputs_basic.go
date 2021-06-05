package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	/* fmt.Print("Enter Your Name:\t")
	scanner.Scan()
	input := scanner.Text()
	fmt.Printf("Hello %v", input) */
	fmt.Print("Enter Your Year Of Birth:\t")
	scanner.Scan()
	input, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {

		fmt.Println(err)
	} else {

		fmt.Printf("You are %v years old", 2021-input)
	}

}
