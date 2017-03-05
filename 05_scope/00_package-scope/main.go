package main

import "fmt"

	var a string = "this is stored in the variable a"
	var b, c string = "stored in b", "stored in c"
	var d string

func main(){
	d = "stored in d"
	var e int = 42
	f := 43
	g := "stored in g"
	h, i := "stored in h", "stored in i"
	j, k, l, m := 44.7, true, false, 'm'
	n := "n"
	o := `o`

	fmt.Printf("%T \n", a)
	fmt.Printf("%v \n", b)
	fmt.Println("c - ", c)
	fmt.Println("d - ", d)
	fmt.Println("e - ", e)
	fmt.Println("f - ", f)
	fmt.Println("g - ", g)
	fmt.Println("h - ", h)
	fmt.Println("i - ", i)
	fmt.Printf("%T \n", j)
	fmt.Printf("%T \n", k)
	fmt.Println("l - ", l)
	fmt.Printf("%T \n", m)
	fmt.Println("n - ", n)
	fmt.Printf("%T \n", o)

}
