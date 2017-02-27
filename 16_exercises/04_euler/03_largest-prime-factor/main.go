package main

import "fmt"

func main() {
	for i := 1; i < 600851475143; i++ {
			if 600851475143 % i == 0 {
				fmt.Println(i)

			}
	}
}