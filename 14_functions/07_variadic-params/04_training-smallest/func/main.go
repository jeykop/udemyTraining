package main

import ("fmt")

func smallest(numbers ...float64) float64 {
	var petite float64
	for i, v := range numbers {
				if v < petite || i == 0 {
			petite = v
		}
fmt.Println(v, i)
	}
	return petite
}

func main() {
	answer := smallest(70, 10, 20, 500)
	fmt.Println(answer)
}
