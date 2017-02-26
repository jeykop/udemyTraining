package main

import "fmt"

func largest(numbers ...float64) float64 {
	var biggest float64
	for _, v := range numbers {
		if v > biggest {
			biggest = v
		}
	}
return biggest
}

func main() {
	n := largest(0, 1, 5, 10, 50, 50.5)
	fmt.Println(n)
}