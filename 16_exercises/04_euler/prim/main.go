package main

import "fmt"

func main() {
	var top int
	for i := 1; 10; i++ {
		k := 0
		top = i
		for j := 1; top; j++{

			if i%j == 0{
				k++
				if k == 2{
					fmt.Println(i)
				}
			}
		}

	}
}