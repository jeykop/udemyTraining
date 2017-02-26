package main

import "fmt"

func main() {

	mySlice := []int{1, 2, 3, 4, 5, 6, 7}
	myOtherSlice := []int{25, 9, 10, 11, 12}

	mySlice = append(mySlice, myOtherSlice...)

	fmt.Println(mySlice)
}
