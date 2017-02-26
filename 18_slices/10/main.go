package main

import "fmt"

func main() {

	mySlice := []string{"monday", "tuesday", "wednesday", "thursday"}
	myOtherSlice := []string{"friday", "saturday", "sunday"}

	mySlice = append(mySlice, myOtherSlice...)
	fmt.Println(mySlice)

	mySlice = append(mySlice[:0], mySlice[3:]...)
	fmt.Println(mySlice)
}
