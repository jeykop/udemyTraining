package main

import "fmt"

func main() {

	greet("John")
	greet("Jane")
	greet("3")
}

func greet(name string) {
	fmt.Println(name)
}

// Greet was declared with a parameter
// When calling greet(), pass in an argument