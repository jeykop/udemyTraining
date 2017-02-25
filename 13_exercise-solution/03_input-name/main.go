package main

import "fmt"

func main () {
	var yourName string
	fmt.Printf("Enter your name \t",)
	fmt.Scan(&yourName)
	fmt.Printf("Hello %s \t", yourName)

}