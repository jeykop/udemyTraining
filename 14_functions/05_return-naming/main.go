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

// Occasionally named returns are useful. Read this article for more information:
// https://www.goinggo.net/2013/10/functions-and-naked-returns-in-go.html