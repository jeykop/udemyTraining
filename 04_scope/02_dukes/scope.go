package main

import "fmt"

var X string = "Boss Hogg"

func oneFunc(){
	fmt.Println(X)
}

func twoFunc(){
	fmt.Println(X)
}

func main(){
	oneFunc()
	twoFunc()
}
