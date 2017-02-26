package main

import "fmt"


func average(numbers ...float64) float64 {
	var total float64
	for _, v := range numbers {
		total += v
	}
	return total / float64(len(numbers))
}

func main() {
	n := average(0, 1, 5, 10, 50, 50.5)
	fmt.Println(n)
}