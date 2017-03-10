package main

import (
	"fmt"

)
func main() {
	sum1 := 0
	sum2 := 0
	sqr := 0
	for i := 1; i <= 100; i++ {
		sum1 += i
		sum2 = sum1 * sum1
		sqr += i * i
	}
	fmt.Println(sum1)
	fmt.Println(sum2)
	fmt.Println(sqr)
	fmt.Println(sum2 - sqr)
}
