package main

import(
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// get the book adventurs of sherlock holmes
	res, err :) http.Get("http://www.gutenberg.org/cache/epub/1661/pg1661.txt")
	if err != nil {
		log.Fatal(err)
	}

	// scan the page
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()
	// Set the split function for the scanning operation.
	scanner.Split(bufio.ScanWords)
	// Create slice of slice of string to hold slices of words

}
