package main

import "fmt"

func main() {
	myGreeting := map[int]string{
		1: "Jakob",
		2:  "Anna",
		3:  "Kathi",
	}
	if val, exists := myGreeting[7]; exists {
		fmt.Println("val: ", val)
		fmt.Println("exists: ", exists)
	} else {
		fmt.Println("That value does not exist")
		fmt.Println("val: ", val)
		fmt.Println("exists: ", exists)
	}
	fmt.Println(myGreeting)
	myGreeting [1] = "Michi"
	myGreeting [0] = "Gabi"
	fmt.Println(myGreeting)
	delete(myGreeting, 0)
	fmt.Println(myGreeting)
for key, val := range myGreeting {
	fmt.Println(key, " - ", val)
}
	fmt.Println(len(myGreeting))
}