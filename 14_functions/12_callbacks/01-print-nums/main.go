package main

import"fmt"

func visit(numbers []int, callback func(int)) {
	for _, n := range numbers {
		callback(n)
	}
}

func main() {
	visit([]int{5, 1, 2, 3, 4}, func(n int) {
		fmt.Println(n)
			})
}

// callback passing func as an argument