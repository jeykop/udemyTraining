package main

import "fmt"

func reverse(s string) string {
	a := []rune(s)
	for i, j := 0, len(a)-1; i < len(a)/2;i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	fmt.Println(string(a))
	return string(a)

}

func main() {
	reverse("Möstl Michael Jakob")
}
