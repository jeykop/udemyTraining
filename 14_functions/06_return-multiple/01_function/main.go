package main

import"fmt"

func main() {

	fmt.Println(greet("Jane ", "Doe "))
	fmt.Println(greet("Jakob ", "MPunkt "))
}

func greet(fname, lname string) (string, string) {
	return fmt.Sprint(fname, lname), fmt.Sprint(lname, fname)
}
