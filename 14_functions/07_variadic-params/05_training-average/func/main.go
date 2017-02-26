package main

import "fmt"

func average(numbers ...float64)float64{
	var total float64
	for _, v := range numbers {
	total += v
	}
	return total / float64(len(numbers))
}

func main() {
	n := average(1,10,100,1000,10000)
	fmt.Println(n)
}