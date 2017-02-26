package main

import "fmt"

func main() {
	mySlice := make([]int, 3,5)
	mySlice[0] = 2
	mySlice[1] =4
	mySlice[2] = 6
	mySlice[3] = 8

	fmt.Println(mySlice)
}
