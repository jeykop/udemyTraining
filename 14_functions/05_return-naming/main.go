package main

import "fmt"

func main() {

	fmt.Println(greet("Jane ", "Doe"))
	fmt.Println(greet("Jakob ", "MPunkt"))
}

func greet(fname string, lname string) (s string) {
	s = fmt.Sprint(fname, lname)
	return
}
