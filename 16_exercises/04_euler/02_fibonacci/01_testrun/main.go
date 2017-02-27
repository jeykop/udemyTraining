package main

import "fmt"

func main() {
	fiboneu := 1
	fiboalt := 1
	var fibosum int
	for i := 1; i <= 4000000; i++ {

				if i == fiboneu {

				fiboneu = fiboneu + fiboalt
					fiboalt = fiboneu - fiboalt
				if fiboneu%2 == 0 {
					fibosum = fibosum + fiboneu
				}
				fmt.Println(fiboneu)
			}

			}
	fmt.Println(fibosum)
	}


