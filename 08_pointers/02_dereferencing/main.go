package main

import "fmt"

func main() {

	a := "Jeykop"

	fmt.Println(a)
	fmt.Println(&a)

	var b = &a

	fmt.Println(b)
	fmt.Printf("%T \n", b)
	fmt.Println(*b)
	fmt.Printf("%T \n", *b)
}
