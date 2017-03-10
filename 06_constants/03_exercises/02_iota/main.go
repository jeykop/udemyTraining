package main

import "fmt"

const (
	_ = iota
	A = 1 << (iota)
	B = 1 << (iota)
	C = 1 << (iota)
)

func main() {
	fmt.Printf("%b\n",A)
	fmt.Printf("%d\n",A)
	fmt.Printf("%b\n",B)
	fmt.Printf("%d\n",B)
	fmt.Printf("%b\n",C)
	fmt.Printf("%d\n",C)
}

