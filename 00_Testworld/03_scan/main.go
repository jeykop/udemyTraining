package main

import "fmt"

func main() {
	var s, t string
	fmt.Println("Firstname Lastname")
	fmt.Scanln(&s, &t)
	fmt.Println("Firstname is \b",s,"\n", "Lastname is \b", t, "\n hallo")
	return
}