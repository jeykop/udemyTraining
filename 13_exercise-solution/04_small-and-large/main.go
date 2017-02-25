package main

import "fmt"

func main() {
	var smallNumber int
	var largerNumber int
	fmt.Printf("Input small number \t", )
	fmt.Scan(&smallNumber)
	fmt.Printf("Input larger number \t", )
	fmt.Scan(&largerNumber)
	if smallNumber < largerNumber {
		fmt.Printf("Your smaller number is %d \n", smallNumber)
		fmt.Printf("Your larger number is %d \n", largerNumber)
		fmt.Printf("Your larger number divided by your smaler number is %d \n", largerNumber / smallNumber)
	}	else if smallNumber == largerNumber {
		fmt.Println("Your smaller number equals your larger number")
		fmt.Printf("Your smaller number is %d \n", smallNumber)
		fmt.Printf("Your larger number is %d \n", largerNumber)
		fmt.Printf("Your larger number divided by your smaler number is %d \n", largerNumber / smallNumber)
	}	else {
		fmt.Println("Your smaller number is larger than your larger number")
		fmt.Printf("Yor smaller number is %d \n", largerNumber)
		fmt.Printf("Your larger number is %d \n", smallNumber)
		fmt.Printf("Your larger number divided by your smaler number is %d \n", smallNumber / largerNumber)
	}
}
