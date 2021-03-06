package main

import "fmt"

func average(sf ...float64) (float64) {
	var total float64
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}

func main() {

	n := average(43, 56, 87, 12, 45, 57, 50)
	fmt.Println(n)
}
