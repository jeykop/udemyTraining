package main

import "fmt"

func main() {

	for i := 1; i <= 1000000000; i++ {
counter := 0
		for j := 1; j <= 20; j++{
			if i % j == 0 {
				counter++
				if counter == 20 {
					fmt.Println("i - ",i)
					fmt.Println("j - ",j)
					fmt.Println("counter - ",counter)
				}
			}

		}
	}

}
