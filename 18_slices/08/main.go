package main

import "fmt"

func main() {

	costumerNumber := make([]int, 3)
	// 3 is length and capacity
	// // length - number of elements referred to by a slice
	// // capacity - number of elements in the underlying array
	costumerNumber[0] = 7
	costumerNumber[1] = 10
	costumerNumber[2] = 15

	fmt.Println(costumerNumber[0])
	fmt.Println(costumerNumber[1])
	fmt.Println(costumerNumber[2])

	greeting := make([]string, 3, 5)
	// 3 is length - number of elements referred to by a slice
	// 5 is capacity - number of elements in the underlying array
	// you could also do it like this
	greeting[0] = "Good morning!"
	greeting[1] = "Bonjour!"
	greeting[2] = "dias!"
	greeting = append(greeting, "suprabadram!")
	greeting = append(greeting, "Ohayou gozaimasu!")
	greeting = append(greeting, "gidday!")

	fmt.Println(greeting[5])
	fmt.Println(len(greeting))
	fmt.Println(cap(greeting))

}