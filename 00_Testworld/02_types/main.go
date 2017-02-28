package main

import (
	"fmt"
	"reflect"
)

func main() {
	fname := "Jeykop"
	lname := "Möstl"
	street := "Fuchsgrabengasse 24"
	zipcode := "5555"
	city := "Kogelhof"
	fmt.Println("\n",fname, "\b",lname, "\n",street, "\n",zipcode, "\b",city )
	fmt.Println(len(fname),len(lname),len(street),len(zipcode),len(city))
	get := 'A'
	fmt.Println(reflect.TypeOf(get), " - ", get)
	fmt.Println(fname + " - " + lname, street)
}
