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
							}

			}
	fmt.Println(fibosum)
	}

/*
Each new term in the Fibonacci sequence is generated by adding the previous two terms.
By starting with 1 and 2, the first 10 terms will be:
1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...
By considering the terms in the Fibonacci sequence whose values do not exceed four million,
find the sum of the even-valued terms.
*/
