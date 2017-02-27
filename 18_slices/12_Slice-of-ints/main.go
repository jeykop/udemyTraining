package main

import "fmt"

func main() {

	transactions := make([][]int, 0, 3)

	for i := 0; i < 5; i++ {
		transaction := make([]int, 0)
		for j := 0; j < 10; j++ {
			transaction = append(transaction, j)
		}
		transactions = append(transactions, transaction)
	}
	fmt.Println(transactions)
}
