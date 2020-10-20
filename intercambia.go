package main

import "fmt"

func intercambia(a, b *int) {
	t := *a
	*a = *b
	*b = t
}

func main() {
	var a, b int
	fmt.Scan(&a)
	fmt.Scan(&b)
	intercambia(&a, &b)
	fmt.Println("a =", a, ", b = ", b)
}