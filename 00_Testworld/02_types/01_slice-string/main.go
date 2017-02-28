package main

import "fmt"

func main() {
	word := "Fuchsgrabengasse 24"
	fmt.Println(word[3])
	fmt.Println(word[:3])
	fmt.Println(word[2:5])

	for i := 0; i < len(word); i += 1 {
	fmt.Println(string(word[:i]))
}
}
