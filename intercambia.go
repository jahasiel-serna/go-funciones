package main

import "fmt"

func intercambia(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	var a, b int
	fmt.Scan(&a)
	fmt.Scan(&b)
	intercambia(&a, &b)
	fmt.Println("a =", a, ", b = ", b)
}