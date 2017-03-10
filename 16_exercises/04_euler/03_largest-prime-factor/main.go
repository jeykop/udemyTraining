package main

import "fmt"
var x int = 600851475143
func main() {
	for i := 1; i <= x; i++ {
			if 600851475143  % i == 0 {
				fmt.Println(i)
				x = x / i
			}
	}
}