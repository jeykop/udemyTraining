package main

import "fmt"

func main() {

	data := [] float64{43, 56, 12, 45, 57, 50}
	n := highest(data...)
	fmt.Println(n)
}

func highest(sf ...float64) float64 {
	fmt.Println(sf)
	fmt.Println(len(sf))
	fmt.Printf("%T, \n", sf)
	var highest float64
	for _, v := range sf {
		if v > highest {
			highest = v
		}
	}

	return highest
}