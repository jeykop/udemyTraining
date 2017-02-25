package main

import "fmt"

func main() {
		x := 0
	for i := 1; i <= 999; i++ {
		if i%3 == 0 {
			x += (i)
		} else if i%5 == 0 {
			x += (i)
		}
	}
	fmt.Println(x)
}
