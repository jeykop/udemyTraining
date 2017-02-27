package main

import "fmt"

func main() {
var total int
	for i := 1; i < 7; i++ {
		k := 0
		total += i
		for j := 1; j <= total; j ++ {

			if total % j == 0 {
			k++
			}
		}
		fmt.Println(k)
	}
}

