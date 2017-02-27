package main

import "fmt"

func main() {

	myGreeting := map[string]string{
		"Tim": "Good mornin!",
		"Jenny": "Bonjour!",
	}

	myGreeting["Harleen"] = "Howdy"

	fmt.Println(len(myGreeting))
}
