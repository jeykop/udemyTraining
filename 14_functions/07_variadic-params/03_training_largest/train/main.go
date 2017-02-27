package main

import "fmt"

func highest(sf ...float64) float64 {
	var higher float64
	for _, v := range sf {
		if v > higher {
			higher = v
		}

	}
	return higher
	}

func main(){
	n := highest(1, 10, 50, 7, 101)
	fmt.Println(n)
}