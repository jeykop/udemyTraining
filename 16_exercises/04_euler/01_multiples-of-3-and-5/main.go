package main

import "fmt"

func main() {
	var total int
	for i := 1; i < 1000; i++ {

		if i%3 == 0 {
			total += i
		}
		if i%5 == 0 {
			total += i
		}

	}
	fmt.Println(total)
}
