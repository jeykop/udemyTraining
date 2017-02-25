package main

import"fmt"

func main() {
	greet("Jane", "Doe")
	greet("Jakob", "MPunkt")
}

func greet(fname string, lname string) {
	fmt.Println(fname, lname)
}
