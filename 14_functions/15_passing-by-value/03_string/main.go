package main

import "fmt"

func main() {

	name := "Jeykop"
	fmt.Println(name) //Jeykop

	changeMe(name)

	fmt.Println(name) // Jeykop
}

func changeMe(z string) {
	fmt.Println(z) // Jeykop
	z = "Adrien"
	fmt.Println(z) // Adrien
}
