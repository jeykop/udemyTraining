package main

import "fmt"

func main() {

	age := 44
	fmt.Println(&age) // 0xc0420381c0

	changeMe(&age)

	fmt.Println(&age) // 0xc0420381c0
	fmt.Println(age) // 24
}

func changeMe(z *int) {
	fmt.Println(z) // 0xc0420381c0
	fmt.Println(*z) // 44
	*z = 24
	fmt.Println(z) // 0xc0420381c0
	fmt.Println(*z) // 24
}