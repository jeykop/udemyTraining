package main

import "fmt"

func main() {
	fmt.Println(multiple("John ", "Daily "))
}

func multiple(fname string, lname string) (string, string) {
	return fmt.Sprint(fname, lname), fmt.Sprint(lname, fname)
}

