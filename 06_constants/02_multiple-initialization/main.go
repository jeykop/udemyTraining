package main

import "fmt"

const (
	_	= 	iota
	B	= 	1 << (iota * 10)
	C	=	1 << (iota * 10)
)

const (
	D	= iota
	E
	F
)

func main() {

	fmt.Printf("%b\t", B)
	fmt.Printf("%d\n", B)
	fmt.Printf("%b\t", C)
	fmt.Printf("%d\n", C)
	fmt.Println(D)
	fmt.Println(E)
	fmt.Println(F)
}