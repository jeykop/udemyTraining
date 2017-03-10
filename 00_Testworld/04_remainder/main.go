package main

import "fmt"

func main() {
	var x int
	var y int
	var rem int
	fmt.Scanln(&x)
	fmt.Scanln(&y)
	rem = (x % y)
	fmt.Println(rem)

}
