package main

import "fmt"

func zero(z int) {
	z = 0
	fmt.Printf("%p\n", &z)
	fmt.Println(&z)
}

func main() {
	x := 5
	zero(x)
	fmt.Printf("%p\n", &x)
	fmt.Println(&x)
	fmt.Println(x)
}
