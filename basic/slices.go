package main

import "fmt"

func main() {

	arr := [5]int{1, 2, 3, 4, 5}
	slice := make([]int, 5)
	var sli []int = []int{1, 2, 3, 4, 5, 8, 4, 7, 9, 12, 45, 66}
	var s []int = arr[1:3]
	sli = append(sli, 20)
	fmt.Println(cap(s))
	fmt.Println(cap(arr))
	fmt.Println(cap(sli))
	fmt.Println(sli)
	fmt.Println(slice)
	for i, ele := range sli {
		fmt.Println(i, ":", ele)

	}
	for i, ele := range sli {
		for j, ele2 := range sli {
			if ele == ele2 && j > i {
				fmt.Println(ele)

			}
		}

	}
}
