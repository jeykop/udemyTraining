package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// get the book moby dick
	res, err := http.Get("http://www.gutenberg.org/cache/epub/1661/pg1661.txt")
	if err != nil {
		log.Fatal(err)
	}

	// scan the page
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()
	// Set the split function for the scanning operation.
	scanner.Split(bufio.ScanWords)
	// Creat slice to hold counts
	buckets := make([][]string, 12)
	// Loop over words
	for scanner.Scan() {
		word := scanner.Text()
		n := hashBucket(scanner.Text(), 12)
		buckets[n] = append(buckets[n], word)
	}
	// Print len of each bucker
	for i := 0; i < 12; i++ {
		fmt.Println(i, " - ", len(buckets[i]))
	}
	//Print the word in one of the buckets
	// fmt.Println(buckets[6])
	fmt.Println(len(buckets))
	fmt.Println(cap(buckets))
}

func hashBucket(word string, buckets int) int {
	var sum int
	for _, v := range word {
		sum += int(v)
	}
	return sum % buckets
}