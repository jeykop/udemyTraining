package main

import "fmt"

func main() {

	name := "Jeykop"
	fmt.Println(&name) // 0xc0420381c0

	changeMe(&name)

	fmt.Println(&name) //  0xc0420381c0
	fmt.Println(name) // Adrien

}

func changeMe(z *string) {
	fmt.Println(*z) // 0xc0420381c0
	fmt.Println(z) // Jeykop
	*z = "Adrien"
	fmt.Println(*z) // 0xc0420381c0
	fmt.Println(z) // Adrien
}
