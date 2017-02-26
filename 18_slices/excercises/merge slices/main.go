package main

import "fmt"

func main() {

	sliceOfInts := make([]int, 5, 10)
sliceOfInts[0] = 1
	sliceOfInts[1] = 2
	sliceOfInts[2] = 3
	sliceOfInts[3] = 4
	sliceOfInts[4] = 5

	fmt.Println(sliceOfInts)
}
